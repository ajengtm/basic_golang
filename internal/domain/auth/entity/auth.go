package entity

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}
