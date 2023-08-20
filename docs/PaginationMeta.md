# PaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **int32** |  | 
**Start** | **int32** |  | 
**End** | **int32** |  | 
**Pages** | **int32** |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewPaginationMeta

`func NewPaginationMeta(results int32, start int32, end int32, pages int32, links PaginationLinks, ) *PaginationMeta`

NewPaginationMeta instantiates a new PaginationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationMetaWithDefaults

`func NewPaginationMetaWithDefaults() *PaginationMeta`

NewPaginationMetaWithDefaults instantiates a new PaginationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginationMeta) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginationMeta) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginationMeta) SetResults(v int32)`

SetResults sets Results field to given value.


### GetStart

`func (o *PaginationMeta) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PaginationMeta) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PaginationMeta) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *PaginationMeta) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PaginationMeta) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PaginationMeta) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetPages

`func (o *PaginationMeta) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PaginationMeta) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PaginationMeta) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetLinks

`func (o *PaginationMeta) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginationMeta) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginationMeta) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


