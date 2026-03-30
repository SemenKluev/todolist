package main

import (
	"bufio"
	"fmt"
	"os"
	"to-do-list/control"
	"to-do-list/methods"
)

func main() {
	//создание инструментов
	scanner := bufio.NewScanner(os.Stdin)
	userInput := 0
	userInputTasks := ""
	tasksName := []string{}

	fmt.Println("---ДОБРО ПОЖАЛОВАТЬ В TODO LIST---\n")

	for {
		control.Menu()
		fmt.Print("Выберете номер действия: ")
		fmt.Scan(&userInput)

		if userInput != 1 && userInput != 2 && userInput != 3 && userInput != 4 && userInput != 5 {
			fmt.Println("Вы ввели не корректный номер действия")
			continue
		}

		switch userInput {
		case 1:
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			userInputTasks = scanner.Text()

			methods.AddTasks(userInputTasks, tasksName)
		case 2:
			var err error = methods.DeleteTasks(userInput, tasksName)

			if err != nil {
				fmt.Println("Ошибка :( Причина:", err)
			}
		case 3:
			methods.CheckTasksList(tasksName)
		case 4:

		}
	}
}
