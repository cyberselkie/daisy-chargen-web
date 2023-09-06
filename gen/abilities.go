package gen

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ajson, aerr = os.ReadFile("gen/data/abilities.json")

func flipWeapon(p *Profile) string {
	var s = []string{p.W1, p.W2}
	n := randNum(2)

	return s[n]
}

func abilitySelect(p *Profile) string {
	var ability string
	w := flipWeapon(p)
	tier := "t1"

	//a := unmarshal(ajson, fmt.Sprintf("%s.%s.#", w, tier)).Int()
	i := randNum(int(3))

	abilityN := unmarshal(ajson, fmt.Sprintf("%s.%s.%s.name", w, tier, strconv.Itoa(i))).Str
	abilityP := unmarshal(ajson, fmt.Sprintf("%s.%s.%s.prose", w, tier, strconv.Itoa(i))).Str

	ability = fmt.Sprintf("%s: %s\n\n", abilityN, abilityP)
	return ability
}

// listing the abilities as a string
func Abilities(amt int, p *Profile) string {
	var ability string

	for i := 0; i < amt; i++ {
		add := abilitySelect(p)
		if strings.Contains(ability, add) {
			i--
		} else {
			ability += add
		}
	}
	return ability
}
