package webapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type account struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func handleAccountSave() http.HandlerFunc {
	type request struct {
		AccountName        string `json:"name"`
		AccountOldPassword string `json:"oldpassword"`
		AccountNewPassword string `json:"newpassword"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse body. err=%v", err)
			respond(w, r, nil, http.StatusBadRequest)
			return
		}

		service := &AccountServiceImpl{}
		a, err2 := service.GetAccountById(req.AccountName)
		if err2 != nil {
			log.Printf("Cannot get account from DB. err=%v", err2)
			respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		log.Printf(a.String())

		a.password = req.AccountNewPassword
		err = service.UpdateAccount(a)

		if err != nil {
			log.Printf("Cannot save account in DB. err=%v", err)
			respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapAccountToJson(a)
		respond(w, r, resp, http.StatusOK)
	}
}

func handleAccountDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := vars["id"]
		log.Printf("handleAccountDetail %v", id)

		service := &AccountServiceImpl{}
		m, err := service.GetAccountById(id)
		if err != nil {
			log.Printf("Cannot load. err=%v", err)
			respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapAccountToJson(m)
		respond(w, r, resp, http.StatusOK)
	}
}

func mapAccountToJson(a *Account) account {
	return account{
		Id:    a.Id,
		Email: a.Email,
	}
}
