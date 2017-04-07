// Represents the standard structure of a page

package view

type View struct {
	Name string
	Page *PageView
}

type PageView struct {
	MainHeader string
	Title      string
}

func Views() (map[string]*View, error) {
	m := make(map[string]*View)

	m["User"] = &View{Name: "User",
		Page: &PageView{MainHeader: "User", Title: "User Dude"}}

	m["Home"] = &View{Name: "Home",
		Page: &PageView{MainHeader: "Welcome!", Title: "Home"}}

	return m, nil
}
