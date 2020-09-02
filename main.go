package main

//taken from: tutorialedge.net
//needs modification to compare strings
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Users struct which contains
//an array of users
type Users struct {
	Users []User `json:"users"`
}

//User struct which contains a name
//a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

//Social struct which constains a
//list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	//open our jsonFile
	jsonFile, err := os.Open("users.json")
	//if os.Open returns an error then print err
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	//defer the closing of our jsonFile so that we can parse it
	defer jsonFile.Close()

	//read out opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//we initialize our User array
	var users Users

	//we unmarshal our byteArray which contains our
	//jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	//we iterate through every user within our users array and
	//print out the user Type, their name, and their facebook url
	//as an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type:" + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
	}

	//my test to see if string type is the same
	fmt.Println(users.Users[0].Type)
	typeResult := users.Users[0].Type == users.Users[1].Type
	fmt.Println("Are types the same? ", typeResult)

}
