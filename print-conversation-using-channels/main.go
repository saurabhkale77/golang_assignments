package main

import "fmt"

func main() {
	var conversationStr = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	var chanOfAlice chan string = make(chan string)
	var chanOfBob chan string = make(chan string)

	var aliceMessage, bobMessage string
	var isChannelOfAliceOpen, isChannelOfBobOpen = true, true

	go extractTheConversation(conversationStr, chanOfAlice, chanOfBob)

	for {
		select {
		case aliceMessage, isChannelOfAliceOpen = <-chanOfAlice:
			if isChannelOfAliceOpen {
				fmt.Printf("Alice: %s\n", aliceMessage)
			}
		case bobMessage, isChannelOfBobOpen = <-chanOfBob:
			if isChannelOfBobOpen {
				fmt.Printf("Bob: %s\n", bobMessage)
			}
		default:
			// fmt.Println("default")
		}

		if !isChannelOfAliceOpen && !isChannelOfBobOpen {
			break
		}

	}
}

func extractTheConversation(str string, chanOfAlice chan string, chanOfBob chan string) {
	startIndex := 0
	for index, char := range str {
		if char == '$' {
			chanOfAlice <- str[startIndex:index]
			startIndex = index + 1
		} else if char == '#' {
			chanOfBob <- str[startIndex:index]
			startIndex = index + 1
		} else if char == '^' {
			close(chanOfAlice)
			close(chanOfBob)
			break
		}
	}
}
