// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	users := make(map[string]map[string]string, 10)
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "Cat"

	modifyUser(users, "smith")

	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {

	//Check key exists or not
	//v,ok:=users[name]
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "nickname~" + name
	}
}
