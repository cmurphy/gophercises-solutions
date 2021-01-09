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

type tree struct {
	visited  map[string]struct{}
	maxDepth int
}

func formatURL(url string, baseURL *neturl.URL) (string, error) {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	if parsedURL.Scheme == "" {
		return baseURL.String() + strings.TrimPrefix(url, "/"), nil
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("not an HTTP URL")
	}
	return url, nil
}

func isInternal(url string, baseURL *neturl.URL) bool {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		//fmt.Printf("invalid URL %s\n", url)
		return false
	}
	return parsedURL.Host == baseURL.Host
}

func isAnchorLink(url string) bool {
	if strings.HasPrefix(url, "#") {
		return true
	}
	return false
}

func (t *tree) getPageLinks(url string) ([]string, error) {
	r, err := http.Get(url)
	baseURL := &neturl.URL{
		Scheme: r.Request.URL.Scheme,
		Host:   r.Request.URL.Host,
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	t.visited[url] = struct{}{}
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
		url, err = formatURL(l.Href, baseURL)
		if err != nil {
			//fmt.Printf("skipping invalid link %s\n", url)
			continue
		}
		if !isInternal(url, baseURL) {
			//fmt.Printf("skipping external link %s\n", url)
			continue
		}
		if _, ok := t.visited[url]; ok {
			//fmt.Printf("skipping visited link %s\n", url)
			continue
		}
		result = append(result, url)
	}
	return result, nil
}

func (t *tree) followLink(url string, depth int) ([]string, error) {
	//fmt.Printf("GET %s\n", url)
	if t.maxDepth >= 0 && depth > t.maxDepth {
		//fmt.Println("exceeded max depth")
		return nil, nil
	}
	result, err := t.getPageLinks(url)
	if err != nil {
		return nil, err
	}
	for _, l := range result {
		childLinks, err := t.followLink(l, depth+1)
		if err != nil {
			return nil, err
		}
		result = append(result, childLinks...)
	}
	return result, nil
}

func NewSiteMap(url string, maxDepth int) (*SiteMap, error) {
	t := &tree{
		maxDepth: maxDepth,
		visited:  make(map[string]struct{}),
	}
	links, err := t.followLink(url, 0)
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
