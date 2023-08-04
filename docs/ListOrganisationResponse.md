# ListOrganisationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Data** | Pointer to [**[]OrganisationResponse**](OrganisationResponse.md) |  | [optional] 

## Methods

### NewListOrganisationResponse

`func NewListOrganisationResponse() *ListOrganisationResponse`

NewListOrganisationResponse instantiates a new ListOrganisationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganisationResponseWithDefaults

`func NewListOrganisationResponseWithDefaults() *ListOrganisationResponse`

NewListOrganisationResponseWithDefaults instantiates a new ListOrganisationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListOrganisationResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOrganisationResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOrganisationResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOrganisationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *ListOrganisationResponse) GetData() []OrganisationResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOrganisationResponse) GetDataOk() (*[]OrganisationResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOrganisationResponse) SetData(v []OrganisationResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ListOrganisationResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


