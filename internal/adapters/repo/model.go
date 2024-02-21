package repo

type Url struct {
	ID  int    `db:"id"`
	Url string `db:"url"`
}
