package handlers

type AccountStruct struct {
	Id                  string
	Email               string
	Password            string
	Wipe                bool
	Edition             string
	Friends             FriendsStruct
	Matching            MatchingStruct
	FriendRequestInbox  []map[int]interface{}
	FriendRequestOutbox []map[int]interface{}
}

type FriendsStruct struct {
	Friends      []map[int]interface{}
	Ignore       []map[int]interface{}
	InIgnoreList []map[int]interface{}
}

type MatchingStruct struct {
	LookingForGroup bool
}

func AccountRegister() {
	//sessionID := tools.GenerateMongoId()
	//maxTime := 3600 * 24 // 1 day
	//c.SetCookie("PHPSESSID", sessionID, maxTime, "/", "", true, false)
}

func AccountCreate() {}
