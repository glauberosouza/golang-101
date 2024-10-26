package main

import "fmt"

func main()  {
	fmt.Println(ParseCard("ace"))
	fmt.Println(FirstTurn("ace","ace","four"))
}

func ParseCard(card string) int {
    
    switch card {
    	case "ace":
        	return 11
        case "ten", "jack", "queen", "king":
        	return 10
        case "nine":
        	return 9
        case "eight":
        	return 8
        case "seven":
        	return 7
        case "six":
        	return 6
        case "five":
        	return 5
        case "four":
        	return 4
        case "three":
        	return 3
        case "two":
        	return 2
        default:
        	return 0
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
	playerCardA := ParseCard(card1)
    playerCardB := ParseCard(card2)
    dealer := ParseCard(dealerCard)
	playerScore := playerCardA + playerCardB
    if card1 == "ace" && card2 == "ace" {
		return "P" // Split (par de ases)
	}
	if playerScore == 21 {
		if dealer == 10 || dealer == 11 {
			return "S" 
		}
		return "W" 
	}
	if playerScore >= 17 && playerScore <= 20 {
		return "S" 
	}
	if playerScore >= 12 && playerScore <= 16 {
		if dealer >= 7 {
			return "H" 
		}
		return "S" 
	}
	return "H" 
}