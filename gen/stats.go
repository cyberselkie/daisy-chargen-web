package gen

import (
	"github.com/mroth/weightedrand/v2"
)

var STATS = []string{"Charm", "Heart", "Focus", "Power"}

func stat2Int(s string, p *Profile) int {
	stat := 0
	switch s {
	case "Charm":
		stat = p.Charm
	case "Heart":
		stat = p.Heart
	case "Focus":
		stat = p.Focus
	case "Power":
		stat = p.Power
	}
	return stat
}

func statGen(max int, p *Profile) {
	chooser, _ := weightedrand.NewChooser(
		weightedrand.NewChoice(p.Pri, 3),
		weightedrand.NewChoice(p.Sec, 2),
		weightedrand.NewChoice("Charm", 1),
		weightedrand.NewChoice("Heart", 1),
		weightedrand.NewChoice("Focus", 1),
		weightedrand.NewChoice("Power", 1),
	)
	for i := 0; i < max; i++ {
		stat := chooser.Pick()
		switch stat {
		case "Charm":
			if p.Charm == 6 {
				i -= 1
				continue
			} else {
				p.Charm += 1
			}
		case "Heart":
			if p.Heart == 6 {
				i -= 1
				continue
			} else {
				p.Heart += 1
			}
		case "Focus":
			if p.Focus == 6 {
				i -= 1
				continue
			} else {
				p.Focus += 1
			}
		case "Power":
			if p.Power == 6 {
				i -= 1
				continue
			} else {
				p.Power += 1
			}
		}
	}
}

func derivedStats(p *Profile) {
	p.Spells = p.Charm - 1
	p.Speed = p.Focus + stat2Int(p.Pri, p)
	p.HP = p.Heart * 3
}

func Stats(max int, p *Profile) {
	statGen(max, p)
	derivedStats(p)
}
