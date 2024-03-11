package piscine

func FindNextPrime(nb int) int {
	/*for {
		prime := true
		if nb <= 1 {
			nb++
			continue
		}
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return nb
		}
		nb++
	}*/
	for {
		if nb <= 1 {
			return 0
		}
		prime := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return nb
		}
		nb--
	}
}
