package snippets

import (
	"fmt"
	"log"

	"github.com/blufor/glsnip/sources"
	"github.com/spf13/viper"
	gitlab "github.com/xanzy/go-gitlab"
)

func UpdateSnippet(data []sources.Source, cfg *viper.Viper) *Result {
	client := setupClient(cfg)
	files := gatherData(data)

	update := &gitlab.CreateSnippetOptions{
		Title:       gitlab.String(cfg.GetString("title")),
		Description: gitlab.String(cfg.GetString("description")),
		Visibility:  gitlab.Visibility(visibility(cfg)),
		Files:       &files,
	}

	s, resp, err := client.Snippets.CreateSnippet(update)
	if err != nil {
		log.Fatalf("Failed to update snippet: %v", err)
	}

	return &Result{
		Code:    uint(resp.StatusCode),
		Error:   err,
		ID:      s.ID,
		Message: fmt.Sprintf("Snippet id %d updated", s.ID),
		Raw:     s.RawURL,
		URL:     s.WebURL,
	}
}
