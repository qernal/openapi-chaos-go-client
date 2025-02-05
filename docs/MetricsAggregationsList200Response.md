# MetricsAggregationsList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**MetricResourceAggregationResources**](MetricResourceAggregationResources.md) |  | [optional] 
**HttpCodes** | Pointer to [**MetricHttpAggregationHttpCodes**](MetricHttpAggregationHttpCodes.md) |  | [optional] 

## Methods

### NewMetricsAggregationsList200Response

`func NewMetricsAggregationsList200Response() *MetricsAggregationsList200Response`

NewMetricsAggregationsList200Response instantiates a new MetricsAggregationsList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsAggregationsList200ResponseWithDefaults

`func NewMetricsAggregationsList200ResponseWithDefaults() *MetricsAggregationsList200Response`

NewMetricsAggregationsList200ResponseWithDefaults instantiates a new MetricsAggregationsList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *MetricsAggregationsList200Response) GetResources() MetricResourceAggregationResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MetricsAggregationsList200Response) GetResourcesOk() (*MetricResourceAggregationResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MetricsAggregationsList200Response) SetResources(v MetricResourceAggregationResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MetricsAggregationsList200Response) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetHttpCodes

`func (o *MetricsAggregationsList200Response) GetHttpCodes() MetricHttpAggregationHttpCodes`

GetHttpCodes returns the HttpCodes field if non-nil, zero value otherwise.

### GetHttpCodesOk

`func (o *MetricsAggregationsList200Response) GetHttpCodesOk() (*MetricHttpAggregationHttpCodes, bool)`

GetHttpCodesOk returns a tuple with the HttpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCodes

`func (o *MetricsAggregationsList200Response) SetHttpCodes(v MetricHttpAggregationHttpCodes)`

SetHttpCodes sets HttpCodes field to given value.

### HasHttpCodes

`func (o *MetricsAggregationsList200Response) HasHttpCodes() bool`

HasHttpCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


