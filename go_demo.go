// package main
// import "fmt"

// func main() {
// //    var a int = 21
// //    var b int = 10
// //    var c int

// //    fmt.Println(a == b)
// //    // Not equal to
// //    fmt.Println(a != b)
// //    // Greater than
// //    fmt.Println(a > b)
// //    // Less than
// //    fmt.Println(a < b)
// //    // Greater than or equal to
// //    fmt.Println(a >= b)
// //    // Less than or equal to
// //    fmt.Println(a <= b)

// //    c = a + b
// //    fmt.Printf("Line 1 - Value of c is %d\n", c )

// //    c = a - b
// //    fmt.Printf("Line 2 - Value of c is %d\n", c )

// //    c = a * b
// //    fmt.Printf("Line 3 - Value of c is %d\n", c )

// //    c = a / b
// //    fmt.Printf("Line 4 - Value of c is %d\n", c )

// //    c = a % b
// //    fmt.Printf("Line 5 - Value of c is %d\n", c )

// //    a++
// //    fmt.Printf("Line 6 - Value of a is %d\n", a )
//    var a int = 20
//    var b int = 10
//    var c int = 15
//    var d int = 5
//    var e int;

//    e = (a + b) * c / d;      // ( 30 * 15 ) / 5
//    fmt.Printf("Value of (a + b) * c / d is : %d\n",  e );

//    e = ((a + b) * c) / d;    // (30 * 15 ) / 5
//    fmt.Printf("Value of ((a + b) * c) / d is  : %d\n" ,  e );

//    e = (a + b) * (c / d);   // (30) * (15/5)
//    fmt.Printf("Value of (a + b) * (c / d) is  : %d\n",  e );

//    e = a + (b * c) / d;     //  20 + (150/5)
//    fmt.Printf("Value of a + (b * c) / d is  : %d\n" ,  e );
// }

// package main

// import "fmt"

// func main() {
//    /* local variable definition */
//    var a int = 100
//    var b int = 200
//    var ret int

//    /* calling a function to get max value */
//    ret = max(a, b)

//    fmt.Printf( "Max value is : %d\n", ret )
// }

// /* function returning the max between two numbers */
// func max(num1, num2 int) int {
//    /* local variable declaration */
//    var result int

//    if (num1 > num2) {
//       result = num1
//    } else {
//       result = num2
//    }
//    return result
// }
//employees := make([]Employee, 0, count)
//employees = append(employees, emp)

/*package main

import "fmt"

type Employee struct {
	emp_id   int
	emp_name string
	salary   float64
}

func main() {
	var no_of_emp int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&no_of_emp)

	var employees [100]Employee

	for i := 0; i < no_of_emp; i++ {
		fmt.Printf("\nEnter details for Employee %d\n", i+1)
		fmt.Print("Enter Employee ID: ")
		fmt.Scan(&employees[i].emp_id)
		fmt.Print("Enter Employee Name: ")
		fmt.Scan(&employees[i].emp_name)
		fmt.Print("Enter Employee Salary: ")
		fmt.Scan(&employees[i].salary)
	}

	fmt.Println("\n--- Employee Details ---")
	for i := 0; i < no_of_emp; i++ {
		fmt.Printf("Employee ID: %d, \nEmployee Name: %s, \nEmployee Salary: %.2f\n",
			employees[i].emp_id, employees[i].emp_name, employees[i].salary)
		fmt.Printf("----------------------------------------------------------------\n")
	}
}*/

package main

import "fmt"

type Employee struct {
	emp_id   int
	emp_name string
	salary   float64
}

// Function Call by Value
func updateByValue(arr [3]Employee) {

	arr[1].emp_name = "Modified Employee 2"
	arr[1].salary = 123.00

	fmt.Println("\n--- Call by Value ---")
	fmt.Println("---Inside updateArray function (local copy changed)---")
	for _, emp := range arr {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}
}

// Function Call by Reference
func updateByReference(slice []Employee) {

	slice[1].emp_name = "Modified Employee 2 by reference"
	slice[1].salary = 456.00

	fmt.Println("\n--- Call by Reference ---")
	fmt.Println("--- Inside updateSlice (original changed) ---")
	for _, emp := range slice {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}
}
func main() {
	var n int
	fmt.Print("\nEnter number of employees: ")
	fmt.Scan(&n)
	
	arrayEmployees := [3]Employee{
		{101, "AAA", 5000.00},
		{102, "BBB", 6000.00},
		{103, "CCC", 7000.00},
	}
	
	fmt.Println("\n=== Original Employees for Value ===")
	for _, emp := range arrayEmployees {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}

	// Call by Value
	updateByValue(arrayEmployees)

	fmt.Println("After updateByValue (Original Values unchanged):")
	for _, emp := range arrayEmployees {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}

	Employees := make([]Employee, n)
	fmt.Println("\nEnter details for employees :")
	for i := 0; i < n; i++ {
		fmt.Printf("\nEmployee %d:\n", i+1)
		fmt.Print("Enter ID: ")
		fmt.Scan(&Employees[i].emp_id)
		fmt.Print("Enter Name: ")
		fmt.Scan(&Employees[i].emp_name)
		fmt.Print("Enter Salary: ")
		fmt.Scan(&Employees[i].salary)
	}

	fmt.Println("\n=== Original Employees for reference ===")
	for _, emp := range Employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}

	// Call by Reference
	updateByReference(Employees)

	fmt.Println("After updateByReference (Original Values changed):")
	for _, emp := range Employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", emp.emp_id, emp.emp_name, emp.salary)
	}
}