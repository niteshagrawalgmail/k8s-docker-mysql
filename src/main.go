package main
import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
func main() {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/commerce")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Customer struct {
		Id int
		First_Name string
		Last_Name string
	}
	router := gin.Default()
	// GET a customer detail
	router.GET("/customers/:id", func(c *gin.Context) {
		var (
			customer Customer
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, first_name, last_name from customers where id = ?;", id)
			err = row.Scan(&customer.Id, &customer.First_Name, &customer.Last_Name)
			if err != nil {
				// If no results send null
				result = gin.H{
					"result": nil,
					"count": 0,
				}
			} else {
				result = gin.H{
					"result": customer,
					"count": 1,
				}
			}
			c.JSON(http.StatusOK, result)
		})
		// GET all customers
		router.GET("/customers", func(c *gin.Context) {
			var (
				customer Customer
				customers []Customer
			)
			rows, err := db.Query("select id, first_name, last_name from customers;")
				if err != nil {
					fmt.Print(err.Error())
				}
				for rows.Next() {
					err = rows.Scan(&customer.Id, &customer.First_Name, &customer.Last_Name)
					customers = append(customers, customer)
					if err != nil {
						fmt.Print(err.Error())
					}
				}
				defer rows.Close()
				c.JSON(http.StatusOK, gin.H{
					"result": customers,
					"count": len(customers),
				})
			})
			// POST new customer details
			router.POST("/customer", func(c *gin.Context) {
				var buffer bytes.Buffer
				first_name := c.PostForm("first_name")
				last_name := c.PostForm("last_name")
				stmt, err := db.Prepare("insert into customers (first_name, last_name) values(?,?);")
				if err != nil {
					fmt.Print(err.Error())
				}
				_, err = stmt.Exec(first_name, last_name)
				if err != nil {
					fmt.Print(err.Error())
				}
				// Fastest way to append strings
				buffer.WriteString(first_name)
				buffer.WriteString(" ")
				buffer.WriteString(last_name)
				defer stmt.Close()
				name := buffer.String()
				c.JSON(http.StatusOK, gin.H{
					"message": fmt.Sprintf(" %s successfully created", name),
				})
			})
			// PUT – update a customer details
			router.PUT("/customer", func(c *gin.Context) {
				var buffer bytes.Buffer
				id := c.Query("id")
				first_name := c.PostForm("first_name")
				last_name := c.PostForm("last_name")
				stmt, err := db.Prepare("update customers set first_name= ?, last_name= ? where id= ?;")
				if err != nil {
					fmt.Print(err.Error())
				}
				_, err = stmt.Exec(first_name, last_name, id)
				if err != nil {
					fmt.Print(err.Error())
				}
				// Fastest way to append strings
				buffer.WriteString(first_name)
				buffer.WriteString(" ")
				buffer.WriteString(last_name)
				defer stmt.Close()
				name := buffer.String()
				c.JSON(http.StatusOK, gin.H{
					"message": fmt.Sprintf("Customer Successfully updated to %s", name),
				})
			})
			// Delete customer
			router.DELETE("/customer", func(c *gin.Context) {
				id := c.Query("id")
				stmt, err := db.Prepare("delete from customers where id= ?;")
				if err != nil {
					fmt.Print(err.Error())
				}
				_, err = stmt.Exec(id)
				if err != nil {
					fmt.Print(err.Error())
				}
				c.JSON(http.StatusOK, gin.H{
					"message": fmt.Sprintf("Successfully deleted customer: %s", id),
				})
			})
			router.Run(":3000")
		}