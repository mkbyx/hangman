package hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

func Menu() {
	WelcomePlayer()
	fmt.Println("Bienvenu")
	for {
		fmt.Println("Pour jouer, écrire [play]\nPour quitter, écrire [quit]")
		fmt.Scan(&play)
		if play == "play" {
			Game()
			continue
		}
		if play == "quit" {
			break
		}
		fmt.Println("commande invalide")
	}
}

func Game() {
	hangman = nil
	printHangman()
	randomWord()
	word = strings.ToLower(words[rand.Intn(200)])
	wordfind = ""

	tab = nil
	pv = 10
	for i := 0; i < len(word); i++ {
		tab = append(tab, "_")
		tab = append(tab, " ")
		wordfind += "_ "
	}
	fmt.Println(wordfind)
	for {
		win = true
		lose = false
		fmt.Scan(&input)
		if len(input) != 1 && input != word {
			pv--
			PrintNext(pv, input)
		}
		if len(input) != 1 && input == word {
			PrintWin(word)
			break
		}
		for _, i := range word {
			if string(i) == input {
				tab[count] = input
				ishere = true
			}
			count += 2
		}
		count = 0
		if !ishere && len(input) == 1 {
			pv--
			PrintNext(pv, input)
		}
		if pv <= 0 {
			PrintLose(word)
			lose = true

		}
		ishere = false
		wordfind = ""
		for i := 0; i < len(tab); i++ {
			wordfind += tab[i]
		}
		fmt.Println(wordfind)
		for i := 0; i < len(tab); i++ {
			if tab[i] == "_" {
				win = false
			}
		}
		if pv != 10 {
			fmt.Printf(hangman[9-pv])
		}
		if win {
			PrintWin(word)
			break
		}
		if lose {
			break
		}

	}
}
