package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Doctor struct {
	ID                               int
	Name                             string
	Gender                           string
	Address                          string
	City                             string
	Phone                            string
	Specialisation                   string
	Opening_time                     string
	Closing_time                     string
	Availability_time                string
	Availability                     string
	Available_for_home_visit         string
	Available_for_online_consultancy string
	Fees                             int
}

func Connect() {

	// Create Database
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/")
	if err != nil {

		panic(err.Error())

	}
	defer db.Close()

	//  Create Database

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS doctorinfo")

	if err != nil {

		panic(err.Error())

	}

	// Make Database Connection
	db, err = sql.Open("mysql", "root:india@123@tcp(localhost:3306)/doctorinfo")
	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Connected to MySQL database!")

	// Create Docter table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Doctor(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Specialisation VARCHAR(20),Opening_time VARCHAR(10),Closing_time VARCHAR(10),Availability_time VARCHAR(30),Availability VARCHAR(10),Available_for_home_visit VARCHAR(4),Available_for_online_consultancy VARCHAR(4),Fees INT ,PRIMARY KEY (ID) );")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Docter Table Created")

}

func Add_docter() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/doctorInfo")
		if err != nil {
			log.Fatal(err)
		}
		var data Doctor
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Doctor (Name,Gender,Address,City,Phone,Specialisation,Opening_time,Closing_time,Availability_time,Availability,Available_for_home_visit,Available_for_online_consultancy,Fees) VALUES ( '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%d)`, data.Name, data.Gender, data.Address, data.City, data.Phone, data.Specialisation, data.Opening_time, data.Closing_time, data.Availability_time, data.Availability, data.Available_for_home_visit, data.Available_for_online_consultancy, data.Fees)
		fmt.Println(query_data)
		//insert data
		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Doctor added successfully"})

	}
}

func Get_my_profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/doctorinfo")
		if err != nil {
			log.Fatal(err)
		}
		var mob Doctor
		err = c.BindJSON(&mob)
		if err != nil {
			return
		}
		get_detail := fmt.Sprintf("SELECT * FROM Doctor WHERE Phone = '%s'", mob.Phone)
		detail, err := db.Query(get_detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Opening_time string
			var Closing_time string
			var Availability_Time string
			var Availability string
			var Available_for_home_visit string
			var Available_for_online_consultancy string
			var Fees float64
			err = detail.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Doctor details": output})

		}

	}
}

func Update_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/doctorinfo")
		if err != nil {

			log.Fatal(err)

		}

		var data Doctor
		var updateColumns []string
		var args []interface{}

		err = c.BindJSON(&data)

		if err != nil {

			return

		}
		fmt.Println(data)
		if data.Address != "" {
			updateColumns = append(updateColumns, "Address = ?")
			args = append(args, data.Address)
		}
		fmt.Println(updateColumns, args)

		if data.City != "" {
			updateColumns = append(updateColumns, "City = ?")
			args = append(args, data.City)
		}
		fmt.Println(updateColumns, args)
		if data.Phone != "" {
			updateColumns = append(updateColumns, "Phone = ?")
			args = append(args, data.Phone)
		}
		fmt.Println(updateColumns, args)
		if data.Specialisation != "" {
			updateColumns = append(updateColumns, "Specialisation = ?")
			args = append(args, data.Specialisation)
		}

		if data.Opening_time != "" {
			updateColumns = append(updateColumns, "Opening_time = ?")
			args = append(args, data.Opening_time)
		}
		fmt.Println(updateColumns, args)
		if data.Closing_time != "" {
			updateColumns = append(updateColumns, "Closing_time = ?")
			args = append(args, data.Closing_time)
		}
		fmt.Println(updateColumns, args)

		if data.Availability_time != "" {
			updateColumns = append(updateColumns, "Availability_time = ?")
			args = append(args, data.Availability_time)
		}
		fmt.Println(updateColumns, args)
		if data.Availability != "" {
			updateColumns = append(updateColumns, "Availability = ?")
			args = append(args, data.Availability)
		}
		fmt.Println(updateColumns, args)
		if data.Available_for_home_visit != "" {
			updateColumns = append(updateColumns, "Available_for_home_visit = ?")
			args = append(args, data.Available_for_home_visit)
		}
		fmt.Println(updateColumns, args)
		if data.Available_for_online_consultancy != "" {
			updateColumns = append(updateColumns, "Available_for_online_consultancy = ?")
			args = append(args, data.Available_for_online_consultancy)
		}
		fmt.Println(updateColumns, args)
		if data.Fees != 0 {
			updateColumns = append(updateColumns, "Fees = ?")
			args = append(args, data.Fees)
		}
		fmt.Println(updateColumns, args)
		if len(updateColumns) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No update data provided"})
			return
		}
		fmt.Println(updateColumns, args)
		updateQuery := "UPDATE Doctor SET " + strings.Join(updateColumns, ", ") + " WHERE id = ?"
		args = append(args, data.ID)
		fmt.Println(updateQuery)
		stmt, err := db.Prepare(updateQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()
		if _, err := stmt.Exec(args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, data)

		c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})

	}
}

func Delete_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/doctorinfo")
		if err != nil {

			log.Fatal(err)

		}

		var data Doctor

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		delete_query := fmt.Sprintf("DELETE FROM Doctor WHERE ID = %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Doctor Deleted successfully"})

	}
}

func main() {
	Connect()
	router := gin.Default()
	router.POST("/", Add_docter())
	router.GET("/", Get_my_profile())
	router.PUT("/", Update_docter())
	router.DELETE("/", Delete_docter())
	router.Run("localhost:8080")
}
