package main

import "fmt"

func main(){
	fmt.Println(Welcome("Glauber"))
	fmt.Println(HappyBirthday("Glauber", 32))
	fmt.Println(AssignTable("Glauber", 021, "Frank", "direita", 9.2))
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s", name + "!")
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
}