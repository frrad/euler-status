package main

const prizes = 5

var names = [prizes]string{
	"Prime Obsession",
	"Triangle Trophy",
	"Lucky Luke",
	"Decimation II",
	"Ultimate Decimator",
}

var taglines = [prizes]string{
	"prime numbered problems",
	"first triangle numbered problems",
	"lucky numbered problems",
	"rows",
	"rows",
}

var goals = [prizes]int{
	50,
	25,
	50,
	10,
	10,
}

var prizeFns [](func(map[int]bool) (int, map[int]bool)) = [](func(map[int]bool) (int, map[int]bool)){
	//PRIME OBSESSION (Index = 0)
	func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		for i := 1; i <= MAX; i++ {
			if dict[i] {
				if isPrime(i) {
					ans++
				}
			} else if isPrime(i) {
				set[i] = true
			}
		}
		return ans, set
	},

	//TRIANGLE NUMBERS (Index = 1)
	func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		for i := 1; i <= 25; i++ {
			if dict[i*(i+1)/2] {
				ans++
			} else {
				set[i*(i+1)/2] = true
			}
		}
		return
	},

	//LUCKY NUMBER (Index = 2)
	func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)

		seive := luckySeive(MAX)
		for i := 0; i < len(seive); i++ {
			if dict[seive[i]] {
				ans++
			} else {
				set[seive[i]] = true
			}
		}
		return
	},

	//DECIMATION II (Index = 3)
	func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		decStart := 200
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	},

	//ULTIMATE DECIMATOR (Index = 4)
	func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		decStart := 300
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	},
}

func luckySeive(max int) []int {

	luckyseive := make([]int, max)
	for i := 0; i < max; i++ {
		luckyseive[i] = i + 1
	}

	last := -1
	pointer := 1

	for pointer < len(luckyseive) {

		last = luckyseive[pointer]

		for del := last - 1; del < len(luckyseive); del += last {
			luckyseive[del] = 0
		}
		for i := 0; i < len(luckyseive); i++ {
			if luckyseive[i] == 0 {
				luckyseive = append(luckyseive[:i], luckyseive[i+1:]...)
				i--
			}

		}

		if luckyseive[pointer] == last {
			pointer++
		}

	}

	return luckyseive
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
