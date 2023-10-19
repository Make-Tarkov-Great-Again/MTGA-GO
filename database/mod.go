package database

type ModInfo struct {
	NameSpace       string
	ModNameNoSpaces string
	Advanced        struct {
		CustomRoutes bool
	}
	Config map[string]interface{}
}

func (m *ModInfo) GetConfig() map[string]interface{} {
	return m.Config
}
