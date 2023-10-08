# AuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** | User | 
**Name** | **string** | Token name | 
**ExpiryAt** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | OAuth2 client id and client secret used to generate API access token. Client secret can&#39;t be created and must be saved on user side | [optional] 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewAuthToken

`func NewAuthToken(id string, userId string, name string, date Date, ) *AuthToken`

NewAuthToken instantiates a new AuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithDefaults

`func NewAuthTokenWithDefaults() *AuthToken`

NewAuthTokenWithDefaults instantiates a new AuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthToken) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *AuthToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthToken) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *AuthToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthToken) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryAt

`func (o *AuthToken) GetExpiryAt() string`

GetExpiryAt returns the ExpiryAt field if non-nil, zero value otherwise.

### GetExpiryAtOk

`func (o *AuthToken) GetExpiryAtOk() (*string, bool)`

GetExpiryAtOk returns a tuple with the ExpiryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAt

`func (o *AuthToken) SetExpiryAt(v string)`

SetExpiryAt sets ExpiryAt field to given value.

### HasExpiryAt

`func (o *AuthToken) HasExpiryAt() bool`

HasExpiryAt returns a boolean if a field has been set.

### GetToken

`func (o *AuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDate

`func (o *AuthToken) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuthToken) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuthToken) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


