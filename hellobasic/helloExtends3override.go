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

// admin implements notifier overrides (*user).notify()
func (a *admin) notify() {
	fmt.Printf("Sending notification to admin %s<%s>\n", a.name, a.email)
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

	// access inner type's method directly
	ad.user.notify()

	// inner type's method is promoted
	ad.notify()
}
