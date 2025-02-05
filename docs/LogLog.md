# LogLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | Pointer to **string** | Which log stream | [optional] 
**Kind** | Pointer to **string** | If this was an event on the function or a log line | [optional] 
**Labels** | Pointer to **[]string** | An array of labels | [optional] 
**Type** | Pointer to **string** | Log line type | [optional] 
**Line** | Pointer to **string** | Log line | [optional] 
**Timestamp** | Pointer to **string** | The date/time that this log was generated | [optional] 

## Methods

### NewLogLog

`func NewLogLog() *LogLog`

NewLogLog instantiates a new LogLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogLogWithDefaults

`func NewLogLogWithDefaults() *LogLog`

NewLogLogWithDefaults instantiates a new LogLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *LogLog) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *LogLog) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *LogLog) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *LogLog) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetKind

`func (o *LogLog) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogLog) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogLog) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LogLog) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *LogLog) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LogLog) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LogLog) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LogLog) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *LogLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogLog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLine

`func (o *LogLog) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *LogLog) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *LogLog) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *LogLog) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogLog) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogLog) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogLog) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


