package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (db *sql.DB) {
	dsn := "test_user:password@tcp(localhost:3306)/crm"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// using named return value.
	return
}

func GetCustomers() []Customer {
	conn := GetConnection()
	defer conn.Close()

	//var rows *sql.Rows
	rows, err := conn.Query("SELECT * FROM Customer ORDER BY customer_id DESC")
	if err != nil {
		panic(err.Error())
	}
	var customers []Customer
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string

		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		var customer Customer
		customer.CustomerId, customer.CustomerName, customer.SSN = customerId, customerName, ssn

		customers = append(customers, customer)
	}

	return customers
}

func insertCustomer(c Customer) {
	conn := GetConnection()
	defer conn.Close()

	insert, err := conn.Prepare("INSERT INTO CUSTOMER(customer_name, ssn) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := insert.Exec(c.CustomerName, c.SSN)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.LastInsertId())

}
func main() {
	customers := GetCustomers()
	fmt.Println("Customers", customers)

	c := Customer{
		CustomerName: "Rishabh Sharma2",
		SSN:          "126-42-9344"}
	insertCustomer(c)
}
