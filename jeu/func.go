package hangman

import (
	"bufio"
	"fmt"
	"os"
)
//Affiche le Nom du jeu avec des jolies couleurs
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

//Affiche la victoie en cas de victoire
func PrintWin(str string) {
	fmt.Printf("GG tu as trouvé le mot : %s \n", str)
}

//Affiche la défaite en cas de défaite
func PrintLose(str string) {
	fmt.Printf("vous avez perdu\nLe mot était %s.", str)
}

//Affiche une erreur en cas de mauvais entrée et le reste d'essaie qu'il reste
func PrintNext(n int, str string) {
	if len(str) == 1 {
		fmt.Printf("\nla lettre %s n est pas dans le mot\n", str)
	}
	if len(str) > 1 {
		fmt.Printf("\nCe n'est pas le mot :  %s\n", str)
	}
	fmt.Printf("il vous reste %d chances\n", n)
}

func PrintTab(){
	fmt.Println(tabletter)
}

//Affiche le Print du Hangman au fur et à mesure du nombre d'erreur qu'on commet
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

//Sélectione un mot aléatoirement dans un fichier texte
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
