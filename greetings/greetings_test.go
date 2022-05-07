package greetings

import (
	"regexp"
	"testing"
)

func Test_GivenName_WhenHello_ThenReturnGreetings(t *testing.T) {
	// given
	name := "Gladys"

	// when
	msg, err := Hello("Gladys")

	// then
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func Test_GivenEmptyName_WhenHello_ThenReturnError(t *testing.T) {
	// when
	msg, err := Hello("")

	// then
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}

}
