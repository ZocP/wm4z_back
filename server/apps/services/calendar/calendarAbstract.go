package calendar

type Event struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Date  int    `json:"date"`
	Event string `json:"event"`
}
