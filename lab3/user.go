package lab3

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Fullname string   `json:"fullname"`
	Email    []string `json:"email"`
	Address  []string `json:"address"`
}
type Users struct {
	Users []User `json:"users"`
}

var tokens = make(map[string]string)
