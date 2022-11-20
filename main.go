package main

import "fmt"

func main() {

	makeTeams()
	return

	var home, visitor Team
	home.LoadTeam("san_jose_valley_yak.json")
	visitor.LoadTeam("norfolk_ocean_oysters.json")

	for {
		var game Game
		game.Home = home
		game.Visitor = visitor
		game.PlayBall()
		if game.VHits == 0 || game.HHits == 0 {
			fmt.Println("No Hitter!!")
			return
		}
	}
}
