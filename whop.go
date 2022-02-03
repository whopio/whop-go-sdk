package whop

import (
	"context"

	. "github.com/whopio/go-whop-sdk/sdk/whopclient"
)

type Whop struct {
	client *APIClient
}

func NewWhop(bearer string) Whop {
	cfg := NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer "+bearer)
	client := NewAPIClient(cfg)
	return Whop{
		client: client,
	}
}

// Fetch All Products
func (w Whop) GetProducts(ctx context.Context) (GetProductsResponse, error) {
	req := w.client.ProductsApi.GetProducts(ctx)
	resp, _, err := req.Execute()
	return resp, err
}

// Fetch Product
func (w Whop) GetProductById(ctx context.Context, id int64) (GetProductByIdResponse, error) {
	req := w.client.ProductsApi.GetProductById(ctx, id)
	resp, _, err := req.Execute()
	return resp, err
}

// Create Product
func (w Whop) CreateProduct(ctx context.Context, body CreateProductRequest) (CreateProductResponse, error) {
	req := w.client.ProductsApi.CreateProduct(ctx)
	req = req.CreateProductRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Product Creation Confirmation
func (w Whop) ConfirmProduct(ctx context.Context, body ConfirmProductRequest) (ConfirmProductResponse, error) {
	req := w.client.ProductsApi.ConfirmProduct(ctx)
	req = req.ConfirmProductRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Send a push notification
func (w Whop) SendPushNotification(ctx context.Context, body SendPushNotificationRequest) (SendPushNotificationResponse, error) {
	req := w.client.NotificationsApi.SendPushNotification(ctx)
	req = req.SendPushNotificationRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Fetch Links
func (w Whop) GetLinks(ctx context.Context) (GetLinksResponse, error) {
	req := w.client.LinksApi.GetLinks(ctx)
	resp, _, err := req.Execute()
	return resp, err
}

// Create Password Protected Link
func (w Whop) CreateLink(ctx context.Context, body CreateLinkRequest) (CreateLinkResponse, error) {
	req := w.client.LinksApi.CreateLink(ctx)
	req = req.CreateLinkRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Validate Key
func (w Whop) ValidateLicenseByKey(ctx context.Context, key string, body ValidateLicenseByKeyRequest) (ValidateLicenseByKeyResponse, error) {
	req := w.client.LicensesApi.ValidateLicenseByKey(ctx, key)
	req = req.ValidateLicenseByKeyRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Update License
func (w Whop) UpdateLicenseByKey(ctx context.Context, key string, body UpdateLicenseByKeyRequest) (UpdateLicenseByKeyResponse, error) {
	req := w.client.LicensesApi.UpdateLicenseByKey(ctx, key)
	req = req.UpdateLicenseByKeyRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Reset License
func (w Whop) ResetLicenseByKey(ctx context.Context, key string, body ResetLicenseByKeyRequest) (ResetLicenseByKeyResponse, error) {
	req := w.client.LicensesApi.ResetLicenseByKey(ctx, key)
	req = req.ResetLicenseByKeyRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Fetch All Licenses
func (w Whop) GetLicenses(ctx context.Context, discordAccountId *string, page *int64, start *string, end *string) (GetLicensesResponse, error) {
	req := w.client.LicensesApi.GetLicenses(ctx)
	if discordAccountId != nil {
		req = req.DiscordAccountId(*discordAccountId)
	}
	if page != nil {
		req = req.Page(*page)
	}
	if start != nil {
		req = req.Start(*start)
	}
	if end != nil {
		req = req.End(*end)
	}
	resp, _, err := req.Execute()
	return resp, err
}

// Fetch License
func (w Whop) GetLicenseByKey(ctx context.Context, key string) (GetLicenseByKeyResponse, error) {
	req := w.client.LicensesApi.GetLicenseByKey(ctx, key)
	resp, _, err := req.Execute()
	return resp, err
}

// Ban License
func (w Whop) BanLicenseByKey(ctx context.Context, key string, body BanLicenseByKeyRequest) (BanLicenseByKeyResponse, error) {
	req := w.client.LicensesApi.BanLicenseByKey(ctx, key)
	req = req.BanLicenseByKeyRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}

// Fetch Checkout Logs
func (w Whop) GetCheckoutLogs(ctx context.Context, key *string, status *string) (GetCheckoutLogsResponse, error) {
	req := w.client.CheckoutLogsApi.GetCheckoutLogs(ctx)
	if key != nil {
		req = req.Key(*key)
	}
	if status != nil {
		req = req.Status(*status)
	}
	resp, _, err := req.Execute()
	return resp, err
}

// Add Checkout Log
func (w Whop) CreateCheckoutLog(ctx context.Context, body CreateCheckoutLogRequest) (CreateCheckoutLogResponse, error) {
	req := w.client.CheckoutLogsApi.CreateCheckoutLog(ctx)
	req = req.CreateCheckoutLogRequest(body)
	resp, _, err := req.Execute()
	return resp, err
}
