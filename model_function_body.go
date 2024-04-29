/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_chaos_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FunctionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionBody{}

// FunctionBody Function create body
type FunctionBody struct {
	// ID of the project this function belongs to
	ProjectId string `json:"project_id"`
	// Function spec version
	Version string `json:"version"`
	// Name of the function
	Name string `json:"name"`
	// Description of what the function does
	Description string `json:"description"`
	// Path to container image
	Image string `json:"image"`
	Type FunctionType `json:"type"`
	Size FunctionSize `json:"size"`
	// Port the application runs on
	Port int32 `json:"port"`
	// The public route/path to this function, only applicable to http type functions
	Routes []FunctionRoute `json:"routes,omitempty"`
	Scaling FunctionScaling `json:"scaling"`
	// List of deployments for this function
	Deployments []FunctionDeploymentBody `json:"deployments"`
	// List of environment variables for secrets
	Secrets []FunctionEnv `json:"secrets"`
	// Tags to limit deployment
	Compliance []FunctionCompliance `json:"compliance,omitempty"`
}

type _FunctionBody FunctionBody

// NewFunctionBody instantiates a new FunctionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionBody(projectId string, version string, name string, description string, image string, type_ FunctionType, size FunctionSize, port int32, scaling FunctionScaling, deployments []FunctionDeploymentBody, secrets []FunctionEnv) *FunctionBody {
	this := FunctionBody{}
	this.ProjectId = projectId
	this.Version = version
	this.Name = name
	this.Description = description
	this.Image = image
	this.Type = type_
	this.Size = size
	this.Port = port
	this.Scaling = scaling
	this.Deployments = deployments
	this.Secrets = secrets
	return &this
}

// NewFunctionBodyWithDefaults instantiates a new FunctionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionBodyWithDefaults() *FunctionBody {
	this := FunctionBody{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *FunctionBody) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *FunctionBody) SetProjectId(v string) {
	o.ProjectId = v
}

// GetVersion returns the Version field value
func (o *FunctionBody) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FunctionBody) SetVersion(v string) {
	o.Version = v
}

// GetName returns the Name field value
func (o *FunctionBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionBody) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *FunctionBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FunctionBody) SetDescription(v string) {
	o.Description = v
}

// GetImage returns the Image field value
func (o *FunctionBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *FunctionBody) SetImage(v string) {
	o.Image = v
}

// GetType returns the Type field value
func (o *FunctionBody) GetType() FunctionType {
	if o == nil {
		var ret FunctionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetTypeOk() (*FunctionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FunctionBody) SetType(v FunctionType) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *FunctionBody) GetSize() FunctionSize {
	if o == nil {
		var ret FunctionSize
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetSizeOk() (*FunctionSize, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FunctionBody) SetSize(v FunctionSize) {
	o.Size = v
}

// GetPort returns the Port field value
func (o *FunctionBody) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *FunctionBody) SetPort(v int32) {
	o.Port = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *FunctionBody) GetRoutes() []FunctionRoute {
	if o == nil || IsNil(o.Routes) {
		var ret []FunctionRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetRoutesOk() ([]FunctionRoute, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *FunctionBody) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []FunctionRoute and assigns it to the Routes field.
func (o *FunctionBody) SetRoutes(v []FunctionRoute) {
	o.Routes = v
}

// GetScaling returns the Scaling field value
func (o *FunctionBody) GetScaling() FunctionScaling {
	if o == nil {
		var ret FunctionScaling
		return ret
	}

	return o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetScalingOk() (*FunctionScaling, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scaling, true
}

// SetScaling sets field value
func (o *FunctionBody) SetScaling(v FunctionScaling) {
	o.Scaling = v
}

// GetDeployments returns the Deployments field value
func (o *FunctionBody) GetDeployments() []FunctionDeploymentBody {
	if o == nil {
		var ret []FunctionDeploymentBody
		return ret
	}

	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetDeploymentsOk() ([]FunctionDeploymentBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deployments, true
}

// SetDeployments sets field value
func (o *FunctionBody) SetDeployments(v []FunctionDeploymentBody) {
	o.Deployments = v
}

// GetSecrets returns the Secrets field value
func (o *FunctionBody) GetSecrets() []FunctionEnv {
	if o == nil {
		var ret []FunctionEnv
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetSecretsOk() ([]FunctionEnv, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *FunctionBody) SetSecrets(v []FunctionEnv) {
	o.Secrets = v
}

// GetCompliance returns the Compliance field value if set, zero value otherwise.
func (o *FunctionBody) GetCompliance() []FunctionCompliance {
	if o == nil || IsNil(o.Compliance) {
		var ret []FunctionCompliance
		return ret
	}
	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionBody) GetComplianceOk() ([]FunctionCompliance, bool) {
	if o == nil || IsNil(o.Compliance) {
		return nil, false
	}
	return o.Compliance, true
}

// HasCompliance returns a boolean if a field has been set.
func (o *FunctionBody) HasCompliance() bool {
	if o != nil && !IsNil(o.Compliance) {
		return true
	}

	return false
}

// SetCompliance gets a reference to the given []FunctionCompliance and assigns it to the Compliance field.
func (o *FunctionBody) SetCompliance(v []FunctionCompliance) {
	o.Compliance = v
}

func (o FunctionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["version"] = o.Version
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["image"] = o.Image
	toSerialize["type"] = o.Type
	toSerialize["size"] = o.Size
	toSerialize["port"] = o.Port
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["scaling"] = o.Scaling
	toSerialize["deployments"] = o.Deployments
	toSerialize["secrets"] = o.Secrets
	if !IsNil(o.Compliance) {
		toSerialize["compliance"] = o.Compliance
	}
	return toSerialize, nil
}

func (o *FunctionBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
		"version",
		"name",
		"description",
		"image",
		"type",
		"size",
		"port",
		"scaling",
		"deployments",
		"secrets",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFunctionBody := _FunctionBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionBody)

	if err != nil {
		return err
	}

	*o = FunctionBody(varFunctionBody)

	return err
}

type NullableFunctionBody struct {
	value *FunctionBody
	isSet bool
}

func (v NullableFunctionBody) Get() *FunctionBody {
	return v.value
}

func (v *NullableFunctionBody) Set(val *FunctionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionBody(val *FunctionBody) *NullableFunctionBody {
	return &NullableFunctionBody{value: val, isSet: true}
}

func (v NullableFunctionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

