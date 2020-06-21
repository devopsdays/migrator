package organizers

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.md.tmpl",
		FileModTime: time.Unix(1592780053, 0),

		Content: string("+++\nheadless = true\n+++"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "organizer.md.tmpl",
		FileModTime: time.Unix(1592782597, 0),

		Content: string("+++\nTitle = \"{{ .Title }}\"\nTwitter = \"{{ .Twitter }}\"\nLinkedIn = \"{{ .LinkedIn }}\"\nGitHub = \"{{ .GitHub }}\"\nWebsite = \"{{ .Website }}\"\nEmployer = \"{{ .Employer }}\"\nRole = \"{{ .Role }}\"\nImage = \"{{ .Image }}\"\nGitLab = \"{{ .GitLab }}\"\nFacebook = \"{{ .Facebook }}\"\n+++\n{{ .Bio }}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1592782411, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.md.tmpl"
			file3, // "organizer.md.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`organizer-templates`, &embedded.EmbeddedBox{
		Name: `organizer-templates`,
		Time: time.Unix(1592782411, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.md.tmpl":     file2,
			"organizer.md.tmpl": file3,
		},
	})
}
