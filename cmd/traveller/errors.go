package main

// StringError describes a basic error with static information.
type StringError string

// Error returns the string itself.
func (err StringError) Error() string {
	return string(err)
}
