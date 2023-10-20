package main

import (
	"fmt"
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

func main() {
	var patients []DopInfa

	for {
		patient := newDopInfa()
		patients = append(patients, patient)

		var cont string
		fmt.Print("Добавить еще пациента? (да/нет): ")
		fmt.Scanln(&cont)

		if cont != "да" {
			break
		}
	}
	for i, patient := range patients {
		fmt.Printf("\n...Пациент %d...\nИмя: %-10s\nФамилия: %-10s\nВозраст: %-3d\nДиагноз: %-10s\nГород: %-10s\nЕсть дети? : %-3s\n", i+1, patient.Name, patient.Surname, patient.Age, patient.Diagnoz, patient.City, patient.Children)
	}

}
