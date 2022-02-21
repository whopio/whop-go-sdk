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

// GetLicenseByKeyResponse struct for GetLicenseByKeyResponse
type GetLicenseByKeyResponse struct {
	Id NullableInt64 `json:"id,omitempty"`
	Key NullableString `json:"key,omitempty"`
	RentalKey NullableString `json:"rental_key,omitempty"`
	Rented NullableBool `json:"rented,omitempty"`
	RentalEnd NullableString `json:"rental_end,omitempty"`
	Valid NullableBool `json:"valid,omitempty"`
	KeyStatus NullableString `json:"key_status,omitempty"`
	SubscriptionStatus NullableString `json:"subscription_status,omitempty"`
	Quantity NullableInt64 `json:"quantity,omitempty"`
	Banned NullableBool `json:"banned,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	NextRenewalDate NullableString `json:"next_renewal_date,omitempty"`
	RenewalPeriod NullableInt64 `json:"renewal_period,omitempty"`
	CreatedAt NullableString `json:"created_at,omitempty"`
	Discord NullableLicenseDiscord `json:"discord,omitempty"`
	Twitter NullableLicenseTwitter `json:"twitter,omitempty"`
	IsScammer interface{} `json:"is_scammer,omitempty"`
	Plan NullableLicensePlan `json:"plan,omitempty"`
}

// NewGetLicenseByKeyResponse instantiates a new GetLicenseByKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLicenseByKeyResponse() *GetLicenseByKeyResponse {
	this := GetLicenseByKeyResponse{}
	return &this
}

// NewGetLicenseByKeyResponseWithDefaults instantiates a new GetLicenseByKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLicenseByKeyResponseWithDefaults() *GetLicenseByKeyResponse {
	this := GetLicenseByKeyResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *GetLicenseByKeyResponse) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GetLicenseByKeyResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetId() {
	o.Id.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *GetLicenseByKeyResponse) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *GetLicenseByKeyResponse) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetKey() {
	o.Key.Unset()
}

// GetRentalKey returns the RentalKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetRentalKey() string {
	if o == nil || o.RentalKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.RentalKey.Get()
}

// GetRentalKeyOk returns a tuple with the RentalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetRentalKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RentalKey.Get(), o.RentalKey.IsSet()
}

// HasRentalKey returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasRentalKey() bool {
	if o != nil && o.RentalKey.IsSet() {
		return true
	}

	return false
}

// SetRentalKey gets a reference to the given NullableString and assigns it to the RentalKey field.
func (o *GetLicenseByKeyResponse) SetRentalKey(v string) {
	o.RentalKey.Set(&v)
}
// SetRentalKeyNil sets the value for RentalKey to be an explicit nil
func (o *GetLicenseByKeyResponse) SetRentalKeyNil() {
	o.RentalKey.Set(nil)
}

// UnsetRentalKey ensures that no value is present for RentalKey, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetRentalKey() {
	o.RentalKey.Unset()
}

// GetRented returns the Rented field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetRented() bool {
	if o == nil || o.Rented.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Rented.Get()
}

// GetRentedOk returns a tuple with the Rented field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetRentedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Rented.Get(), o.Rented.IsSet()
}

// HasRented returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasRented() bool {
	if o != nil && o.Rented.IsSet() {
		return true
	}

	return false
}

// SetRented gets a reference to the given NullableBool and assigns it to the Rented field.
func (o *GetLicenseByKeyResponse) SetRented(v bool) {
	o.Rented.Set(&v)
}
// SetRentedNil sets the value for Rented to be an explicit nil
func (o *GetLicenseByKeyResponse) SetRentedNil() {
	o.Rented.Set(nil)
}

// UnsetRented ensures that no value is present for Rented, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetRented() {
	o.Rented.Unset()
}

// GetRentalEnd returns the RentalEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetRentalEnd() string {
	if o == nil || o.RentalEnd.Get() == nil {
		var ret string
		return ret
	}
	return *o.RentalEnd.Get()
}

// GetRentalEndOk returns a tuple with the RentalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetRentalEndOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RentalEnd.Get(), o.RentalEnd.IsSet()
}

// HasRentalEnd returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasRentalEnd() bool {
	if o != nil && o.RentalEnd.IsSet() {
		return true
	}

	return false
}

// SetRentalEnd gets a reference to the given NullableString and assigns it to the RentalEnd field.
func (o *GetLicenseByKeyResponse) SetRentalEnd(v string) {
	o.RentalEnd.Set(&v)
}
// SetRentalEndNil sets the value for RentalEnd to be an explicit nil
func (o *GetLicenseByKeyResponse) SetRentalEndNil() {
	o.RentalEnd.Set(nil)
}

// UnsetRentalEnd ensures that no value is present for RentalEnd, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetRentalEnd() {
	o.RentalEnd.Unset()
}

// GetValid returns the Valid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetValid() bool {
	if o == nil || o.Valid.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Valid.Get()
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetValidOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Valid.Get(), o.Valid.IsSet()
}

// HasValid returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasValid() bool {
	if o != nil && o.Valid.IsSet() {
		return true
	}

	return false
}

// SetValid gets a reference to the given NullableBool and assigns it to the Valid field.
func (o *GetLicenseByKeyResponse) SetValid(v bool) {
	o.Valid.Set(&v)
}
// SetValidNil sets the value for Valid to be an explicit nil
func (o *GetLicenseByKeyResponse) SetValidNil() {
	o.Valid.Set(nil)
}

// UnsetValid ensures that no value is present for Valid, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetValid() {
	o.Valid.Unset()
}

// GetKeyStatus returns the KeyStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetKeyStatus() string {
	if o == nil || o.KeyStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.KeyStatus.Get()
}

// GetKeyStatusOk returns a tuple with the KeyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetKeyStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KeyStatus.Get(), o.KeyStatus.IsSet()
}

// HasKeyStatus returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasKeyStatus() bool {
	if o != nil && o.KeyStatus.IsSet() {
		return true
	}

	return false
}

// SetKeyStatus gets a reference to the given NullableString and assigns it to the KeyStatus field.
func (o *GetLicenseByKeyResponse) SetKeyStatus(v string) {
	o.KeyStatus.Set(&v)
}
// SetKeyStatusNil sets the value for KeyStatus to be an explicit nil
func (o *GetLicenseByKeyResponse) SetKeyStatusNil() {
	o.KeyStatus.Set(nil)
}

// UnsetKeyStatus ensures that no value is present for KeyStatus, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetKeyStatus() {
	o.KeyStatus.Unset()
}

// GetSubscriptionStatus returns the SubscriptionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetSubscriptionStatus() string {
	if o == nil || o.SubscriptionStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionStatus.Get()
}

// GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetSubscriptionStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionStatus.Get(), o.SubscriptionStatus.IsSet()
}

// HasSubscriptionStatus returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasSubscriptionStatus() bool {
	if o != nil && o.SubscriptionStatus.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionStatus gets a reference to the given NullableString and assigns it to the SubscriptionStatus field.
func (o *GetLicenseByKeyResponse) SetSubscriptionStatus(v string) {
	o.SubscriptionStatus.Set(&v)
}
// SetSubscriptionStatusNil sets the value for SubscriptionStatus to be an explicit nil
func (o *GetLicenseByKeyResponse) SetSubscriptionStatusNil() {
	o.SubscriptionStatus.Set(nil)
}

// UnsetSubscriptionStatus ensures that no value is present for SubscriptionStatus, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetSubscriptionStatus() {
	o.SubscriptionStatus.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetQuantity() int64 {
	if o == nil || o.Quantity.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetQuantityOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt64 and assigns it to the Quantity field.
func (o *GetLicenseByKeyResponse) SetQuantity(v int64) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *GetLicenseByKeyResponse) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetBanned returns the Banned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetBanned() bool {
	if o == nil || o.Banned.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Banned.Get()
}

// GetBannedOk returns a tuple with the Banned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetBannedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Banned.Get(), o.Banned.IsSet()
}

// HasBanned returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasBanned() bool {
	if o != nil && o.Banned.IsSet() {
		return true
	}

	return false
}

// SetBanned gets a reference to the given NullableBool and assigns it to the Banned field.
func (o *GetLicenseByKeyResponse) SetBanned(v bool) {
	o.Banned.Set(&v)
}
// SetBannedNil sets the value for Banned to be an explicit nil
func (o *GetLicenseByKeyResponse) SetBannedNil() {
	o.Banned.Set(nil)
}

// UnsetBanned ensures that no value is present for Banned, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetBanned() {
	o.Banned.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetLicenseByKeyResponse) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicenseByKeyResponse) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GetLicenseByKeyResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetNextRenewalDate returns the NextRenewalDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetNextRenewalDate() string {
	if o == nil || o.NextRenewalDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextRenewalDate.Get()
}

// GetNextRenewalDateOk returns a tuple with the NextRenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetNextRenewalDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextRenewalDate.Get(), o.NextRenewalDate.IsSet()
}

// HasNextRenewalDate returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasNextRenewalDate() bool {
	if o != nil && o.NextRenewalDate.IsSet() {
		return true
	}

	return false
}

// SetNextRenewalDate gets a reference to the given NullableString and assigns it to the NextRenewalDate field.
func (o *GetLicenseByKeyResponse) SetNextRenewalDate(v string) {
	o.NextRenewalDate.Set(&v)
}
// SetNextRenewalDateNil sets the value for NextRenewalDate to be an explicit nil
func (o *GetLicenseByKeyResponse) SetNextRenewalDateNil() {
	o.NextRenewalDate.Set(nil)
}

// UnsetNextRenewalDate ensures that no value is present for NextRenewalDate, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetNextRenewalDate() {
	o.NextRenewalDate.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetRenewalPeriod() int64 {
	if o == nil || o.RenewalPeriod.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetRenewalPeriodOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableInt64 and assigns it to the RenewalPeriod field.
func (o *GetLicenseByKeyResponse) SetRenewalPeriod(v int64) {
	o.RenewalPeriod.Set(&v)
}
// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *GetLicenseByKeyResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *GetLicenseByKeyResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GetLicenseByKeyResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDiscord returns the Discord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetDiscord() LicenseDiscord {
	if o == nil || o.Discord.Get() == nil {
		var ret LicenseDiscord
		return ret
	}
	return *o.Discord.Get()
}

// GetDiscordOk returns a tuple with the Discord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetDiscordOk() (*LicenseDiscord, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Discord.Get(), o.Discord.IsSet()
}

// HasDiscord returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasDiscord() bool {
	if o != nil && o.Discord.IsSet() {
		return true
	}

	return false
}

// SetDiscord gets a reference to the given NullableLicenseDiscord and assigns it to the Discord field.
func (o *GetLicenseByKeyResponse) SetDiscord(v LicenseDiscord) {
	o.Discord.Set(&v)
}
// SetDiscordNil sets the value for Discord to be an explicit nil
func (o *GetLicenseByKeyResponse) SetDiscordNil() {
	o.Discord.Set(nil)
}

// UnsetDiscord ensures that no value is present for Discord, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetDiscord() {
	o.Discord.Unset()
}

// GetTwitter returns the Twitter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetTwitter() LicenseTwitter {
	if o == nil || o.Twitter.Get() == nil {
		var ret LicenseTwitter
		return ret
	}
	return *o.Twitter.Get()
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetTwitterOk() (*LicenseTwitter, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Twitter.Get(), o.Twitter.IsSet()
}

// HasTwitter returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasTwitter() bool {
	if o != nil && o.Twitter.IsSet() {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given NullableLicenseTwitter and assigns it to the Twitter field.
func (o *GetLicenseByKeyResponse) SetTwitter(v LicenseTwitter) {
	o.Twitter.Set(&v)
}
// SetTwitterNil sets the value for Twitter to be an explicit nil
func (o *GetLicenseByKeyResponse) SetTwitterNil() {
	o.Twitter.Set(nil)
}

// UnsetTwitter ensures that no value is present for Twitter, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetTwitter() {
	o.Twitter.Unset()
}

// GetIsScammer returns the IsScammer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetIsScammer() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.IsScammer
}

// GetIsScammerOk returns a tuple with the IsScammer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetIsScammerOk() (*interface{}, bool) {
	if o == nil || o.IsScammer == nil {
		return nil, false
	}
	return &o.IsScammer, true
}

// HasIsScammer returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasIsScammer() bool {
	if o != nil && o.IsScammer != nil {
		return true
	}

	return false
}

// SetIsScammer gets a reference to the given interface{} and assigns it to the IsScammer field.
func (o *GetLicenseByKeyResponse) SetIsScammer(v interface{}) {
	o.IsScammer = v
}

// GetPlan returns the Plan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetLicenseByKeyResponse) GetPlan() LicensePlan {
	if o == nil || o.Plan.Get() == nil {
		var ret LicensePlan
		return ret
	}
	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetLicenseByKeyResponse) GetPlanOk() (*LicensePlan, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// HasPlan returns a boolean if a field has been set.
func (o *GetLicenseByKeyResponse) HasPlan() bool {
	if o != nil && o.Plan.IsSet() {
		return true
	}

	return false
}

// SetPlan gets a reference to the given NullableLicensePlan and assigns it to the Plan field.
func (o *GetLicenseByKeyResponse) SetPlan(v LicensePlan) {
	o.Plan.Set(&v)
}
// SetPlanNil sets the value for Plan to be an explicit nil
func (o *GetLicenseByKeyResponse) SetPlanNil() {
	o.Plan.Set(nil)
}

// UnsetPlan ensures that no value is present for Plan, not even an explicit nil
func (o *GetLicenseByKeyResponse) UnsetPlan() {
	o.Plan.Unset()
}

func (o GetLicenseByKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.RentalKey.IsSet() {
		toSerialize["rental_key"] = o.RentalKey.Get()
	}
	if o.Rented.IsSet() {
		toSerialize["rented"] = o.Rented.Get()
	}
	if o.RentalEnd.IsSet() {
		toSerialize["rental_end"] = o.RentalEnd.Get()
	}
	if o.Valid.IsSet() {
		toSerialize["valid"] = o.Valid.Get()
	}
	if o.KeyStatus.IsSet() {
		toSerialize["key_status"] = o.KeyStatus.Get()
	}
	if o.SubscriptionStatus.IsSet() {
		toSerialize["subscription_status"] = o.SubscriptionStatus.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if o.Banned.IsSet() {
		toSerialize["banned"] = o.Banned.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.NextRenewalDate.IsSet() {
		toSerialize["next_renewal_date"] = o.NextRenewalDate.Get()
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewal_period"] = o.RenewalPeriod.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Discord.IsSet() {
		toSerialize["discord"] = o.Discord.Get()
	}
	if o.Twitter.IsSet() {
		toSerialize["twitter"] = o.Twitter.Get()
	}
	if o.IsScammer != nil {
		toSerialize["is_scammer"] = o.IsScammer
	}
	if o.Plan.IsSet() {
		toSerialize["plan"] = o.Plan.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetLicenseByKeyResponse struct {
	value *GetLicenseByKeyResponse
	isSet bool
}

func (v NullableGetLicenseByKeyResponse) Get() *GetLicenseByKeyResponse {
	return v.value
}

func (v *NullableGetLicenseByKeyResponse) Set(val *GetLicenseByKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLicenseByKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLicenseByKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLicenseByKeyResponse(val *GetLicenseByKeyResponse) *NullableGetLicenseByKeyResponse {
	return &NullableGetLicenseByKeyResponse{value: val, isSet: true}
}

func (v NullableGetLicenseByKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLicenseByKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


