package internal

import (
	"context"

	"github.com/Salvadego/hac/hac"
)

func New(baseURL, user, pass string, skip bool) *hac.HACClient {
	return hac.NewClient(&hac.Config{
		BaseURL:       baseURL,
		Username:      user,
		Password:      pass,
		SkipTLSVerify: skip,
	})
}

func Login(c *hac.HACClient, ctx context.Context) error {
	return c.Auth.Login(ctx)
}
