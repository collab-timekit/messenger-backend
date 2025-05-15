package keycloak

import (
	"encoding/json"
	"fmt"
	"io"
	"messenger/internal/infra/config"
	"net/http"
	"net/url"
	"strings"
)

// KeycloakClient is a client for interacting with Keycloak's REST API.
type KeycloakClient struct {
	ClientID     string
	ClientSecret string
	Realm        string
	BaseURL      string
}

// User represents a Keycloak user with an ID and email address.
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Username      string            `json:"username"`
	FirstName     string            `json:"firstName"`
	LastName      string            `json:"lastName"`
}

// NewKeycloakClient creates a new instance of KeycloakClient using the provided configuration.
func NewKeycloakClient(cfg config.KeycloakConfig) *KeycloakClient {
	return &KeycloakClient{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Realm:        cfg.Realm,
		BaseURL:      cfg.BaseURL,
	}
}

func (kc *KeycloakClient) getAccessToken() (string, error) {
	data := url.Values{}
	data.Set("client_id", kc.ClientID)
	data.Set("client_secret", kc.ClientSecret)
	data.Set("grant_type", "client_credentials")

	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", kc.BaseURL, kc.Realm)

	// 🔵 LOG REQUEST INFO
	fmt.Println("➡️ Requesting access token from Keycloak:")
	fmt.Printf("URL: %s\n", tokenURL)
	fmt.Printf("Payload: %s\n", data.Encode())

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // loguj body błędu
		return "", fmt.Errorf("failed to get access token: %s\nBody: %s", resp.Status, string(body))
	}

	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	// 🔵 LOG SUCCESS
	fmt.Println("✅ Access token received successfully.")

	return res.AccessToken, nil
}


// GetUsers retrieves a list of users from Keycloak based on the provided filter.
func (kc *KeycloakClient) GetUsers(filter string) ([]User, error) {
	token, err := kc.getAccessToken()
	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s/admin/realms/%s/users?search=%s", kc.BaseURL, kc.Realm, url.QueryEscape(filter))
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// 🔵 LOG REQUEST INFO
	fmt.Println("➡️ Sending request to Keycloak:")
	fmt.Printf("URL: %s\n", fullURL)
	fmt.Printf("Authorization: Bearer %s...\n", token[:10]) // log only start of token for safety
	fmt.Printf("Headers: %+v\n", req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch users: %s", resp.Status)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}