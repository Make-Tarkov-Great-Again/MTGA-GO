package services

import "MT-GO/database"

func IsNicknameAvailable(nickname string) bool {
	profiles := database.GetProfiles()

	for _, profile := range profiles {
		if profile.Character == nil {
			continue
		}

		Nickname, ok := profile.Character.Info["Nickname"].(string)
		if ok {
			if Nickname == nickname {
				return false
			}
		}

	}
	return true
}
