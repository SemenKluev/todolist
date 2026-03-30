package methods

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

func AddTasks(description string, a []string) []string {
	a = append(a, description)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Загрузка данных....")

	time.Sleep(1 * time.Second)
	fmt.Println("Ваша задача успешно добавлена!")

	return a
}

func DeleteTasks(user int, a []string) error {
	// Показываем список задач
	for index, value := range a {
		fmt.Println(index+1, ":", value)
	}

	fmt.Print("Выберите номер задачи для удаления: ")
	fmt.Scan(&user)

	if user < 1 || user > len(a) {
		return errors.New("Неизвестный номер задачи")
	}

	a = slices.Delete(a, user-1, user)

	fmt.Println("Задача успешно удалена")
	return nil
}

func CheckTasksList(a []string) {
	fmt.Println("Загрузка...")
	time.Sleep(1 * time.Second)

	if len(a) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	fmt.Println("\n--- СПИСОК ЗАДАЧ ---")
	for i, value := range a {
		fmt.Printf("%d. %s\n", i+1, value)
	}
}

func CompletingTask(user int, a []string) ([]string, error) {
	for index, value := range a {
		fmt.Println(index+1, ":", value)
	}

	fmt.Print("Выберите номер задачи для выполнения: ")
	fmt.Scan(&user)

	if user < 1 || user > len(a) {
		return a, errors.New("Неизвестный номер задачи")
	}

	a = slices.Delete(a, user-1, user)

	fmt.Println("Ваша задача успешно выполнена!!!")
	return a, nil
}
