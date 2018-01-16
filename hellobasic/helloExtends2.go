package main

import (
	"fmt"
)

// lowercase type name denotes unexported interface
type notifier interface {
	notify()
}

// lowercase type name denotes unexported type
type user struct {
	name  string
	email string
}

// user implements notifier via a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending notification to user %s<%s>\n", u.name, u.email)
}

// admin has-a user, and own property
type admin struct {
	user  // embedded type
	level string
}

// accepts value that implements notifier
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	// creates an admin user
	ad := admin{
		user: user{
			name:  "Mona Lisa",
			email: "monalisa@somewhere.com",
		},
		level: "superuser",
	}

	// the embedded inner type's implementation of notifier interface is promoted
	sendNotification(&ad)
}
