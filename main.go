package main

func main() {
	var home, visitor Team
	home.LoadTeam("san_jose_valley_yak.json")
	visitor.LoadTeam("norfolk_ocean_oysters.json")

	var game Game
	game.Home = home
	game.Visitor = visitor
	game.PlayBall()
}
