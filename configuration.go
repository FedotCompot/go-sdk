/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"net/http"
	"strings"

	"github.com/openfga/go-sdk/credentials"
)

var SdkVersion = "0.0.1"

// RetryParams configures configuration for retry in case of HTTP too many request
type RetryParams struct {
	MaxRetry    int `json:"maxRetry,omitempty"`
	MinWaitInMs int `json:"minWaitInMs,omitempty"`
}

// Configuration stores the configuration of the API client
type Configuration struct {
	ApiScheme      string                   `json:"apiScheme,omitempty"`
	ApiHost        string                   `json:"apiHost,omitempty"`
	StoreId        string                   `json:"storeId,omitempty"`
	Credentials    *credentials.Credentials `json:"credentials,omitempty"`
	DefaultHeaders map[string]string        `json:"defaultHeader,omitempty"`
	UserAgent      string                   `json:"userAgent,omitempty"`
	Debug          bool                     `json:"debug,omitempty"`
	HTTPClient     *http.Client
	RetryParams    *RetryParams
}

// DefaultRetryParams returns the default retry parameters
func DefaultRetryParams() *RetryParams {
	return &RetryParams{
		MaxRetry:    3,
		MinWaitInMs: 100,
	}
}

func GetSdkUserAgent() string {
	userAgent := strings.Replace("openfga-sdk {sdkId}/{packageVersion}", "{sdkId}", "go", -1)
	userAgent = strings.Replace(userAgent, "{packageVersion}", SdkVersion, -1)

	return userAgent
}

// NewConfiguration returns a new Configuration object
func NewConfiguration(config Configuration) (*Configuration, error) {
	cfg := &Configuration{
		ApiScheme:      config.ApiScheme,
		ApiHost:        config.ApiHost,
		StoreId:        config.StoreId,
		Credentials:    config.Credentials,
		DefaultHeaders: make(map[string]string),
		UserAgent:      GetSdkUserAgent(),
		Debug:          false,
		RetryParams:    config.RetryParams,
	}

	if cfg.ApiScheme == "" {
		cfg.ApiScheme = "https"
	}

	err := cfg.ValidateConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeaders[key] = value
}

// ValidateConfig ensures that the given configuration is valid
func (c *Configuration) ValidateConfig() error {
	if c.ApiHost == "" {
		return reportError("Configuration.ApiHost is required")
	}

	if c.ApiScheme == "" {
		return reportError("Configuration.ApiScheme is required")
	}

	if !IsWellFormedUri(c.ApiScheme + "://" + c.ApiHost) {
		return reportError("Configuration.ApiScheme and Configuration.ApiHost (%s) do not generate a valid uri", c.ApiScheme+"://"+c.ApiHost)
	}

	if c.Credentials != nil {
		if err := c.Credentials.ValidateCredentialsConfig(); err != nil {
			return reportError("Credentials are invalid: %v", err)
		}
	}

	if c.RetryParams != nil && c.RetryParams.MaxRetry > 5 {
		return reportError("Configuration.RetryParams.MaxRetry exceeds maximum allowed limit of 5")
	}

	return nil
}
