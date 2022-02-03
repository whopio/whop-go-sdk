/*
Whop API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.10
Contact: support@whop.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package whopclient

import (
	"encoding/json"
)

// CreateLinkResponse struct for CreateLinkResponse
type CreateLinkResponse struct {
	Link NullableLink `json:"link,omitempty"`
}

// NewCreateLinkResponse instantiates a new CreateLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkResponse() *CreateLinkResponse {
	this := CreateLinkResponse{}
	return &this
}

// NewCreateLinkResponseWithDefaults instantiates a new CreateLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkResponseWithDefaults() *CreateLinkResponse {
	this := CreateLinkResponse{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLinkResponse) GetLink() Link {
	if o == nil || o.Link.Get() == nil {
		var ret Link
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLinkResponse) GetLinkOk() (*Link, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *CreateLinkResponse) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableLink and assigns it to the Link field.
func (o *CreateLinkResponse) SetLink(v Link) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *CreateLinkResponse) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *CreateLinkResponse) UnsetLink() {
	o.Link.Unset()
}

func (o CreateLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLinkResponse struct {
	value *CreateLinkResponse
	isSet bool
}

func (v NullableCreateLinkResponse) Get() *CreateLinkResponse {
	return v.value
}

func (v *NullableCreateLinkResponse) Set(val *CreateLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkResponse(val *CreateLinkResponse) *NullableCreateLinkResponse {
	return &NullableCreateLinkResponse{value: val, isSet: true}
}

func (v NullableCreateLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


