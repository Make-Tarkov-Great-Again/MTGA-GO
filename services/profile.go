package services

import (
	"MT-GO/database"
)

func IsNicknameAvailable(nickname string, profiles map[string]*database.Profile) bool {
	for _, profile := range profiles {
		if profile.Character == nil || profile.Character.Info.Nickname == "" {
			continue
		}

		Nickname := profile.Character.Info.Nickname
		if Nickname == nickname {
			return false
		}

	}
	return true
}
