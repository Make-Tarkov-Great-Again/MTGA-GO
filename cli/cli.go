package cli

import (
	"MT-GO/data"
	"MT-GO/tools"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Start() {

	fmt.Println("\nAlright fella, what now?")
	fmt.Println("1. Register an Account")
	fmt.Println("2. Login")
	fmt.Println("\n69. Exit")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

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
	account := new(data.Account)
	profiles := data.GetProfiles()
	var input string
	fmt.Println("What is your username?")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input = scanner.Text()

		if !validateUsername(profiles, input) {
			fmt.Println("Username taken, try again")
			continue
		}
		break
	}
	account.Username = input

	fmt.Println("What is your password?")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	account.Password = scanner.Text()

	fmt.Println("Account Type? 1 - Dev, 2 - Edge of Darkness")
	for account.Edition == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "1":
			account.Edition = "developer"
		case "2":
			account.Edition = "edge of darkness"
		default:
			fmt.Println("Invalid input bozo")
		}
	}

	fmt.Println("Account Language? (ex: en, ru, sk, es-mx)")
	for account.Lang == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		_, err := data.GetLocaleByName(input)
		if err != nil {
			fmt.Println("Invalid language bozo")
			continue
		}
		account.Lang = input
	}

	UID := tools.GenerateMongoID()
	account.UID = UID
	account.AID = len(profiles)

	profiles[UID] = &data.Profile{
		Account: account,
		Character: &data.Character[map[string]data.PlayerTradersInfo]{
			ID: UID,
		},
		Friends: &data.Friends{
			Friends:             []data.FriendRequest{},
			Ignore:              []string{},
			InIgnoreList:        []string{},
			Matching:            data.Matching{},
			FriendRequestInbox:  []any{},
			FriendRequestOutbox: []any{},
		},
		Storage: &data.Storage{
			Suites: []string{},
			Builds: &data.Builds{
				EquipmentBuilds: []*data.EquipmentBuild{},
				WeaponBuilds:    []*data.WeaponBuild{},
			},
			Insurance: []any{},
			Mailbox:   []*data.Notification{},
		},
		Dialogue: new(data.Dialogue),
		Cache:    nil,
	}

	//save profile
	profiles[UID].SaveProfile()

	//login
	fmt.Println("\nAccount created, logging in...")
	loggedIn(profiles[UID].Account)
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
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input = scanner.Text()

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
		scanner.Scan()
		input = scanner.Text()

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
	fmt.Println("\n\nAlright fella, we're at the Login Menu, what now?")
	fmt.Println("\n1. Launch Tarkov")
	fmt.Println("2. Change Account Info")
	fmt.Println("3. Wipe yo ass")
	fmt.Println("\n69. Exit")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			launchTarkov(account)
		case "2":
			editAccountInfo(account)
		case "3":
			wipeYoAss(account)
		case "69":
			fmt.Println("Adios")
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func editAccountInfo(account *data.Account) {
	fmt.Println("\nAlright fella, what do you want to edit?")
	fmt.Println("\n1. Change Escape From Tarkov executable path")
	fmt.Println("69. Go back to Login Menu")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			for {
				fmt.Println("\nSet new Path to Tarkov executable")
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("> ")
				scanner.Scan()
				path := scanner.Text()

				exePath := filepath.Join(path, "EscapeFromTarkov.exe")
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
			fmt.Println("Invalid input")
		}
	}
}

func wipeYoAss(account *data.Account) {
	account.Wipe = true
	profiles := data.GetProfiles()

	profiles[account.UID].Character = &data.Character[map[string]data.PlayerTradersInfo]{}
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

func setTarkovPath() string {
	fmt.Println("EscapeFromTarkov not found")
	fmt.Println("Input the folder/directory path to your 'EscapeFromTarkov.exe'")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		path := scanner.Text()

		exePath := filepath.Join(path, "BepInEx")
		if !tools.FileExist(exePath) {
			fmt.Println("This folder doesn't contain the 'BepInEx' directory, set path to your non-live 'EscapeFromTarkov' directory")
			continue
		}

		exePath = filepath.Join(path, "EscapeFromTarkov.exe")
		if !tools.FileExist(exePath) {
			fmt.Println("Invalid path, does not contain 'EscapeFromTarkov.exe', try again")
			continue
		}

		fmt.Println("Valid path to 'EscapeFromTarkov.exe' has been set")
		return exePath
	}
}

func checkIfValidPath(path string) bool {
	exePath := filepath.Join(path, "BepInEx")
	if !tools.FileExist(exePath) {
		fmt.Println("This folder doesn't contain the 'BepInEx' directory, set path to your non-live 'EscapeFromTarkov' directory")
		return false
	}

	exePath = filepath.Join(path, "EscapeFromTarkov.exe")
	if !tools.FileExist(exePath) || path != exePath {
		fmt.Println("Invalid path, does not contain 'EscapeFromTarkov.exe'")
		return false
	}

	return true
}

func launchTarkov(account *data.Account) {
	if !checkIfValidPath(account.TarkovPath) {
		account.TarkovPath = setTarkovPath()
		if err := account.SaveAccount(); err != nil {
			log.Fatalln(err)
		}
	}

	cmdArgs := []string{
		fmt.Sprintf(email, account.Username, account.Password),
		fmt.Sprintf(token, account.UID),
		fmt.Sprintf(config, data.GetMainAddress()),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, account.TarkovPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Println("Client has been closed with an error:", err)
		os.Exit(1)
	}
}
