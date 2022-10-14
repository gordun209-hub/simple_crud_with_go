package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

var (
	Users = map[int]*User{}
	Seq   = 1
)
