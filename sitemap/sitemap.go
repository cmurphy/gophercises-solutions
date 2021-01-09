package sitemap

import (
	"encoding/xml"
	"fmt"
	"github.com/cmurphy/gophercises/link"
	"net/http"
	neturl "net/url"
	"strings"
)

type SiteMap struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Url     []Url
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type site struct {
	url      string
	domain   string
	visited  map[string]bool
	maxDepth int
}

func (s *site) formatURL(url string) (string, error) {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	if parsedURL.Scheme == "" {
		return s.url + strings.TrimPrefix(url, "/"), nil
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("not an HTTP URL")
	}
	return url, nil
}

func (s *site) isInternal(url string) bool {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		//fmt.Printf("invalid URL %s\n", url)
		return false
	}
	return parsedURL.Host == s.domain
}

func isAnchorLink(url string) bool {
	if strings.HasPrefix(url, "#") {
		return true
	}
	return false
}

func (s *site) getPageLinks(url string) ([]string, error) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	s.visited[url] = true
	links, err := link.Parse(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	result := make([]string, 0)
	for _, l := range links {
		if isAnchorLink(l.Href) {
			continue
		}
		url, err = s.formatURL(l.Href)
		if err != nil {
			//fmt.Printf("skipping invalid link %s\n", url)
			continue
		}
		if !s.isInternal(url) {
			//fmt.Printf("skipping external link %s\n", url)
			continue
		}
		if _, ok := s.visited[url]; ok {
			//fmt.Printf("skipping visited link %s\n", url)
			continue
		}
		result = append(result, url)
	}
	return result, nil
}

func (s *site) followLink(url string, depth int) ([]string, error) {
	//fmt.Printf("GET %s\n", url)
	if s.maxDepth >= 0 && depth > s.maxDepth {
		//fmt.Println("exceeded max depth")
		return nil, nil
	}
	result, err := s.getPageLinks(url)
	if err != nil {
		return nil, err
	}
	for _, l := range result {
		childLinks, err := s.followLink(l, depth+1)
		if err != nil {
			return nil, err
		}
		result = append(result, childLinks...)
	}
	return result, nil
}

func newSite(url string, maxDepth int) (*site, error) {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	parsedURL, err := neturl.Parse(url)
	s := site{}
	if err != nil {
		return &s, err
	}
	s.url = url
	s.domain = parsedURL.Host
	s.visited = make(map[string]bool)
	s.maxDepth = maxDepth
	return &s, nil
}

func NewSiteMap(url string, maxDepth int) (*SiteMap, error) {
	s, err := newSite(url, maxDepth)
	if err != nil {
		return nil, err
	}
	links, err := s.followLink(url, 0)
	if err != nil {
		return nil, err
	}
	sitemap := SiteMap{}
	for _, l := range links {
		sitemap.Url = append(sitemap.Url, Url{Loc: l})
	}
	return &sitemap, nil
}

func (s SiteMap) String() string {
	out, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return ""
	}
	return xml.Header + string(out)
}
