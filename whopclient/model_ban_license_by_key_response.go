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

// BanLicenseByKeyResponse struct for BanLicenseByKeyResponse
type BanLicenseByKeyResponse struct {
	Message NullableString `json:"message,omitempty"`
}

// NewBanLicenseByKeyResponse instantiates a new BanLicenseByKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanLicenseByKeyResponse() *BanLicenseByKeyResponse {
	this := BanLicenseByKeyResponse{}
	return &this
}

// NewBanLicenseByKeyResponseWithDefaults instantiates a new BanLicenseByKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanLicenseByKeyResponseWithDefaults() *BanLicenseByKeyResponse {
	this := BanLicenseByKeyResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanLicenseByKeyResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanLicenseByKeyResponse) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *BanLicenseByKeyResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *BanLicenseByKeyResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *BanLicenseByKeyResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *BanLicenseByKeyResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o BanLicenseByKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBanLicenseByKeyResponse struct {
	value *BanLicenseByKeyResponse
	isSet bool
}

func (v NullableBanLicenseByKeyResponse) Get() *BanLicenseByKeyResponse {
	return v.value
}

func (v *NullableBanLicenseByKeyResponse) Set(val *BanLicenseByKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBanLicenseByKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBanLicenseByKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanLicenseByKeyResponse(val *BanLicenseByKeyResponse) *NullableBanLicenseByKeyResponse {
	return &NullableBanLicenseByKeyResponse{value: val, isSet: true}
}

func (v NullableBanLicenseByKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanLicenseByKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


