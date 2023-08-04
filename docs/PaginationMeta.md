# PaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **float32** |  | 
**Start** | **float32** |  | 
**End** | **float32** |  | 
**Pages** | **float32** |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewPaginationMeta

`func NewPaginationMeta(results float32, start float32, end float32, pages float32, links PaginationLinks, ) *PaginationMeta`

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

`func (o *PaginationMeta) GetResults() float32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginationMeta) GetResultsOk() (*float32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginationMeta) SetResults(v float32)`

SetResults sets Results field to given value.


### GetStart

`func (o *PaginationMeta) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PaginationMeta) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PaginationMeta) SetStart(v float32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *PaginationMeta) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PaginationMeta) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PaginationMeta) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetPages

`func (o *PaginationMeta) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PaginationMeta) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PaginationMeta) SetPages(v float32)`

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


