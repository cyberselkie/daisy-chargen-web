package main

import (
	"daisychainsaw-web/gen"
	"fmt"
	"net/http"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func webChar() string {
	p := &gen.Profile{
		W1:     gen.RandWeapon(),
		W2:     gen.RandWeapon(),
		Pri:    "",
		Sec:    "",
		Charm:  1,
		Heart:  1,
		Focus:  1,
		Power:  1,
		Spells: 0,
		Speed:  0,
		HP:     0,
	}
	// primary and secondary stats
	p.Pri = gen.Primary(p)
	p.Sec = gen.Secondary(p)

	// list random quirks & traits
	quirks := gen.Quirks(p)
	traits := gen.Traits(p)

	gen.Stats(6, p)

	ability := gen.Abilities(2, p)
	spells := gen.Spells(p)

	//divider := "\n ------ \n"
	weapons := "Main Hand: " + p.W1 + " | " + "Offhand: " + p.W2
	statPrio := "Primary: " + p.Pri + " | " + "Secondary: " + p.Sec
	charm := gen.StylePlainBorder.Render("Charm: " + strconv.Itoa(p.Charm))
	heart := gen.StylePlainBorder.Render("Heart: " + strconv.Itoa(p.Heart))
	focus := gen.StylePlainBorder.Render("Focus: " + strconv.Itoa(p.Focus))
	power := gen.StylePlainBorder.Render("Power: " + strconv.Itoa(p.Power))

	hp := gen.StylePlainBorder.Render("HP: " + strconv.Itoa(p.HP))
	speed := gen.StylePlainBorder.Render("Speed: " + strconv.Itoa(p.Speed))
	spn := gen.StylePlainBorder.Render("Spells: " + strconv.Itoa(p.Spells))

	tr := "Traits\n" + traits
	qu := "Quirks\n" + quirks

	ab := "Abilities\n" + ability
	spdis := "Spells\n" + spells

	if p.Spells != 0 {

	}
	var statBlock = lipgloss.JoinHorizontal(lipgloss.Center, charm, heart, focus, power)
	var primInfo = lipgloss.JoinVertical(lipgloss.Center, weapons, statPrio, statBlock, lipgloss.JoinHorizontal(lipgloss.Center, hp, speed, spn, "\n"))
	var moreInfo = lipgloss.JoinVertical(lipgloss.Left, tr, qu, ab, spdis)

	var charSheet = gen.StyleCuteBorder.Render(lipgloss.JoinVertical(lipgloss.Center, primInfo, moreInfo))
	block := lipgloss.PlaceHorizontal(80, lipgloss.Center, charSheet)

	return block
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, webChar())
	})

	http.ListenAndServe(":9990", nil)
}
