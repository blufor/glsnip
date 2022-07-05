package snippets

import (
	"log"

	"github.com/blufor/glsnip/sources"
	"github.com/spf13/viper"
	gitlab "github.com/xanzy/go-gitlab"
)

// Reads the configured sources and presents them as the API
// request types.
//
func gatherData(src []sources.Source) (snips []*gitlab.SnippetFile) {
	for _, f := range src {
		path := f.Path()

		content, err := f.Content()
		if err != nil {
			log.Printf("Failed to read content from %s: %v", path, err)
			continue
		}

		err = f.Close()
		if err != nil {
			log.Printf("Failed to close the file %s: %v", path, err)
			continue
		}

		snips = append(snips, &gitlab.SnippetFile{
			FilePath: &path,
			Content:  &content,
		})
	}
	return
}

// Translates CLI options to the `gitlab.VisibilityValue`
// interface types.
//
func visibility(c *viper.Viper) gitlab.VisibilityValue {
	switch {
	case c.GetBool("internal"):
		return gitlab.InternalVisibility
	case c.GetBool("public"):
		return gitlab.PublicVisibility
	default:
		return gitlab.PrivateVisibility
	}
}
