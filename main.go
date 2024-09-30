package main

import (
	"fmt"
)

func PrintWin(str string) {
	fmt.Printf("GG tu as trouvé le mot : %s \n", str)
}

func PrintLose(str string) {
	fmt.Printf("vous avez perdu\nLe mot était %s.", str)
}

func PrintNext(n int, str string) {
	if len(str) == 1 {
		fmt.Printf("\nla lettre %s n est pas dans le mot\n", str)
	}
	if len(str) != 1 {
		fmt.Printf("\nCe n'est pas le mot :  %s\n", str)
	}
	fmt.Printf("il vous reste %d chances\n", n)
}

func main() {
	var mot string
	var test string
	var mot2 string
	var pv int = 10
	var cmt int
	var estla bool = false
	var win bool
	var tab []string

	mot = "hello"
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
		if win {
			PrintWin(mot)
			break
		}
	}
}
