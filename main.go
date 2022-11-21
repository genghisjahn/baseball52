package main

import "fmt"

func main() {

	// makeTeams()
	// return

	var home, visitor Team
	home.LoadTeam("san_jose_valley_yaks.json")
	visitor.LoadTeam("mt._airy_fightin_hawks.json")

	// for _, v := range visitor.Players {
	// 	fmt.Println(v.AnnouceName())
	// }
	// return

	// for {
	var game Game
	game.Home = home
	game.Visitor = visitor
	game.PlayBall()
	if game.VHits == 0 || game.HHits == 0 {
		fmt.Println("No Hitter!!")
		return
	}
	// }
}
