package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"net/url"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson" 
	"go.mongodb.org/mongo-driver/mongo/options"

)

// Employee model
type Employee struct {
	EmployeeID int     `bson:"employee_id"`
	Name       string  `bson:"name"`
	Salary     float64 `bson:"salary"`
}

// Department model
type Department struct {
	EmployeeID int    `bson:"employee_id"`
	Department string `bson:"department"`
}

// Developer model
type Developer struct {
	EmployeeID int    `bson:"employee_id"`
	programmingLanguage      string `bson:"planguage"`
}

// Tester model
type Tester struct {
	EmployeeID int    `bson:"employee_id"`
	programmingLanguage       string `bson:"planguage"`
}

// Connect to MongoDB
func connectMongoDB(username, password string) (*mongo.Client, context.Context, context.CancelFunc) {
	// Replace <cluster-url> with your MongoDB Atlas cluster address
	encodedUser := url.QueryEscape(username)
	encodedPass := url.QueryEscape(password)

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.im3q7gc.mongodb.net/", encodedUser, encodedPass)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping failed:", err)
	}
	fmt.Println("Connected to MongoDB")
	return client, ctx, cancel
}

// Insert employee data
func insertEmployee(client *mongo.Client, ctx context.Context) {
	var emp Employee
	var dept Department
	var dev Developer
	var tester Tester

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&emp.EmployeeID)
	fmt.Print("Enter Name: ")
	fmt.Scan(&emp.Name)
	fmt.Print("Enter Salary: ")
	fmt.Scan(&emp.Salary)

	dept.EmployeeID = emp.EmployeeID
	fmt.Print("Enter Department: ")
	fmt.Scan(&dept.Department)

	dev.EmployeeID = emp.EmployeeID
	fmt.Print("Enter Developer Programming Language: ")
	fmt.Scan(&dev.programmingLanguage)

	tester.EmployeeID = emp.EmployeeID
	fmt.Print("Enter Tester Programming Language: ")
	fmt.Scan(&tester.programmingLanguage)

	db := client.Database("Magna_support")
	_, err := db.Collection("Employee").InsertOne(ctx, emp)
	if err != nil {
		log.Fatal("Insert Employee failed:", err)
	}
	_, _ = db.Collection("Department").InsertOne(ctx, dept)
	_, _ = db.Collection("Developer").InsertOne(ctx, dev)
	_, _ = db.Collection("Tester").InsertOne(ctx, tester)

	fmt.Println("Employee inserted successfully")
}

// Update employee name
func updateEmployee(client *mongo.Client, ctx context.Context) {
	var empID int
	var newName string

	fmt.Print("Enter Employee ID to update: ")
	fmt.Scan(&empID)
	fmt.Print("Enter new Name: ")
	fmt.Scan(&newName)

	employeeColl := client.Database("Magna_support").Collection("Employee")
	filter := bson.M{"employee_id": empID}
	update := bson.M{"$set": bson.M{"name": newName}}

	_, err := employeeColl.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("Update failed:", err)
	}
	fmt.Println("Employee updated successfully")
}

// Delete employee
func deleteEmployee(client *mongo.Client, ctx context.Context) {
	var empID int
	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scan(&empID)

	db := client.Database("Magna_support")
	collections := []string{"Employee", "Department", "Developer", "Tester"}
	for _, coll := range collections {
		_, err := db.Collection(coll).DeleteOne(ctx, bson.M{"employee_id": empID})
		if err != nil {
			log.Fatal("Delete failed:", err)
		}
	}
	fmt.Println("Employee deleted successfully")
}

// Read employees with $lookup
func readEmployeeDataWithLookup(client *mongo.Client, ctx context.Context) {
	db := client.Database("Magna_support")
	employeeColl := db.Collection("Employee")

	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Department"},
			{Key: "localField", Value: "employee_id"},
			{Key: "foreignField", Value: "employee_id"},
			{Key: "as", Value: "department"},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Developer"},
			{Key: "localField", Value: "employee_id"},
			{Key: "foreignField", Value: "employee_id"},
			{Key: "as", Value: "developer"},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Tester"},
			{Key: "localField", Value: "employee_id"},
			{Key: "foreignField", Value: "employee_id"},
			{Key: "as", Value: "tester"},
		}}},
	}

	cursor, err := employeeColl.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal("Aggregation failed:", err)
	}
	defer cursor.Close(ctx)

	fmt.Println("\n=== Final Employee Records ===")
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
		log.Fatal("Decode failed:", err)
	}

	// Convert result to JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal("JSON Marshal failed:", err)
	}

	fmt.Println(string(jsonData))
}
}

func main() {
	// Pass your MongoDB Atlas username & password
	var user_name, password string
	fmt.Print("Enter username for Mongo DB: ")
	fmt.Scan(&user_name)
	fmt.Print("Enter password for Mongo DB: ")
	fmt.Scan(&password)

	client, ctx, cancel := connectMongoDB(user_name, password)
	defer cancel()
	defer client.Disconnect(ctx)

	for {
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. Insert Employee")
		fmt.Println("2. Update Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Read Employees")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			insertEmployee(client, ctx)
		case 2:
			updateEmployee(client, ctx)
		case 3:
			deleteEmployee(client, ctx)
		case 4:
			readEmployeeDataWithLookup(client, ctx)
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}