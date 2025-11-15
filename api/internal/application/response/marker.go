package response

type Marker struct {
	ID           int64    `json:"id"`
	ChatID       int64    `json:"chatId"`
	Lat          float64  `json:"lat"`
	Lon          float64  `json:"lon"`
	Reports      []Report `json:"reports"`
	ReportsCount int      `json:"reportsCount"`
	Type         string   `json:"type"`
	Title        string   `json:"title"`
}

type Report struct {
	ID      int64   `json:"id"`
	Comment string  `json:"comment"`
	Photos  []int64 `json:"photos"`
}
