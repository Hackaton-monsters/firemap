package geo_ip

import (
	"context"
	"encoding/json"
	"firemap/internal/infrastructure/config"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type LocationIQResponse struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`

	Address struct {
		Road        string `json:"road"`
		Name        string `json:"name"`
		HouseNumber string `json:"house_number"`
		Suburb      string `json:"suburb"`
		County      string `json:"county"`
		City        string `json:"city"`
		Town        string `json:"town"`
		Village     string `json:"village"`
		State       string `json:"state"`
		Postcode    string `json:"postcode"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
	} `json:"address"`
}

type Address struct {
	DisplayName string

	Street      string
	HouseNumber string
	District    string
	City        string
	State       string
	Postcode    string
	Country     string
	CountryCode string
}

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

type InfoGetter interface {
	GetDisplayNameByCoordinate(ctx context.Context, lat, lon float64) (string, error)
}

type Client struct {
	config     *config.Config
	httpClient *http.Client
}

func NewClient(config *config.Config) InfoGetter {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

func (c *Client) GetDisplayNameByCoordinate(ctx context.Context, lat, lon float64) (string, error) {
	u, err := url.Parse(c.config.GeoIP.GeoIPUrl)
	if err != nil {
		return "", fmt.Errorf("parse base url: %w", err)
	}

	q := u.Query()
	q.Set("key", c.config.GeoIP.GeoIPKey)
	q.Set("lat", fmt.Sprintf("%f", lat))
	q.Set("lon", fmt.Sprintf("%f", lon))
	q.Set("format", "json")
	q.Set("normalizeaddress", "1")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("User-Agent", "firemap/1.0 (firemap@test.dev)")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("locationiq status: %s", resp.Status)
	}

	var liq LocationIQResponse
	if err := json.NewDecoder(resp.Body).Decode(&liq); err != nil {
		return "", fmt.Errorf("decode json: %w", err)
	}

	return liq.Address.Suburb, nil
}
