package main

import "fmt"

func main() {
	testRun("a", "a", true)
	testRun("a+", "aaa", true)
	testRun("ba+", "baa", true)
	testRun("ba+", "baab", false)
	testRun("ba+b", "baab", true)
	testRun("ba+b", "baa", false)
	testRun("ba+b+", "baabbb", true)
	testRun("ba+b+a", "baabbb", false)
	testRun("ba+b+a", "baabbba", true)

	// fmt.Println(splitExpToParts("a"))
	// fmt.Println(splitExpToParts("a+"))
	// fmt.Println(splitExpToParts("ba+"))
	// fmt.Println(splitExpToParts("ba+b"))

	// fmt.Println(isPlusExp("b"))
	// fmt.Println(isPlusExp("b+"))

}

func testRun(exp, word string, want bool) {
	got := isMatch(exp, word)
	if got != want {
		fmt.Printf("**** ERROR **** %s, %s, %t, %t\n", exp, word, want, got)
	}
}

func isMatch(exp, word string) bool {
	return isMatchHelper([]rune(exp), []rune(word), false)
}

func isMatchHelper(exp, word []rune, plusSatisfied bool) bool {
	fmt.Printf("*** exp=%s word=%s\n", string(exp), string(word))

	// Exp is empty
	if len(exp) == 0 {
		if len(word) == 0 {
			return true
		}
		return false
	}

	var nextExp = exp[0]
	var isPlus = len(exp) > 1 && exp[1] == '+'

	if len(word) == 0 {
		if !isPlus {
			return false
		}
		if !plusSatisfied {
			return false
		}
		return isMatchHelper(exp[2:], word, false)
	}

	var nextChar = word[0]

	// Unequal characters
	if nextChar != nextExp {
		if !isPlus {
			return false
		}
		if !plusSatisfied {
			return false
		}

		return isMatchHelper(exp[2:], word, false)

	}

	// Equal Characters
	if isPlus {
		return isMatchHelper(exp, word[1:], true)
	}

	return isMatchHelper(exp[1:], word[1:], false)

}
