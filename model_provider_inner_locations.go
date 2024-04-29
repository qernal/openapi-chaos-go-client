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
)

// checks if the ProviderInnerLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderInnerLocations{}

// ProviderInnerLocations Locations at varying levels this provider operates within
type ProviderInnerLocations struct {
	Continents []string `json:"continents,omitempty"`
	Countries []string `json:"countries,omitempty"`
	Cities []string `json:"cities,omitempty"`
}

// NewProviderInnerLocations instantiates a new ProviderInnerLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderInnerLocations() *ProviderInnerLocations {
	this := ProviderInnerLocations{}
	return &this
}

// NewProviderInnerLocationsWithDefaults instantiates a new ProviderInnerLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderInnerLocationsWithDefaults() *ProviderInnerLocations {
	this := ProviderInnerLocations{}
	return &this
}

// GetContinents returns the Continents field value if set, zero value otherwise.
func (o *ProviderInnerLocations) GetContinents() []string {
	if o == nil || IsNil(o.Continents) {
		var ret []string
		return ret
	}
	return o.Continents
}

// GetContinentsOk returns a tuple with the Continents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInnerLocations) GetContinentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Continents) {
		return nil, false
	}
	return o.Continents, true
}

// HasContinents returns a boolean if a field has been set.
func (o *ProviderInnerLocations) HasContinents() bool {
	if o != nil && !IsNil(o.Continents) {
		return true
	}

	return false
}

// SetContinents gets a reference to the given []string and assigns it to the Continents field.
func (o *ProviderInnerLocations) SetContinents(v []string) {
	o.Continents = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *ProviderInnerLocations) GetCountries() []string {
	if o == nil || IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInnerLocations) GetCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *ProviderInnerLocations) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *ProviderInnerLocations) SetCountries(v []string) {
	o.Countries = v
}

// GetCities returns the Cities field value if set, zero value otherwise.
func (o *ProviderInnerLocations) GetCities() []string {
	if o == nil || IsNil(o.Cities) {
		var ret []string
		return ret
	}
	return o.Cities
}

// GetCitiesOk returns a tuple with the Cities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInnerLocations) GetCitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cities) {
		return nil, false
	}
	return o.Cities, true
}

// HasCities returns a boolean if a field has been set.
func (o *ProviderInnerLocations) HasCities() bool {
	if o != nil && !IsNil(o.Cities) {
		return true
	}

	return false
}

// SetCities gets a reference to the given []string and assigns it to the Cities field.
func (o *ProviderInnerLocations) SetCities(v []string) {
	o.Cities = v
}

func (o ProviderInnerLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderInnerLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Continents) {
		toSerialize["continents"] = o.Continents
	}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !IsNil(o.Cities) {
		toSerialize["cities"] = o.Cities
	}
	return toSerialize, nil
}

type NullableProviderInnerLocations struct {
	value *ProviderInnerLocations
	isSet bool
}

func (v NullableProviderInnerLocations) Get() *ProviderInnerLocations {
	return v.value
}

func (v *NullableProviderInnerLocations) Set(val *ProviderInnerLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderInnerLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderInnerLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderInnerLocations(val *ProviderInnerLocations) *NullableProviderInnerLocations {
	return &NullableProviderInnerLocations{value: val, isSet: true}
}

func (v NullableProviderInnerLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderInnerLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

