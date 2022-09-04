package status

type Code int64

const (
	OK Code = iota
	NoLogin
	TokenError
	TokenExpired
	TooLong
	UsernameExists
	UsernamePasswordNotMatch
	UserNotExists
)

func (c Code) Message() string {
	switch c {
	case OK:
		return "OK"
	case NoLogin:
		return "未登录"
	case TokenError:
		return "不正确的token"
	case TokenExpired:
		return "token已过期"
	case TooLong:
		return "用户名或密码最长32个字符"
	case UsernameExists:
		return "用户名重名"
	case UsernamePasswordNotMatch:
		return "用户名或密码错误"
	case UserNotExists:
		return "用户不存在"
	default:
		return "Unknown error"
	}
}
