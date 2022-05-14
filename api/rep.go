package api

const (
	// SourceBuiltIn 来自程序内部的固定内容
	SourceBuiltIn = "built-in"
)

// Rep 基本返回体
type Rep struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Source string      `json:"source"`
	Data   interface{} `json:"data"`
}
