// Package main is a package declaration

package main

import (
	"MT-GO/srv"
	"MT-GO/user/mods"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"MT-GO/data"
	"MT-GO/tools"
)

func main() {
	startTime := time.Now()
	done := make(chan bool)
	//TODO: Squeeze MS where possible, investigate TraderIndex if possible

	data.SetDatabase()
	mods.Init()

	data.LoadBundleManifests()
	data.LoadCustomItems()

	go func() {
		data.SetWeaponMasteries()
		done <- true
	}()
	go func() {
		data.SetServerConfig()
		done <- true
	}()
	go func() {
		data.SetTraderIndex()
		done <- true
	}()
	go func() {
		data.SetProfiles()
		done <- true
	}()
	for i := 0; i < 4; i++ {
		<-done
	}

	srv.SetServer()

	endTime := time.Now()
	fmt.Printf("\nDatabase initialized in %s\n\n", endTime.Sub(startTime))

	startHome()
}

func startHome() {

	fmt.Println("Alright fella, what now?")
	fmt.Println("1. Register an Account")
	fmt.Println("2. Login")
	fmt.Println("\n69. Exit")
	var input string
	for {
		fmt.Printf("> ")
		_, _ = fmt.Scanln(&input)

		switch input {
		case "1":
			registerAccount()
		case "2":
			login()
		case "69":
			fmt.Println("Adios fella")
			os.Exit(1)
		default:
			fmt.Println("Invalid input, intellectually less able fella")
		}
	}
}

func registerAccount() {
	account := data.Account{}
	profiles := data.GetProfiles()
	var input string

	fmt.Println("What is your username?")
	for {
		fmt.Printf("> ")
		_, _ = fmt.Scanln(&input)
		if !validateUsername(profiles, input) {
			fmt.Println("Username taken, try again")
			continue
		}
		break
	}
	account.Username = input

	fmt.Println("What is your password?")
	_, _ = fmt.Scanln(&input)
	fmt.Printf("> ")
	account.Password = input

	UID := tools.GenerateMongoID()
	account.UID = UID
	account.AID = len(profiles)

	profiles[UID] = &data.Profile{}
	profiles[UID].Account = &account
	profiles[UID].Character = &data.Character{
		ID: UID,
	}
	profiles[UID].Storage = &data.Storage{
		Suites: []string{},
		Builds: &data.Builds{
			EquipmentBuilds: []*data.EquipmentBuild{},
			WeaponBuilds:    []*data.WeaponBuild{},
		},
		Insurance: []any{},
		Mailbox:   []*data.Notification{},
	}
	profiles[UID].Dialogue = &data.Dialogue{}
	profiles[UID].Friends = &data.Friends{
		Friends:             []data.FriendRequest{},
		Ignore:              []string{},
		InIgnoreList:        []string{},
		Matching:            data.Matching{},
		FriendRequestInbox:  []any{},
		FriendRequestOutbox: []any{},
	}

	//save account
	profiles[UID].SaveProfile()

	//login
	fmt.Println("\nAccount created, logging in...")
	loggedIn(&account)
}

func validateUsername(profiles map[string]*data.Profile, username string) bool {
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
	var account *data.Account
	profiles := data.GetProfiles()
	if len(profiles) == 0 {
		fmt.Println("No profiles, redirecting to Account Register...")
		registerAccount()
	}

	for {
		fmt.Println("What is your username?")
		fmt.Printf("> ")
		_, _ = fmt.Scanln(&input)

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
		_, _ = fmt.Scanln(&input)

		if account.Password != input {
			fmt.Println("Invalid password, try again moron")
			continue
		}

		fmt.Println("Logging in...")

		loggedIn(profiles[account.UID].Account)
		break
	}

}

func loggedIn(account *data.Account) {
	fmt.Println("\nAlright fella, we're at the Login Menu, what now?")
	fmt.Println("\n1. Launch Tarkov")
	fmt.Println("2. Change Account Info")
	fmt.Println("3. Wipe yo ass")
	fmt.Println("\n69. Exit")

	for {
		var input string

		fmt.Printf("> ")
		_, _ = fmt.Scanln(&input)
		switch input {
		case "1":
			launchTarkov(account)
		case "2":
			editAccountInfo(account)
		case "3":
			wipeYoAss(account)
		case "69":
			fmt.Println("Adios faggot")
			return
		default:
			fmt.Println("Invalid input, retard")
		}
	}
}

func editAccountInfo(account *data.Account) {
	fmt.Println("\nAlright fella, what do you want to edit?")
	fmt.Println("\n1. Change Escape From Tarkov executable path")
	fmt.Println("69. Go back to Login Menu")

	for {
		var input string

		fmt.Printf("> ")
		_, _ = fmt.Scanln(&input)

		switch input {
		case "1":
			for {
				var tarkovPath string

				fmt.Println("\nSet new Path to Tarkov executable")
				fmt.Printf("> ")
				_, _ = fmt.Scanln(&tarkovPath)
				exePath := filepath.Join(tarkovPath, "EscapeFromTarkov.exe")
				if tools.FileExist(exePath) && exePath != account.TarkovPath {
					account.TarkovPath = exePath
					fmt.Println("Path has been set")

					if err := account.SaveAccount(); err != nil {
						log.Println(err)
						return
					}
					break
				}
				fmt.Println("Invalid path, try again")
			}
			editAccountInfo(account)
		case "69":
			loggedIn(account)
		default:
			fmt.Println("Invalid input, retard")
		}
	}
}

func wipeYoAss(account *data.Account) {
	account.Wipe = true
	profiles := data.GetProfiles()

	profiles[account.UID].Character = &data.Character{}
	profiles[account.UID].Storage = &data.Storage{
		Suites: []string{},
		Builds: &data.Builds{
			EquipmentBuilds: []*data.EquipmentBuild{},
			WeaponBuilds:    []*data.WeaponBuild{},
		},
		Insurance: []any{},
		Mailbox:   []*data.Notification{},
	}
	profiles[account.UID].Dialogue = &data.Dialogue{}
	fmt.Println("Yo ass is clean")
	profiles[account.UID].SaveProfile()
	loggedIn(account)
}

// tarkovPath + ' -bC5vLmcuaS5u={"email":"' + userAccount.email + '","password":"' + userAccount.password + '","toggle":true,"timestamp":0} -token=' + sessionID + ' -config={"BackendUrl":"https://' + serverConfig.ip + ':' + serverConfig.mainPort + '","Version":"live"}'
const (
	config = "-config={'BackendUrl':'%s','Version':'live'}"
	token  = "-token=%s"
	email  = "-bC5vLmcuaS5u={'email':'%s','password': '%s','toggle':true,'timestamp':0}"
)

func launchTarkov(account *data.Account) {
	if account.TarkovPath == "" || !tools.FileExist(account.TarkovPath) {
		fmt.Println("EscapeFromTarkov not found")
		fmt.Println("Input the folder/directory path to your 'EscapeFromTarkov.exe'")
		for {
			var tarkovPath string

			fmt.Printf("> ")
			_, _ = fmt.Scanln(&tarkovPath)
			if !tools.FileExist(filepath.Join(tarkovPath, "BepInEx")) {
				fmt.Println("This folder doesn't contain the 'BepInEx' directory, set path to your non-live 'EscapeFromTarkov' directory")
				continue
			}

			account.TarkovPath = filepath.Join(tarkovPath, "EscapeFromTarkov.exe")
			if !tools.FileExist(account.TarkovPath) {
				fmt.Println("Invalid path, does not contain 'EscapeFromTarkov.exe', try again")
				continue
			}
			if err := account.SaveAccount(); err != nil {
				log.Println(err)
				return
			}
			fmt.Println("Valid path to 'EscapeFromTarkov.exe' has been set")

			break
		}
	}

	cmdArgs := []string{
		fmt.Sprintf(email, account.Username, account.Password),
		fmt.Sprintf(token, account.UID),
		fmt.Sprintf(config, data.GetMainAddress()),
	}

	cmd := exec.Command(account.TarkovPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Client has been closed")
		//data.GetProfileByUID(account.UID).SaveProfile()
		//os.Exit(0)
	}
}
