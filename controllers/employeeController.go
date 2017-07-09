package controllers

import(
  "thrifty/gen-go/person"
)

type EmployeeController struct{}

func(employeeController *EmployeeController) Get() (*person.Person, error) {
  return &person.Person{}, nil
}
