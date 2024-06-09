# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of provider | 
**Locations** | [**ProviderLocations**](ProviderLocations.md) |  | 

## Methods

### NewProvider

`func NewProvider(id string, name string, locations ProviderLocations, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.


### GetLocations

`func (o *Provider) GetLocations() ProviderLocations`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Provider) GetLocationsOk() (*ProviderLocations, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Provider) SetLocations(v ProviderLocations)`

SetLocations sets Locations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


