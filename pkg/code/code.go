package code

type Code struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCode(code int, msg string) Code {
	return Code{code, msg}
}

func (c Code) Error() string {
	return c.Msg
}
