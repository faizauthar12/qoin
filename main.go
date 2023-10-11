package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	NumDice    int
	RolledDice []int
	Point      int
}

const (
	minDiceNumber = 1
	maxDiceNumber = 6
)

func RollTheDice() int {
	return rand.Intn(maxDiceNumber) + minDiceNumber
}

func InitPlayer(simulatePlayer int, simulateDice int) []Player {

	var players []Player
	for i := 0; i < simulatePlayer; i++ {
		player := Player{
			NumDice:    simulateDice,
			RolledDice: []int{},
			Point:      0,
		}

		players = append(players, player)
	}

	return []Player{}
}

func RollPlayerDice(players []Player) []Player {

	newPlayers := []Player(players)

	for index, item := range players {
		var rolledDices []int
		if item.NumDice != 0 {
			for i := 0; i < item.NumDice; i++ {
				rolledDice := rand.Intn(maxDiceNumber) + minDiceNumber

				rolledDices = append(rolledDices, rolledDice)
			}

			newPlayers[index].RolledDice = rolledDices
		}
	}

	return newPlayers
}

func EvaluatePlayerDice(players []Player) []Player {

	newPlayers := []Player(players)

	for indexPlayer, itemPlayer := range players {
		if itemPlayer.RolledDice != nil {
			for indexDice, itemDice := range itemPlayer.RolledDice {
				dices := itemPlayer.RolledDice
				givenDice := itemDice

				if itemDice == 1 {
					var nextPlayerIndex int

					if indexPlayer == len(players)-1 {
						nextPlayerIndex = 0
					} else {
						nextPlayerIndex = indexPlayer + 1
					}

					// remove current dice from player
					newPlayers[indexPlayer].RolledDice = append(dices[:indexDice], dices[indexDice+1:]...)
					newPlayers[indexPlayer].NumDice = players[indexPlayer].NumDice - 1

					// add dice to next player
					newPlayers[nextPlayerIndex].RolledDice = append(newPlayers[nextPlayerIndex].RolledDice, givenDice)
					newPlayers[nextPlayerIndex].NumDice = players[nextPlayerIndex].NumDice + 1

				} else if itemDice == 6 {
					newPlayers[indexPlayer].RolledDice = append(dices[:indexDice], dices[indexDice+1:]...)
					newPlayers[indexPlayer].NumDice = players[indexPlayer].NumDice - 1

					newPlayers[indexPlayer].Point = players[indexPlayer].Point + 1
				}
			}
		}
	}

	return newPlayers
}

func main() {

	simulatePlayer := 4
	simulateDice := 4

	players := make([]Player, simulatePlayer)

	for index := range players {
		players[index].NumDice = simulateDice
	}

	fmt.Printf("Pemain = %d, Dadu = %d\n", len(players), simulateDice)

	players = RollPlayerDice(players)

	fmt.Println("Lempar Dadu")
	for index, item := range players {
		fmt.Printf("Pemain #%d (%d): %d\n", index, item.Point, item.RolledDice)
	}

	players = EvaluatePlayerDice(players)

	fmt.Println("Evaluasi")
	for index, item := range players {
		fmt.Printf("Pemain #%d (%d): %d\n", index, item.Point, item.RolledDice)
	}
}
