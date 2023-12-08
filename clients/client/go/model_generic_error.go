/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.5
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GenericError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericError{}

// GenericError Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
type GenericError struct {
	// The status code
	Code *int64 `json:"code,omitempty"`
	// Debug information  This field is often not exposed to protect against leaking sensitive information.
	Debug *string `json:"debug,omitempty"`
	// Further error details
	Details interface{} `json:"details,omitempty"`
	Error *GenericErrorContent `json:"error,omitempty"`
	// The error ID  Useful when trying to identify various errors in application logic.
	Id *string `json:"id,omitempty"`
	// Error message  The error's message.
	Message string `json:"message"`
	// A human-readable reason for the error
	Reason *string `json:"reason,omitempty"`
	// The request ID  The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID.
	Request *string `json:"request,omitempty"`
	// The status description
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericError GenericError

// NewGenericError instantiates a new GenericError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericError(message string) *GenericError {
	this := GenericError{}
	this.Message = message
	return &this
}

// NewGenericErrorWithDefaults instantiates a new GenericError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorWithDefaults() *GenericError {
	this := GenericError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GenericError) GetCode() int64 {
	if o == nil || IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GenericError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *GenericError) SetCode(v int64) {
	o.Code = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *GenericError) GetDebug() string {
	if o == nil || IsNil(o.Debug) {
		var ret string
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetDebugOk() (*string, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *GenericError) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given string and assigns it to the Debug field.
func (o *GenericError) SetDebug(v string) {
	o.Debug = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericError) GetDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericError) GetDetailsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GenericError) HasDetails() bool {
	if o != nil && IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given interface{} and assigns it to the Details field.
func (o *GenericError) SetDetails(v interface{}) {
	o.Details = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GenericError) GetError() GenericErrorContent {
	if o == nil || IsNil(o.Error) {
		var ret GenericErrorContent
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetErrorOk() (*GenericErrorContent, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GenericError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given GenericErrorContent and assigns it to the Error field.
func (o *GenericError) SetError(v GenericErrorContent) {
	o.Error = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GenericError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GenericError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GenericError) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value
func (o *GenericError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GenericError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GenericError) SetMessage(v string) {
	o.Message = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GenericError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GenericError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GenericError) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *GenericError) GetRequest() string {
	if o == nil || IsNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetRequestOk() (*string, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *GenericError) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *GenericError) SetRequest(v string) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GenericError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GenericError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GenericError) SetStatus(v string) {
	o.Status = &v
}

func (o GenericError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericError) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varGenericError := _GenericError{}

	err = json.Unmarshal(bytes, &varGenericError)

	if err != nil {
		return err
	}

	*o = GenericError(varGenericError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "debug")
		delete(additionalProperties, "details")
		delete(additionalProperties, "error")
		delete(additionalProperties, "id")
		delete(additionalProperties, "message")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "request")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericError struct {
	value *GenericError
	isSet bool
}

func (v NullableGenericError) Get() *GenericError {
	return v.value
}

func (v *NullableGenericError) Set(val *GenericError) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericError) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericError(val *GenericError) *NullableGenericError {
	return &NullableGenericError{value: val, isSet: true}
}

func (v NullableGenericError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


