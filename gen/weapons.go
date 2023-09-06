package gen

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/yektadev/rango"
)

var (
	WEAPONS   = []string{"Baseball Bat", "Brass Knuckles", "Chainsaw", "Gun", "Knives", "Hat", "Katana", "Knives", "Microphone", "Roller Skates", "Whip"}
	json, err = os.ReadFile("gen/data/weapons.json")
)

func randNum(i int) int {
	w := rango.RnInt(0, i)
	return w
}

// random selection
func randomSelect(s []string) string {
	w := randNum(len(s))
	return s[w]
}

// random weapon
func RandWeapon() string {
	return randomSelect(WEAPONS)
}

func unmarshal(json []byte, s string) gjson.Result {
	return gjson.GetBytes(json, s)
}

// finding primary statistic
func Primary(p *Profile) string {
	return unmarshal(json, fmt.Sprint(p.W1, ".primary")).Str
}

// finding secondary statistic
func Secondary(p *Profile) string {
	// secondary statistic check
	secondary := unmarshal(json, fmt.Sprint(p.W2, ".primary")).Str

	if Primary(p) == unmarshal(json, fmt.Sprint(p.W2, ".primary")).Str {
		secondary = unmarshal(json, fmt.Sprint(p.W2, ".secondary")).Str
	}
	return secondary
}

// checking if both weapons are the same
func (p Profile) WeapCompare() bool {
	// same weapon check
	var same bool = false
	if p.W1 == p.W2 {
		same = true
	}
	return same
}

// listing the traits as a string
func Traits(p *Profile) string {
	// traits
	var traits string
	var i int64

	t1 := unmarshal(json, fmt.Sprint(p.W1, ".traits.#")).Int()
	for i = 0; i < t1; i++ {
		name := unmarshal(json, fmt.Sprint(p.W1, ".traits.", i, ".name")).Str
		prose := unmarshal(json, fmt.Sprint(p.W1, ".traits.", i, ".prose")).Str
		traits += fmt.Sprintf("%s: %s\n\n", name, prose)
	}
	if p.WeapCompare() == false { // if you have two different weapons
		t2 := unmarshal(json, fmt.Sprint(p.W2, ".traits.#")).Int()
		for i = 0; i < t2; i++ {
			name := unmarshal(json, fmt.Sprint(p.W2, ".traits.", i, ".name")).Str
			prose := unmarshal(json, fmt.Sprint(p.W2, ".traits.", i, ".prose")).Str
			traits += fmt.Sprintf("%s: %s\n\n", name, prose)
		}
	}
	return traits
}

// selecting random quirk from W1, and 2 random quirks if same weapon
func Quirks(p *Profile) string {
	// quirks
	q := unmarshal(json, fmt.Sprint(p.W1, ".quirks")).Str
	fmt.Print(q)
	// selecting the random quirk
	i := randNum(4)
	quirkN := unmarshal(json, fmt.Sprintf("%s.quirks.%s.name", p.W1, strconv.Itoa(i))).Str
	quirkP := unmarshal(json, fmt.Sprintf("%s.quirks.%s.prose", p.W1, strconv.Itoa(i))).Str

	quirk := fmt.Sprintf("%s: %s", quirkN, quirkP)

	// if the weapon is chosen twice
	if p.WeapCompare() == true {
		i2 := randNum(4)
		for i == i2 {
			i2 = randNum(4)
		}

		quirkN2 := unmarshal(json, fmt.Sprintf("%s.quirks.%s.name", p.W1, strconv.Itoa(i2))).Str
		quirkP2 := unmarshal(json, fmt.Sprintf("%s.quirks.%s.prose", p.W1, strconv.Itoa(i2))).Str

		quirk += fmt.Sprintf("\n%s: %s", quirkN2, quirkP2)
	}
	return BasicPadding.Render(quirk)
}
