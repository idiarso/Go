package services

import "github.com/idiarso/belajar-git/src/api/models"

var students []models.Student

func AddStudent(name string, age int, class string) models.Student {
    student := models.Student{Name: name, Age: age, Class: class}
    students = append(students, student)
    return student
}

func GetStudents() []models.Student {
    return students
}
