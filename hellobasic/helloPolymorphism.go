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

type admin struct {
	name  string
	email string
}

// admin implements notifier via a pointer receiver
func (a *admin) notify() {
	fmt.Printf("Sending notification to admin %s<%s>\n", a.name, a.email)
}

// accepts value that implements notifier
func sendNotification(n notifier) {
	n.notify()
}

func main() {

	// both user and admin implement notifier

	mona := user{"Mona", "mona@somewhere.com"}
	sendNotification(&mona)

	lisa := admin{"Lisa", "lisa@nowhere.com"}
	sendNotification(&lisa)
}
