package signup_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/opencodeco/validgen-samples/signup"
)

func TestUserValidate(t *testing.T) {
	valid := &signup.User{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Age:       36,
		Username:  "adalovelace",
	}
	if errs := signup.UserValidate(valid); len(errs) != 0 {
		t.Fatalf("valid user: %v", errs)
	}

	edges := []*signup.User{
		{FirstName: "Ada", LastName: "Lovelace", Email: "ada@example.com", Age: 18, Username: "ada18"},
		{FirstName: "Ada", LastName: "Lovelace", Email: "ada@example.com", Age: 130, Username: strings.Repeat("b", 20)},
	}
	for _, user := range edges {
		if errs := signup.UserValidate(user); len(errs) != 0 {
			t.Fatalf("age %d username %q: %v", user.Age, user.Username, errs)
		}
	}

	assertErrors(t, signup.UserValidate(&signup.User{}), []string{
		"FirstName is required",
		"LastName is required",
		"Email is required",
		"Email must be a valid email",
		"Age must be >= 18",
		"Username length must be >= 5",
	})

	bad := &signup.User{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada.example.com",
		Age:       131,
		Username:  strings.Repeat("a", 21),
	}
	assertErrors(t, signup.UserValidate(bad), []string{
		"Email must be a valid email",
		"Age must be <= 130",
		"Username length must be <= 20",
	})
}

func TestUnmarshalJSON(t *testing.T) {
	var user signup.User
	raw := []byte(`{"FirstName":"Ada","LastName":"Lovelace","Email":"ada@example.com","Age":36,"Username":"adalovelace"}`)
	if err := json.Unmarshal(raw, &user); err != nil {
		t.Fatal(err)
	}
	if user.Email != "ada@example.com" {
		t.Fatalf("email = %s", user.Email)
	}

	var missing signup.User
	err := json.Unmarshal([]byte(`{"LastName":"Lovelace","Email":"ada@example.com","Age":36,"Username":"adalovelace"}`), &missing)
	assertErrors(t, unwrapJoin(t, err), []string{
		"FirstName is required",
	})

	var broken signup.User
	err = json.Unmarshal([]byte(`{"Age":`), &broken)
	if err == nil {
		t.Fatal("expected a decode error")
	}
	if _, ok := err.(interface{ Unwrap() []error }); ok {
		t.Fatalf("decode error unwrapped as validation errors: %v", err)
	}
}

func unwrapJoin(t *testing.T, err error) []error {
	t.Helper()
	joined, ok := err.(interface{ Unwrap() []error })
	if !ok {
		t.Fatalf("expected joined validation errors, got %v", err)
	}
	return joined.Unwrap()
}

func assertErrors(t *testing.T, errs []error, want []string) {
	t.Helper()
	got := make([]string, len(errs))
	for i, err := range errs {
		got[i] = err.Error()
	}
	if len(got) != len(want) {
		t.Fatalf("got %q, want %q", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %q, want %q", got, want)
		}
	}
}
