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
	"time"
	"fmt"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message struct for Message
type Message struct {
	Body string `json:"body"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt time.Time `json:"created_at"`
	// Dispatches store information about the attempts of delivering a message May contain an error if any happened, or just the `success` state.
	Dispatches []MessageDispatch `json:"dispatches,omitempty"`
	Id string `json:"id"`
	Recipient string `json:"recipient"`
	SendCount int64 `json:"send_count"`
	Status CourierMessageStatus `json:"status"`
	Subject string `json:"subject"`
	//  recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid
	TemplateType string `json:"template_type"`
	Type CourierMessageType `json:"type"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(body string, createdAt time.Time, id string, recipient string, sendCount int64, status CourierMessageStatus, subject string, templateType string, type_ CourierMessageType, updatedAt time.Time) *Message {
	this := Message{}
	this.Body = body
	this.CreatedAt = createdAt
	this.Id = id
	this.Recipient = recipient
	this.SendCount = sendCount
	this.Status = status
	this.Subject = subject
	this.TemplateType = templateType
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetBody returns the Body field value
func (o *Message) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Message) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *Message) SetBody(v string) {
	o.Body = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Message) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Message) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDispatches returns the Dispatches field value if set, zero value otherwise.
func (o *Message) GetDispatches() []MessageDispatch {
	if o == nil || IsNil(o.Dispatches) {
		var ret []MessageDispatch
		return ret
	}
	return o.Dispatches
}

// GetDispatchesOk returns a tuple with the Dispatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetDispatchesOk() ([]MessageDispatch, bool) {
	if o == nil || IsNil(o.Dispatches) {
		return nil, false
	}
	return o.Dispatches, true
}

// HasDispatches returns a boolean if a field has been set.
func (o *Message) HasDispatches() bool {
	if o != nil && !IsNil(o.Dispatches) {
		return true
	}

	return false
}

// SetDispatches gets a reference to the given []MessageDispatch and assigns it to the Dispatches field.
func (o *Message) SetDispatches(v []MessageDispatch) {
	o.Dispatches = v
}

// GetId returns the Id field value
func (o *Message) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Message) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Message) SetId(v string) {
	o.Id = v
}

// GetRecipient returns the Recipient field value
func (o *Message) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *Message) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *Message) SetRecipient(v string) {
	o.Recipient = v
}

// GetSendCount returns the SendCount field value
func (o *Message) GetSendCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SendCount
}

// GetSendCountOk returns a tuple with the SendCount field value
// and a boolean to check if the value has been set.
func (o *Message) GetSendCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendCount, true
}

// SetSendCount sets field value
func (o *Message) SetSendCount(v int64) {
	o.SendCount = v
}

// GetStatus returns the Status field value
func (o *Message) GetStatus() CourierMessageStatus {
	if o == nil {
		var ret CourierMessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Message) GetStatusOk() (*CourierMessageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Message) SetStatus(v CourierMessageStatus) {
	o.Status = v
}

// GetSubject returns the Subject field value
func (o *Message) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Message) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *Message) SetSubject(v string) {
	o.Subject = v
}

// GetTemplateType returns the TemplateType field value
func (o *Message) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *Message) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *Message) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetType returns the Type field value
func (o *Message) GetType() CourierMessageType {
	if o == nil {
		var ret CourierMessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Message) GetTypeOk() (*CourierMessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Message) SetType(v CourierMessageType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Message) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Message) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Message) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Dispatches) {
		toSerialize["dispatches"] = o.Dispatches
	}
	toSerialize["id"] = o.Id
	toSerialize["recipient"] = o.Recipient
	toSerialize["send_count"] = o.SendCount
	toSerialize["status"] = o.Status
	toSerialize["subject"] = o.Subject
	toSerialize["template_type"] = o.TemplateType
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"created_at",
		"id",
		"recipient",
		"send_count",
		"status",
		"subject",
		"template_type",
		"type",
		"updated_at",
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

	varMessage := _Message{}

	err = json.Unmarshal(bytes, &varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "dispatches")
		delete(additionalProperties, "id")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "send_count")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "template_type")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


