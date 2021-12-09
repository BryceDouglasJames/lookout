package web_interface

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPageParse(t *testing.T) {
	resp, err := http.Get("https://Wikipedia.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc, _ := html.Parse(strings.NewReader(string(b)))

	Parse_Href_Tags(16, doc, "WIKI")
}

func TestInstDelete(t *testing.T) {
	Add_User_To_Link_Pool(16)
	Create_Search_Instance_Entry(16, "Tester")
	get_user_links(16)
	add_link(16, "Tester", "LINK", "TAG_TYPE")
	delete_inst_entry(16, "Tester")
	if len(get_user_links(16).link_map) != 0 {
		t.Errorf("The instance link queue should be empty but is not.")
	}
}

func TestLinkDelete(t *testing.T) {
	Add_User_To_Link_Pool(16)
	Create_Search_Instance_Entry(16, "Tester")
	get_user_links(16)
	add_link(16, "Tester", "LINK_1", "TAG_TYPE_1")
	add_link(16, "Tester", "LINK_2", "TAG_TYPE_2")
	delete_link(16, "Tester", "LINK_2")
	if !(len(get_user_links(16).link_map) == 1 && get_user_links(16).link_map["Tester"][0].link_literal == "LINK_1") {
		t.Errorf("Something went wrong. Here is the link list: %+v", get_user_links(16).link_map)
	}
}

func TestLinkHit(t *testing.T) {
	Add_User_To_Link_Pool(16)
	Create_Search_Instance_Entry(16, "Tester")
	get_user_links(16)
	add_link(16, "Tester", "LINK_1", "TAG_TYPE_1")
	increase_frequency(16, "Tester", "LINK_1")
	if get_user_links(16).link_map["Tester"][0].hit_count != 1 {
		t.Error("Error increasing hit count of link")
	}
}
