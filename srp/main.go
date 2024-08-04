package main      

import "fmt"
import "srp/user"

func main() {
  myUser := user.User{}
  myUser.CreateUser("lal badshah", "abc@xyz.in")
  fmt.Println(myUser.GetFullName())
}
