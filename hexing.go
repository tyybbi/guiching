package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Hexagram struct {
	Id    int       `json:"id"`
	Lines [6]string `json:"lines"`
	Name  string    `json:"name"`
	Desc  string    `json:"desc"`
}

type Hexagrams struct {
	Hexagrams []Hexagram `json:"hexagrams"`
}

func findHexagram(a [6]string, b [6]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func importHexagramData() Hexagrams {
	var hexagrams Hexagrams

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &hexagrams)

	return hexagrams
}

func generateHexagram() [6]string {
	var freshHexagram [6]string

	marbles := [16]string{
		"--- X ---",
		"----O----", "----O----", "----O----",
		"---------", "---------", "---------", "---------", "---------",
		"---   ---", "---   ---", "---   ---", "---   ---", "---   ---",
		"---   ---", "---   ---"}

	for i := 0; i < len(freshHexagram); i++ {
		line := rand.Intn(len(marbles))
		freshHexagram[i] = marbles[line]
	}

	return freshHexagram
}

func createShapes(w http.ResponseWriter, r *http.Request) {
	h := importHexagramData()
	rand.Seed(time.Now().UnixNano())

	var relating bool = false
	var initialHxgrm, primaryShape, relatingShape [6]string
	pageVars := PageVars{}

	phex := Hexagram{}
	rhex := Hexagram{}

	initialHxgrm = generateHexagram()

	for i := 0; i < len(initialHxgrm); i++ {
		if initialHxgrm[i] == "--- X ---" {
			primaryShape[i] = "---   ---"
			relatingShape[i] = "---------"
			relating = true
		} else if initialHxgrm[i] == "----O----" {
			primaryShape[i] = "---------"
			relatingShape[i] = "---   ---"
			relating = true
		} else if initialHxgrm[i] == "---------" {
			primaryShape[i] = "---------"
			relatingShape[i] = "---------"
		} else if initialHxgrm[i] == "---   ---" {
			primaryShape[i] = "---   ---"
			relatingShape[i] = "---   ---"
		}
	}

	for match := true; match; match = false {
		for i := 0; i < len(h.Hexagrams); i++ {
			match := findHexagram(primaryShape, h.Hexagrams[i].Lines)
			if match {
				phex.Id = h.Hexagrams[i].Id
				phex.Name = h.Hexagrams[i].Name
				if relating {
					phex.Lines = initialHxgrm
				} else {
					phex.Lines = h.Hexagrams[i].Lines
				}
				phex.Desc = h.Hexagrams[i].Desc
				break
			}
		}
	}

	for i := 0; i < len(phex.Lines); i++ {
		phex.Lines[i] = strings.Replace(phex.Lines[i], " ", "\u00a0", -1)
	}

	if relating {
		for match := true; match; match = false {
			for i := 0; i < len(h.Hexagrams); i++ {
				match := findHexagram(relatingShape, h.Hexagrams[i].Lines)
				if match {
					rhex.Id = h.Hexagrams[i].Id
					rhex.Name = h.Hexagrams[i].Name
					rhex.Lines = h.Hexagrams[i].Lines
					rhex.Desc = h.Hexagrams[i].Desc
					break
				}
			}
		}
		for i := 0; i < len(rhex.Lines); i++ {
			rhex.Lines[i] = strings.Replace(rhex.Lines[i], " ", "\u00a0", -1)
		}
	}

	pageVars = PageVars{
		PId:    phex.Id,
		RId:    rhex.Id,
		PName:  phex.Name,
		RName:  rhex.Name,
		PLines: phex.Lines,
		RLines: rhex.Lines,
		PDesc:  phex.Desc,
		RDesc:  rhex.Desc,
		Rel:    relating,
	}

	render(w, "index.html", pageVars)
}
