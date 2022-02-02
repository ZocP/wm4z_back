package tour

type Position struct {
	ID          int    `json:"ID"`
	Picture     string `json:"Picture"`
	FloorNumber int    `json:"FloorNumber"`
	Map         string `json:"Map"`
}

type Floor struct {
	FloorNumber int        `json:"FloorNumber"`
	Positions   []Position `json:"Positions"`
}
