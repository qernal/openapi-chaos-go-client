# MetricResourceAggregationResourcesBucketsInnerHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]MetricResourceAggregationResourcesBucketsInnerHistogramBucketsInner**](MetricResourceAggregationResourcesBucketsInnerHistogramBucketsInner.md) | Array of resource usage by interval  &gt; Note: A metric will have either a &#x60;counter&#x60; or &#x60;gauge&#x60; value  | [optional] 

## Methods

### NewMetricResourceAggregationResourcesBucketsInnerHistogram

`func NewMetricResourceAggregationResourcesBucketsInnerHistogram() *MetricResourceAggregationResourcesBucketsInnerHistogram`

NewMetricResourceAggregationResourcesBucketsInnerHistogram instantiates a new MetricResourceAggregationResourcesBucketsInnerHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResourceAggregationResourcesBucketsInnerHistogramWithDefaults

`func NewMetricResourceAggregationResourcesBucketsInnerHistogramWithDefaults() *MetricResourceAggregationResourcesBucketsInnerHistogram`

NewMetricResourceAggregationResourcesBucketsInnerHistogramWithDefaults instantiates a new MetricResourceAggregationResourcesBucketsInnerHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MetricResourceAggregationResourcesBucketsInnerHistogram) GetBuckets() []MetricResourceAggregationResourcesBucketsInnerHistogramBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MetricResourceAggregationResourcesBucketsInnerHistogram) GetBucketsOk() (*[]MetricResourceAggregationResourcesBucketsInnerHistogramBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MetricResourceAggregationResourcesBucketsInnerHistogram) SetBuckets(v []MetricResourceAggregationResourcesBucketsInnerHistogramBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MetricResourceAggregationResourcesBucketsInnerHistogram) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


