package piscine
<<<<<<< HEAD
  
func Atoi(s string) int {
      myResult := 0 
      ponc := 1
  
for a, b := range s {
      if a == 0 && (b == '+' || b == '-') {
          if b == '-' {
               ponc = -1
           }
      } else if b < '0' || b > '9' {
        return 0
      } else {
        
         myResult = myResult*10 + int(b-'0')
      }  
    
    } 
    return myResult * ponc
  }
=======

func Atoi(s string) int {
	myResult := 0
	ponctuation := 1
	for a, numb := range s {
		if a == 0 && (numb == '+' || numb == '-') {
			if numb == '-' {
				ponctuation = -1
			}
		} else if numb < '0' || numb > '9' {
			return 0
		} else {

			myResult = myResult*10 + int(numb-'0')
		}

	}
	return myResult * ponctuation
}

// on peut utiliser continue si on veut simplifier le code à la place de tous les "else"
>>>>>>> 7c8927e (added)
