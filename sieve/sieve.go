package sieve

func Sieve(limit int) []int {
	initList, myPrimes := generateSliceFromZeroToN(limit)
	finalList := []int{}

	for i := 0; i <= len(initList)-1; i++ {
		if i == 0 || i == 1 {
			myPrimes[i] = Res{assigned: true, value: false}
			continue
		}
		if i == 2 {
			myPrimes[i] = Res{assigned: true, value: true}
		}

		if myPrimes[i].value {

			for o := i + 1; o <= len(initList)-1; o++ {

				if initList[o]%initList[i] == 0 {
					myPrimes[o] = Res{assigned: true, value: false}
				} else if myPrimes[o].assigned && !myPrimes[o].value && (initList[o]%initList[i] == 0) {
					myPrimes[o] = Res{assigned: true, value: true}
				} else if !myPrimes[o].assigned {
					myPrimes[o] = Res{assigned: true, value: true}
				}

			}
		}

	}
	for i, v := range myPrimes {
		if v.value {
			finalList = append(finalList, initList[i])
		}
	}
	return finalList
}

type Res struct {
	assigned bool
	value    bool
}

func generateSliceFromZeroToN(n int) ([]int, []Res) {
	result := make([]int, n+1)

	boolRes := make([]Res, n+1)

	for i := 0; i <= n; i++ {
		result[i] = i
	}

	return result, boolRes
}
