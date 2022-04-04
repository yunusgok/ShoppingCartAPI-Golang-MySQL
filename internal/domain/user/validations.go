package user

import "regexp"

// A valid username should start with an alphabet so, [A-Za-z].
// All other characters can be alphabets, numbers or an underscore so, [A-Za-z0-9_].
// Minimum 8 characters
// Maximum 30 characters
var usernameRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,29}$")

//Minimum eight characters, at least one letter and one number:
var passwordRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{7,29}$`)

func ValidateUserName(name string) bool {
	return usernameRegex.MatchString(name)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)

}
