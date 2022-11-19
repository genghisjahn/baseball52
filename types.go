package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type InningPart struct {
	Outs  int    `json:"outs"`
	Plays []Play `json:"plays"`
	Hits  int    `json:"hits"`
	Runs  int    `json:"runs"`
}

func (t *Team) AddPlayer() Player {
	if t.numbers == nil {
		t.numbers = make(map[int]string)
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p := Player{}
	for {
		p.Number = r1.Intn(100)
		if _, ok := t.numbers[p.Number]; !ok {
			break
		}
	}
	p.FirstName = first[r1.Intn(len(first))]
	p.MiddleName = first[r1.Intn(len(middle))]
	p.LastName = last[r1.Intn(len(last))]
	p.NickName = ""
	nickcheck := r1.Intn(50)
	if nickcheck == 1 {
		p.NickName = nick[r1.Intn(len(nick))]
	}
	switch len(t.Players) {
	case 0:
		p.Position = "first base"
	case 1:
		p.Position = "second base"
	case 2:
		p.Position = "third base"
	case 3:
		p.Position = "short stop"
	case 4:
		p.Position = "catcher"
	case 5:
		p.Position = "left field"
	case 6:
		p.Position = "center field"
	case 7:
		p.Position = "right field"
	case 8:
		p.Position = "pitcher 1"
		p.IsPitcher = true
	case 9:
		p.Position = "pitcher 2"
		p.IsPitcher = true

	case 10:
		p.Position = "pitcher 3"
		p.IsPitcher = true

	case 11:
		p.Position = "pitcher 4"
		p.IsPitcher = true

	case 12:
		p.Position = "pitcher 5"
		p.IsPitcher = true

	case 13:
		p.Position = "pitcher 6"
		p.IsPitcher = true

	case 14:
		p.Position = "pitcher 7"
		p.IsPitcher = true

	case 15:
		p.Position = "pitcher 8"
		p.IsPitcher = true
	}

	t.Players = append(t.Players, p)
	if nickcheck == 1 {
		fmt.Println(p.FullName())
	}
	return p
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Team struct {
	City    string `json:"city"`
	Mascot  string `json:"mascot"`
	numbers map[int]string
	Players []Player `json:"players"`
	BO      int      `json:"bo"`
}

func (t *Team) LoadTeam(name string) {
	data, err := os.ReadFile("teams/" + name) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	errJ := json.Unmarshal(data, t)
	if errJ != nil {
		log.Fatal(err)
	}
}

type Player struct {
	Number     int    `json:"number"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	NickName   string `json:"nick_name"`
	LastName   string `json:"last_name"`
	Position   string `json:"position"`
	IsPitcher  bool   `json:"is_pitcher"`
}

func (p Player) FullName() string {
	if p.NickName == "" {
		return fmt.Sprintf("%s %s %s", p.FirstName, p.MiddleName, p.LastName)
	}
	return fmt.Sprintf("%s '%s' %s %s", p.FirstName, p.NickName, p.MiddleName, p.LastName)
}

func (p Player) FirstLast() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p Player) AnnouceName() string {
	return fmt.Sprintf("%s, (%v) %s %s ", p.Position, p.Number, p.FirstName, p.LastName)
}
