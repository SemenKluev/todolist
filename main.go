package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	//переменная считывает значения пользователя
	var a int
	taskSlice := []string{}

	fmt.Println("---Добро пожаловать в TO DO LIST---")

	for {
		menu()
		fmt.Print("Выберите действие: ")
		fmt.Scan(&a)

		switch a {
		case 1:
			addTask(&taskSlice)
		case 2:
			getAllTasks(taskSlice)
		case 3:
			removeTask(&taskSlice)
		case 4:
			exit()
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func menu() {
	fmt.Println("")
	fmt.Println("1 - Добавить дело")
	fmt.Println("2 - Просмотреть список дел")
	fmt.Println("3 - Закрыть задачу")
	fmt.Println("4 - Выход")
}

func addTask(taskSlice *[]string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите задачу: ")
	scanner.Scan()
	*taskSlice = append(*taskSlice, scanner.Text())

	fmt.Println("Ваша задача добавлена!!!")
}

func getAllTasks(taskSlice []string) {
	fmt.Println("---Все ваши задачи---")

	if len(taskSlice) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	for _, task := range taskSlice {
		fmt.Println(task)
	}
}

func removeTask(taskSlice *[]string) {
	number := 0

	for index, task := range *taskSlice {
		fmt.Println(index, ":", task)
	}

	fmt.Print("Введите номер задачи которую хотите удалить: ")
	fmt.Scan(&number)

	if number < 0 || number >= len(*taskSlice) {
		fmt.Println("Вы ввели некорректный номер")
		return
	}

	*taskSlice = slices.Delete(*taskSlice, number, number+1)
	fmt.Println("Задача закрыта")
}

func exit() {
	fmt.Println("Выход...")
}
