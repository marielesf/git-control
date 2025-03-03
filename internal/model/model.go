package model

type Commits struct {
	Timestamp  string `csv:"timestamp"` // .csv column headers
	User       string `csv:"username"`
	Repository string `csv:"repository"`
	Files      int    `csv:"files"`
	Additions  int    `csv:"additions"`
	Deletions  int    `csv:"deletions"`
}
