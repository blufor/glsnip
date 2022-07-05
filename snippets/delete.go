package snippets

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func DeleteSnippet(cfg *viper.Viper) *Result {
	client := setupClient(cfg)

	resp, err := client.Snippets.DeleteSnippet(int(cfg.GetUint("delete")))
	if err != nil {
		log.Fatalf("Failed to delete snippet: %v", err)
	}

	return &Result{
		Code:    uint(resp.StatusCode),
		Error:   err,
		ID:      int(cfg.GetUint("delete")),
		Message: fmt.Sprintf("Snippet ID %d deleted", cfg.GetUint("delete")),
	}
}
