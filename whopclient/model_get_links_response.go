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

// GetLinksResponse struct for GetLinksResponse
type GetLinksResponse struct {
	Links []Link `json:"links,omitempty"`
}

// NewGetLinksResponse instantiates a new GetLinksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLinksResponse() *GetLinksResponse {
	this := GetLinksResponse{}
	return &this
}

// NewGetLinksResponseWithDefaults instantiates a new GetLinksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLinksResponseWithDefaults() *GetLinksResponse {
	this := GetLinksResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLinksResponse) GetLinks() []Link {
	if o == nil  {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLinksResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetLinksResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *GetLinksResponse) SetLinks(v []Link) {
	o.Links = v
}

func (o GetLinksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableGetLinksResponse struct {
	value *GetLinksResponse
	isSet bool
}

func (v NullableGetLinksResponse) Get() *GetLinksResponse {
	return v.value
}

func (v *NullableGetLinksResponse) Set(val *GetLinksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinksResponse(val *GetLinksResponse) *NullableGetLinksResponse {
	return &NullableGetLinksResponse{value: val, isSet: true}
}

func (v NullableGetLinksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

