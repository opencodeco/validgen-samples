package signup

//go:generate go run github.com/opencodeco/validgen@latest -unmarshal-json .

// User is validated in this package. ValidGen writes UserValidate and UnmarshalJSON into validator__.go.
type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Age       uint8  `validate:"gte=18,lte=130"`
	Username  string `validate:"min=5,max=20"`
}

// Note is parsed and skipped because it has no valid or validate tag.
type Note struct {
	Body string
}
