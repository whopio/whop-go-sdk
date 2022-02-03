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

// LicensePlan struct for LicensePlan
type LicensePlan struct {
	Id NullableInt64 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	InitialPrice NullableInt64 `json:"initial_price,omitempty"`
	RenewalPrice NullableInt64 `json:"renewal_price,omitempty"`
	BillingPeriod NullableInt64 `json:"billing_period,omitempty"`
	LicenseType NullableString `json:"license_type,omitempty"`
}

// NewLicensePlan instantiates a new LicensePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensePlan() *LicensePlan {
	this := LicensePlan{}
	return &this
}

// NewLicensePlanWithDefaults instantiates a new LicensePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensePlanWithDefaults() *LicensePlan {
	this := LicensePlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *LicensePlan) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *LicensePlan) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *LicensePlan) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *LicensePlan) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LicensePlan) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LicensePlan) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LicensePlan) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LicensePlan) UnsetTitle() {
	o.Title.Unset()
}

// GetInitialPrice returns the InitialPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetInitialPrice() int64 {
	if o == nil || o.InitialPrice.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InitialPrice.Get()
}

// GetInitialPriceOk returns a tuple with the InitialPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetInitialPriceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitialPrice.Get(), o.InitialPrice.IsSet()
}

// HasInitialPrice returns a boolean if a field has been set.
func (o *LicensePlan) HasInitialPrice() bool {
	if o != nil && o.InitialPrice.IsSet() {
		return true
	}

	return false
}

// SetInitialPrice gets a reference to the given NullableInt64 and assigns it to the InitialPrice field.
func (o *LicensePlan) SetInitialPrice(v int64) {
	o.InitialPrice.Set(&v)
}
// SetInitialPriceNil sets the value for InitialPrice to be an explicit nil
func (o *LicensePlan) SetInitialPriceNil() {
	o.InitialPrice.Set(nil)
}

// UnsetInitialPrice ensures that no value is present for InitialPrice, not even an explicit nil
func (o *LicensePlan) UnsetInitialPrice() {
	o.InitialPrice.Unset()
}

// GetRenewalPrice returns the RenewalPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetRenewalPrice() int64 {
	if o == nil || o.RenewalPrice.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RenewalPrice.Get()
}

// GetRenewalPriceOk returns a tuple with the RenewalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetRenewalPriceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenewalPrice.Get(), o.RenewalPrice.IsSet()
}

// HasRenewalPrice returns a boolean if a field has been set.
func (o *LicensePlan) HasRenewalPrice() bool {
	if o != nil && o.RenewalPrice.IsSet() {
		return true
	}

	return false
}

// SetRenewalPrice gets a reference to the given NullableInt64 and assigns it to the RenewalPrice field.
func (o *LicensePlan) SetRenewalPrice(v int64) {
	o.RenewalPrice.Set(&v)
}
// SetRenewalPriceNil sets the value for RenewalPrice to be an explicit nil
func (o *LicensePlan) SetRenewalPriceNil() {
	o.RenewalPrice.Set(nil)
}

// UnsetRenewalPrice ensures that no value is present for RenewalPrice, not even an explicit nil
func (o *LicensePlan) UnsetRenewalPrice() {
	o.RenewalPrice.Unset()
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetBillingPeriod() int64 {
	if o == nil || o.BillingPeriod.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillingPeriod.Get()
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetBillingPeriodOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingPeriod.Get(), o.BillingPeriod.IsSet()
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *LicensePlan) HasBillingPeriod() bool {
	if o != nil && o.BillingPeriod.IsSet() {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given NullableInt64 and assigns it to the BillingPeriod field.
func (o *LicensePlan) SetBillingPeriod(v int64) {
	o.BillingPeriod.Set(&v)
}
// SetBillingPeriodNil sets the value for BillingPeriod to be an explicit nil
func (o *LicensePlan) SetBillingPeriodNil() {
	o.BillingPeriod.Set(nil)
}

// UnsetBillingPeriod ensures that no value is present for BillingPeriod, not even an explicit nil
func (o *LicensePlan) UnsetBillingPeriod() {
	o.BillingPeriod.Unset()
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicensePlan) GetLicenseType() string {
	if o == nil || o.LicenseType.Get() == nil {
		var ret string
		return ret
	}
	return *o.LicenseType.Get()
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicensePlan) GetLicenseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LicenseType.Get(), o.LicenseType.IsSet()
}

// HasLicenseType returns a boolean if a field has been set.
func (o *LicensePlan) HasLicenseType() bool {
	if o != nil && o.LicenseType.IsSet() {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given NullableString and assigns it to the LicenseType field.
func (o *LicensePlan) SetLicenseType(v string) {
	o.LicenseType.Set(&v)
}
// SetLicenseTypeNil sets the value for LicenseType to be an explicit nil
func (o *LicensePlan) SetLicenseTypeNil() {
	o.LicenseType.Set(nil)
}

// UnsetLicenseType ensures that no value is present for LicenseType, not even an explicit nil
func (o *LicensePlan) UnsetLicenseType() {
	o.LicenseType.Unset()
}

func (o LicensePlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.InitialPrice.IsSet() {
		toSerialize["initial_price"] = o.InitialPrice.Get()
	}
	if o.RenewalPrice.IsSet() {
		toSerialize["renewal_price"] = o.RenewalPrice.Get()
	}
	if o.BillingPeriod.IsSet() {
		toSerialize["billing_period"] = o.BillingPeriod.Get()
	}
	if o.LicenseType.IsSet() {
		toSerialize["license_type"] = o.LicenseType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLicensePlan struct {
	value *LicensePlan
	isSet bool
}

func (v NullableLicensePlan) Get() *LicensePlan {
	return v.value
}

func (v *NullableLicensePlan) Set(val *LicensePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensePlan(val *LicensePlan) *NullableLicensePlan {
	return &NullableLicensePlan{value: val, isSet: true}
}

func (v NullableLicensePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

