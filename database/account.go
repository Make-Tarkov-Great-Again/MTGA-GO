package database

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
)

func GetAccountByID(uid string) (*Account, error) {
	profile, err := GetProfileByUID(uid)
	if err != nil {
		return nil, err
	}

	if profile.Account != nil {
		return profile.Account, nil
	}

	return nil, fmt.Errorf("Account for ", uid, " does not exist")
}

func (a *Account) SaveAccount() {
	accountFilePath := filepath.Join(profilesPath, a.UID, "account.json")

	err := tools.WriteToFile(accountFilePath, a)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Account saved")
}

// #region Account structs

type Account struct {
	AID        int    `json:"aid"`
	UID        string `json:"uid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Wipe       bool   `json:"wipe"`
	Edition    string `json:"edition"`
	TarkovPath string `json:"tarkovPath"`
	Lang       string `json:"lang"`
}

// #endregion
