package database

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
)

const friendsNotExist string = "Friends for %s do not exist"

func GetFriendsByID(uid string) (*Friends, error) {
	profile, err := GetProfileByUID(uid)
	if err != nil {
		log.Fatalln(err)
	}

	if profile.Friends != nil {
		return profile.Friends, nil
	}
	return nil, fmt.Errorf(friendsNotExist, uid)
}

func (friends *Friends) SaveFriends(sessionID string) {
	friendsFilePath := filepath.Join(profilesPath, sessionID, "friends.json")

	err := tools.WriteToFile(friendsFilePath, friends)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Friends saved")
}
