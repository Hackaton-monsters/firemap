package command

type CreateMarker struct {
	Lat     float64 `json:"lat" binding:"required"`
	Lon     float64 `json:"lon" binding:"required"`
	Type    string  `json:"type" binding:"required"`
	Comment string  `json:"comment"`
	Photos  []int64 `json:"photos"`
}
