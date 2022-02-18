package entity

type Resource struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	ParsedDate   string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type ResourceResponse struct {
	UUID         string  `json:"uuid"`
	Komoditas    string  `json:"komoditas"`
	AreaProvinsi string  `json:"area_provinsi"`
	AreaKota     string  `json:"area_kota"`
	Size         string  `json:"size"`
	Price        float64 `json:"price"`
	PriceUSD     float64 `json:"usd_price"`
	ParsedDate   string  `json:"tgl_parsed"`
	Timestamp    string  `json:"timestamp"`
}

type ResourceAgregationResponse struct {
	Min    []AggregateResources `json:"min"`
	Max    []AggregateResources `json:"max"`
	Median []AggregateResources `json:"median"`
	Avg    []AggregateResources `json:"avg"`
}

type AggregateResources struct {
	AreaProvinsi string `json:"area_provinsi"`
	WeekBegining string `json:"week_begining"`
}
