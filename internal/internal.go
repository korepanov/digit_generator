package internal

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// находит решения
func FindSolutions(s string) ([]string, error) {
	var res []string
	solutions := generateStrings(s)

	for _, solution := range solutions {
		r, err := findResult(lexicalAnalyze(solution))
		if nil != err {
			return res, err
		}
		if 200 == r {
			res = append(res, solution)
		}
	}

	return res, nil
}

// генерирует все возможные строки
func generateStrings(s string) (res []string) {

	variants := permWithRepetitions()
	var newS string
	newS = s

	for _, variant := range variants {
		i := 0
		j := 9

		for {
			newS = newInsert(newS, variant[i], j)
			i++
			j--

			if i > 8 {
				break
			}
		}
		res = append(res, newS)
		newS = s
	}

	return res
}

// находит размещения с повторениями для множества ["+", "-", ""] из 3 по 9
func permWithRepetitions() [][9]string {

	var res [][9]string
	var el [9]string

	var d [9]int

	for i := 0; i < 9; i++ {
		d[i] = 1
	}

	d[8] = 0
	var j int

	j = 9

	for {
		if d[j-1] == 3 {
			j--
			if j == 0 {
				return res
			}
		} else {
			d[j-1]++
			for j < 9 {
				d[j] = 1
				j++
			}

			for i := 0; i < len(d); i++ {
				if 1 == d[i] {
					el[i] = "+"
				} else if 2 == d[i] {
					el[i] = "-"
				} else if 3 == d[i] {
					el[i] = ""
				} else {
					panic("permWithRepetitions: error in the algorithm")
				}
			}

			res = append(res, el)
		}
	}
}

// лексический анализ строки
func lexicalAnalyze(s string) (res [][]string) {
	var numberS string
	var el []string

	for _, ch := range s {
		if unicode.IsDigit(rune(ch)) {
			numberS += string(ch)
		} else {
			el = append(el, "NUM")
			el = append(el, numberS)
			res = append(res, el)
			el = nil
			el = append(el, "OP")
			el = append(el, string(ch))
			res = append(res, el)
			numberS = ""
			el = nil
		}
	}

	el = append(el, "NUM")
	el = append(el, numberS)
	res = append(res, el)

	return res
}

// посчитать результат, на вход берет результат лексического анализа
func findResult(s [][]string) (int, error) {

	if 1 == len(s) {
		return strconv.Atoi(s[0][1])
	}

	leftOperand, err := strconv.Atoi(s[0][1])
	if nil != err {
		return 0, err
	}
	rightOperand, err := strconv.Atoi(s[2][1])
	if nil != err {
		return 0, err
	}

	var numberS string

	if "+" == s[1][1] {
		numberS = strconv.Itoa(leftOperand + rightOperand)
	} else if "-" == s[1][1] {
		numberS = strconv.Itoa(leftOperand - rightOperand)
	} else {
		return 0, errors.New("wrong syntax")
	}

	s = s[3:]
	s = append([][]string{{"NUM", numberS}}, s...)

	return findResult(s)
}

// вставить в строке s операцию op после цифры num
func newInsert(s string, op string, num int) string {

	for i := 0; i < len(s); i++ {
		if fmt.Sprintf("%d", num) == string(s[i]) {
			return s[:i+1] + op + s[i+1:]
		}
	}

	return s
}
