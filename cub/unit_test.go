package main

import "testing"

func Test_Email(t *testing.T) {
	SendEmail("support@dailywire.com", "me@here.com", "In Nomine Patris...", "Et Filii...", make(chan error))
}
