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

// CreateCheckoutLogResponse struct for CreateCheckoutLogResponse
type CreateCheckoutLogResponse struct {
	Checkout NullableCheckoutLog `json:"checkout,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewCreateCheckoutLogResponse instantiates a new CreateCheckoutLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCheckoutLogResponse() *CreateCheckoutLogResponse {
	this := CreateCheckoutLogResponse{}
	return &this
}

// NewCreateCheckoutLogResponseWithDefaults instantiates a new CreateCheckoutLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCheckoutLogResponseWithDefaults() *CreateCheckoutLogResponse {
	this := CreateCheckoutLogResponse{}
	return &this
}

// GetCheckout returns the Checkout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCheckoutLogResponse) GetCheckout() CheckoutLog {
	if o == nil || o.Checkout.Get() == nil {
		var ret CheckoutLog
		return ret
	}
	return *o.Checkout.Get()
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCheckoutLogResponse) GetCheckoutOk() (*CheckoutLog, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Checkout.Get(), o.Checkout.IsSet()
}

// HasCheckout returns a boolean if a field has been set.
func (o *CreateCheckoutLogResponse) HasCheckout() bool {
	if o != nil && o.Checkout.IsSet() {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given NullableCheckoutLog and assigns it to the Checkout field.
func (o *CreateCheckoutLogResponse) SetCheckout(v CheckoutLog) {
	o.Checkout.Set(&v)
}
// SetCheckoutNil sets the value for Checkout to be an explicit nil
func (o *CreateCheckoutLogResponse) SetCheckoutNil() {
	o.Checkout.Set(nil)
}

// UnsetCheckout ensures that no value is present for Checkout, not even an explicit nil
func (o *CreateCheckoutLogResponse) UnsetCheckout() {
	o.Checkout.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateCheckoutLogResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutLogResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateCheckoutLogResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateCheckoutLogResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o CreateCheckoutLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checkout.IsSet() {
		toSerialize["checkout"] = o.Checkout.Get()
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCheckoutLogResponse struct {
	value *CreateCheckoutLogResponse
	isSet bool
}

func (v NullableCreateCheckoutLogResponse) Get() *CreateCheckoutLogResponse {
	return v.value
}

func (v *NullableCreateCheckoutLogResponse) Set(val *CreateCheckoutLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCheckoutLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCheckoutLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCheckoutLogResponse(val *CreateCheckoutLogResponse) *NullableCreateCheckoutLogResponse {
	return &NullableCreateCheckoutLogResponse{value: val, isSet: true}
}

func (v NullableCreateCheckoutLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCheckoutLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


