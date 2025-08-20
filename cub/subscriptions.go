package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
	PrettyPrint("PLANS", plans)

	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: is_authenticated,
		Data:          map[string]any{"plans": plans[:len(plans)-1]},
	}
	Render(w, "plans.html", data)
}

func (app *Bridge) SubscribeUser(w http.ResponseWriter, r *http.Request) {
	is_authenticated := false
	token, err := r.Cookie("token")
	if err == nil {
		is_authenticated = true
	}

	if !is_authenticated {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	id := token.Value
	user := User{}
	res := app.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		panic("user should not be authenticated if not in db")
	}

	plan_id, err := strconv.Atoi(r.PathValue("plan_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request")
		return
	}

	user.PlanId = plan_id
	app.DB.Save(user)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
