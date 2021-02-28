package main

import (
	"fmt"
	"strconv"
)

func main() {
	staffList := make([]StaffMember, 6)
	staffList[0] = NewExecutive("Sam", "123 Main Line", "555-0469", "123-45-6789", 2423.07)
	staffList[1] = NewEmployee("Carla", "456 Off Line", "555-0101", "987-65-4321", 1246.15)
	staffList[2] = NewEmployee("Woody", "789 Off Rocker", "555-0000", "010-20-3040", 1169.23)
	staffList[3] = NewHourly("Diane", "678 Fifth Ave.", "555-0690", "958-47-3625", 10.55)
	staffList[4] = NewVolunteer("Norm", "987 Suds Blvd.", "555-8374")
	staffList[5] = NewVolunteer("Cliff", "321 Duds Lane", "555-7282")

	staffList[0].(*Executive).awardBonus(500.00)
	staffList[3].(*Hourly).addHours(40)

	var amount float64

	for count := 0; count < len(staffList); count++ {
		fmt.Println(staffList[count].toString())
		amount = staffList[count].pay()
		if amount == 0.0 {
			fmt.Println("Thanks!")
		} else {
			fmt.Printf("Paid: %f \n", amount)
			fmt.Println("-----------------------------------")
		}
	}
}

//StaffMemner interface
type StaffMember interface {
	pay() float64
	toString() string
}

//AbstractStaffMember is gonna help us create the default toString and constructor
type AbstractStaffMember struct {
	StaffMember
	name    string
	address string
	phone   string
}

//default constructor
func NewAbstractStaffMember(eName string, eAddress string, ePhone string) *AbstractStaffMember {
	asm := AbstractStaffMember{name: eName, address: eAddress, phone: ePhone}
	return &asm
}

//default toString
func (asm *AbstractStaffMember) toString() string {
	result := "Name: " + asm.name + "\n"
	result += "Adress: " + asm.address + "\n"
	result += "Phone: " + asm.phone
	return result
}

func (asm *AbstractStaffMember) pay() float64 {
	return 0.0
}

//************************************************

//Volunteer
type Volunteer struct {
	AbstractStaffMember
}

// Volunteer constructor
func NewVolunteer(eName string, eAddress string, ePhone string) *Volunteer {
	v := Volunteer{*NewAbstractStaffMember(eName, eAddress, ePhone)}
	return &v
}

// Volunteer pay method
func (v *Volunteer) pay() float64 {
	return 0.0
}

//************************************************

//Employee
type Employee struct {
	AbstractStaffMember
	socialSecurityNumber string
	payRate              float64
}

//Employee constructor
func NewEmployee(eName string, eAddress string, ePhone string, socSecNumber string, rate float64) *Employee {
	emp := Employee{AbstractStaffMember: *NewAbstractStaffMember(eName, eAddress, ePhone), socialSecurityNumber: socSecNumber, payRate: rate}
	return &emp
}

//Employee toString method
func (emp *Employee) toString() string {
	result := emp.AbstractStaffMember.toString()
	result += "\nSocial Security Number: " + emp.socialSecurityNumber
	return result
}

//Employee pay
func (emp *Employee) pay() float64 {
	return emp.payRate
}

//************************************************

//Executive
type Executive struct {
	Employee
	bonus float64
}

//Executive constructor
func NewExecutive(eName string, eAddress string, ePhone string, socSecNumber string, rate float64) *Executive {
	exec := Executive{Employee: *NewEmployee(eName, eAddress, ePhone, socSecNumber, rate), bonus: 0.0}
	return &exec
}

//Executive awardBonus method
func (exec *Executive) awardBonus(execBonus float64) {
	exec.bonus = execBonus
}

//Executive pay method
func (exec *Executive) pay() float64 {
	payment := exec.Employee.pay() + exec.bonus
	exec.bonus = 0
	return payment
}

//************************************************

//Hourly
type Hourly struct {
	Employee
	hoursWorked int
}

//Hourly constructor
func NewHourly(eName string, eAddress string, ePhone string, socSecNumber string, rate float64) *Hourly {
	hourly := Hourly{Employee: *NewEmployee(eName, eAddress, ePhone, socSecNumber, rate), hoursWorked: 0}
	return &hourly
}

//Hourly to String
func (hourly *Hourly) toString() string {
	result := hourly.Employee.toString()
	i := strconv.Itoa(hourly.hoursWorked)
	result += "\nCurrent hours: " + i
	return result
}

//Hourly addHours
func (hourly *Hourly) addHours(moreHours int) {
	hourly.hoursWorked += moreHours
}

//Hourly pay
func (hourly *Hourly) pay() float64 {
	payment := hourly.payRate * float64(hourly.hoursWorked)
	hourly.hoursWorked = 0
	return payment
}
