# ListProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 
**Data** | [**[]ProjectResponse**](ProjectResponse.md) |  | 

## Methods

### NewListProjectResponse

`func NewListProjectResponse(meta PaginationMeta, data []ProjectResponse, ) *ListProjectResponse`

NewListProjectResponse instantiates a new ListProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProjectResponseWithDefaults

`func NewListProjectResponseWithDefaults() *ListProjectResponse`

NewListProjectResponseWithDefaults instantiates a new ListProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListProjectResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListProjectResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListProjectResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *ListProjectResponse) GetData() []ProjectResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListProjectResponse) GetDataOk() (*[]ProjectResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListProjectResponse) SetData(v []ProjectResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


