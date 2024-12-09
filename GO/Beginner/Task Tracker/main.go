package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "slices"
)

type Task struct {
  Id            int     `json:"id"`
  Description   string  `json:"description"`
  Done          bool  `json:"status"`
}



func showTasks(data []Task)([]Task) {
  for _, task := range data {
    fmt.Println("--------------------")
    statusLine := "[ ]"
    if task.Done {
      statusLine = "[x]"
    }
    fmt.Println(statusLine, task.Description, "ID: ", task.Id)
  }

  return data
  
}

func createTask(split []string, data []Task, counter int)([]Task, int) {
  data = append(data, Task{
    Id: counter,
    Description: strings.Join(split[1:], " "), // Join the slice to get the description
    Done: false,
  })
  counter++
  fmt.Print("Task created\n")
  return data, counter
}

func removeTask(split []string, data []Task, id int)([]Task) {
  index, err := strconv.Atoi(split[1]) // Convert the string to an integer
  if (err != nil) {
    panic(err)
  }
  for i, task := range data {
    if task.Id == index {
      data = slices.Delete(data, i, i+1)
      break
    }
  }
  return data
}

func doneTask(split []string, data []Task)([]Task) {
  index, err := strconv.Atoi(split[1]) // Convert the string to an integer
  if (err != nil) {
    panic(err)
  }
  for i, task := range data {
    if task.Id == index {
      data[i].Done = true
      break
    }
  }
  return data
}

func main() {
  counter := 0 // Counter to generate the task Id
  data := make([]Task, 0) // Slice to store the tasks

  // Create a new scanner to read from the command line
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print("CMD: show create remove done\n")

  // Scan for the next token
  for scanner.Scan() {
    cmd := scanner.Text() // Get the token
    split := strings.Split(cmd, " ") // Split the token into a slice to get the command and the argument

    // Check the command
    switch split[0] {
      case "create":
        data, counter = createTask(split, data, counter)
      case "show":
        showTasks(data)
      case "remove":
        data = removeTask(split, data, counter)
      case "done":
        data = doneTask(split, data) 
      default:
        fmt.Print("Invalid command\n")
    } 
  }

}
