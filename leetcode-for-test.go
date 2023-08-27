package main

import "fmt"

type stack struct {
	vals []rune
}

func (st *stack) remove(char rune) bool {
	if len(st.vals) == 0 {
		return false
	}

	if char == ')' {
		if st.vals[len(st.vals)-1] != '(' {
			return false
		}
	} else if char == '}' {
		if st.vals[len(st.vals)-1] != '{' {
			return false
		}

	} else if char == ']' {
		if st.vals[len(st.vals)-1] != '[' {
			return false
		}
	}

	st.vals = st.vals[:len(st.vals)-1]

	return true
}

func isValid(s string) bool {
	if (len(s) % 2) != 0 {
		return false
	}

	if len(s) == 2 {
		return (s[1]-s[0] < 3) && (s[1]-s[0] > 0)
	}

	st := stack{
		vals: []rune{},
	}

	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			st.vals = append(st.vals, char)
		} else {
			if ok := st.remove(char); !ok {
				return false
			}
		}

	}

	return len(st.vals) == 0
}

func main() {
	fmt.Println("Enter a string of brackets")
	var userInput string
	fmt.Scanln(&userInput)
	fmt.Println(isValid(userInput))
}
