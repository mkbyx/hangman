package hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

func Jeu() {
	WelcomePlayer()
	afficherPendu()
	motAlea()
	mot = strings.ToLower(mots[rand.Intn(200)])
	for i := 0; i < len(mot); i++ {
		tab = append(tab, "_")
		tab = append(tab, " ")
		mot2 += "_ "
	}
	fmt.Println(mot2)
	for {
		win = true
		fmt.Scan(&test)
		if len(test) != 1 && test != mot {
			pv--
			PrintNext(pv, test)
		}
		if len(test) != 1 && test == mot {
			PrintWin(mot)
			break
		}
		for _, i := range mot {
			if string(i) == test {
				tab[cmt] = test
				estla = true
			}
			cmt += 2
		}
		cmt = 0
		if !estla && len(test) == 1 {
			pv--
			PrintNext(pv, test)
		}
		if pv <= 0 {
			PrintLose(mot)
			break
		}
		estla = false
		mot2 = ""
		for i := 0; i < len(tab); i++ {
			mot2 += tab[i]
		}
		fmt.Println(mot2)
		for i := 0; i < len(tab); i++ {
			if tab[i] == "_" {
				win = false
			}
		}
		if pv != 10 {
			fmt.Printf(pendu[9-pv])
		}
		if win {
			PrintWin(mot)
			break
		}
	}
}
