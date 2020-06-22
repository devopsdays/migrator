package sponsors

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "sponsor.md.tmpl",
		FileModTime: time.Unix(1592654846, 0),

		Content: string("+++\nname = \"{{ .Name }}\"\nwebsite = \"{{ .Website }}\"\ntwitter = \"{{ .Twitter }}\"\n+++"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1592654846, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "sponsor.md.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`sponsor-templates`, &embedded.EmbeddedBox{
		Name: `sponsor-templates`,
		Time: time.Unix(1592654846, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"sponsor.md.tmpl": file2,
		},
	})
}
