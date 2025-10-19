package main

import "fmt"

func main() {
	var num int64

	fmt.Print("Введите целое число: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Ошибка ввода! Пожалуйста, введите целое число.")
		return
	}

	if num >= 12307 {
		fmt.Printf("Введенное число %d уже больше или равно 12307\n", num)
		fmt.Printf("Финальный результат: %d\n", num)
		return
	}

	for num < 12307 {

		if num < 0 {
			num = num * -1
			fmt.Printf("Число отрицательное, умножаем на -1: %d\n", num)
		} else if num%7 == 0 {
			num = num * 39
			fmt.Printf("Число кратно 7, умножаем на 39: %d\n", num)
		} else if num%9 == 0 {
			num = num*13 + 1
			fmt.Printf("Число кратно 9, умножаем на 13 и прибавляем 1: %d\n", num)
			continue
		} else {
			num = (num + 2) * 3
			fmt.Printf("Во всех остальных случаях: (n+2)*3 = %d\n", num)
		}

		if num%13 == 0 && num%9 == 0 {
			fmt.Printf("Число %d одновременно кратно 13 и 9\n", num)
			fmt.Printf("service error\n")
			return
		} else {
			oldNum := num
			num += 1
			fmt.Printf("Прибавляем 1: %d + 1 = %d\n", oldNum, num)
		}
	}

	fmt.Printf("Финальное число: %d\n", num)
}
