// Package structs contains all structs used throughout the database
package structs

type Usernames map[string]string

type Account struct {
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
