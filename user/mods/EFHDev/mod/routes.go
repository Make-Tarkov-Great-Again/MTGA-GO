package mod

import "net/http"

var Routes = map[string]http.HandlerFunc{
	"/howdoyado": func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://youtu.be/xvFZjo5PgG0", http.StatusSeeOther)
	},
}
