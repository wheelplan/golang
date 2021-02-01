package main

import (
    "fmt"
)

type SalaryCal interface {
    CalSalary() float64
}

type Team1 struct {
    id          int
    basicpay    float64
    bonus       float64
}

type Team2 struct {
    id          int
    basicpay    float64
}


func main() {
    p1 := Team1{1, 12645.91, 100.01}
    p2 := Team1{2, 17391.15, 210.99}
    p3 := Team2{3, 20009.13}
    employees := []SalaryCal{p1, p2, p3}
    totalExpense(employees)
}

func totalExpense(s []SalaryCal) {
    expense := 0.0
    for _, v := range s {
        expense += v.CalSalary()
    }

    fmt.Printf("TotalExpense is $%7.2f", expense)
}

func (t1 Team1) CalSalary() float64 {
    return t1.basicpay + t1.bonus
}


func (t2 Team2) CalSalary() float64 {
    return t2.basicpay
}
