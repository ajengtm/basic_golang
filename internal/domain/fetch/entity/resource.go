package entity

type Resource struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"usd_price"`
	ParsedDate   string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type ResourceData struct {
	AreaProvinsi string `json:"area_provinsi"`
	Profit       map[string]map[string]int
	Max          float64 `json:"max_profit"`
	Min          float64 `json:"min_profit"`
	Avg          float64 `json:"average_profit"`
	Median       float64 `json:"median_profit"`
}
