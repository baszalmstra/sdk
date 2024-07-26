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
	"time"
	"fmt"
)

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoice{}

// Invoice struct for Invoice
type Invoice struct {
	// The ID of the subscription
	Id string `json:"id"`
	InvoicedAt time.Time `json:"invoiced_at"`
	// Type is the type of the invoice. usage InvoiceTypeUsage base InvoiceTypeBase
	Type string `json:"type"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	V1 *InvoiceDataV1 `json:"v1,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Invoice Invoice

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice(id string, invoicedAt time.Time, type_ string) *Invoice {
	this := Invoice{}
	this.Id = id
	this.InvoicedAt = invoicedAt
	this.Type = type_
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value
func (o *Invoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invoice) SetId(v string) {
	o.Id = v
}

// GetInvoicedAt returns the InvoicedAt field value
func (o *Invoice) GetInvoicedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.InvoicedAt
}

// GetInvoicedAtOk returns a tuple with the InvoicedAt field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetInvoicedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoicedAt, true
}

// SetInvoicedAt sets field value
func (o *Invoice) SetInvoicedAt(v time.Time) {
	o.InvoicedAt = v
}

// GetType returns the Type field value
func (o *Invoice) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Invoice) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invoice) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invoice) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invoice) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetV1 returns the V1 field value if set, zero value otherwise.
func (o *Invoice) GetV1() InvoiceDataV1 {
	if o == nil || IsNil(o.V1) {
		var ret InvoiceDataV1
		return ret
	}
	return *o.V1
}

// GetV1Ok returns a tuple with the V1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetV1Ok() (*InvoiceDataV1, bool) {
	if o == nil || IsNil(o.V1) {
		return nil, false
	}
	return o.V1, true
}

// HasV1 returns a boolean if a field has been set.
func (o *Invoice) HasV1() bool {
	if o != nil && !IsNil(o.V1) {
		return true
	}

	return false
}

// SetV1 gets a reference to the given InvoiceDataV1 and assigns it to the V1 field.
func (o *Invoice) SetV1(v InvoiceDataV1) {
	o.V1 = &v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["invoiced_at"] = o.InvoicedAt
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.V1) {
		toSerialize["v1"] = o.V1
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Invoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"invoiced_at",
		"type",
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

	varInvoice := _Invoice{}

	err = json.Unmarshal(data, &varInvoice)

	if err != nil {
		return err
	}

	*o = Invoice(varInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "invoiced_at")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "v1")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


