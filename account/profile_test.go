package account_test

import (
	"testing"

	"github.com/opencodeco/validgen-samples/account"
)

func TestProfileValidate(t *testing.T) {
	ok := &account.Profile{FirstName: "Ada", LastName: "Lovelace", Age: 36}
	if errs := account.ProfileValidate(ok); len(errs) != 0 {
		t.Fatal(errs)
	}

	assertErrors(t, account.ProfileValidate(&account.Profile{}), []string{
		"FirstName is required",
		"LastName is required",
		"Age is required",
	})

	named := &account.Profile{FirstName: "Ada", LastName: "Lovelace"}
	assertErrors(t, account.ProfileValidate(named), []string{
		"Age is required",
	})
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
