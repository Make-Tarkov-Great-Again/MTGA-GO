package mod

import "net/http"

func Routes(mux http.ServeMux) {
	mux.HandleFunc("/howdoyado", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://youtu.be/xvFZjo5PgG0", http.StatusSeeOther)
	})
}
