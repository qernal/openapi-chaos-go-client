# MetricHttpAggregationHttpCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]MetricHttpAggregationHttpCodesBucketsInner**](MetricHttpAggregationHttpCodesBucketsInner.md) | Array of unique http status codes | [optional] 
**DocCountErrorUpperBound** | Pointer to **int32** | Upper bound of error in document count | [optional] 
**SumOtherDocCount** | Pointer to **int32** | Sum of other document counts | [optional] 

## Methods

### NewMetricHttpAggregationHttpCodes

`func NewMetricHttpAggregationHttpCodes() *MetricHttpAggregationHttpCodes`

NewMetricHttpAggregationHttpCodes instantiates a new MetricHttpAggregationHttpCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricHttpAggregationHttpCodesWithDefaults

`func NewMetricHttpAggregationHttpCodesWithDefaults() *MetricHttpAggregationHttpCodes`

NewMetricHttpAggregationHttpCodesWithDefaults instantiates a new MetricHttpAggregationHttpCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MetricHttpAggregationHttpCodes) GetBuckets() []MetricHttpAggregationHttpCodesBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MetricHttpAggregationHttpCodes) GetBucketsOk() (*[]MetricHttpAggregationHttpCodesBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MetricHttpAggregationHttpCodes) SetBuckets(v []MetricHttpAggregationHttpCodesBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MetricHttpAggregationHttpCodes) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetDocCountErrorUpperBound

`func (o *MetricHttpAggregationHttpCodes) GetDocCountErrorUpperBound() int32`

GetDocCountErrorUpperBound returns the DocCountErrorUpperBound field if non-nil, zero value otherwise.

### GetDocCountErrorUpperBoundOk

`func (o *MetricHttpAggregationHttpCodes) GetDocCountErrorUpperBoundOk() (*int32, bool)`

GetDocCountErrorUpperBoundOk returns a tuple with the DocCountErrorUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCountErrorUpperBound

`func (o *MetricHttpAggregationHttpCodes) SetDocCountErrorUpperBound(v int32)`

SetDocCountErrorUpperBound sets DocCountErrorUpperBound field to given value.

### HasDocCountErrorUpperBound

`func (o *MetricHttpAggregationHttpCodes) HasDocCountErrorUpperBound() bool`

HasDocCountErrorUpperBound returns a boolean if a field has been set.

### GetSumOtherDocCount

`func (o *MetricHttpAggregationHttpCodes) GetSumOtherDocCount() int32`

GetSumOtherDocCount returns the SumOtherDocCount field if non-nil, zero value otherwise.

### GetSumOtherDocCountOk

`func (o *MetricHttpAggregationHttpCodes) GetSumOtherDocCountOk() (*int32, bool)`

GetSumOtherDocCountOk returns a tuple with the SumOtherDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumOtherDocCount

`func (o *MetricHttpAggregationHttpCodes) SetSumOtherDocCount(v int32)`

SetSumOtherDocCount sets SumOtherDocCount field to given value.

### HasSumOtherDocCount

`func (o *MetricHttpAggregationHttpCodes) HasSumOtherDocCount() bool`

HasSumOtherDocCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


