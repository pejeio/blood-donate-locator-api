package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pejeio/blood-donate-locator-api/internal/utils"
)

type Client struct {
	GC           *gocloak.GoCloak
	ClientID     string
	ClientSecret string
	Realm        string
}

func NewClient(baseURL, clientID, clientSecret, realm string) *Client {
	return &Client{
		GC:           gocloak.NewClient(baseURL),
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Realm:        realm,
	}
}

func (c *Client) CheckScopesAllowedOnResource(scopes []string, resource string, accessToken string) bool {
	gt := "urn:ietf:params:oauth:grant-type:uma-ticket"
	rm := "permissions"
	ctx := context.Background()

	res, err := c.GC.GetRequestingPartyPermissions(ctx, accessToken, c.Realm, gocloak.RequestingPartyTokenOptions{
		GrantType:    &gt,
		Audience:     &c.ClientID,
		ResponseMode: &rm,
	})

	if err != nil {
		return false
	}

	for _, perm := range *res {
		if *perm.ResourceName == resource {
			if utils.ContainsAllString(*perm.Scopes, scopes) {
				return true
			}
		}
	}
	return false
}
