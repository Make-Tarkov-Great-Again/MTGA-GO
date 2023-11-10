package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"
)

const (
	friendsNotExist string = "Friends for %s do not exist"
	friendsNotSaved string = "Friends for %s was not saved: %s"
)

func setFriends(path string) *Friends {
	output := new(Friends)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, &output)
	if err != nil {
		log.Println(err)
	}

	return output
}

func GetFriendsByID(uid string) (*Friends, error) {
	profile, err := GetProfileByUID(uid)
	if err != nil {
		log.Println(err)
	}

	if profile.Friends != nil {
		return profile.Friends, nil
	}
	return nil, fmt.Errorf(friendsNotExist, uid)
}

func (friends *Friends) SaveFriends(sessionID string) error {
	friendsFilePath := filepath.Join(profilesPath, sessionID, "friends.json")

	err := tools.WriteToFile(friendsFilePath, friends)
	if err != nil {
		return fmt.Errorf(friendsNotSaved, sessionID, err)
	}
	log.Println("Friends saved")
	return nil
}
