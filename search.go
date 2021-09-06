package gogpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchRequest represents a request structure for search API
type SearchRequest struct {
	Documents      []*string `json:"documents,omitempty"`
	File           *string   `json:"file,omitempty"`
	MaxRerank      *int      `json:"max_rerank,omitempty"`
	Query          string    `json:"query"`
	ReturnMetadata *bool     `json:"return_metadata,omitempty"`
	UserID         *string   `json:"user,omitempty"`
}

// SearchResult represents single result from search API
type SearchResult struct {
	Document int     `json:"document"`
	Metadata *string `json:"metadata,omitempty"`
	Object   string  `json:"object"`
	Score    float32 `json:"score"`
	Text     string  `json:"text"`
}

// SearchResponse represents a response structure for search API
type SearchResponse struct {
	SearchResults []SearchResult `json:"data"`
}

// Search — perform a semantic search api call over a list of documents.
func (c *Client) Search(ctx context.Context, engineID string, request SearchRequest) (response SearchResponse, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	urlSuffix := fmt.Sprintf("/engines/%s/search", engineID)
	req, err := http.NewRequest("POST", c.fullURL(urlSuffix), bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &response)
	return
}
