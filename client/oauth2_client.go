// Copyright 2025 Nametag Inc.
//
// All information contained herein is the property of Nametag Inc.. The
// intellectual and technical concepts contained herein are proprietary, trade
// secrets, and/or confidential to Nametag, Inc. and may be covered by U.S.
// and Foreign Patents, patents in process, and are protected by trade secret or
// copyright law. Reproduction or distribution, in whole or in part, is
// forbidden except by express written permission of Nametag, Inc.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

func NewOAUTH2Client() (*http.Client, error) {
	tokenURL := os.Getenv("WORKDAY_TOKEN_URL")
	clientID := os.Getenv("WORKDAY_CLIENT_ID")
	clientSecret := os.Getenv("WORKDAY_CLIENT_SECRET")
	clientRefreshToken := os.Getenv("WORKDAY_CLIENT_REFRESH_TOKEN")

	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
	}

	tokenSource := oauth2.ReuseTokenSource(nil, &RefreshTokenSource{
		config:       conf,
		refreshToken: clientRefreshToken,
		ctx:          ctx,
		cachePath:    "./.oauth2_token_cache.json",
	})

	return oauth2.NewClient(ctx, tokenSource), nil
}

type RefreshTokenSource struct {
	config       *oauth2.Config
	refreshToken string
	ctx          context.Context
	cachePath    string
	mu           sync.Mutex
}

func (r *RefreshTokenSource) Token() (*oauth2.Token, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Try to load token from cache
	if r.cachePath != "" {
		if f, err := os.Open(r.cachePath); err == nil {
			defer f.Close()
			var cached oauth2.Token
			if err := json.NewDecoder(f).Decode(&cached); err == nil && cached.Valid() {
				fmt.Printf("Using cached access token: %s\n", cached.AccessToken[1:8]+"..."+cached.AccessToken[len(cached.AccessToken)-8:])
				return &cached, nil
			}
		}
	}

	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", r.refreshToken)
	form.Set("client_id", r.config.ClientID)
	form.Set("client_secret", r.config.ClientSecret)

	res, err := http.PostForm(r.config.Endpoint.TokenURL, form)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	fmt.Printf("[%d](%s) Got new access token: %s\n",
		res.StatusCode, res.Status,
		token.AccessToken[1:8]+"..."+token.AccessToken[len(token.AccessToken)-8:])

	// Is there an expiration attached?
	// If not, assume default expiry of minimum 1 hour per Workday docs
	// https://workday.my.site.com/customercenter/article?no=000020655&redirect=false
	if token.Expiry.IsZero() || token.Expiry.Before(time.Now()) || token.ExpiresIn < 1 {
		token.Expiry = time.Now().Add(1 * time.Hour)
	}

	// Save token to cache
	if r.cachePath != "" {
		if f, err := os.Create(r.cachePath); err == nil {
			_ = json.NewEncoder(f).Encode(&token)
			f.Close()
		}
	}

	return &token, nil
}
