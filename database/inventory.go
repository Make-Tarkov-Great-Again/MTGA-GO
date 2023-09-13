package database

type Inventory struct {
	Items              []map[string]interface{} `json:"items"`
	Equipment          string                   `json:"equipment"`
	Stash              string                   `json:"stash"`
	SortingTable       string                   `json:"sortingTable"`
	QuestRaidItems     string                   `json:"questRaidItems"`
	QuestStashItems    string                   `json:"questStashItems"`
	FastPanel          interface{}              `json:"fastPanel"`
	HideoutAreaStashes interface{}              `json:"hideoutAreaStashes"`
}

type Containers struct {
	Equipment    Container
	Stash        Container
	SortingTable Container
}

type Container struct{}
