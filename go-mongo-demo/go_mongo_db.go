package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"net/url"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson" 
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	EmpID  int     `bson:"emp_id"`
	Name   string  `bson:"emp_name"`
	Salary float64 `bson:"salary"`
}

// connection function
func connectMongoDB(username, password string) (*mongo.Client, context.Context, context.CancelFunc) {
	// Replace <cluster-url> with your MongoDB Atlas cluster address
	encodedUser := url.QueryEscape(username)
	encodedPass := url.QueryEscape(password)

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.pw9wdb5.mongodb.net/", encodedUser, encodedPass)

	// Create a client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating Mongo client:", err)
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	// Test connection with Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB Atlas")
	return client, ctx, cancel
}

func main() {
	// Pass your MongoDB Atlas username & password
	username := "jagan"
	password := "Jagan@1433"

	client, ctx, cancel := connectMongoDB(username, password)
	defer cancel()
	defer client.Disconnect(ctx)

	// Select database & collection
	db := client.Database("Magna_support")
	collection := db.Collection("Employee")

	// Get input of employee details from user
	var n int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	employees := make([]interface{}, n)
	for i := 0; i < n; i++ {
		var emp Employee
		fmt.Printf("\nEnter details for Employee %d\n", i+1)
		fmt.Print("Enter Employee ID: ")
		fmt.Scan(&emp.EmpID)
		fmt.Print("Enter Employee Name: ")
		fmt.Scan(&emp.Name)
		fmt.Print("Enter Employee Salary: ")
		fmt.Scan(&emp.Salary)
		employees[i] = emp
	}

	// Insert employees into MongoDB
	insertResult, err := collection.InsertMany(ctx, employees)
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	fmt.Println("\nInserted Employee IDs:")
	for _, id := range insertResult.InsertedIDs {
		fmt.Println(id)
	}

	fmt.Println("\n=== Employees in MongoDB ===")

	// Retrieve and display all employees details from MongoDB
	cursor, err := db.Collection("Employee").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Find failed:", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var emp Employee
		if err := cursor.Decode(&emp); err != nil {
			log.Fatal("Decode failed:", err)
		}
		fmt.Printf("ID: %d, \nName: %s, \nSalary: %.2f\n", emp.EmpID, emp.Name, emp.Salary)
	}

	// Check if loop ended due to error
	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error:", err)
	}
}