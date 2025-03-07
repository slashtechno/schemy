package airtable

import (
	"errors"

	"github.com/charmbracelet/log"
	"resty.dev/v3"
)

type client struct {
	// ApiKey is exported so assuming NewClient has been called, the calling package can read what API key is being used
	// ApiKey  string
	baseUrl string
	client  *resty.Client
}

type errorResponse struct {
	Error struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *client) GetBaseSchema(baseId string) (*Schema, error) {
	defer c.client.Close()

	resp, err := c.client.R().
		SetResult(&Schema{}).
		Get("meta/bases/" + baseId + "/tables")
	if err != nil {
		return nil, err
	} else if resp.StatusCode() != 200 {
		log.Error("error getting base schema", "error", resp.Error().(*errorResponse))
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}
	return resp.Result().(*Schema), err
}

func NewClient(apiKey string) *client {
	return &client{
		// ApiKey: apiKey,
		baseUrl: "https://api.airtable.com/v0/",
		client: resty.New().
			SetBaseURL("https://api.airtable.com/v0/").
			SetAuthToken(apiKey).
			SetError(&errorResponse{}),
	}
}

type Schema struct {
	Tables []struct {
		Description string `json:"description,omitempty"`
		Fields      []struct {
			Description string `json:"description,omitempty"`
			// ID          string `json:"id"`
			Name    string                 `json:"name"`
			Type    string                 `json:"type"`
			Options map[string]interface{} `json:"options,omitempty"`
		} `json:"fields"`
		// ID   string `json:"id"`
		Name string `json:"name"`
		// PrimaryFieldID string `json:"primaryFieldId"`
		// Views []struct {
		// 	// ID   string `json:"id"`
		// 	Name string `json:"name"`
		// 	Type string `json:"type"`
		// } `json:"views"`
	} `json:"tables"`
}
