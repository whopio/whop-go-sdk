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

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	Success NullableBool `json:"success,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Errors []*string `json:"errors,omitempty"`
}

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorResponse) GetSuccess() bool {
	if o == nil || o.Success.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Success.Get()
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorResponse) GetSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Success.Get(), o.Success.IsSet()
}

// HasSuccess returns a boolean if a field has been set.
func (o *ErrorResponse) HasSuccess() bool {
	if o != nil && o.Success.IsSet() {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given NullableBool and assigns it to the Success field.
func (o *ErrorResponse) SetSuccess(v bool) {
	o.Success.Set(&v)
}
// SetSuccessNil sets the value for Success to be an explicit nil
func (o *ErrorResponse) SetSuccessNil() {
	o.Success.Set(nil)
}

// UnsetSuccess ensures that no value is present for Success, not even an explicit nil
func (o *ErrorResponse) UnsetSuccess() {
	o.Success.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ErrorResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ErrorResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ErrorResponse) UnsetMessage() {
	o.Message.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorResponse) GetErrors() []*string {
	if o == nil  {
		var ret []*string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorResponse) GetErrorsOk() (*[]*string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []*string and assigns it to the Errors field.
func (o *ErrorResponse) SetErrors(v []*string) {
	o.Errors = v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success.IsSet() {
		toSerialize["success"] = o.Success.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


