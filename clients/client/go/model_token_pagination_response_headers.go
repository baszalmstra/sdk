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
)

// checks if the TokenPaginationResponseHeaders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPaginationResponseHeaders{}

// TokenPaginationResponseHeaders The `Link` HTTP header contains multiple links (`first`, `next`, `last`, `previous`) formatted as: `<https://{project-slug}.projects.oryapis.com/admin/clients?page_size={limit}&page_token={offset}>; rel=\"{page}\"`  For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
type TokenPaginationResponseHeaders struct {
	// The Link HTTP Header  The `Link` header contains a comma-delimited list of links to the following pages:  first: The first page of results. next: The next page of results. prev: The previous page of results. last: The last page of results.  Pages are omitted if they do not exist. For example, if there is no next page, the `next` link is omitted. Examples:  </clients?page_size=5&page_token=0>; rel=\"first\",</clients?page_size=5&page_token=15>; rel=\"next\",</clients?page_size=5&page_token=5>; rel=\"prev\",</clients?page_size=5&page_token=20>; rel=\"last\"
	Link *string `json:"link,omitempty"`
	// The X-Total-Count HTTP Header  The `X-Total-Count` header contains the total number of items in the collection.
	XTotalCount *int64 `json:"x-total-count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPaginationResponseHeaders TokenPaginationResponseHeaders

// NewTokenPaginationResponseHeaders instantiates a new TokenPaginationResponseHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPaginationResponseHeaders() *TokenPaginationResponseHeaders {
	this := TokenPaginationResponseHeaders{}
	return &this
}

// NewTokenPaginationResponseHeadersWithDefaults instantiates a new TokenPaginationResponseHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPaginationResponseHeadersWithDefaults() *TokenPaginationResponseHeaders {
	this := TokenPaginationResponseHeaders{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *TokenPaginationResponseHeaders) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPaginationResponseHeaders) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *TokenPaginationResponseHeaders) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *TokenPaginationResponseHeaders) SetLink(v string) {
	o.Link = &v
}

// GetXTotalCount returns the XTotalCount field value if set, zero value otherwise.
func (o *TokenPaginationResponseHeaders) GetXTotalCount() int64 {
	if o == nil || IsNil(o.XTotalCount) {
		var ret int64
		return ret
	}
	return *o.XTotalCount
}

// GetXTotalCountOk returns a tuple with the XTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPaginationResponseHeaders) GetXTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.XTotalCount) {
		return nil, false
	}
	return o.XTotalCount, true
}

// HasXTotalCount returns a boolean if a field has been set.
func (o *TokenPaginationResponseHeaders) HasXTotalCount() bool {
	if o != nil && !IsNil(o.XTotalCount) {
		return true
	}

	return false
}

// SetXTotalCount gets a reference to the given int64 and assigns it to the XTotalCount field.
func (o *TokenPaginationResponseHeaders) SetXTotalCount(v int64) {
	o.XTotalCount = &v
}

func (o TokenPaginationResponseHeaders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPaginationResponseHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.XTotalCount) {
		toSerialize["x-total-count"] = o.XTotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPaginationResponseHeaders) UnmarshalJSON(data []byte) (err error) {
	varTokenPaginationResponseHeaders := _TokenPaginationResponseHeaders{}

	err = json.Unmarshal(data, &varTokenPaginationResponseHeaders)

	if err != nil {
		return err
	}

	*o = TokenPaginationResponseHeaders(varTokenPaginationResponseHeaders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "link")
		delete(additionalProperties, "x-total-count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenPaginationResponseHeaders struct {
	value *TokenPaginationResponseHeaders
	isSet bool
}

func (v NullableTokenPaginationResponseHeaders) Get() *TokenPaginationResponseHeaders {
	return v.value
}

func (v *NullableTokenPaginationResponseHeaders) Set(val *TokenPaginationResponseHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPaginationResponseHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPaginationResponseHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPaginationResponseHeaders(val *TokenPaginationResponseHeaders) *NullableTokenPaginationResponseHeaders {
	return &NullableTokenPaginationResponseHeaders{value: val, isSet: true}
}

func (v NullableTokenPaginationResponseHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPaginationResponseHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


