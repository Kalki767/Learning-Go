package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// a function to accept string input from the user until a new line is encountered and removes leading and trailing spaces
func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	return strings.TrimSpace(inp)
}

// a function to accept input and convert it into integer and handles errors
func getIntInput() int {
	input := stringInput()
	num, err := strconv.Atoi(input)

	for err != nil {
		fmt.Println(err)
		fmt.Println("Please try again.")
		fmt.Print("Insert value: ")
		input = stringInput()
		num, err = strconv.Atoi(input)
	}

	return num
}

// a function to accept the total number of subjects from the user
func GetNumberOfSubjects() int {
	var total_subjects int
	fmt.Print("Enter the total number of subjects: ")
	total_subjects = getIntInput()

	return total_subjects
}

// a function to accept subjects associated with the grades and return a map of each subject with the grade
func GetSubjects(total_subjects int) map[string]float64 {

	subject_grade := make(map[string]float64)

	for i := 0; i < total_subjects; i++ {

		var subject string
		fmt.Print("Enter the subject Name: ")
		subject = stringInput()

		for {
			if _, exist := subject_grade[subject]; exist {
				fmt.Println("You have already entered this subject. Make sure to Enter unique subjects")
				fmt.Print("Please enter the subject: ")
				fmt.Scan(&subject)
			} else {
				break
			}
		}

		var grade float64
		fmt.Print("Enter the grade associated with the current subject: ")
		grade, err := strconv.ParseFloat(stringInput(), 64)
		for err != nil || grade < 0 || grade > 100 {
			fmt.Println("A student grade must be in the range between 0 & 100. Make sure to enter the correct grade.")
			fmt.Print("Enter the student Grade: ")
			grade, err = strconv.ParseFloat(stringInput(), 64)
		}

		for grade < 0 || grade > 100 {
			fmt.Println("A student grade must be in the range between 0 & 100. Make sure to Enter the correct grade.")
			fmt.Print("Enter the student Grade: ")
			fmt.Scan(&grade)
		}
		subject_grade[subject] = grade
	}
	return subject_grade
}

//main function

func main() {

	var total_number_of_students int
	fmt.Print("Enter the total number of students: ")
	total_number_of_students = getIntInput()
	student_grade := make(map[string]map[string]float64)
	student_average := make(map[string]float64)
	
	for i := 0; i < total_number_of_students; i++ {
		fmt.Print("Enter the student name: ")
		name := stringInput()
		total_subjects := GetNumberOfSubjects()
		subject_grade := GetSubjects(total_subjects)
		student_grade[name] = subject_grade
		total_sum := 0.0
		for _, grade := range subject_grade {
			total_sum += grade
		}
		student_average[name] = total_sum / float64(total_subjects)
	}

	for name, average := range student_average {
		fmt.Print(name)
		fmt.Print("      ")
		fmt.Println(average)
	}

}
