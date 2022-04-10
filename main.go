package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Color int
	Number int 
}
type User struct {
	Name string
	Deck []Card
}
func (user *User)Add_Card(data Card) {
	user.Deck = append(user.Deck, data)
}
func (user *User)Remove_Card(index int) {
	user.Deck = append(user.Deck[:index], user.Deck[index+1:]...)
}
func (user *User)Control(Turn_card Card) int {
	for i,data := range user.Deck {
		if data.Color == Turn_card.Color || data.Number == Turn_card.Number {
			return i
		}
	}
	return -1
}
type Game struct {
	Player_1 *User	
	Player_2 *User
	Deck []Card
	Turn_card Card
	Turn_count int
}
func (game *Game)Card_get() Card{
	card := game.Deck[0]
	game.Deck = game.Deck[1:]
	return card
}
func (game *Game)Game_set(){
	for color := 0;color < 3;color++ { // 색은 3개(0부터 셈)
		for number := 0;number < 12;number++ { // 숫자는 12개(0부터 셈)
			game.Deck = append(game.Deck, Card{color,number})
		}
	}
	for i := range game.Deck { // 덱 섞기
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		game.Deck[i], game.Deck[j] = game.Deck[j], game.Deck[i]
	}
	for i:=0;i < 7;i++ {
		game.Player_1.Add_Card(game.Card_get())
		game.Player_2.Add_Card(game.Card_get())
	}
}
func (game *Game)Game_Control(Control int, Player *User){
	if Control == -1 {
		Player.Add_Card(game.Card_get())
	} else {
		game.Deck = append(game.Deck, game.Turn_card)
		game.Turn_card = Player.Deck[Control]
		Player.Remove_Card(Control)	
	}
}
func main() {
	game := Game{}
	game.Player_1 = &User{Name: "user1"}
	game.Player_2 = &User{Name: "user2"}
	game.Game_set()
	game.Turn_card = game.Card_get()
	for len(game.Player_1.Deck) != 0 || len(game.Player_1.Deck) != 0 {
		if len(game.Player_1.Deck) == 0 {
			break
		}	
		game.Game_Control(game.Player_1.Control(game.Turn_card),game.Player_1)
		if len(game.Player_2.Deck) == 0 {
			break
		}	
		game.Game_Control(game.Player_2.Control(game.Turn_card),game.Player_2)	
	}
	if len(game.Player_1.Deck) == 0 {
		fmt.Println("player 1 won")
	}else {
		fmt.Println("player 2 won")
	}
}