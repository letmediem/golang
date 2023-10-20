package main

import (
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Diagnoz string
}

func newPerson() Person {
	var name, surname, diagnoz string
	var age int

	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите фамилию: ")
	fmt.Scanln(&surname)
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&age)
	fmt.Print("Введите диагноз: ")
	fmt.Scanln(&diagnoz)

	return Person{
		Name:    name,
		Surname: surname,
		Age:     age,
		Diagnoz: diagnoz,
	}
}

type DopInfa struct {
	Person
	City     string
	Children string
}

func newDopInfa() DopInfa {
	var city, children string

	person := newPerson()

	fmt.Print("Введите город: ")
	fmt.Scanln(&city)
	fmt.Print("Есть ли дети?: ")
	fmt.Scanln(&children)

	return DopInfa{
		Person:   person,
		City:     city,
		Children: children,
	}
}

func findPatientsByCriteria(patients []DopInfa, criteria string, value interface{}) []DopInfa {
	var filteredPatients []DopInfa
	for _, patient := range patients {
		switch criteria {
		case "age":
			if patient.Age == value.(int) {
				filteredPatients = append(filteredPatients, patient)
			}
		case "city":
			if patient.City == value.(string) {
				filteredPatients = append(filteredPatients, patient)
			}
		}
	}
	return filteredPatients
}

func writeToFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var patients []DopInfa

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Добавить пациента")
		if len(patients) > 0 {
			fmt.Println("2. Посмотреть всех пациента")
			fmt.Println("3. Найти пациентов по критерию")
		}
		fmt.Println("4. Завершить программу")

		var choice int
		fmt.Print("Выберите оперцию: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			patient := newDopInfa()
			patients = append(patients, patient)
			fmt.Println("Пациент добавлен")

			content := fmt.Sprintf("Имя: %s, Фамилия: %s, Возраст: %d, Диагноз: %s, Город: %s, Дети: %s\n", patient.Name, patient.Surname, patient.Age, patient.Diagnoz, patient.City, patient.Children)
			writeToFile("patients.txt", content)
		case 2:
			for i, patient := range patients {
				fmt.Printf("\n...Пациент %d...\nИмя: %-10s\nФамилия: %-10s\nВозраст: %-3d\nДиагноз: %-10s\nГород: %-10s\nЕсть дети? : %-3s\n", i+1, patient.Name, patient.Surname, patient.Age, patient.Diagnoz, patient.City, patient.Children)
			}

		case 3:
			var criteria string
			fmt.Println("Введите критерий для поиска(age/city): ")
			fmt.Scanln(&criteria)

			var value interface{}
			fmt.Printf("Введите значение для критерия %s: ", criteria)
			if criteria == "age" {
				var age int
				fmt.Scanln(&age)
				value = age
			} else if criteria == "city" {
				var city string
				fmt.Scanln(&city)
				value = city
			} else {
				fmt.Println("Вы выбрали несуществующий критерий")
			}

			filteredPatients := findPatientsByCriteria(patients, criteria, value)

			for i, patient := range filteredPatients {
				fmt.Printf("\n...Пациент %d...\nИмя: %-10s\nФамилия: %-10s\nВозраст: %-3d\nДиагноз: %-10s\nГород: %-10s\nЕсть дети? : %-3s\n", i+1, patient.Name, patient.Surname, patient.Age, patient.Diagnoz, patient.City, patient.Children)
			}

		case 4:
			fmt.Println("Программа завершена")
			return

		default:
			fmt.Println("Неверный вариант. Пожалуйста, выберите снова")

		}
	}
}
