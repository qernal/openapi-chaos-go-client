# MetricResourceAggregationResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]MetricResourceAggregationResourcesBucketsInner**](MetricResourceAggregationResourcesBucketsInner.md) | Array of unqiue resources | [optional] 
**DocCountErrorUpperBound** | Pointer to **int32** | Upper bound of error in document count | [optional] 
**SumOtherDocCount** | Pointer to **int32** | Sum of other document counts | [optional] 

## Methods

### NewMetricResourceAggregationResources

`func NewMetricResourceAggregationResources() *MetricResourceAggregationResources`

NewMetricResourceAggregationResources instantiates a new MetricResourceAggregationResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResourceAggregationResourcesWithDefaults

`func NewMetricResourceAggregationResourcesWithDefaults() *MetricResourceAggregationResources`

NewMetricResourceAggregationResourcesWithDefaults instantiates a new MetricResourceAggregationResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MetricResourceAggregationResources) GetBuckets() []MetricResourceAggregationResourcesBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MetricResourceAggregationResources) GetBucketsOk() (*[]MetricResourceAggregationResourcesBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MetricResourceAggregationResources) SetBuckets(v []MetricResourceAggregationResourcesBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MetricResourceAggregationResources) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetDocCountErrorUpperBound

`func (o *MetricResourceAggregationResources) GetDocCountErrorUpperBound() int32`

GetDocCountErrorUpperBound returns the DocCountErrorUpperBound field if non-nil, zero value otherwise.

### GetDocCountErrorUpperBoundOk

`func (o *MetricResourceAggregationResources) GetDocCountErrorUpperBoundOk() (*int32, bool)`

GetDocCountErrorUpperBoundOk returns a tuple with the DocCountErrorUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCountErrorUpperBound

`func (o *MetricResourceAggregationResources) SetDocCountErrorUpperBound(v int32)`

SetDocCountErrorUpperBound sets DocCountErrorUpperBound field to given value.

### HasDocCountErrorUpperBound

`func (o *MetricResourceAggregationResources) HasDocCountErrorUpperBound() bool`

HasDocCountErrorUpperBound returns a boolean if a field has been set.

### GetSumOtherDocCount

`func (o *MetricResourceAggregationResources) GetSumOtherDocCount() int32`

GetSumOtherDocCount returns the SumOtherDocCount field if non-nil, zero value otherwise.

### GetSumOtherDocCountOk

`func (o *MetricResourceAggregationResources) GetSumOtherDocCountOk() (*int32, bool)`

GetSumOtherDocCountOk returns a tuple with the SumOtherDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOtherDocCount

`func (o *MetricResourceAggregationResources) SetSumOtherDocCount(v int32)`

SetSumOtherDocCount sets SumOtherDocCount field to given value.

### HasSumOtherDocCount

`func (o *MetricResourceAggregationResources) HasSumOtherDocCount() bool`

HasSumOtherDocCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


