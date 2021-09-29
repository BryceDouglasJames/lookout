package web_interface

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func PageRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return ``, errors.New("cannot request url")
	}

	doc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ``, errors.New("cannot parse response")
	}

	return string(doc), nil
}

func GetDesiredTag(doc *html.Node, elm string) (*html.Node, error) {
	var tag *html.Node
	var crawl func(node *html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == elm {
			tag = node
			return
		}
		for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
			crawl(childNode)
		}
	}

	crawl(doc)
	if tag != nil {
		return tag, nil
	}

	return nil, errors.New("error tryin to find header tags")
}

func Format(node *html.Node) string {
	var buf bytes.Buffer
	write := io.Writer(&buf)
	html.Render(write, node)
	return buf.String()
}
