package main

import "fmt"

func main(){
	fmt.Println("Can Fast Attack?", CanFastAttack(false))
	fmt.Println("Can Spy?", CanSpy(false, true, false))
	fmt.Println("Can Signal Prisoner?", CanSignalPrisoner(true, false))
	fmt.Println("Can Free Prisoner?", CanFreePrisoner(false, false, false, true))
}

func CanFastAttack(knightIsAwake bool) bool {
    return !knightIsAwake 
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return !archerIsAwake && petDogIsPresent || prisonerIsAwake && !knightIsAwake && !archerIsAwake
}