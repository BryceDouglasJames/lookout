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

func Parse_Href_Tags(id int16, doc *html.Node, search_key string) error {
	var crawl func(node *html.Node)

	//grab user link node from pool
	user, found_user := user_links_pool[id]
	if !found_user {
		Add_User_To_Link_Pool(id)
		user = user_links_pool[id]
	}

	//grab the instance link slice
	//Create_Search_Instance_Entry(id, search_key)
	link_slice, found_links := user.link_map[search_key]
	if !found_links {
		Create_Search_Instance_Entry(id, search_key)
		link_slice = user.link_map[search_key]
	}

	//crawl through the doc and add <a href=""> values
	crawl = func(node *html.Node) {
		for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
			if childNode.Data == "a" {
				for _, attr := range childNode.Attr {
					if attr.Key == "href" {
						temp_link := &link{
							instance_id:  search_key,
							link_literal: attr.Val,
							link_type:    childNode.Data,
							hit_count:    1,
						}
						link_slice = append(link_slice, temp_link)
					}
				}
			}
			crawl(childNode)
		}
	}
	crawl(doc)
	user.link_map[search_key] = link_slice

	//uncomment to see link objects
	/*for _, s := range user.link_map[search_key] {
		fmt.Println(s)
	}*/

	return errors.New("")
}

func Link_Format(node *html.Node) string {
	var buf bytes.Buffer
	write := io.Writer(&buf)
	html.Render(write, node)
	return buf.String()
}
