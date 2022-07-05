package snippets

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	gitlab "github.com/xanzy/go-gitlab"
)

type Result struct {
	Code    uint
	Error   error
	Message string
	ID      int
	Raw     string
	URL     string
}

func setupClient(cfg *viper.Viper) (client *gitlab.Client) {
	client, err := gitlab.NewClient(
		cfg.GetString("token"),
		gitlab.WithBaseURL(fmt.Sprintf("%s/api/v4", cfg.GetString("url"))),
	)
	if err != nil {
		log.Fatalf("Failed to initialize GitLab client: %v", err)
	}
	return
}
