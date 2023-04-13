package model

type Mahasiswa struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Major string `json:"major"`
}
