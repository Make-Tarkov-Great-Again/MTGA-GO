package database

import (
	"fmt"
	"log"
	"path/filepath"

	"MT-GO/tools"
)

func GetAccountByUID(uid string) *Account {
	profile := GetProfileByUID(uid)
	if profile.Account != nil {
		return profile.Account
	}

	fmt.Println("Profile with UID ", uid, " does not have an account, how the fuck did you get here????!?!?!?!?!?")
	return nil
}

func (a *Account) SaveAccount() {
	accountFilePath := filepath.Join(profilesPath, a.UID, "account.json")

	err := tools.WriteToFile(accountFilePath, a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Account saved")
}

// #region Account structs

type Account struct {
	AID                 int           `json:"aid"`
	UID                 string        `json:"uid"`
	Username            string        `json:"username"`
	Password            string        `json:"password"`
	Wipe                bool          `json:"wipe"`
	Edition             string        `json:"edition"`
	Friends             Friends       `json:"friends"`
	Matching            Matching      `json:"Matching"`
	FriendRequestInbox  []interface{} `json:"friendRequestInbox"`
	FriendRequestOutbox []interface{} `json:"friendRequestOutbox"`
	TarkovPath          string        `json:"tarkovPath"`
	Lang                string        `json:"lang"`
}
type Friends struct {
	Friends      []FriendRequest `json:"Friends"`
	Ignore       []string        `json:"Ignore"`
	InIgnoreList []string        `json:"InIgnoreList"`
}
type Matching struct {
	LookingForGroup bool `json:"LookingForGroup"`
}

type FriendRequest struct {
	ID      string               `json:"_id"`
	From    string               `json:"from"`
	To      string               `json:"to"`
	Date    int32                `json:"date"`
	Profile FriendRequestProfile `json:"profile"`
}

type FriendRequestProfile struct {
	ID   int32
	Info struct {
		Nickname       string         `json:"Nickname"`
		Side           string         `json:"Side"`
		Level          int8           `json:"Level"`
		MemberCategory MemberCategory `json:"MemberCategory"`
	}
}

// #endregion
