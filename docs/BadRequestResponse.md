# BadRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Fields** | [**BadRequestResponseFields**](BadRequestResponseFields.md) |  | 

## Methods

### NewBadRequestResponse

`func NewBadRequestResponse(message string, fields BadRequestResponseFields, ) *BadRequestResponse`

NewBadRequestResponse instantiates a new BadRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestResponseWithDefaults

`func NewBadRequestResponseWithDefaults() *BadRequestResponse`

NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BadRequestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFields

`func (o *BadRequestResponse) GetFields() BadRequestResponseFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BadRequestResponse) GetFieldsOk() (*BadRequestResponseFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BadRequestResponse) SetFields(v BadRequestResponseFields)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


