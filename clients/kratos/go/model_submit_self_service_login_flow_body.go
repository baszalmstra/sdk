/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.8
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SubmitSelfServiceLoginFlowBody - struct for SubmitSelfServiceLoginFlowBody
type SubmitSelfServiceLoginFlowBody struct {
	SubmitSelfServiceLoginFlowWithPasswordMethodBody *SubmitSelfServiceLoginFlowWithPasswordMethodBody
}

// SubmitSelfServiceLoginFlowWithPasswordMethodBodyAsSubmitSelfServiceLoginFlowBody is a convenience function that returns SubmitSelfServiceLoginFlowWithPasswordMethodBody wrapped in SubmitSelfServiceLoginFlowBody
func SubmitSelfServiceLoginFlowWithPasswordMethodBodyAsSubmitSelfServiceLoginFlowBody(v *SubmitSelfServiceLoginFlowWithPasswordMethodBody) SubmitSelfServiceLoginFlowBody {
	return SubmitSelfServiceLoginFlowBody{
		SubmitSelfServiceLoginFlowWithPasswordMethodBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceLoginFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceLoginFlowWithPasswordMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
	if err == nil {
		jsonSubmitSelfServiceLoginFlowWithPasswordMethodBody, _ := json.Marshal(dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
		if string(jsonSubmitSelfServiceLoginFlowWithPasswordMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceLoginFlowWithPasswordMethodBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceLoginFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceLoginFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceLoginFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceLoginFlowWithPasswordMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceLoginFlowWithPasswordMethodBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceLoginFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceLoginFlowWithPasswordMethodBody != nil {
		return obj.SubmitSelfServiceLoginFlowWithPasswordMethodBody
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceLoginFlowBody struct {
	value *SubmitSelfServiceLoginFlowBody
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowBody) Get() *SubmitSelfServiceLoginFlowBody {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowBody) Set(val *SubmitSelfServiceLoginFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowBody(val *SubmitSelfServiceLoginFlowBody) *NullableSubmitSelfServiceLoginFlowBody {
	return &NullableSubmitSelfServiceLoginFlowBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


