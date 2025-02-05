# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | Container ID the log line is for | [optional] 
**Function** | Pointer to **string** | Function ID the log line is for | [optional] 
**Project** | Pointer to **string** | Project ID the log line is for | [optional] 
**Organisation** | Pointer to **string** | Organisation ID the log line is for | [optional] 
**Group** | Pointer to **string** | Group ID the log line is for | [optional] 
**Log** | Pointer to [**LogLog**](LogLog.md) |  | [optional] 

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *Log) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *Log) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *Log) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *Log) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetFunction

`func (o *Log) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Log) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Log) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *Log) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetProject

`func (o *Log) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Log) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Log) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Log) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganisation

`func (o *Log) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *Log) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *Log) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *Log) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetGroup

`func (o *Log) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Log) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Log) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Log) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLog

`func (o *Log) GetLog() LogLog`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *Log) GetLogOk() (*LogLog, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *Log) SetLog(v LogLog)`

SetLog sets Log field to given value.

### HasLog

`func (o *Log) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


