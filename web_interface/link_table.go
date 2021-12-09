package web_interface

var (
	user_links_pool = make(map[int16]*user_links)
	//err  error HANDLE THESE ONCE MORE INTEGRATED
)

type link struct {
	link_literal string
	link_type    string
	instance_id  string //index by search keyword
	hit_count    int
}

type user_links struct {
	user_id  int16
	link_map map[string][]*link
}

func Add_User_To_Link_Pool(id int16) {
	new_user := &user_links{
		user_id:  id,
		link_map: make(map[string][]*link),
	}
	user_links_pool[id] = new_user
}

func Create_Search_Instance_Entry(id int16, search_key string) {
	links := get_user_links(id)
	links.link_map[search_key] = make([]*link, 0)
}

func add_link(id int16, inst string, l_string string, l_type string) {
	link := &link{
		link_literal: l_string,
		link_type:    l_type,
		instance_id:  inst,
		hit_count:    0,
	}
	user := get_user_links(id)
	user.link_map[inst] = append(user.link_map[inst], link)
}

func increase_frequency(id int16, inst string, literal string) {
	user := get_user_links(id)
	links := user.link_map[inst]
	for _, val := range links {
		if val.link_literal == literal {
			val.hit_count += 1
			break
		}
	}
}

func delete_inst_entry(id int16, inst string) {
	user := get_user_links(id)
	delete(user.link_map, inst)
}

//TODO go into instance search and delete link by literal
func delete_link(id int16, inst string, literal string) {
	user := get_user_links(id)
	links := user.link_map[inst]
	for i, val := range links {
		if val.link_literal == literal {
			user.link_map[inst] = fix_link_slice(links, i)
			break
		}
	}
}

func get_user_links(id int16) *user_links {
	return user_links_pool[id]
}

func fix_link_slice(slice []*link, index int) []*link {
	piece := append(slice[:index], slice[index+1:]...)
	return piece
}

/*
func convert_string(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("ERROR: cannot convert this key")
	}
	return value
}
*/
