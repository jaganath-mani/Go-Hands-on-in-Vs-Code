package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client and database reference
var db *mongo.Database

// ==================== MONGODB CONNECTION ====================
// connectMongoDB establishes a connection to MongoDB and returns the client, context, cancel function, and any error.
func connectMongoDB(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// Create a context with a 10-second timeout for connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}

	return client, ctx, cancel, nil
}

// ==================== GET EMPLOYEES ====================
// getEmployees fetches all employees along with their department and developer info
func getEmployees(c *gin.Context) {
	employeeColl := db.Collection("Employee")

	// MongoDB aggregation pipeline:
	// 1. $lookup to join Department collection
	// 2. $lookup to join Developer collection
	// 3. $project to select only required fields
	pipeline := mongo.Pipeline{
		// Join Department collection
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Department"},
			{Key: "localField", Value: "emp_id"},
			{Key: "foreignField", Value: "emp_id"},
			{Key: "as", Value: "department"},
		}}},
		// Flatten the department array
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$department"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},
		// Join Developer collection
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Developer"},
			{Key: "localField", Value: "emp_id"},
			{Key: "foreignField", Value: "emp_id"},
			{Key: "as", Value: "developer"},
		}}},
		// Flatten the developer array
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$developer"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},
		// Project only necessary fields
		{{Key: "$project", Value: bson.D{
			{Key: "emp_id", Value: 1},
			{Key: "emp_name", Value: 1},
			{Key: "dept_name", Value: "$department.dept_name"},
			{Key: "planguage", Value: "$developer.planguage"},
		}}},
	}

	// Run aggregation
	cursor, err := employeeColl.Aggregate(context.Background(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	// Decode results
	var employees []map[string]interface{}
	if err := cursor.All(context.Background(), &employees); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Safe type handling before returning
	safeEmployees := make([]gin.H, 0, len(employees))
	for _, e := range employees {
		emp := gin.H{
			"emp_id":    e["emp_id"],
			"emp_name":  e["emp_name"],
			"dept_name": e["dept_name"],
			"planguage": e["planguage"],
		}
		safeEmployees = append(safeEmployees, emp)
	}

	c.JSON(http.StatusOK, safeEmployees)
}

// ==================== ADD EMPLOYEE ====================
// addEmployee adds a new employee to the Employee collection
// It uses an auto-increment logic for emp_id based on the current max emp_id in the collection
func addEmployee(c *gin.Context) {
	fmt.Println("Inside the addEmployee function")
	employeeColl := db.Collection("Employee")
	departmentColl := db.Collection("Department")
	developerColl := db.Collection("Developer")

	var input struct {
		EmpName         string `json:"emp_name"`
		Salary          int    `json:"salary"`
		DeptName        string `json:"dept_name"`
		ProgramLanguage string `json:"planguage"`
	}

	// Bind JSON input from frontend
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Auto-increment logic: find the current max emp_id
	var result bson.M
	opts := options.FindOne().SetSort(bson.D{{Key: "emp_id", Value: -1}})
	err := employeeColl.FindOne(context.Background(), bson.D{}, opts).Decode(&result)
	newEmpID := 1
	if err == nil {
		if id, ok := result["emp_id"].(int32); ok {
			newEmpID = int(id) + 1
		} else if id, ok := result["emp_id"].(int64); ok {
			newEmpID = int(id) + 1
		}
	}

	// Insert new employee document
	_, err = employeeColl.InsertOne(context.Background(), bson.D{
		{Key: "emp_id", Value: newEmpID},
		{Key: "emp_name", Value: input.EmpName},
		{Key: "salary", Value: input.Salary},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert employee"})
		return
	}

	_, _ = departmentColl.InsertOne(context.Background(), bson.D{
		{Key: "emp_id", Value: newEmpID},
		{Key: "dept_name", Value: input.DeptName},
	})

	_, _ = developerColl.InsertOne(context.Background(), bson.D{
		{Key: "emp_id", Value: newEmpID},
		{Key: "planguage", Value: input.ProgramLanguage},
	})

	c.JSON(http.StatusOK, gin.H{"message": "Employee added successfully", "emp_id": newEmpID})
}

// ==================== MAIN FUNCTION ====================
func main() {
	// Connect to MongoDB
	client, ctx, cancel, err := connectMongoDB("mongodb+srv://jagan:Jagan%401433@cluster0.ih6wvra.mongodb.net/")
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	// Set database reference
	db = client.Database("Magna_support")

	// Setup Gin router
	r := gin.Default()

	// Enable CORS for Vue frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vue dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	r.GET("/employees", getEmployees) // Fetch all employees
	r.POST("/employees", addEmployee) // Add new employee

	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
