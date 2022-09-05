package request

type User struct {
	LoginID int64
	UserID  int64
}

type Register struct {
	Username string
	Password string
}

type Login struct {
	Username string
	Password string
}
