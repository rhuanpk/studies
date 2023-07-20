package typeexample

type newType string

const (
	// FirstConst is the first constant.
	FirstConst newType = "hello"
	// SecondConst is the second constant.
	SecondConst newType = "world"
)

// AnyFunction is any random function.
func AnyFunction(str newType) {
	println(str)
}
