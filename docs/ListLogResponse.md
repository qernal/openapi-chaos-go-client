# ListLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 
**Data** | [**[]Log**](Log.md) |  | 

## Methods

### NewListLogResponse

`func NewListLogResponse(meta PaginationMeta, data []Log, ) *ListLogResponse`

NewListLogResponse instantiates a new ListLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogResponseWithDefaults

`func NewListLogResponseWithDefaults() *ListLogResponse`

NewListLogResponseWithDefaults instantiates a new ListLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListLogResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLogResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLogResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *ListLogResponse) GetData() []Log`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLogResponse) GetDataOk() (*[]Log, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLogResponse) SetData(v []Log)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


