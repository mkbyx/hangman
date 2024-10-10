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
	randLetter = rand.Intn(len(word) - 1)
	tab = nil
	pv = 10
	for _, runes := range word {
		wordtab = append(wordtab, string(runes))
	}
	for i := 0; i < len(word)-1; i++ {
		if i == randLetter {
			tab = append(tab, wordtab[i])
			tab = append(tab, " ")
			wordfind += wordtab[i]
			wordfind += " "
		}
		tab = append(tab, "_")
		tab = append(tab, " ")
		wordfind += "_ "
	}
	fmt.Println(wordfind)
	fmt.Println(tab)
	for {
		win = true
		lose = false
		fmt.Scan(&input)
		if len(input) != 1 && input != word {
			pv -= 2
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
			displayGameOver()
			break

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
			displayVictory()
			break
		}
		if lose {
			break
		}

	}
}
