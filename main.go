package main
import "fmt"


func welcomePlayer() {
	asciiArt := `	 _   _    _    _   _  ____ __  __    _    _   _`
	asciiArt2 := `	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |`
	asciiArt3 := `	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |`
	asciiArt4 := `	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |`
	asciiArt5 := `	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|`
	var Green = "\033[32m"
	var Reset = "\033[0m"
	fmt.Println(Green + asciiArt + Reset)
	var Red = "\033[31m"
	fmt.Println(Red + asciiArt2 + Reset)
	var Magenta = "\033[35m"
	fmt.Println(Magenta + asciiArt3 + Reset)
	var Yellow = "\033[33m"
	fmt.Println(Yellow + asciiArt4 + Reset)
	var Blue = "\033[34m"
	fmt.Println(Blue + asciiArt5 + Reset)
}
func displayGameOver() {
	asciiArt := `	__     ______  _    _   _      ____   _____ _______ `
	asciiArt2 := `	\ \   / / __ \| |  | | | |    / __ \ / ____|__   __|`
	asciiArt3 := `	 \ \_/ / |  | | |  | | | |   | |  | | (___    | |   `
	asciiArt4 := `	  \   /| |  | | |  | | | |   | |  | |\___ \   | |   `
	asciiArt5 := `	   | | | |__| | |__| | | |___| |__| |____) |  | |   `
	asciiArt6 := `	   |_|  \____/ \____/  |______\____/|_____/   |_|   `
	var Red = "\033[31m"
	var Reset = "\033[0m"
	fmt.Println(Red + asciiArt + Reset)
	fmt.Println(Red + asciiArt2 + Reset)
	fmt.Println(Red + asciiArt3 + Reset)
	fmt.Println(Red + asciiArt4 + Reset)
	fmt.Println(Red + asciiArt5 + Reset)
	fmt.Println(Red + asciiArt6 + Reset)
}
func main() {
	var mot string
	mot = "hello"
	fmt.Println("_ _ _ _ _ ")
	fmt.Println(mot)
	fmt.Scan(&mot)
	fmt.Println(mot)
	var test string
	var mot2 string
	var pv int = 10
	var cmt int
	var estla bool = false
	var win bool
	var tab []string
	welcomePlayer()
	mot = "Mahan"
	for i := 0; i < len(mot); i++ {
		tab = append(tab, "_")
		tab = append(tab, " ")
		mot2 += "_ "
	}
	fmt.Println(mot2)
	// Main game loop
	for {
		win = true
		fmt.Scan(&test)
		if len(test) != 1 && test != mot {
			pv--
			fmt.Printf("❤️ il vous reste %d chances\n", pv)
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
			fmt.Printf("❤️ il vous reste %d chances\n", pv)
		}
		if pv <= 0 {
			displayGameOver()
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
			var Green = "\033[32m"
			var Reset = "\033[0m"
			fmt.Println(Green + "You Won Congrats!" + Reset)
			break
		}
	}
}
