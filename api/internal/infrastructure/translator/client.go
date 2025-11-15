package translator

import (
	"context"
	"encoding/json"
	"errors"
	"firemap/internal/infrastructure/config"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrInvalidTranslationFormat    = errors.New("invalid translation format")
	ErrInvalidTranslationStructure = errors.New("invalid translation structure")
	ErrTranslationNotString        = errors.New("translation text is not a string")
	ErrEmptyResponse               = errors.New("empty response from translator")
)

type Translator interface {
	Translate(ctx context.Context, text, targetLang string) (string, error)
}

type Client struct {
	config *config.Config
}

func NewClient(config *config.Config) Translator {
	return &Client{
		config: config,
	}
}

func (t *Client) Translate(ctx context.Context, text string, targetLang string) (string, error) {
	baseURL := t.config.Translator.URL

	params := url.Values{}
	params.Set("client", "gtx")
	params.Set("sl", "auto")
	params.Set("tl", targetLang)
	params.Set("dt", "t")
	params.Set("q", text)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("translation failed with status: %d", resp.StatusCode)
	}

	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return t.extractTranslation(result)
}

func (t *Client) extractTranslation(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", ErrEmptyResponse
	}

	translations, ok := result[0].([]interface{})
	if !ok || len(translations) == 0 {
		return "", ErrInvalidTranslationFormat
	}

	firstTranslation, ok := translations[0].([]interface{})
	if !ok || len(firstTranslation) == 0 {
		return "", ErrInvalidTranslationStructure
	}

	translatedText, ok := firstTranslation[0].(string)
	if !ok {
		return "", ErrTranslationNotString
	}

	return translatedText, nil
}
