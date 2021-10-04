package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var taskList []Task

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option: (a - add task, t-task list, x - exit): ", reader)
	switch opt {
	case "a":
		fmt.Println("You chose a")
		promptOptions()
	case "t":
		showTaskList()
		promptOptions()
	case "x":
		fmt.Println("You chose exit (x)")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions()
	}
}

func showTaskList() {
	for _, value := range taskList {
		fmt.Printf("ID: %-10vName: %-40vStatus: %-15vDeadline: %-20v\n", value.ID, value.Name, value.Status, value.Deadline.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	director := Director{"Main director"}
	director.giveTask("Do some stuff")
	director.giveTask("Make a connection with database")
	// fmt.Println(taskList)
	developer := Developer{"Nodir"}
	developer.develop(81, "dev")
	promptOptions()
	// fmt.Println(task.Deadline.Format("2006-01-02 15:04:05"))
}
