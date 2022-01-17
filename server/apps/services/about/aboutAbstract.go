package about

type Abouts struct {
	About     []*AboutContent
	TimeStamp int
}

type AboutContent struct {
	Title   string `json:"title"`
	Intro   string `json:"intro"`
	Picture string `json:"pictures"`
	Number  int    `json:"number"`
}
