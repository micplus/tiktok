package request

type User struct {
	ID    int64
	Token string
}

type Register struct {
	Username string
	Password string
}

type Login struct {
	Username string
	Password string
}
