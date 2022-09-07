package list

func List(args *Request) (*Response, error) {
	return nil, nil
}

type Request struct {
}

type Response struct {
}

type status int32

const (
	statusOK status = iota
)

func (s status) msg() string {
	switch s {
	case statusOK:
		return "OK"
	default:
		return "未知错误"
	}
}
