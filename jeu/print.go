package hangman

import "fmt"

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
