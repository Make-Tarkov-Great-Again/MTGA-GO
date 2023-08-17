package structs

type Locations struct {
	Locations map[string]interface{} `json:"locations"`
	Paths     []interface{}          `json:"paths"`
}
