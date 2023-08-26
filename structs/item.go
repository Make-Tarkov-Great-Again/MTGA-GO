package structs

type DatabaseItem struct {
	ID         string                 `json:"_id"`
	Name       string                 `json:"_name"`
	Parent     string                 `json:"_parent"`
	Type       string                 `json:"_type"`
	Properties map[string]interface{} `json:"_props"`
	Proto      string                 `json:"_proto"`
}
