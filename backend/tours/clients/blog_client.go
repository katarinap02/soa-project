package clients

import (
	"bytes"
	"context"
	"database-example/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BlogClient interface {
	CreateBlogPost(ctx context.Context, tour *model.Tour, authorID string) error
}

type HTTPBlogClient struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPBlogClient(baseURL string, client *http.Client) *HTTPBlogClient {
	if client == nil {
		client = &http.Client{}
	}
	return &HTTPBlogClient{
		BaseURL: baseURL,
		Client:  client,
	}
}

func (c *HTTPBlogClient) CreateBlogPost(ctx context.Context, tour *model.Tour, authorID string) error {
	payload := map[string]interface{}{
		"username":    authorID,
		"title":       tour.Name,
		"description": tour.Description,
		"date":        time.Now(),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal blog payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/blog/create-post", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("blog service request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("blog service returned status %d", resp.StatusCode)
	}

	return nil
}
