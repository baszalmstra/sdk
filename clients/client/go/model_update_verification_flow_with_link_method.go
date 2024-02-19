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

// checks if the UpdateVerificationFlowWithLinkMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVerificationFlowWithLinkMethod{}

// UpdateVerificationFlowWithLinkMethod Update Verification Flow with Link Method
type UpdateVerificationFlowWithLinkMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
	Email string `json:"email"`
	// Method is the method that should be used for this verification flow  Allowed values are `link` and `code` link VerificationStrategyLink code VerificationStrategyCode
	Method string `json:"method"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVerificationFlowWithLinkMethod UpdateVerificationFlowWithLinkMethod

// NewUpdateVerificationFlowWithLinkMethod instantiates a new UpdateVerificationFlowWithLinkMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVerificationFlowWithLinkMethod(email string, method string) *UpdateVerificationFlowWithLinkMethod {
	this := UpdateVerificationFlowWithLinkMethod{}
	this.Email = email
	this.Method = method
	return &this
}

// NewUpdateVerificationFlowWithLinkMethodWithDefaults instantiates a new UpdateVerificationFlowWithLinkMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVerificationFlowWithLinkMethodWithDefaults() *UpdateVerificationFlowWithLinkMethod {
	this := UpdateVerificationFlowWithLinkMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithLinkMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithLinkMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithLinkMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateVerificationFlowWithLinkMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value
func (o *UpdateVerificationFlowWithLinkMethod) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithLinkMethod) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UpdateVerificationFlowWithLinkMethod) SetEmail(v string) {
	o.Email = v
}

// GetMethod returns the Method field value
func (o *UpdateVerificationFlowWithLinkMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithLinkMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateVerificationFlowWithLinkMethod) SetMethod(v string) {
	o.Method = v
}

func (o UpdateVerificationFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVerificationFlowWithLinkMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["email"] = o.Email
	toSerialize["method"] = o.Method

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateVerificationFlowWithLinkMethod) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"method",
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

	varUpdateVerificationFlowWithLinkMethod := _UpdateVerificationFlowWithLinkMethod{}

	err = json.Unmarshal(bytes, &varUpdateVerificationFlowWithLinkMethod)

	if err != nil {
		return err
	}

	*o = UpdateVerificationFlowWithLinkMethod(varUpdateVerificationFlowWithLinkMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "email")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVerificationFlowWithLinkMethod struct {
	value *UpdateVerificationFlowWithLinkMethod
	isSet bool
}

func (v NullableUpdateVerificationFlowWithLinkMethod) Get() *UpdateVerificationFlowWithLinkMethod {
	return v.value
}

func (v *NullableUpdateVerificationFlowWithLinkMethod) Set(val *UpdateVerificationFlowWithLinkMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVerificationFlowWithLinkMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVerificationFlowWithLinkMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVerificationFlowWithLinkMethod(val *UpdateVerificationFlowWithLinkMethod) *NullableUpdateVerificationFlowWithLinkMethod {
	return &NullableUpdateVerificationFlowWithLinkMethod{value: val, isSet: true}
}

func (v NullableUpdateVerificationFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVerificationFlowWithLinkMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


