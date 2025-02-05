# MetricHttpAggregationHttpCodesBucketsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocCount** | Pointer to **int32** | Number of documents in the bucket | [optional] 
**Histogram** | Pointer to [**MetricHttpAggregationHttpCodesBucketsInnerHistogram**](MetricHttpAggregationHttpCodesBucketsInnerHistogram.md) |  | [optional] 
**Key** | Pointer to **string** | HTTP status code, typical values will be;  - http-2xx - http-3xx - http-4xx - http-5xx  &gt; Note: the &#39;xx&#39; is intentional and literal, all status codes within that range will be grouped  | [optional] 

## Methods

### NewMetricHttpAggregationHttpCodesBucketsInner

`func NewMetricHttpAggregationHttpCodesBucketsInner() *MetricHttpAggregationHttpCodesBucketsInner`

NewMetricHttpAggregationHttpCodesBucketsInner instantiates a new MetricHttpAggregationHttpCodesBucketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricHttpAggregationHttpCodesBucketsInnerWithDefaults

`func NewMetricHttpAggregationHttpCodesBucketsInnerWithDefaults() *MetricHttpAggregationHttpCodesBucketsInner`

NewMetricHttpAggregationHttpCodesBucketsInnerWithDefaults instantiates a new MetricHttpAggregationHttpCodesBucketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocCount

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *MetricHttpAggregationHttpCodesBucketsInner) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *MetricHttpAggregationHttpCodesBucketsInner) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetHistogram

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetHistogram() MetricHttpAggregationHttpCodesBucketsInnerHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetHistogramOk() (*MetricHttpAggregationHttpCodesBucketsInnerHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *MetricHttpAggregationHttpCodesBucketsInner) SetHistogram(v MetricHttpAggregationHttpCodesBucketsInnerHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *MetricHttpAggregationHttpCodesBucketsInner) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetKey

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricHttpAggregationHttpCodesBucketsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricHttpAggregationHttpCodesBucketsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricHttpAggregationHttpCodesBucketsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


