package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name  string
	age   int
	grade int
}

var studentsInfo = make(map[string]*student)

func newStudent(infoStr string) (st student, err error) {
	data := strings.Fields(infoStr)

	if len(data) != 3 {
		err = fmt.Errorf("Некорректный ввод, введите данные студентов в формате Имя Возраст Курс:")
		return st, err
	}

	st.name = data[0]
	st.age, err = strconv.Atoi(data[1])
	if err != nil {
		fmt.Println("Не удалось считать возраст:")
		return st, err
	}

	st.grade, err = strconv.Atoi(data[2])
	if err != nil {
		fmt.Println("Не удалось считать курс:")
		return st, err
	}

	return st, err
}

func (st student) put() {
	studentsInfo[st.name] = &st
}

func (st student) get() student {
	return *studentsInfo[st.name]
}

func main() {
	scStudentInfo := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите данные студентов в формате \"Имя Возраст Курс\":")
	for scStudentInfo.Scan() {
		infoStr := scStudentInfo.Text()
		st, err := newStudent(infoStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		st.put()
	}
	fmt.Println("Студенты из хранилища:")
	for _, students := range studentsInfo {
		fmt.Println(students.get())
	}
}
