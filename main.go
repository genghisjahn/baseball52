package main

import "fmt"

func main() {
	var home, visitor Team
	home.LoadTeam("san_jose_valley_yak.json")
	visitor.LoadTeam("norfolk_ocean_oysters.json")

	for {
		var game Game
		game.Home = home
		game.Visitor = visitor
		game.PlayBall()
		if game.VHits > 5 || game.HHits > 5 {
			fmt.Println("5 Hitter!!")
		}
	}
}
