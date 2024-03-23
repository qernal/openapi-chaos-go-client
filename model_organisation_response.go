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

// checks if the OrganisationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganisationResponse{}

// OrganisationResponse Organisation response
type OrganisationResponse struct {
	// Organisation id
	Id string `json:"id"`
	// User id
	UserId string `json:"user_id"`
	// Organisation name
	Name string `json:"name"`
	Date Date `json:"date"`
}

type _OrganisationResponse OrganisationResponse

// NewOrganisationResponse instantiates a new OrganisationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationResponse(id string, userId string, name string, date Date) *OrganisationResponse {
	this := OrganisationResponse{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Date = date
	return &this
}

// NewOrganisationResponseWithDefaults instantiates a new OrganisationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationResponseWithDefaults() *OrganisationResponse {
	this := OrganisationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganisationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganisationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganisationResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *OrganisationResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OrganisationResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *OrganisationResponse) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *OrganisationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganisationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganisationResponse) SetName(v string) {
	o.Name = v
}

// GetDate returns the Date field value
func (o *OrganisationResponse) GetDate() Date {
	if o == nil {
		var ret Date
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *OrganisationResponse) GetDateOk() (*Date, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *OrganisationResponse) SetDate(v Date) {
	o.Date = v
}

func (o OrganisationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganisationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

func (o *OrganisationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOrganisationResponse := _OrganisationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganisationResponse)

	if err != nil {
		return err
	}

	*o = OrganisationResponse(varOrganisationResponse)

	return err
}

type NullableOrganisationResponse struct {
	value *OrganisationResponse
	isSet bool
}

func (v NullableOrganisationResponse) Get() *OrganisationResponse {
	return v.value
}

func (v *NullableOrganisationResponse) Set(val *OrganisationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationResponse(val *OrganisationResponse) *NullableOrganisationResponse {
	return &NullableOrganisationResponse{value: val, isSet: true}
}

func (v NullableOrganisationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


