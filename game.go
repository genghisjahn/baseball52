package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Game struct {
	Home    Team     `json:"home_team"`
	Visitor Team     `json:"visitor_team"`
	Innings []Inning `json:"innings"`
	VHits   int
	HHits   int
}

func (g *Game) BoxScore() string {
	vname := g.Visitor.City
	hname := g.Home.City
	s := len(vname) - len(hname)
	if s > 0 {
		hname += strings.Repeat(" ", s)
	} else {
		vname += strings.Repeat(" ", s*-1)
	}
	vscore := 0
	hscore := 0
	vhits := 0
	hhits := 0
	vval, hval := "", ""
	for _, v := range g.Innings {

		vval += " " + fmt.Sprintf("%v", v.Top.Runs)
		hval += " " + fmt.Sprintf("%v", v.Bottom.Runs)
		vscore += v.Top.Runs
		vhits += v.Top.Hits
		hscore += v.Bottom.Runs
		hhits += v.Bottom.Hits

	}

	result := ""
	result = vname + " | " + vval + "|" + "Runs:" + fmt.Sprintf("%v", vscore) + "|Hits: " + fmt.Sprintf("%v", vhits) + "\n"
	result += hname + " | " + hval + "|" + "Runs:" + fmt.Sprintf("%v", hscore) + "|Hits: " + fmt.Sprintf("%v", hhits) + "\n"
	g.HHits = hhits
	g.VHits = vhits
	return result
}

type Inning struct {
	Number int        `json:"number"`
	Top    InningPart `json:"top"`
	Bottom InningPart `json:"bottom"`
}

type Play struct {
	Description string  `json:"description"`
	AtBat       *Player `json:"at_bat"`
	FirstBase   *Player `json:"first,omitempty"`
	SecondBase  *Player `json:"second,omitempty"`
	Thirdbase   *Player `json:"third,omitempty"`
}

func (p Play) Diamond() string {
	t := "At bat, playing %s...number %d %s!\nFB:%d %s\nSB:%d %s\nTB:%d %s\n"

	var aname, apos, fname, sname, tname string
	var anum, fnum, snum, tnum int
	if p.AtBat != nil {
		anum, aname, apos = p.AtBat.Number, p.AtBat.FullName(), p.AtBat.Position
	}
	if p.FirstBase != nil {
		fnum, fname = p.FirstBase.Number, p.FirstBase.LastName
	}
	if p.SecondBase != nil {
		snum, sname = p.SecondBase.Number, p.SecondBase.LastName
	}
	if p.Thirdbase != nil {
		tnum, tname = p.Thirdbase.Number, p.Thirdbase.LastName
	}

	return fmt.Sprintf(t, apos, anum, aname, fnum, fname, snum, sname, tnum, tname)
}

func (g *Game) PlayBall() {

	gameover := false

	for !gameover {
		i := Inning{}
		i.Number = len(g.Innings) + 1

		fmt.Println("Top of inning " + fmt.Sprintf("%v", i.Number) + ". The " + g.Visitor.City + " " + g.Visitor.Mascot + " are up.")
		for i.Top.Outs < 3 {
			inningProcess(&i.Top, &g.Visitor)
		}
		if len(g.Innings) >= 8 {
			vscore := 0
			hscore := 0
			for _, v := range g.Innings {
				vscore += v.Top.Runs
				hscore += v.Bottom.Runs
			}
			if vscore != hscore && hscore > vscore {
				gameover = true
				showFinal(g.Visitor.Mascot, g.Home.Mascot, vscore, hscore)
				//g.ShowPlays()
			}
		}
		if !gameover {

			fmt.Println("Bottom of inning " + fmt.Sprintf("%v", i.Number) + ". The " + g.Home.City + " " + g.Home.Mascot + " are up.")
			for i.Bottom.Outs < 3 {
				inningProcess(&i.Bottom, &g.Home)
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Println(g.BoxScore())
			reader.ReadString('\n')

			if len(g.Innings) >= 8 {
				vscore := 0
				hscore := 0
				for _, v := range g.Innings {
					vscore += v.Top.Runs
					hscore += v.Bottom.Runs
				}
				if vscore != hscore {
					gameover = true
					showFinal(g.Visitor.Mascot, g.Home.Mascot, vscore, hscore)
				}
			}

		}
		g.Innings = append(g.Innings, i)
		fmt.Println(g.BoxScore())
		//g.ShowPlays()
	}

}
func showFinal(vname, hname string, vscore, hscore int) {
	fmt.Println("Final Score...")
	fmt.Println(vname + " " + fmt.Sprintf("%v", vscore))
	fmt.Println(hname + " " + fmt.Sprintf("%v", hscore))

}

func (g Game) ShowPlays() {
	for _, v := range g.Innings {

		for _, v2 := range v.Top.Plays {
			fmt.Println(v.Number, g.Visitor.City, v2.AtBat.FirstLast(), v2.Description)
		}
		for _, v2 := range v.Bottom.Plays {
			fmt.Println(v.Number, g.Home.City, v2.AtBat.FirstLast(), v2.Description)
		}
	}
}

func inningProcess(ip *InningPart, t *Team) {
	play := Play{}

	if len(ip.Plays) > 0 {
		play = ip.Plays[len(ip.Plays)-1]
	}

	player := t.Players[t.BO%9]
	play.AtBat = &player
	fmt.Println("-------")
	fmt.Println(play.Diamond())

	r := result()
	if r == 0 {
		play.Description = "out"
		ip.Outs++
	}
	if r == 1 {
		play.Description = "single"
		ip.Hits++
		if play.Thirdbase != nil {
			ip.Runs++
			play.Thirdbase = nil
		}
		if play.SecondBase != nil {
			play.Thirdbase = play.SecondBase
			play.SecondBase = nil
		}
		if play.FirstBase != nil {
			play.SecondBase = play.FirstBase
			play.FirstBase = nil
		}
		play.FirstBase = &player
		ip.Plays = append(ip.Plays, play)
	}
	if r == 2 {
		play.Description = "double"
		ip.Hits++
		if play.Thirdbase != nil {
			play.Thirdbase = nil
			ip.Runs++
		}
		if play.SecondBase != nil {
			ip.Runs++
			play.SecondBase = nil
		}
		if play.FirstBase != nil {
			play.FirstBase = play.Thirdbase
			play.FirstBase = nil
		}
		play.SecondBase = &player
		ip.Plays = append(ip.Plays, play)
	}
	if r == 3 {
		play.Description = "triple"
		ip.Hits++
		if play.Thirdbase != nil {
			play.Thirdbase = nil
			ip.Runs++
		}
		if play.SecondBase != nil {
			ip.Runs++
			play.SecondBase = nil
		}
		if play.FirstBase != nil {
			ip.Runs++
			play.FirstBase = nil
		}
		play.Thirdbase = &player
		ip.Plays = append(ip.Plays, play)
	}
	if r == 4 {
		if play.FirstBase != nil && play.SecondBase != nil && play.Thirdbase != nil {
			fmt.Println("Grand Slam!")
		}
		ip.Hits++
		play.Description = "home run"
		if play.Thirdbase != nil {
			play.Thirdbase = nil
			ip.Runs++

		}
		if play.SecondBase != nil {
			ip.Runs++
			play.SecondBase = nil
		}
		if play.FirstBase != nil {
			ip.Runs++
			play.FirstBase = nil
		}
		ip.Runs++
		play.FirstBase = nil
		play.SecondBase = nil
		play.Thirdbase = nil
		ip.Plays = append(ip.Plays, play)
	}
	fmt.Println(play.AtBat.LastName, " ", play.Description)
	fmt.Println(" ")
	play.AtBat = nil
	fmt.Printf("Outs %d,Hits %d, Runs %d\n", ip.Outs, ip.Hits, ip.Runs)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	t.BO++
}

func result() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	c := r1.Intn(52)
	result := 0
	if c >= 37 {
		result = 1
	}
	if c >= 41 {
		result = 2
	}
	if c >= 45 {
		result = 3
	}
	if c >= 49 {
		result = 4
	}
	return result
}
