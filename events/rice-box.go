package events

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "_index.md.tmpl",
		FileModTime: time.Unix(1592841925, 0),

		Content: string("+++\ndate = \"2020-06-11T02:11:48-05:00\"\ndescription = \"{{ .EventDescription }}\"\ntitle = \"devopsdays {{ .City }} {{ .Year }}\"\ntype = \"new-event\"\ncity = \"{{ .City }}\"\nyear = \"{{ .Year }}\"\n{{- with .EventTwitter }}\nevent_twitter = \"{{ . }}\"\n{{- end }}\n{{- with .EventGroup }}\nevent_group = \"{{ . }}\"\n{{- end }}\n{{- with .MastheadBackground }}\nmasthead_background = \"{{ . }}\"\n{{- end }}\n{{- with .GoogleAnalytics }}\nga_tracking_id = \"{{ . }}\"\n{{- end }}\n{{- with .SpeakersVerbose }}\nspeakers_verbose = \"{{ . }}\"\n{{- end }}\n{{- with .Cancel }}\ncancel = \"{{ . }}\"\n{{- end }}\n{{- with .StartDate }}\nstartdate = \"{{ . }}\"\n{{- end }}\n{{- with .EndDate }}\nenddate = \"{{ . }}\"\n{{- end }}\n{{- with .CFPDateStart }}\ncfp_date_start = \"{{ . }}\"\n{{- end }}\n{{- with .CFPDateEnd }}\ncfp_date_end = \"{{ . }}\"\n{{- end }}\n{{- with .CFPDateAnnounce }}\ncfp_data_announce = \"{{ . }}\"\n{{- end }}\n{{- with .CFPLink }}\ncfp_link = \"{{ . }}\"\n{{- end }}\n{{- with .RegistrationDateStart }}\nregistration_date_start = \"{{ . }}\"\n{{- end }}\n{{- with .RegistrationDateEnd }}\nregistration_date_end = \"{{ . }}\"\n{{- end }}\n{{- with .RegistrationClosed }}\nregistration_closed = \"{{ . }}\"\n{{- end }}\n{{- with .RegistrationLink }}\nregistration_link = \"{{ . }}\"\n{{- end }}\n{{- with .SponsorLink }}\nsponsor_link = \"{{ . }}\"\n{{- end }}\n{{- with .SharingImage }}\nsharing_image = \"{{ . }}\"\n{{- end }}\n{{- with .Coordinates }}\ncoordinates = \"{{ . }}\"\n{{- end }}\n{{- with .Location }}\nlocation = \"{{ . }}\"\n{{- end }}\n{{- with .LocationAddress }}\nlocation_address = \"{{ . }}\"\n{{- end }}\n{{- with .OrganizerEmail }}\norganizer_email = \"{{ . }}\"\n{{- end }}\n{{- with .ProposalEmail }}\nproposal_email = \"{{ . }}\"\n{{- end }}\n{{- with .SponsorsAccepted }}\nsponsors_accepted = \"{{ . }}\"\n{{- end }}\n{{- with .SponsorLevels }}\nsponsor_levels = [\n    {{- range . }}\n    { id = \"{{ .id }}\", label = \"{{ .label }}\"{{ with .max}}, max = {{.}}{{end}} },\n    {{- end }}\n]\n{{- end }}\n{{- with .Sponsors }}\nsponsors = [\n    {{- range . }}\n    { name = \"{{ .id}}\", level = \"{{ .level }}\"{{ with .url}}, url = \"{{.}}\"{{end}} },\n    {{- end }}\n]\n{{- end }}\n{{- with .NavElements }}\nnavigation = [\n    {{- range . }}\n    { name = \"{{ .name }}\"{{ with .url}}, url = \"{{.}}\"{{end}} },\n    {{- end }}\n]\n{{- end }}\n+++\n{{ .Content }}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "content_file.md.tmpl",
		FileModTime: time.Unix(1592837657, 0),

		Content: string("+++\ndescription = \"{{ .Description }}\"\ntitle = \"{{ .Title }}\"\n{{ with .Type}}type = \"{{ . }}\"{{end}}\n{{- with .Aliases }}\naliases = [\n    {{- range . }}\n        \"{{ . }}\",\n    {{- end }}\n]\n{{- end }}\n{{- with .SharingImage }}\nsharing_image = \"{{ . }}\"\n{{- end }}\n{{- with .Speakers }}\nspeakers = [\n    {{- range . }}\n        \"{{ . }}\",\n    {{- end }}\n]\n{{- end }}\n{{- with .YouTube }}\nyoutube = \"{{ . }}\"\n{{- end }}\n{{- with .Vimeo }}\nvimeo = \"{{ . }}\"\n{{- end }}\n{{- with .Slideslive }}\nslideslive = \"{{ . }}\"\n{{- end }}\n{{- with .Speakerdeck }}\nspeakerdeck = \"{{ . }}\"\n{{- end }}\n{{- with .Slideshare }}\nslideshare = \"{{ . }}\"\n{{- end }}\n{{- with .GoogleSlides }}\ngoogleslides = \"{{ . }}\"\n{{- end }}\n{{- with .PDF }}\npdf = \"{{ . }}\"\n{{- end }}\n{{- with .Notist }}\nnotist = \"{{ . }}\"\n{{- end }}\n{{- with .Slides }}\nslides = \"{{ . }}\"\n{{- end }}\n{{- with .Website }}\nwebsite = \"{{ . }}\"\n{{- end }}\n{{- with .Twitter }}\ntwitter = \"{{ . }}\"\n{{- end }}\n{{- with .Facebook }}\nfacebook = \"{{ . }}\"\n{{- end }}\n{{- with .LinkedIn }}\nlinkedin = \"{{ . }}\"\n{{- end }}\n{{- with .GitHub }}\ngithub = \"{{ . }}\"\n{{- end }}\n{{- with .GitLab }}\ngitlab = \"{{ . }}\"\n{{- end }}\n{{- with .Image }}\nimage = \"{{ . }}\"\n{{- end }}\n{{- with .Icons }}\nicons = \"{{ . }}\"\n{{- end }}\n{{- with .LinkTitle }}\nlinktitle = \"{{ . }}\"\n{{- end }}\n+++\n{{ .Content }}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1592831256, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "_index.md.tmpl"
			file3, // "content_file.md.tmpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1592831256, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"_index.md.tmpl":       file2,
			"content_file.md.tmpl": file3,
		},
	})
}
