package entity

type Resource struct {
	UUID         int64  `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         int64  `json:"size"`
	Price        int64  `json:"price"`
	ParsedDate   string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}
