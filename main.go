package main

import (
	"context"
	"go-templ/views"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	users := []string{"robson", "andre", "maria"}

	http.Handle("/", templ.Handler(views.Layout(views.Home())))

	http.Handle("/users", templ.Handler(views.Layout(views.Users(users))))

	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("id")
		ctx := context.WithValue(r.Context(), "user", user)
		views.Layout(views.User()).Render(ctx, w)
	})

	http.ListenAndServe(":3000", nil)
}
