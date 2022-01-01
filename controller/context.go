package controller

// Context gin.Contextのインターフェース
type Context interface {
	Param(string) string
	Query(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
