package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/blufor/glsnip/snippets"
	"github.com/blufor/glsnip/sources"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "glsnip",
	Short: "GitLab Snippet CLI util",
	Args:  cobra.RangeArgs(0, 10),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
	},
	Run: run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Config initialization
//
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().UintP("update", "U", 0, "Update snippet by its ID")
	rootCmd.PersistentFlags().UintP("delete", "D", 0, "Delete snippet by its ID")

	rootCmd.PersistentFlags().String("token", "", "GitLab authentication token (use GITLAB_TOKEN in ENV)")
	rootCmd.PersistentFlags().String("url", "", "GitLab URL (use GITLAB_URL in ENV)")
	rootCmd.PersistentFlags().StringP("title", "t", "", "title for the snippet")
	rootCmd.PersistentFlags().StringP("description", "d", "", "description of the snippet")
	rootCmd.PersistentFlags().BoolP("internal", "i", false, "internal visibility")
	rootCmd.PersistentFlags().BoolP("public", "p", false, "public visibility")

	rootCmd.MarkFlagsMutuallyExclusive("internal", "public")
	rootCmd.MarkFlagsMutuallyExclusive("update", "delete")
}

// ENV variables binding initialization
//
func initConfig() {
	viper.SetEnvPrefix("gitlab")
	viper.AutomaticEnv()
}

// Primary run function for `cobra.Command` root
//
func run(cmd *cobra.Command, args []string) {
	var snips []sources.Source
	var snippet *snippets.Result

	if len(args) == 0 {
		snips = append(snips, sources.NewPipe())
	} else {
		for _, f := range args {
			s, err := sources.NewFile(f)
			if err != nil {
				log.Printf("Failed to read file: %v", err)
				continue
			}
			snips = append(snips, s)
		}
	}

	switch {
	case viper.GetUint("delete") > 0:
		snippet = snippets.DeleteSnippet(viper.GetViper())
	case viper.GetUint("update") > 0:
		snippet = snippets.UpdateSnippet(snips, viper.GetViper())
	default:
		snippet = snippets.CreateSnippet(snips, viper.GetViper())
	}

	log.Println(snippet.Message)

	if snippet.URL != "" {
		fmt.Println(snippet.URL)
	}
}
