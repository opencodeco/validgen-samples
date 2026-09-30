package account

//go:generate go run github.com/opencodeco/validgen@latest .

// Profile lives in this package. Callers import this module and call ProfileValidate.
type Profile struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"required"`
}
