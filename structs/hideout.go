package structs

type Hideout struct {
	Areas       []interface{}   `json:"areas"`
	Productions []interface{}   `json:"production"`
	QTE         []interface{}   `json:"qte"`
	ScavCase    []interface{}   `json:"scavcase"`
	Settings    HideoutSettings `json:"settings"`
}

type HideoutSettings struct {
	GeneratorSpeedWithoutFuel float64 `json:"generatorSpeedWithoutFuel"`
	GeneratorFuelFlowRate     float64 `json:"generatorFuelFlowRate"`
	AirFilterUnitFlowRate     float64 `json:"airFilterUnitFlowRate"`
	GPUBoostRate              float64 `json:"gpuBoostRate"`
}
