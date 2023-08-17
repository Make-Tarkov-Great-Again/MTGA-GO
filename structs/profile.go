package structs

type Profile struct {
	Account   *Account
	Character *PlayerTemplate
	Storage   *Storage
	Dialogue  map[string]interface{}
}

type Storage struct {
	ID        string                 `json:"_id"`
	Suites    []string               `json:"suites"`
	Builds    map[string]interface{} `json:"builds"`
	Insurance []interface{}          `json:"insurance"`
	Mailbox   []interface{}          `json:"mailbox"`
}
