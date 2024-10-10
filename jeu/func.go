package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func WelcomePlayer() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	asciiArt := `
	 _   _    _    _   _  ____ __  __    _    _   _`
	asciiArt2 := `	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |`
	asciiArt3 := `	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |`
	asciiArt4 := `	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |`
	asciiArt5 := `	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|`

	fmt.Println(Magenta + asciiArt + Reset)
	fmt.Println(Red + asciiArt2 + Reset)
	fmt.Println(Green + asciiArt3 + Reset)
	fmt.Println(Yellow + asciiArt4 + Reset)
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

func displayVictory() {
	green := "\033[32m"
	reset := "\033[0m"

	fmt.Println(green + `
 _     ______  _    _  __          _______ _   _ 
 \ \   / / __ \| |  | | \ \        / /_   _| \ | |
  \ \_/ / |  | | |  | |  \ \  /\  / /  | | |  \| |
   \   /| |  | | |  | |   \ \/  \/ /   | | | .   |
    | | | |__| | |__| |    \  /\  /   _| |_| |\  |
    |_|  \____/ \____/      \/  \/   |_____|_| \_|` + reset)
}

func PrintWin(str string) {
	fmt.Printf("GG tu as trouvé le mot : %s \n", str)
}

func PrintLose(str string) {
	fmt.Printf("vous avez perdu\nLe mot était %s.", str)
}
func PrintList(tab string) {
fmt.Printf("la lettre est déjà utiliser \n inserer une lettre")

}
func PrintNext(n int, str string) {
	if len(str) == 1 {
		fmt.Printf("\nla lettre %s n est pas dans le mot\n", str)
	}
	if len(str) != 1 {
		fmt.Printf("\nCe n'est pas le mot :  %s\n", str)
	}
	fmt.Printf("❤️ il vous reste %d chances\n", n)
}

func printHangman() {
	fichier, err := os.Open("data/hangman.txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	i := 0
	j := -1
	for scanner.Scan() {
		if i%8 == 0 {
			hangman = append(hangman, scanner.Text())
			j++
		}
		if i%8 != 0 {
			hangman[j] = hangman[j] + "\n" + scanner.Text()
		}
		i++
	}
	if len(hangman) == 0 {
		fmt.Println("le fichier ne contient rien")
		return
	}
}

func randomWord() {
	fichier, err := os.Open("data/motsimple.txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if len(words) == 0 {
		fmt.Println("le fichier ne contient rien")
		return
	}

}
