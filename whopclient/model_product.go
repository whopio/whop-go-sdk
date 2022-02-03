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

// Product struct for Product
type Product struct {
	Id NullableInt64 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	LicenseType NullableString `json:"license_type,omitempty"`
	InitialPrice NullableFloat32 `json:"initial_price,omitempty"`
	RecurringPrice NullableFloat32 `json:"recurring_price,omitempty"`
	BillingPeriod NullableInt64 `json:"billing_period,omitempty"`
	DiscordRole NullableString `json:"discord_role,omitempty"`
	Stock NullableInt64 `json:"stock,omitempty"`
	Test NullableBool `json:"test,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Product) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Product) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Product) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Product) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Product) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Product) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Product) UnsetTitle() {
	o.Title.Unset()
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetLicenseType() string {
	if o == nil || o.LicenseType.Get() == nil {
		var ret string
		return ret
	}
	return *o.LicenseType.Get()
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetLicenseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LicenseType.Get(), o.LicenseType.IsSet()
}

// HasLicenseType returns a boolean if a field has been set.
func (o *Product) HasLicenseType() bool {
	if o != nil && o.LicenseType.IsSet() {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given NullableString and assigns it to the LicenseType field.
func (o *Product) SetLicenseType(v string) {
	o.LicenseType.Set(&v)
}
// SetLicenseTypeNil sets the value for LicenseType to be an explicit nil
func (o *Product) SetLicenseTypeNil() {
	o.LicenseType.Set(nil)
}

// UnsetLicenseType ensures that no value is present for LicenseType, not even an explicit nil
func (o *Product) UnsetLicenseType() {
	o.LicenseType.Unset()
}

// GetInitialPrice returns the InitialPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetInitialPrice() float32 {
	if o == nil || o.InitialPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.InitialPrice.Get()
}

// GetInitialPriceOk returns a tuple with the InitialPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetInitialPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitialPrice.Get(), o.InitialPrice.IsSet()
}

// HasInitialPrice returns a boolean if a field has been set.
func (o *Product) HasInitialPrice() bool {
	if o != nil && o.InitialPrice.IsSet() {
		return true
	}

	return false
}

// SetInitialPrice gets a reference to the given NullableFloat32 and assigns it to the InitialPrice field.
func (o *Product) SetInitialPrice(v float32) {
	o.InitialPrice.Set(&v)
}
// SetInitialPriceNil sets the value for InitialPrice to be an explicit nil
func (o *Product) SetInitialPriceNil() {
	o.InitialPrice.Set(nil)
}

// UnsetInitialPrice ensures that no value is present for InitialPrice, not even an explicit nil
func (o *Product) UnsetInitialPrice() {
	o.InitialPrice.Unset()
}

// GetRecurringPrice returns the RecurringPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetRecurringPrice() float32 {
	if o == nil || o.RecurringPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.RecurringPrice.Get()
}

// GetRecurringPriceOk returns a tuple with the RecurringPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetRecurringPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecurringPrice.Get(), o.RecurringPrice.IsSet()
}

// HasRecurringPrice returns a boolean if a field has been set.
func (o *Product) HasRecurringPrice() bool {
	if o != nil && o.RecurringPrice.IsSet() {
		return true
	}

	return false
}

// SetRecurringPrice gets a reference to the given NullableFloat32 and assigns it to the RecurringPrice field.
func (o *Product) SetRecurringPrice(v float32) {
	o.RecurringPrice.Set(&v)
}
// SetRecurringPriceNil sets the value for RecurringPrice to be an explicit nil
func (o *Product) SetRecurringPriceNil() {
	o.RecurringPrice.Set(nil)
}

// UnsetRecurringPrice ensures that no value is present for RecurringPrice, not even an explicit nil
func (o *Product) UnsetRecurringPrice() {
	o.RecurringPrice.Unset()
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetBillingPeriod() int64 {
	if o == nil || o.BillingPeriod.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillingPeriod.Get()
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetBillingPeriodOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingPeriod.Get(), o.BillingPeriod.IsSet()
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *Product) HasBillingPeriod() bool {
	if o != nil && o.BillingPeriod.IsSet() {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given NullableInt64 and assigns it to the BillingPeriod field.
func (o *Product) SetBillingPeriod(v int64) {
	o.BillingPeriod.Set(&v)
}
// SetBillingPeriodNil sets the value for BillingPeriod to be an explicit nil
func (o *Product) SetBillingPeriodNil() {
	o.BillingPeriod.Set(nil)
}

// UnsetBillingPeriod ensures that no value is present for BillingPeriod, not even an explicit nil
func (o *Product) UnsetBillingPeriod() {
	o.BillingPeriod.Unset()
}

// GetDiscordRole returns the DiscordRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetDiscordRole() string {
	if o == nil || o.DiscordRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.DiscordRole.Get()
}

// GetDiscordRoleOk returns a tuple with the DiscordRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetDiscordRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiscordRole.Get(), o.DiscordRole.IsSet()
}

// HasDiscordRole returns a boolean if a field has been set.
func (o *Product) HasDiscordRole() bool {
	if o != nil && o.DiscordRole.IsSet() {
		return true
	}

	return false
}

// SetDiscordRole gets a reference to the given NullableString and assigns it to the DiscordRole field.
func (o *Product) SetDiscordRole(v string) {
	o.DiscordRole.Set(&v)
}
// SetDiscordRoleNil sets the value for DiscordRole to be an explicit nil
func (o *Product) SetDiscordRoleNil() {
	o.DiscordRole.Set(nil)
}

// UnsetDiscordRole ensures that no value is present for DiscordRole, not even an explicit nil
func (o *Product) UnsetDiscordRole() {
	o.DiscordRole.Unset()
}

// GetStock returns the Stock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetStock() int64 {
	if o == nil || o.Stock.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Stock.Get()
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetStockOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Stock.Get(), o.Stock.IsSet()
}

// HasStock returns a boolean if a field has been set.
func (o *Product) HasStock() bool {
	if o != nil && o.Stock.IsSet() {
		return true
	}

	return false
}

// SetStock gets a reference to the given NullableInt64 and assigns it to the Stock field.
func (o *Product) SetStock(v int64) {
	o.Stock.Set(&v)
}
// SetStockNil sets the value for Stock to be an explicit nil
func (o *Product) SetStockNil() {
	o.Stock.Set(nil)
}

// UnsetStock ensures that no value is present for Stock, not even an explicit nil
func (o *Product) UnsetStock() {
	o.Stock.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetTest() bool {
	if o == nil || o.Test.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetTestOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *Product) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableBool and assigns it to the Test field.
func (o *Product) SetTest(v bool) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *Product) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *Product) UnsetTest() {
	o.Test.Unset()
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.LicenseType.IsSet() {
		toSerialize["license_type"] = o.LicenseType.Get()
	}
	if o.InitialPrice.IsSet() {
		toSerialize["initial_price"] = o.InitialPrice.Get()
	}
	if o.RecurringPrice.IsSet() {
		toSerialize["recurring_price"] = o.RecurringPrice.Get()
	}
	if o.BillingPeriod.IsSet() {
		toSerialize["billing_period"] = o.BillingPeriod.Get()
	}
	if o.DiscordRole.IsSet() {
		toSerialize["discord_role"] = o.DiscordRole.Get()
	}
	if o.Stock.IsSet() {
		toSerialize["stock"] = o.Stock.Get()
	}
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


