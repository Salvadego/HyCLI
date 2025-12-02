package internal

import (
	"context"

	"github.com/Salvadego/hac/hac"
)

func New(baseURL, username, password string, skip bool) *hac.HACClient {
	return hac.NewClient(&hac.Config{
		BaseURL:       baseURL,
		Username:      username,
		Password:      password,
		SkipTLSVerify: skip,
	})
}

func Login(c *hac.HACClient, ctx context.Context) error {
	return c.Auth.Login(ctx)
}
