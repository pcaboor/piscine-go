package piscine

import("github.com/01-edu/z01")

// boucle forte

func PrintComb(){	
	for x := '0'; x<= '9'; x++ {
      for y := x+1; y<= '9'; y++ {
         for z := y+1; z<= '9'; z++ {
		z01.PrintRune(x) // afficher un nombre de 0 à 9
        z01.PrintRune(y) // afficher un nombre de 0 à 9
        z01.PrintRune(z) // afficher un nombre de 0 à 9    
        z01.PrintRune(44) // afficher des virgules
        z01.PrintRune(32) // afficher des espace 
	}}}
 z01.PrintRune('\n') // afficher en ligne 
}
