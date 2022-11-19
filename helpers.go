package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var first, teams, middle, nick, last []string

var err error

func makeTeams() {

	first, err = readLines("firstnames.txt")
	if err != nil {
		log.Fatal(err)
	}

	teams, err = readLines("teamnames.txt")
	if err != nil {
		log.Fatal(err)
	}

	middle, err = readLines("middlenames.txt")
	if err != nil {
		log.Fatal(err)
	}
	nick, err = readLines("nicknames.txt")
	if err != nil {
		log.Fatal(err)
	}
	last, err = readLines("lastnames.txt")
	if err != nil {
		log.Fatal(err)
	}
	//https://www.sbnation.com/2016/4/25/11502620/minor-league-baseball-franchise-generator
	for _, v := range teams {
		team := Team{}
		parts := strings.Split(v, " ")
		max := len(parts)
		team.Mascot = parts[max-2] + " " + parts[max-1]
		if max == 4 {
			team.City = parts[0] + " " + parts[1]
		} else {
			team.City = parts[0]
		}

		for i := 0; i < 16; i++ {
			team.AddPlayer()
		}
		data, _ := json.Marshal(team)
		err = os.WriteFile("teams/"+strings.ToLower(strings.ReplaceAll(v, " ", "_"))+".json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
