package main

import "fmt"

func main() {
	var mot string
	var test string
	var pv int = 10
	mot = "hello"
	for {
		fmt.Scan(&test)
		if len(test)!=1 && test != mot{
			pv--
			fmt.Printf("il vous reste %d chances\n", pv)
		}
		if len(test)!=1 && test == mot{
			fmt.Printf("GG tu as trouvé le mot : %s \n", mot)
			break
		}
	}
}
