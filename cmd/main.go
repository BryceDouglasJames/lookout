package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	browser "github.com/brycedouglasjames/lookout"
	"golang.org/x/net/html"
)

func pageRequest(url string) (string, error) {
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

func getDesiredTag(doc *html.Node, elm string) (*html.Node, error) {
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

	return nil, errors.New("error tryin gto find header tags")
}

func format(node *html.Node) string {
	var buf bytes.Buffer
	write := io.Writer(&buf)
	html.Render(write, node)
	return buf.String()
}

func main() {
	request, err := pageRequest("https://amazon.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	document, _ := html.Parse(strings.NewReader(request))
	Node, err := getDesiredTag(document, "li")
	if err != nil {
		fmt.Println(err)
		return
	}

	head := format(Node)
	fmt.Println(head)

	test := browser.Browser_test()
	fmt.Println(test)

}
