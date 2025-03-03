/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ListWorkspaces type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorkspaces{}

// ListWorkspaces struct for ListWorkspaces
type ListWorkspaces struct {
	HasNextPage bool `json:"has_next_page"`
	NextPageToken string `json:"next_page_token"`
	Workspaces []Workspace `json:"workspaces"`
	AdditionalProperties map[string]interface{}
}

type _ListWorkspaces ListWorkspaces

// NewListWorkspaces instantiates a new ListWorkspaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaces(hasNextPage bool, nextPageToken string, workspaces []Workspace) *ListWorkspaces {
	this := ListWorkspaces{}
	this.HasNextPage = hasNextPage
	this.NextPageToken = nextPageToken
	this.Workspaces = workspaces
	return &this
}

// NewListWorkspacesWithDefaults instantiates a new ListWorkspaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspacesWithDefaults() *ListWorkspaces {
	this := ListWorkspaces{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *ListWorkspaces) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaces) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *ListWorkspaces) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *ListWorkspaces) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaces) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ListWorkspaces) SetNextPageToken(v string) {
	o.NextPageToken = v
}

// GetWorkspaces returns the Workspaces field value
func (o *ListWorkspaces) GetWorkspaces() []Workspace {
	if o == nil {
		var ret []Workspace
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaces) GetWorkspacesOk() ([]Workspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *ListWorkspaces) SetWorkspaces(v []Workspace) {
	o.Workspaces = v
}

func (o ListWorkspaces) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorkspaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_next_page"] = o.HasNextPage
	toSerialize["next_page_token"] = o.NextPageToken
	toSerialize["workspaces"] = o.Workspaces

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWorkspaces) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_next_page",
		"next_page_token",
		"workspaces",
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

	varListWorkspaces := _ListWorkspaces{}

	err = json.Unmarshal(data, &varListWorkspaces)

	if err != nil {
		return err
	}

	*o = ListWorkspaces(varListWorkspaces)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "has_next_page")
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "workspaces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWorkspaces struct {
	value *ListWorkspaces
	isSet bool
}

func (v NullableListWorkspaces) Get() *ListWorkspaces {
	return v.value
}

func (v *NullableListWorkspaces) Set(val *ListWorkspaces) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaces) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaces(val *ListWorkspaces) *NullableListWorkspaces {
	return &NullableListWorkspaces{value: val, isSet: true}
}

func (v NullableListWorkspaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


