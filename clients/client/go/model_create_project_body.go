/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.6.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateProjectBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectBody{}

// CreateProjectBody Create Project Request Body
type CreateProjectBody struct {
	// The environment of the project. prod Production dev Development
	Environment string `json:"environment"`
	// The name of the project to be created
	Name string `json:"name"`
	WorkspaceId NullableString `json:"workspace_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateProjectBody CreateProjectBody

// NewCreateProjectBody instantiates a new CreateProjectBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectBody(environment string, name string) *CreateProjectBody {
	this := CreateProjectBody{}
	this.Environment = environment
	this.Name = name
	return &this
}

// NewCreateProjectBodyWithDefaults instantiates a new CreateProjectBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectBodyWithDefaults() *CreateProjectBody {
	this := CreateProjectBody{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *CreateProjectBody) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CreateProjectBody) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CreateProjectBody) SetEnvironment(v string) {
	o.Environment = v
}

// GetName returns the Name field value
func (o *CreateProjectBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectBody) SetName(v string) {
	o.Name = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectBody) GetWorkspaceId() string {
	if o == nil || IsNil(o.WorkspaceId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkspaceId.Get()
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectBody) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceId.Get(), o.WorkspaceId.IsSet()
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *CreateProjectBody) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given NullableString and assigns it to the WorkspaceId field.
func (o *CreateProjectBody) SetWorkspaceId(v string) {
	o.WorkspaceId.Set(&v)
}
// SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil
func (o *CreateProjectBody) SetWorkspaceIdNil() {
	o.WorkspaceId.Set(nil)
}

// UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
func (o *CreateProjectBody) UnsetWorkspaceId() {
	o.WorkspaceId.Unset()
}

func (o CreateProjectBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["name"] = o.Name
	if o.WorkspaceId.IsSet() {
		toSerialize["workspace_id"] = o.WorkspaceId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProjectBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateProjectBody := _CreateProjectBody{}

	err = json.Unmarshal(bytes, &varCreateProjectBody)

	if err != nil {
		return err
	}

	*o = CreateProjectBody(varCreateProjectBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "workspace_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProjectBody struct {
	value *CreateProjectBody
	isSet bool
}

func (v NullableCreateProjectBody) Get() *CreateProjectBody {
	return v.value
}

func (v *NullableCreateProjectBody) Set(val *CreateProjectBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectBody(val *CreateProjectBody) *NullableCreateProjectBody {
	return &NullableCreateProjectBody{value: val, isSet: true}
}

func (v NullableCreateProjectBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


