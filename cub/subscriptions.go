package main

import "net/http"

func (app *Bridge) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	is_authenticated := false
	_, err := r.Cookie("token")
	if err == nil {
		is_authenticated = true
	}

	if !is_authenticated {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	plans := []Plan{}
	res := app.DB.Find(&plans)
	if res.Error != nil {
		panic("could not get plans")
	}

	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: is_authenticated,
		Data:          map[string]any{"plans": plans},
	}
	Render(w, "plans.html", data)
}
