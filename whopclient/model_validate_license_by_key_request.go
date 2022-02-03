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

// ValidateLicenseByKeyRequest struct for ValidateLicenseByKeyRequest
type ValidateLicenseByKeyRequest struct {
	Metadata *UpdateLicenseByKeyRequestMetadata `json:"metadata,omitempty"`
}

// NewValidateLicenseByKeyRequest instantiates a new ValidateLicenseByKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateLicenseByKeyRequest() *ValidateLicenseByKeyRequest {
	this := ValidateLicenseByKeyRequest{}
	return &this
}

// NewValidateLicenseByKeyRequestWithDefaults instantiates a new ValidateLicenseByKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateLicenseByKeyRequestWithDefaults() *ValidateLicenseByKeyRequest {
	this := ValidateLicenseByKeyRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ValidateLicenseByKeyRequest) GetMetadata() UpdateLicenseByKeyRequestMetadata {
	if o == nil || o.Metadata == nil {
		var ret UpdateLicenseByKeyRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateLicenseByKeyRequest) GetMetadataOk() (*UpdateLicenseByKeyRequestMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ValidateLicenseByKeyRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given UpdateLicenseByKeyRequestMetadata and assigns it to the Metadata field.
func (o *ValidateLicenseByKeyRequest) SetMetadata(v UpdateLicenseByKeyRequestMetadata) {
	o.Metadata = &v
}

func (o ValidateLicenseByKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableValidateLicenseByKeyRequest struct {
	value *ValidateLicenseByKeyRequest
	isSet bool
}

func (v NullableValidateLicenseByKeyRequest) Get() *ValidateLicenseByKeyRequest {
	return v.value
}

func (v *NullableValidateLicenseByKeyRequest) Set(val *ValidateLicenseByKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateLicenseByKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateLicenseByKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateLicenseByKeyRequest(val *ValidateLicenseByKeyRequest) *NullableValidateLicenseByKeyRequest {
	return &NullableValidateLicenseByKeyRequest{value: val, isSet: true}
}

func (v NullableValidateLicenseByKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateLicenseByKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

