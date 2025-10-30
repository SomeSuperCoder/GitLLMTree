package llm

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/SomeSuperCoder/GitLLMTree/config"
	openai "github.com/sashabaranov/go-openai"
)

func formProxyURLString(ac *config.AppConfig) string {
	return fmt.Sprintf("http://%s:%s@%s:%s", ac.ProxyUsername, ac.ProxyPassword, ac.ProxyIP, ac.ProxyPort)
}

func formProxyURL(ac *config.AppConfig) (*url.URL, error) {
	urlString := formProxyURLString(ac)
	return url.Parse(urlString)
}

func NewClient(appConfig *config.AppConfig) (*openai.Client, error) {
	proxyURL, err := formProxyURL(appConfig)
	if err != nil {
		return nil, err
	}

	cutsomHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	config := openai.DefaultConfig(appConfig.OpenAIAPIKey)
	config.HTTPClient = cutsomHTTPClient

	client := openai.NewClientWithConfig(config)

	return client, nil
}
