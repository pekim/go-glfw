package tagfile

import (
	_ "embed"
	"encoding/xml"
	"net/url"
)

//go:embed tagfile.xml
var tagfileContent []byte

const baseURL = "https://www.glfw.org/docs/latest"

type DocAnchor struct {
	File   string `xml:"file,attr"`
	Title  string `xml:"title,attr"`
	Anchor string `xml:",chardata"`
}

type Member struct {
	Kind string `xml:"kind,attr"`

	Type       string `xml:"type"`
	Name       string `xml:"name"`
	Anchorfile string `xml:"anchorfile"`
	Anchor     string `xml:"anchor"`
	Arglist    string `xml:"arglist"`
}

func (member Member) URL() string {
	url, err := url.JoinPath(baseURL, member.Anchorfile)
	if err != nil {
		panic(err)
	}
	return url + "#" + member.Anchor
}

type Compound struct {
	Kind string `xml:"kind,attr"`

	Name       string      `xml:"name"`
	Title      string      `xml:"title"`
	Path       string      `xml:"path"`
	Filename   string      `xml:"filename"`
	DocAnchors []DocAnchor `xml:"docanchor"`
	Members    []Member    `xml:"member"`
}

func (compound Compound) findMember(name string) (*Member, bool) {
	for _, member := range compound.Members {
		if member.Name == name {
			return &member, true
		}
	}

	return nil, false
}

type Tagfile struct {
	Compounds []Compound `xml:"compound"`
}

func (tagfile Tagfile) findCompound(kind string, name string) (*Compound, bool) {
	for _, compound := range tagfile.Compounds {
		if compound.Kind == kind && compound.Name == name {
			return &compound, true
		}
	}

	return nil, false
}

type URL struct {
	Title string
	URL   string
}

func EntityURLs() map[string]string {
	var tagfile Tagfile
	err := xml.Unmarshal(tagfileContent, &tagfile)
	if err != nil {
		panic(err) // should not be possible
	}

	compound, foundCompound := tagfile.findCompound("file", "glfw3.h")
	if !foundCompound {
		panic("failed to find compound for glfw3.h") // should not be possible
	}

	urlMap := make(map[string]string)
	for _, member := range compound.Members {
		urlMap[member.Name] = member.URL()
	}

	return urlMap
}

func PageSectionURLs() map[string]URL {
	var tagfile Tagfile
	err := xml.Unmarshal(tagfileContent, &tagfile)
	if err != nil {
		panic(err) // should not be possible
	}

	urlMap := make(map[string]URL)

	for _, compound := range tagfile.Compounds {
		if compound.Kind != "page" {
			continue
		}

		for _, docanchor := range compound.DocAnchors {
			url, err := url.JoinPath(baseURL, compound.Filename)
			if err != nil {
				panic(err)
			}

			title := docanchor.Title
			if title == "" {
				title = docanchor.Anchor
			}
			url += "#" + docanchor.Anchor

			urlMap[docanchor.Anchor] = URL{
				Title: title,
				URL:   url,
			}
		}
	}

	return urlMap
}

func GroupURLs() map[string]URL {
	var tagfile Tagfile
	err := xml.Unmarshal(tagfileContent, &tagfile)
	if err != nil {
		panic(err) // should not be possible
	}

	urlMap := make(map[string]URL)

	for _, compound := range tagfile.Compounds {
		if compound.Kind != "group" {
			continue
		}

		url, err := url.JoinPath(baseURL, compound.Filename)
		if err != nil {
			panic(err)
		}
		urlMap[compound.Name] = URL{
			Title: compound.Title,
			URL:   url,
		}
	}

	return urlMap
}
