package main

import "fmt"

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
			fmt.Printf("il vous reste %d chances\n", pv)
		}
		if len(test) != 1 && test == mot {
			fmt.Printf("GG tu as trouvé le mot : %s \n", mot)
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
		if !estla {
			pv--
			fmt.Printf("\nla lettre %s n est pas dans le mot\n", test)
			fmt.Printf("il vous reste %d chances\n", pv)

		}
		if pv <= 0 {
			fmt.Println("vous avez perdu")
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
			fmt.Println("GG vous avez gagné\n")
			break
		}
	}
}
