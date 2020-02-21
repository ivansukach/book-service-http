package handlers

type BookModel struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	Edition       string `json:"edition"`
	NumberOfPages int32  `json:"numberOfPages"`
	Year          int32  `json:"year"`
	Amount        int32  `json:"amount"`
	IsPopular     bool   `json:"isPopular"`
	InStock       bool   `json:"inStock"`
}
