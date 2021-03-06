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

// CreateProductRequest struct for CreateProductRequest
type CreateProductRequest struct {
	// Defines what occurs when a user's license is canceled. Can be: kick_user, remove_roles, remove_all_roles, no_action.
	CancelAction *string `json:"cancel_action,omitempty"`
	// Title of the product.
	Title *string `json:"title,omitempty"`
	// If you would like this product to expire for users who are on it, you can set an expiration days. After this many days, the users license will be terminated.
	ExpirationDays *int64 `json:"expiration_days,omitempty"`
	// 3 letter currency code.
	Currency *string `json:"currency,omitempty"`
	// One time payments charge the user at purchase Renewal charge the user every set period. Can be: one_time, renewal.
	LicenseType *string `json:"license_type,omitempty"`
	// The initial purchase price.
	InitialPrice *float32 `json:"initial_price,omitempty"`
	// The renewal price.
	Price *float32 `json:"price,omitempty"`
	// If the product is renewal, you can pass in (days) of how often the period should renew.
	BillingPeriod *int64 `json:"billing_period,omitempty"`
	// Amount of inventory available.
	Stock *int64 `json:"stock,omitempty"`
	Transferable *bool `json:"transferable,omitempty"`
	// Show time of the trial period in days.
	CustomTrialPeriod *int64 `json:"custom_trial_period,omitempty"`
}

// NewCreateProductRequest instantiates a new CreateProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductRequest() *CreateProductRequest {
	this := CreateProductRequest{}
	return &this
}

// NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductRequestWithDefaults() *CreateProductRequest {
	this := CreateProductRequest{}
	return &this
}

// GetCancelAction returns the CancelAction field value if set, zero value otherwise.
func (o *CreateProductRequest) GetCancelAction() string {
	if o == nil || o.CancelAction == nil {
		var ret string
		return ret
	}
	return *o.CancelAction
}

// GetCancelActionOk returns a tuple with the CancelAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetCancelActionOk() (*string, bool) {
	if o == nil || o.CancelAction == nil {
		return nil, false
	}
	return o.CancelAction, true
}

// HasCancelAction returns a boolean if a field has been set.
func (o *CreateProductRequest) HasCancelAction() bool {
	if o != nil && o.CancelAction != nil {
		return true
	}

	return false
}

// SetCancelAction gets a reference to the given string and assigns it to the CancelAction field.
func (o *CreateProductRequest) SetCancelAction(v string) {
	o.CancelAction = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateProductRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateProductRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateProductRequest) SetTitle(v string) {
	o.Title = &v
}

// GetExpirationDays returns the ExpirationDays field value if set, zero value otherwise.
func (o *CreateProductRequest) GetExpirationDays() int64 {
	if o == nil || o.ExpirationDays == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDays
}

// GetExpirationDaysOk returns a tuple with the ExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetExpirationDaysOk() (*int64, bool) {
	if o == nil || o.ExpirationDays == nil {
		return nil, false
	}
	return o.ExpirationDays, true
}

// HasExpirationDays returns a boolean if a field has been set.
func (o *CreateProductRequest) HasExpirationDays() bool {
	if o != nil && o.ExpirationDays != nil {
		return true
	}

	return false
}

// SetExpirationDays gets a reference to the given int64 and assigns it to the ExpirationDays field.
func (o *CreateProductRequest) SetExpirationDays(v int64) {
	o.ExpirationDays = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateProductRequest) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateProductRequest) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateProductRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *CreateProductRequest) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *CreateProductRequest) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *CreateProductRequest) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetInitialPrice returns the InitialPrice field value if set, zero value otherwise.
func (o *CreateProductRequest) GetInitialPrice() float32 {
	if o == nil || o.InitialPrice == nil {
		var ret float32
		return ret
	}
	return *o.InitialPrice
}

// GetInitialPriceOk returns a tuple with the InitialPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetInitialPriceOk() (*float32, bool) {
	if o == nil || o.InitialPrice == nil {
		return nil, false
	}
	return o.InitialPrice, true
}

// HasInitialPrice returns a boolean if a field has been set.
func (o *CreateProductRequest) HasInitialPrice() bool {
	if o != nil && o.InitialPrice != nil {
		return true
	}

	return false
}

// SetInitialPrice gets a reference to the given float32 and assigns it to the InitialPrice field.
func (o *CreateProductRequest) SetInitialPrice(v float32) {
	o.InitialPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateProductRequest) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateProductRequest) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *CreateProductRequest) SetPrice(v float32) {
	o.Price = &v
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *CreateProductRequest) GetBillingPeriod() int64 {
	if o == nil || o.BillingPeriod == nil {
		var ret int64
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetBillingPeriodOk() (*int64, bool) {
	if o == nil || o.BillingPeriod == nil {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *CreateProductRequest) HasBillingPeriod() bool {
	if o != nil && o.BillingPeriod != nil {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given int64 and assigns it to the BillingPeriod field.
func (o *CreateProductRequest) SetBillingPeriod(v int64) {
	o.BillingPeriod = &v
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *CreateProductRequest) GetStock() int64 {
	if o == nil || o.Stock == nil {
		var ret int64
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetStockOk() (*int64, bool) {
	if o == nil || o.Stock == nil {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *CreateProductRequest) HasStock() bool {
	if o != nil && o.Stock != nil {
		return true
	}

	return false
}

// SetStock gets a reference to the given int64 and assigns it to the Stock field.
func (o *CreateProductRequest) SetStock(v int64) {
	o.Stock = &v
}

// GetTransferable returns the Transferable field value if set, zero value otherwise.
func (o *CreateProductRequest) GetTransferable() bool {
	if o == nil || o.Transferable == nil {
		var ret bool
		return ret
	}
	return *o.Transferable
}

// GetTransferableOk returns a tuple with the Transferable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetTransferableOk() (*bool, bool) {
	if o == nil || o.Transferable == nil {
		return nil, false
	}
	return o.Transferable, true
}

// HasTransferable returns a boolean if a field has been set.
func (o *CreateProductRequest) HasTransferable() bool {
	if o != nil && o.Transferable != nil {
		return true
	}

	return false
}

// SetTransferable gets a reference to the given bool and assigns it to the Transferable field.
func (o *CreateProductRequest) SetTransferable(v bool) {
	o.Transferable = &v
}

// GetCustomTrialPeriod returns the CustomTrialPeriod field value if set, zero value otherwise.
func (o *CreateProductRequest) GetCustomTrialPeriod() int64 {
	if o == nil || o.CustomTrialPeriod == nil {
		var ret int64
		return ret
	}
	return *o.CustomTrialPeriod
}

// GetCustomTrialPeriodOk returns a tuple with the CustomTrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductRequest) GetCustomTrialPeriodOk() (*int64, bool) {
	if o == nil || o.CustomTrialPeriod == nil {
		return nil, false
	}
	return o.CustomTrialPeriod, true
}

// HasCustomTrialPeriod returns a boolean if a field has been set.
func (o *CreateProductRequest) HasCustomTrialPeriod() bool {
	if o != nil && o.CustomTrialPeriod != nil {
		return true
	}

	return false
}

// SetCustomTrialPeriod gets a reference to the given int64 and assigns it to the CustomTrialPeriod field.
func (o *CreateProductRequest) SetCustomTrialPeriod(v int64) {
	o.CustomTrialPeriod = &v
}

func (o CreateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CancelAction != nil {
		toSerialize["cancel_action"] = o.CancelAction
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ExpirationDays != nil {
		toSerialize["expiration_days"] = o.ExpirationDays
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.LicenseType != nil {
		toSerialize["license_type"] = o.LicenseType
	}
	if o.InitialPrice != nil {
		toSerialize["initial_price"] = o.InitialPrice
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.BillingPeriod != nil {
		toSerialize["billing_period"] = o.BillingPeriod
	}
	if o.Stock != nil {
		toSerialize["stock"] = o.Stock
	}
	if o.Transferable != nil {
		toSerialize["transferable"] = o.Transferable
	}
	if o.CustomTrialPeriod != nil {
		toSerialize["custom_trial_period"] = o.CustomTrialPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProductRequest struct {
	value *CreateProductRequest
	isSet bool
}

func (v NullableCreateProductRequest) Get() *CreateProductRequest {
	return v.value
}

func (v *NullableCreateProductRequest) Set(val *CreateProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductRequest(val *CreateProductRequest) *NullableCreateProductRequest {
	return &NullableCreateProductRequest{value: val, isSet: true}
}

func (v NullableCreateProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


