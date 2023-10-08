# ListAuthTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 
**Data** | [**[]AuthTokenMeta**](AuthTokenMeta.md) |  | 

## Methods

### NewListAuthTokens

`func NewListAuthTokens(meta PaginationMeta, data []AuthTokenMeta, ) *ListAuthTokens`

NewListAuthTokens instantiates a new ListAuthTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthTokensWithDefaults

`func NewListAuthTokensWithDefaults() *ListAuthTokens`

NewListAuthTokensWithDefaults instantiates a new ListAuthTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListAuthTokens) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAuthTokens) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAuthTokens) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *ListAuthTokens) GetData() []AuthTokenMeta`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAuthTokens) GetDataOk() (*[]AuthTokenMeta, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAuthTokens) SetData(v []AuthTokenMeta)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


