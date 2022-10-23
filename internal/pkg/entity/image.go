package entity

type Image struct {
	Title string `db:"title" json:"title"`
	Date  string `db:"date" json:"date"`
	URL   string `db:"url" json:"url"`
	Data  []byte `db:"data" json:"data"`
}
