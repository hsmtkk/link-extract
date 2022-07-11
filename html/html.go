package html

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Parse(htmlString string) ([]string, error) {
	tokenizer, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML; %w", err)
	}
	links := []string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(tokenizer)
	return links, nil
}
