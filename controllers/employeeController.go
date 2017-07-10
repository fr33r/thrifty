package controllers

import(
  "thrifty/gen-go/thrifty"
)

type EmployeeController struct{}

func(employeeController *EmployeeController) Get() (*thrifty.Person, error) {
  return &thrifty.Person{}, nil
}
