/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_chaos_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthToken{}

// AuthToken API auth token
type AuthToken struct {
	// Auth token uuid
	Id string `json:"id"`
	// User
	UserId string `json:"user_id"`
	// Name of token
	Name string `json:"name"`
	// When the token expires
	ExpiryAt *string `json:"expiry_at,omitempty"`
	// Combined token required for requesting an access token, this field is only returned once on creation or update (during regeneration).
	Token *string `json:"token,omitempty"`
	Date Date `json:"date"`
}

type _AuthToken AuthToken

// NewAuthToken instantiates a new AuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthToken(id string, userId string, name string, date Date) *AuthToken {
	this := AuthToken{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Date = date
	return &this
}

// NewAuthTokenWithDefaults instantiates a new AuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenWithDefaults() *AuthToken {
	this := AuthToken{}
	return &this
}

// GetId returns the Id field value
func (o *AuthToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthToken) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *AuthToken) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AuthToken) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *AuthToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthToken) SetName(v string) {
	o.Name = v
}

// GetExpiryAt returns the ExpiryAt field value if set, zero value otherwise.
func (o *AuthToken) GetExpiryAt() string {
	if o == nil || IsNil(o.ExpiryAt) {
		var ret string
		return ret
	}
	return *o.ExpiryAt
}

// GetExpiryAtOk returns a tuple with the ExpiryAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthToken) GetExpiryAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryAt) {
		return nil, false
	}
	return o.ExpiryAt, true
}

// HasExpiryAt returns a boolean if a field has been set.
func (o *AuthToken) HasExpiryAt() bool {
	if o != nil && !IsNil(o.ExpiryAt) {
		return true
	}

	return false
}

// SetExpiryAt gets a reference to the given string and assigns it to the ExpiryAt field.
func (o *AuthToken) SetExpiryAt(v string) {
	o.ExpiryAt = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthToken) SetToken(v string) {
	o.Token = &v
}

// GetDate returns the Date field value
func (o *AuthToken) GetDate() Date {
	if o == nil {
		var ret Date
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetDateOk() (*Date, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *AuthToken) SetDate(v Date) {
	o.Date = v
}

func (o AuthToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	if !IsNil(o.ExpiryAt) {
		toSerialize["expiry_at"] = o.ExpiryAt
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

func (o *AuthToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"name",
		"date",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthToken := _AuthToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthToken)

	if err != nil {
		return err
	}

	*o = AuthToken(varAuthToken)

	return err
}

type NullableAuthToken struct {
	value *AuthToken
	isSet bool
}

func (v NullableAuthToken) Get() *AuthToken {
	return v.value
}

func (v *NullableAuthToken) Set(val *AuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthToken(val *AuthToken) *NullableAuthToken {
	return &NullableAuthToken{value: val, isSet: true}
}

func (v NullableAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


