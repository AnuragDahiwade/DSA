package main

import "fmt"

func IsOperand(char byte) bool {
	if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char <= 9 {
		return true
	}
	return false
}

func IsOpeningBracket(char byte) bool {
	if char == '(' || char == '{' || char == '[' {
		return true
	}
	return false
}

func IsClosingBracket(char byte) bool {
	if char == ')' || char == '}' || char == ']' {
		return true
	}
	return false
}

func getPriority(char byte) int {
	if char == '^' {
		return 3
	} else if char == '*' || char == '/' {
		return 2
	} else if char == '+' || char == '-' {
		return 1
	} else {
		return -1
	}
}

func InfixToPostfix(str string) string {
	res := ""
	stack := Stacck[byte]{}

	for i := 0; i < len(str); i++ {
		char := str[i]

		if IsOperand(char) {
			res += string(char)
		} else if IsOpeningBracket(char) {
			stack.push(char)
		} else if IsClosingBracket(char) {
			for !stack.IsEmpty() {
				if ch, ok := stack.peek(); ok && !IsOpeningBracket(ch) {
					res += string(ch)
					stack.pop()
				} else {
					break
				}
			}
			stack.pop()
		} else {

			for !stack.IsEmpty() {
				if ch, ok := stack.peek(); ok && getPriority(char) <= getPriority(ch) {
					res += string(ch)
					stack.pop()
				} else {
					break
				}

			}
			stack.push(char)
		}

	}
	for !stack.IsEmpty() {
		if ch, ok := stack.peek(); ok {
			res += string(ch)
			stack.pop()
		} else {
			break
		}
	}

	return res
}

func conversionMain() {
	str := `K+L-M*N+(O^P)*W/U/V*T+Q`
	fmt.Println("Infix : ", str)
	fmt.Println("Postfix : ", InfixToPostfix(str))
}
