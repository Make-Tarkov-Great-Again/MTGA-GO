package structs

type Hideout struct {
	Index    HideoutIndex
	Areas    []map[string]interface{}
	Recipes  []map[string]interface{}
	QTE      []map[string]interface{}
	ScavCase []map[string]interface{}
	Settings HideoutSettings
}

type HideoutIndex struct {
	Areas    map[int8]int8
	ScavCase map[string]int8
	Recipes  map[string]int16
}

type HideoutSettings struct {
	GeneratorSpeedWithoutFuel float64 `json:"generatorSpeedWithoutFuel"`
	GeneratorFuelFlowRate     float64 `json:"generatorFuelFlowRate"`
	AirFilterUnitFlowRate     float64 `json:"airFilterUnitFlowRate"`
	GPUBoostRate              float64 `json:"gpuBoostRate"`
}
