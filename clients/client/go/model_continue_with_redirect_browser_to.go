/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.14.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ContinueWithRedirectBrowserTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueWithRedirectBrowserTo{}

// ContinueWithRedirectBrowserTo Indicates, that the UI flow could be continued by showing a recovery ui
type ContinueWithRedirectBrowserTo struct {
	// Action will always be `redirect_browser_to` redirect_browser_to ContinueWithActionRedirectBrowserToString
	Action string `json:"action"`
	// The URL to redirect the browser to
	RedirectBrowserTo string `json:"redirect_browser_to"`
	AdditionalProperties map[string]interface{}
}

type _ContinueWithRedirectBrowserTo ContinueWithRedirectBrowserTo

// NewContinueWithRedirectBrowserTo instantiates a new ContinueWithRedirectBrowserTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithRedirectBrowserTo(action string, redirectBrowserTo string) *ContinueWithRedirectBrowserTo {
	this := ContinueWithRedirectBrowserTo{}
	this.Action = action
	this.RedirectBrowserTo = redirectBrowserTo
	return &this
}

// NewContinueWithRedirectBrowserToWithDefaults instantiates a new ContinueWithRedirectBrowserTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithRedirectBrowserToWithDefaults() *ContinueWithRedirectBrowserTo {
	this := ContinueWithRedirectBrowserTo{}
	return &this
}

// GetAction returns the Action field value
func (o *ContinueWithRedirectBrowserTo) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContinueWithRedirectBrowserTo) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContinueWithRedirectBrowserTo) SetAction(v string) {
	o.Action = v
}

// GetRedirectBrowserTo returns the RedirectBrowserTo field value
func (o *ContinueWithRedirectBrowserTo) GetRedirectBrowserTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectBrowserTo
}

// GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field value
// and a boolean to check if the value has been set.
func (o *ContinueWithRedirectBrowserTo) GetRedirectBrowserToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectBrowserTo, true
}

// SetRedirectBrowserTo sets field value
func (o *ContinueWithRedirectBrowserTo) SetRedirectBrowserTo(v string) {
	o.RedirectBrowserTo = v
}

func (o ContinueWithRedirectBrowserTo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueWithRedirectBrowserTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["redirect_browser_to"] = o.RedirectBrowserTo

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContinueWithRedirectBrowserTo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"redirect_browser_to",
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

	varContinueWithRedirectBrowserTo := _ContinueWithRedirectBrowserTo{}

	err = json.Unmarshal(data, &varContinueWithRedirectBrowserTo)

	if err != nil {
		return err
	}

	*o = ContinueWithRedirectBrowserTo(varContinueWithRedirectBrowserTo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "redirect_browser_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContinueWithRedirectBrowserTo struct {
	value *ContinueWithRedirectBrowserTo
	isSet bool
}

func (v NullableContinueWithRedirectBrowserTo) Get() *ContinueWithRedirectBrowserTo {
	return v.value
}

func (v *NullableContinueWithRedirectBrowserTo) Set(val *ContinueWithRedirectBrowserTo) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithRedirectBrowserTo) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithRedirectBrowserTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithRedirectBrowserTo(val *ContinueWithRedirectBrowserTo) *NullableContinueWithRedirectBrowserTo {
	return &NullableContinueWithRedirectBrowserTo{value: val, isSet: true}
}

func (v NullableContinueWithRedirectBrowserTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithRedirectBrowserTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


