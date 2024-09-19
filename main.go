package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RomeToArab = map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100}
var ArabToRome = [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}

func main() {
	//Чтение из консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значение:")
	text, _ := reader.ReadString('\n') //в саму строку вкл-ся ещё сам разделитель (delim)
	//TrimSpace удаляет все начальные и конечные пробелы https://pkg.go.dev/strings#TrimSpace
	text = strings.TrimSpace(text) //в этом случае удаляет сам разделитель (delim) с конца, который яв-ся особенностью использования ReadString, у Scanner такого нет
	//Делим строку на 3 значения – 2 значения (числа) и знак
	SplitText := strings.Split(text, " ")
	if len(SplitText) > 3 || len(SplitText) < 3 {
		panic("Ошибка, формат заданной строки не удовлетворяет заданию – два аргумента и один оператор (+, -, /, *)")
	}

	a := SplitText[0]
	b := SplitText[2]
	firstNumber := 0
	secondNumber := 0

	sign := SplitText[1]
	if !strings.ContainsAny(sign, "+-/*") {
		panic("Ошибка, в качестве математического оператора могут быть использованы только +, -, /, *")
	}
	if strings.ContainsAny(a, "IVX") && strings.ContainsAny(b, "IVX") {
		firstNumber = changeValues(a)
		checkValue(firstNumber, a) //Проверка корректности введенного первого римсского значения
		secondNumber = changeValues(b)
		checkValue(secondNumber, b) //Проверка корректности введеного второго римского значения
		fmt.Println(convArabToRoman(calculate(firstNumber, secondNumber, sign)))
	} else if strings.ContainsAny(a, "0123456789") && strings.ContainsAny(b, "0123456789") {
		firstNumber = changeValues(a)
		secondNumber = changeValues(b)
		fmt.Println(calculate(firstNumber, secondNumber, sign))
	} else {
		panic("Ошибка, используются одновременно разные системы счисления")
	}

}

func changeValues(a string) int {
	Number := 0
	err := error(nil)
	//Если в подстроке содержатся "IVX", то считается число через конвертацию из римских
	if strings.ContainsAny(a, "IVX") {
		Number = convRomanToArab(a)
	} else {
		//Преобразование значений из типа string в int (strconv), сли обычное число (не в римском виде) типа string
		Number, err = strconv.Atoi(a)
	}
	//Ограничение числа в диапазоне от 1 до 10
	if err != nil || Number < 1 || Number > 10 {
		panic("Ошибка, некорректное значение одного из аргументов. Число должно быть целым и в диапазоне от 1 до 10")
	}
	return Number
}

// Проверка на корректность введенного римского аргумента, чтобы не ввели IIII,IIIX,IVI и т.д.
func checkValue(number int, roman string) string {
	var sliceRoman []string
	a := convArabToRoman(number)
	if a != roman {
		panic("Ошибка, некорректное значение одного из аргументов в римском варианте.\nВы ввели: " + roman + "\n" + "Возможно, вы хотели ввести: " + a)
	}
	return strings.Join(sliceRoman, "")
}

func calculate(a, b int, sign string) int {
	result := 0
	if len(sign) > 0 {
		switch sign {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		}
	}
	return result
}

func convRomanToArab(a string) int {
	res := 0
	if RomeToArab[a] != 0 {
		return RomeToArab[a]
	} else {
		str := strings.Split(a, "")
		for _, v := range str {
			if RomeToArab[v] != 0 {
				res += RomeToArab[v]
			}
		}
	}
	return res
}

// Перевод результата из арабского в римский число
func convArabToRoman(res int) string {

	var result []string
	if res <= 0 {
		panic("Ошибка, результат ≤ 0, в римской системе нет нуля и отрицательных чисел")
	} else {
		for res > 0 {
			for _, value := range ArabToRome {
				for v := value; v <= res; res = res - value {
					for key, val := range RomeToArab {
						if val == value {
							result = append(result, key)
						}
					}
				}
			}
		}
	}
	return strings.Join(result, "")
}
