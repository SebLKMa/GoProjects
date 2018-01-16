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
	fmt.Printf("Sending notification to %s<%s>\n", u.name, u.email)
}

// accepts value that implements notifier
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	mona := user{"Mona", "mona@somewhere.com"}

	// compile error!
	// user does not implement notifier (notify method has pointer receiver)
	//sendNotification(mona)

	// implementation expects a pointer (u *user) notify()
	// we will pass in address of object
	sendNotification(&mona)
}
