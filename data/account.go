package data

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

	return nil, fmt.Errorf(accountNotExist, uid)
}

func (a *Account) SaveAccount() error {
	accountFilePath := filepath.Join(profilesPath, a.UID, "account.json")

	if err := tools.WriteToFile(accountFilePath, a); err != nil {
		return fmt.Errorf(accountNotSaved, a.UID, err)
	}
	log.Println("Account saved")
	return nil
}

const (
	accountNotSaved string = "Account for %s was not saved: %s"
	accountNotExist string = "Account for %s does not exist"
)

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
