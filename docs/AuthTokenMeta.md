# AuthTokenMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** | User | 
**Name** | **string** | Name of token | 
**ExpiryAt** | Pointer to **string** |  | [optional] 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewAuthTokenMeta

`func NewAuthTokenMeta(id string, userId string, name string, date Date, ) *AuthTokenMeta`

NewAuthTokenMeta instantiates a new AuthTokenMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenMetaWithDefaults

`func NewAuthTokenMetaWithDefaults() *AuthTokenMeta`

NewAuthTokenMetaWithDefaults instantiates a new AuthTokenMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthTokenMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthTokenMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthTokenMeta) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *AuthTokenMeta) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthTokenMeta) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthTokenMeta) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *AuthTokenMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthTokenMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthTokenMeta) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryAt

`func (o *AuthTokenMeta) GetExpiryAt() string`

GetExpiryAt returns the ExpiryAt field if non-nil, zero value otherwise.

### GetExpiryAtOk

`func (o *AuthTokenMeta) GetExpiryAtOk() (*string, bool)`

GetExpiryAtOk returns a tuple with the ExpiryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAt

`func (o *AuthTokenMeta) SetExpiryAt(v string)`

SetExpiryAt sets ExpiryAt field to given value.

### HasExpiryAt

`func (o *AuthTokenMeta) HasExpiryAt() bool`

HasExpiryAt returns a boolean if a field has been set.

### GetDate

`func (o *AuthTokenMeta) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuthTokenMeta) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuthTokenMeta) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


