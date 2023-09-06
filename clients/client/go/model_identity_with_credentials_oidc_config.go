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

// IdentityWithCredentialsOidcConfig struct for IdentityWithCredentialsOidcConfig
type IdentityWithCredentialsOidcConfig struct {
	Config *IdentityWithCredentialsPasswordConfig `json:"config,omitempty"`
	// A list of OpenID Connect Providers
	Providers []IdentityWithCredentialsOidcConfigProvider `json:"providers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityWithCredentialsOidcConfig IdentityWithCredentialsOidcConfig

// NewIdentityWithCredentialsOidcConfig instantiates a new IdentityWithCredentialsOidcConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsOidcConfig() *IdentityWithCredentialsOidcConfig {
	this := IdentityWithCredentialsOidcConfig{}
	return &this
}

// NewIdentityWithCredentialsOidcConfigWithDefaults instantiates a new IdentityWithCredentialsOidcConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsOidcConfigWithDefaults() *IdentityWithCredentialsOidcConfig {
	this := IdentityWithCredentialsOidcConfig{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentityWithCredentialsOidcConfig) GetConfig() IdentityWithCredentialsPasswordConfig {
	if o == nil || o.Config == nil {
		var ret IdentityWithCredentialsPasswordConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsOidcConfig) GetConfigOk() (*IdentityWithCredentialsPasswordConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentityWithCredentialsOidcConfig) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentityWithCredentialsPasswordConfig and assigns it to the Config field.
func (o *IdentityWithCredentialsOidcConfig) SetConfig(v IdentityWithCredentialsPasswordConfig) {
	o.Config = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityWithCredentialsOidcConfig) GetProviders() []IdentityWithCredentialsOidcConfigProvider {
	if o == nil || o.Providers == nil {
		var ret []IdentityWithCredentialsOidcConfigProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsOidcConfig) GetProvidersOk() ([]IdentityWithCredentialsOidcConfigProvider, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityWithCredentialsOidcConfig) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IdentityWithCredentialsOidcConfigProvider and assigns it to the Providers field.
func (o *IdentityWithCredentialsOidcConfig) SetProviders(v []IdentityWithCredentialsOidcConfigProvider) {
	o.Providers = v
}

func (o IdentityWithCredentialsOidcConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityWithCredentialsOidcConfig) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityWithCredentialsOidcConfig := _IdentityWithCredentialsOidcConfig{}

	if err = json.Unmarshal(bytes, &varIdentityWithCredentialsOidcConfig); err == nil {
		*o = IdentityWithCredentialsOidcConfig(varIdentityWithCredentialsOidcConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "providers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityWithCredentialsOidcConfig struct {
	value *IdentityWithCredentialsOidcConfig
	isSet bool
}

func (v NullableIdentityWithCredentialsOidcConfig) Get() *IdentityWithCredentialsOidcConfig {
	return v.value
}

func (v *NullableIdentityWithCredentialsOidcConfig) Set(val *IdentityWithCredentialsOidcConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsOidcConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsOidcConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsOidcConfig(val *IdentityWithCredentialsOidcConfig) *NullableIdentityWithCredentialsOidcConfig {
	return &NullableIdentityWithCredentialsOidcConfig{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsOidcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsOidcConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


