package pongo2

type Number interface {
	Add(value interface{}) Number
	Sub(value interface{}) Number
	Mul(value interface{}) Number
	Div(value interface{}) Number

	Less(value interface{}) Number
	LessOrEqual(value interface{}) Number
	Equal(value interface{}) Number
	NotEqual(value interface{}) Number
	Greater(value interface{}) Number
	GreaterOrEqual(value interface{}) Number

	Bool() Number

	AsNumber(value interface{}) (Number, error)
}
