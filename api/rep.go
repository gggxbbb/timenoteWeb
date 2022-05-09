package api

const (
	SourceBuiltIn = "built-in"
)

// Rep 基本返回体
type Rep struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Source string      `json:"source"`
	Data   interface{} `json:"data"`
}
