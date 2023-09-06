/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateLoginFlowWithLookupSecretMethod Update Login Flow with Lookup Secret Method
type UpdateLoginFlowWithLookupSecretMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// The lookup secret.
	LookupSecret string `json:"lookup_secret"`
	// Method should be set to \"lookup_secret\" when logging in using the lookup_secret strategy.
	Method string `json:"method"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLoginFlowWithLookupSecretMethod UpdateLoginFlowWithLookupSecretMethod

// NewUpdateLoginFlowWithLookupSecretMethod instantiates a new UpdateLoginFlowWithLookupSecretMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoginFlowWithLookupSecretMethod(lookupSecret string, method string) *UpdateLoginFlowWithLookupSecretMethod {
	this := UpdateLoginFlowWithLookupSecretMethod{}
	this.LookupSecret = lookupSecret
	this.Method = method
	return &this
}

// NewUpdateLoginFlowWithLookupSecretMethodWithDefaults instantiates a new UpdateLoginFlowWithLookupSecretMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoginFlowWithLookupSecretMethodWithDefaults() *UpdateLoginFlowWithLookupSecretMethod {
	this := UpdateLoginFlowWithLookupSecretMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithLookupSecretMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithLookupSecretMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithLookupSecretMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateLoginFlowWithLookupSecretMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetLookupSecret returns the LookupSecret field value
func (o *UpdateLoginFlowWithLookupSecretMethod) GetLookupSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupSecret
}

// GetLookupSecretOk returns a tuple with the LookupSecret field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithLookupSecretMethod) GetLookupSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupSecret, true
}

// SetLookupSecret sets field value
func (o *UpdateLoginFlowWithLookupSecretMethod) SetLookupSecret(v string) {
	o.LookupSecret = v
}

// GetMethod returns the Method field value
func (o *UpdateLoginFlowWithLookupSecretMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithLookupSecretMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateLoginFlowWithLookupSecretMethod) SetMethod(v string) {
	o.Method = v
}

func (o UpdateLoginFlowWithLookupSecretMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["lookup_secret"] = o.LookupSecret
	}
	if true {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateLoginFlowWithLookupSecretMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateLoginFlowWithLookupSecretMethod := _UpdateLoginFlowWithLookupSecretMethod{}

	if err = json.Unmarshal(bytes, &varUpdateLoginFlowWithLookupSecretMethod); err == nil {
		*o = UpdateLoginFlowWithLookupSecretMethod(varUpdateLoginFlowWithLookupSecretMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "lookup_secret")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLoginFlowWithLookupSecretMethod struct {
	value *UpdateLoginFlowWithLookupSecretMethod
	isSet bool
}

func (v NullableUpdateLoginFlowWithLookupSecretMethod) Get() *UpdateLoginFlowWithLookupSecretMethod {
	return v.value
}

func (v *NullableUpdateLoginFlowWithLookupSecretMethod) Set(val *UpdateLoginFlowWithLookupSecretMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowWithLookupSecretMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowWithLookupSecretMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowWithLookupSecretMethod(val *UpdateLoginFlowWithLookupSecretMethod) *NullableUpdateLoginFlowWithLookupSecretMethod {
	return &NullableUpdateLoginFlowWithLookupSecretMethod{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowWithLookupSecretMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowWithLookupSecretMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


