# AuthTokenPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Token name | [optional] 
**ExpiryDuration** | Pointer to **int32** | Token expiration duration in days. 0 - token will never expire | [optional] 

## Methods

### NewAuthTokenPatch

`func NewAuthTokenPatch() *AuthTokenPatch`

NewAuthTokenPatch instantiates a new AuthTokenPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenPatchWithDefaults

`func NewAuthTokenPatchWithDefaults() *AuthTokenPatch`

NewAuthTokenPatchWithDefaults instantiates a new AuthTokenPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthTokenPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthTokenPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthTokenPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthTokenPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiryDuration

`func (o *AuthTokenPatch) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *AuthTokenPatch) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *AuthTokenPatch) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *AuthTokenPatch) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


