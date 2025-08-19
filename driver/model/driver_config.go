package model

type DriverConfig struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	IndexUrl string `json:"indexUrl"`
	ConfJson string `json:"confJson"`
}
