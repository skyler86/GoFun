package serializer

// 通用的基础序列化器结构体
type Response struct {
	Status int `json:"status"`		// 返回一个状态
	Data interface{} `json:"data"`
	Msg string `json:"msg"`			// 返回一个信息
	Error string `json:"error"`		// 返回一个错误
}

// TokenData带有token的Data结构体
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}