package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

// method of User
func (u User) info() {
	fmt.Printf("User name : %v\n", u.Username)
	fmt.Printf("Full name : %v\n", u.FullName)
	fmt.Printf("Profile Picture URL : %v\n", u.ProfilePicUrl)
}

// general method
// func info(u User) {
// 	fmt.Printf("User name : %v\n", u.Username)
// 	fmt.Printf("Full name : %v\n", u.FullName)
// 	fmt.Printf("Profile Picture URL : %v\n", u.ProfilePicUrl)
// }

func main() {
	u := User{
		Username:      "golang",
		FullName:      "Basic Golang",
		ProfilePicUrl: "http://knowhere.fake/gopher.jpg",
	}
	// info(u)
	u.info()

}
