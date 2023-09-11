package webapi

import (
	"fmt"
	"log"
)

type AccountService interface {
	Open() error
	Close() error
	GetAccountById(id string) (*Account, error)
	CreateAccount(m *Account) error
	UpdateAccount(m *Account) error
	AuthenticateAccount(username string, password string) (bool, error)
}

type AccountServiceImpl struct {
	// TODO add a pointer to technical connection
}

func (acc *AccountServiceImpl) Open() error {

	log.Println("Connected to DB")

	return nil
}

func (acc *AccountServiceImpl) Close() error {
	fmt.Println("close")
	return nil
}

func (acc *AccountServiceImpl) GetAccountById(id string) (*Account, error) {
	var movie = &Account{}
	log.Println("GetAccountById " + id)

	return movie, nil
}

func (acc *AccountServiceImpl) CreateAccount(a *Account) error {
	log.Println("CreateAccount ")
	return nil
}

func (acc *AccountServiceImpl) UpdateAccount(a *Account) error {
	log.Println("UpdateAccount " + a.String())
	return nil
}

func (store *AccountServiceImpl) AuthenticateAccount(username string, password string) (bool, error) {

	return true, nil
}
