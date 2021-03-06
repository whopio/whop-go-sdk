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

// GetCheckoutLogsResponse struct for GetCheckoutLogsResponse
type GetCheckoutLogsResponse struct {
	Checkouts *[]CheckoutLog `json:"checkouts,omitempty"`
	Total NullableInt64 `json:"total,omitempty"`
}

// NewGetCheckoutLogsResponse instantiates a new GetCheckoutLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCheckoutLogsResponse() *GetCheckoutLogsResponse {
	this := GetCheckoutLogsResponse{}
	return &this
}

// NewGetCheckoutLogsResponseWithDefaults instantiates a new GetCheckoutLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCheckoutLogsResponseWithDefaults() *GetCheckoutLogsResponse {
	this := GetCheckoutLogsResponse{}
	return &this
}

// GetCheckouts returns the Checkouts field value if set, zero value otherwise.
func (o *GetCheckoutLogsResponse) GetCheckouts() []CheckoutLog {
	if o == nil || o.Checkouts == nil {
		var ret []CheckoutLog
		return ret
	}
	return *o.Checkouts
}

// GetCheckoutsOk returns a tuple with the Checkouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCheckoutLogsResponse) GetCheckoutsOk() (*[]CheckoutLog, bool) {
	if o == nil || o.Checkouts == nil {
		return nil, false
	}
	return o.Checkouts, true
}

// HasCheckouts returns a boolean if a field has been set.
func (o *GetCheckoutLogsResponse) HasCheckouts() bool {
	if o != nil && o.Checkouts != nil {
		return true
	}

	return false
}

// SetCheckouts gets a reference to the given []CheckoutLog and assigns it to the Checkouts field.
func (o *GetCheckoutLogsResponse) SetCheckouts(v []CheckoutLog) {
	o.Checkouts = &v
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetCheckoutLogsResponse) GetTotal() int64 {
	if o == nil || o.Total.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetCheckoutLogsResponse) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *GetCheckoutLogsResponse) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt64 and assigns it to the Total field.
func (o *GetCheckoutLogsResponse) SetTotal(v int64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *GetCheckoutLogsResponse) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *GetCheckoutLogsResponse) UnsetTotal() {
	o.Total.Unset()
}

func (o GetCheckoutLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checkouts != nil {
		toSerialize["checkouts"] = o.Checkouts
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetCheckoutLogsResponse struct {
	value *GetCheckoutLogsResponse
	isSet bool
}

func (v NullableGetCheckoutLogsResponse) Get() *GetCheckoutLogsResponse {
	return v.value
}

func (v *NullableGetCheckoutLogsResponse) Set(val *GetCheckoutLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCheckoutLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCheckoutLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCheckoutLogsResponse(val *GetCheckoutLogsResponse) *NullableGetCheckoutLogsResponse {
	return &NullableGetCheckoutLogsResponse{value: val, isSet: true}
}

func (v NullableGetCheckoutLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCheckoutLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


