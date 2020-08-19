package pongo2

type Number interface {
	Add(value interface{}) Number
	Sub(value interface{}) Number
	Mul(value interface{}) Number
	Div(value interface{}) Number

	AsNumber(value interface{}) (Number, error)
}
