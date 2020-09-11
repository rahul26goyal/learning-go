package main

import "fmt"

/*
Define a User Struct
*/
type User struct {
	Id             int
	Name, Location string
}

/*
 Define a Function for User Struct to print the content.
*/
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %d %s from %s", u.Id,
		u.Name, u.Location)
}

/*
 Define a Player Struct which Embeds User Object inline
*/
type Player struct {
        User
	GameId int
}

/*
 Define a Player Struct which embeds an User object Pointer.
*/
type Player3 struct {
        *User
	GameId int
}

/*
 Define a Player Struct which embeds an User object Pointer in a variable.
*/
type Player4 struct {
   u *User
   GameId int
}


func main() {
	// Create an Object
  //p1 = Player{}
	var p1 Player = Player{}
  // Since the User Struct in contained inline, we can access all fields directly.
	p1.Id = 123
	p1.Name = "Rahul"
	p1.Location = "LDG"
	p1.GameId = 1233

	fmt.Printf("Player1 : %+v Game: %d", p1.Greetings(), p1.GameId)
  // Approach 2
	u := User{1, "Khush", "RPR"}
	p2 := Player{u, 12344}
	fmt.Printf("\nPlayer2 - %+v Game: %d", p2.Greetings(), p2.GameId)

	//update ID
	u.Id = 12
	// does not reflect in p2 object as it was pass by argument.
	fmt.Printf("\nPlayer2 - %+v Game: %d", p2.Greetings(), p2.GameId)
	//User{123, "rahul", "LDG"}, 1234


	p3 := Player3{&u, 12}
	u.Id = 1222
  // Reflects as Pointer is passed.
	fmt.Printf("\nPlayer3 - %+v Game: %d", p3.Greetings(), p3.GameId)
	p3.Id = 123333
	fmt.Printf("\nPlayer3 - %+v Game: %d", p3.Greetings(), p3.GameId)
	fmt.Printf("\nUpdated Id : %d", u.Id)

  p4 := Player4{u: &u, GameId: 12}
  u.Id = 11000
	fmt.Printf("\nPlayer4 - %+v Game: %d", p4.u.Greetings(), p4.GameId)
}
