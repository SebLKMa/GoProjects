package main

import (
	"fmt"
)

type User struct {
	name  string
	email string
}

// A method wuth a value receiver (u User)
func (u User) notify() {
	fmt.Printf("Sending notification to %s<%s>\n", u.name, u.email)
}

// A method wuth a pointer receiver (u *User)
func (u *User) changeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	// Values of type User can be used to call methods
	// declared with a value receiver
	mona := User{"Mona", "mona@somewhere.com"}
	mona.notify()

	// lisa is declared and initialized as a Pointer to an address of type User
	lisa := &User{"Lisa", "lisa@nowhere.com"}
	lisa.notify()

	// Values of type User can be used to call methods declared with
	// a pointer receiver (u *User) changeEmail()
	mona.changeEmail("mona@nowhere.com")
	mona.notify()

	// Pointers of type User can be used to call methods declared with
	// a value receiver (u User) notify()
	lisa.changeEmail("mona@somewhere.com")
	lisa.notify()

}
