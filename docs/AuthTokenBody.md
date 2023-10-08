# AuthTokenBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of token | 
**ExpiryDuration** | **int32** | Token expiration duration in days. 0 - token will never expire | 

## Methods

### NewAuthTokenBody

`func NewAuthTokenBody(name string, expiryDuration int32, ) *AuthTokenBody`

NewAuthTokenBody instantiates a new AuthTokenBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenBodyWithDefaults

`func NewAuthTokenBodyWithDefaults() *AuthTokenBody`

NewAuthTokenBodyWithDefaults instantiates a new AuthTokenBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthTokenBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthTokenBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthTokenBody) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDuration

`func (o *AuthTokenBody) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *AuthTokenBody) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *AuthTokenBody) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


