package model

type Mahasiswa struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Major      string `json:"major"`
	UserName   string `json:"user_name"`
	C_Password string `json:"c_password"`
	Username   string `json:"username"`
}
