# GetOrganisationsPageParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **int32** | Offset | [optional] 
**After** | Pointer to **int32** | Offset | [optional] 
**Size** | Pointer to **int32** | Limit | [optional] 

## Methods

### NewGetOrganisationsPageParameter

`func NewGetOrganisationsPageParameter() *GetOrganisationsPageParameter`

NewGetOrganisationsPageParameter instantiates a new GetOrganisationsPageParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganisationsPageParameterWithDefaults

`func NewGetOrganisationsPageParameterWithDefaults() *GetOrganisationsPageParameter`

NewGetOrganisationsPageParameterWithDefaults instantiates a new GetOrganisationsPageParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *GetOrganisationsPageParameter) GetBefore() int32`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *GetOrganisationsPageParameter) GetBeforeOk() (*int32, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *GetOrganisationsPageParameter) SetBefore(v int32)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *GetOrganisationsPageParameter) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetAfter

`func (o *GetOrganisationsPageParameter) GetAfter() int32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *GetOrganisationsPageParameter) GetAfterOk() (*int32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *GetOrganisationsPageParameter) SetAfter(v int32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *GetOrganisationsPageParameter) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetSize

`func (o *GetOrganisationsPageParameter) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetOrganisationsPageParameter) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetOrganisationsPageParameter) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetOrganisationsPageParameter) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


