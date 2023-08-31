// Package main is a package declaration
package main

import (
	"MT-GO/database"
	"MT-GO/server"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	database.InitializeDatabase()
	startHome()
}

func startHome() {

	fmt.Println("Alright fella, what now?")
	fmt.Println("1. Register an Account")
	fmt.Println("2. Login")
	fmt.Println()
	fmt.Println("69. Exit")
	var input string
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			registerAccount()
			break
		case "2":
			login()
			break
		case "69":
			fmt.Println("Adios faggot")
			return
		default:
			fmt.Println("Invalid input, intellectually less able fella")
		}
	}
}

func registerAccount() {
	account := structs.Account{}
	profiles := database.GetProfiles()
	var input string

	fmt.Println("What is your username?")
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)
		if !validateUsername(profiles, input) {
			fmt.Println("Username taken, try again")
			continue
		}
		break
	}
	account.Username = input

	fmt.Println("What is your password?")
	fmt.Scanln(&input)
	fmt.Printf("> ")
	account.Password = input

	UID, err := tools.GenerateMongoID()
	if err != nil {
		panic(err)
	}
	account.UID = UID
	account.AID = len(profiles)

	account.Friends = structs.Friends{
		Friends:      []structs.FriendRequest{},
		Ignore:       []string{},
		InIgnoreList: []string{},
	}

	account.FriendRequestInbox = []interface{}{}
	account.FriendRequestOutbox = []interface{}{}

	profiles[UID] = &structs.Profile{}
	profiles[UID].Account = &account
	profiles[UID].Character = &structs.PlayerTemplate{}
	profiles[UID].Storage = &structs.Storage{
		Suites: []string{},
		Builds: structs.Builds{
			EquipmentBuilds: []structs.EquipmentBuild{},
			WeaponBuilds:    []interface{}{},
		},
		Insurance: []interface{}{},
		Mailbox:   []interface{}{},
	}
	profiles[UID].Dialogue = map[string]interface{}{}

	//save account
	fmt.Println()
	services.SaveProfile(profiles[UID])
	fmt.Println()

	//login
	fmt.Println("Account created, logging in...")
	fmt.Println()

	loggedIn(&account)
}

func validateUsername(profiles map[string]*structs.Profile, username string) bool {
	for _, profile := range profiles {
		if profile.Account.Username == username {
			return false
		}
	}
	return true
}

func login() {
	fmt.Println()
	var input string
	var account *structs.Account
	profiles := database.GetProfiles()
	if len(profiles) == 0 {
		fmt.Println("No profiles, redirecting to Account Register...")
		registerAccount()
	}

	for {
		fmt.Println("What is your username?")
		fmt.Printf("> ")
		fmt.Scanln(&input)

		for _, profile := range profiles {
			if profile.Account.Username == input {
				account = profile.Account
			}
		}

		if account == nil {
			fmt.Println("Invalid username, try again moron")
			continue
		}

		fmt.Println("What is your password?")
		fmt.Printf("> ")
		fmt.Scanln(&input)

		if account.Password != input {
			fmt.Println("Invalid password, try again moron")
			continue
		}

		server.SetHTTPSServer()
		fmt.Println("Logging in...")
		fmt.Println()

		loggedIn(profiles[account.UID].Account)
		break
	}

}

func loggedIn(account *structs.Account) {

	fmt.Println("Alright fella, we're at the Login Menu, what now?")
	fmt.Println()

	fmt.Println("1. Launch Tarkov")
	fmt.Println("2. Change Account Info")
	fmt.Println()
	fmt.Println("69. Exit")

	for {
		var input string

		fmt.Printf("> ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			launchTarkov(account)
			fmt.Println()
		case "2":
			fmt.Println()
			editAccountInfo(account)
		case "69":
			fmt.Println("Adios faggot")
			return
		default:
			fmt.Println("Invalid input, retard")
		}
	}
}

func editAccountInfo(account *structs.Account) {
	fmt.Println("Alright fella, what do you want to edit?")
	fmt.Println()
	fmt.Println("1. Change Escape From Tarkov executable path")
	fmt.Println()
	fmt.Println("69. Go back to Login Menu")

	for {
		var input string

		fmt.Printf("> ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			for {
				var tarkovPath string

				fmt.Println()
				fmt.Println("Set new Path to Tarkov executable")
				fmt.Printf("> ")
				fmt.Scanln(&tarkovPath)
				exePath := filepath.Join(tarkovPath, "EscapeFromTarkov.exe")
				if tools.FileExist(exePath) && exePath != account.TarkovPath {
					account.TarkovPath = exePath
					fmt.Println("Path has been set")

					services.SaveAccount(*account)
					break
				}
				fmt.Println("Invalid path, try again")
			}
			editAccountInfo(account)
		case "69":
			fmt.Println()
			loggedIn(account)
		default:
			fmt.Println("Invalid input, retard")
		}
	}
}

//tarkovPath + ' -bC5vLmcuaS5u={"email":"' + userAccount.email + '","password":"' + userAccount.password + '","toggle":true,"timestamp":0} -token=' + sessionID + ' -config={"BackendUrl":"https://' + serverConfig.ip + ':' + serverConfig.mainPort + '","Version":"live"}'

func launchTarkov(account *structs.Account) {
	if account.TarkovPath == "" || !tools.FileExist(account.TarkovPath) {
		fmt.Println("EscapeFromTarkov not found")
		fmt.Println("Input the folder/directory path to your 'EscapeFromTarkov.exe'")
		for {
			var tarkovPath string

			fmt.Printf("> ")
			fmt.Scanln(&tarkovPath)
			if !tools.FileExist(filepath.Join(tarkovPath, "BepInEx")) {
				fmt.Println("This folder doesn't contain the 'BepInEx' directory, set path to your non-live 'EscapeFromTarkov' directory")
				continue
			}

			account.TarkovPath = filepath.Join(tarkovPath, "EscapeFromTarkov.exe")
			if !tools.FileExist(account.TarkovPath) {
				fmt.Println("Invalid path, does not contain 'EscapeFromTarkov.exe', try again")
				continue
			}

			fmt.Println("Valid path to 'EscapeFromTarkov.exe' has been set")
			services.SaveAccount(*account)
			break
		}
	}

	cmdArgs := []string{
		"-bC5vLmcuaS5u={'email':'" + account.Username + "','password':'" + account.Password + "','toggle':true,'timestamp':0}",
		"-token=" + account.UID,
		"-config={'BackendUrl':'" + database.GetMainAddress() + "','Version':'live'}",
	}

	cmd := exec.Command(account.TarkovPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}
