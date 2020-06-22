package events

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "_index.md.tmpl",
		FileModTime: time.Unix(1592787414, 0),

		Content: string("+++\ndate = \"2020-06-11T02:11:48-05:00\"\ndescription = \"{{ .EventDescription }}\"\ntitle = \"devopsdays {{ .City }} {{ .Year }}\"\ntype = \"new-event\"\ncity = \"{{ .City }}\"\nyear = \"{{ .Year }}\"\n#event_group = \"Down Under\"\nevent_twitter = \"{{ .EventTwitter }}\"\nstartdate = \"{{ .StartDate }}\"\nenddate = \"{{ .EndDate }}\"\nmasthead_background = \"{{ .MastheadBackground }}\"\nsponsors_accepted = \"{{ .SponsorsAccepted }}\"\nsponsor_levels = [\n    {{- range .SponsorLevels }}\n    { id = \"{{ .id }}\", label = \"{{ .label }}\"{{ with .max}}, max = {{.}}{{end}} },\n    {{- end }}\n]\nsponsors = [\n    {{- range .Sponsors }}\n    { name = \"{{ .id}}\", level = \"{{ .level }}\"{{ with .url}}, url = \"{{.}}\"{{end}} },\n    {{- end }}\n]\nnavigation = [\n    {{- range .NavElements }}\n    { name = \"{{ .name }}\"{{ with .url}}, url = \"{{.}}\"{{end}} },\n    {{- end }}\n]\n+++\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1592664276, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "_index.md.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1592664276, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"_index.md.tmpl": file2,
		},
	})
}
