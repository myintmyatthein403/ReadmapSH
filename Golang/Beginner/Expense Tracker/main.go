package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "time"
  "strconv"
)

type Expense struct {
  Id            int     `json:"id"`
  Description   string  `json:"description"`
  Amount        float64 `json:"amount"`
  Date          string  `json:"date"`
}

func addExpense(split []string, data *[]Expense, counter int) int {
  var description string
  var amount float64

  for i := 0; i < len(split); i++ {
    switch split[i] {
    case "--description":
      if i+1 < len(split) {
        description = split[i+1]
      }
    case "--amount":
      if i+1 < len(split) {
        amount, _ = strconv.ParseFloat(split[i+1], 64)
      }
    }
  }

  fmt.Println("Adding expense")
  *data = append(*data, Expense{
    Id:          counter,
    Description: description,
    Amount:      amount,
    Date:        time.Now().Format("2006-01-02"),
  })
  counter++
  fmt.Printf("Expense added successfully (ID: %d)\n", counter)
  return counter
}

func listExpenses(data []Expense) {
  fmt.Println("Listing expenses")
  for _, expense := range data {
    fmt.Println("--------------------")
    fmt.Println("ID: ", expense.Id)
    fmt.Println("Description: ", expense.Description)
    fmt.Println("Amount: ", expense.Amount)
    fmt.Println("Date: ", expense.Date)
  }
}

func getSummary(data []Expense, split []string) {
  if len(split) == 1 {
    fmt.Println("Getting summary")
    total := 0.0
    for _, expense := range data {
      total += expense.Amount
    }
    fmt.Println("Total: ", total)
  } else if len(split) == 3 && split[1] == "--month" {
    fmt.Println("Getting summary for a specific month")
    month := split[2]
    total := 0.0
    for _, expense := range data {
      expenseMonth := strings.Split(expense.Date, "-")[1]
      if expenseMonth == month {
        total += expense.Amount
      }
    }
    fmt.Println("Total for month", month, ": ", total)
  } else if len(split) == 3 && split[1] == "--year" {
    fmt.Println("Getting summary for a specific year")
    year := split[2]
    total := 0.0
    for _, expense := range data {
      expenseYear := strings.Split(expense.Date, "-")[0]
      if expenseYear == year {
        total += expense.Amount
      }
    }
    fmt.Println("Total for year", year, ": ", total)
  } else {
    fmt.Println("Invalid command")
  }
}

func deleteExpense(split []string, data *[]Expense) {
  index := 0
  for i := 0; i < len(split); i++ {
    switch split[i] {
    case "--id":
      if i+1 < len(split) {
        fmt.Sscanf(split[i+1], "%d", &index)
      }
    }
  }

  fmt.Println("Deleting expense")
  for i, expense := range *data {
    if expense.Id == index {
      *data = append((*data)[:i], (*data)[i+1:]...)
      fmt.Println("Expense deleted successfully")
      return
    }
  }
  fmt.Println("Expense not found")
}



func main() {
  fmt.Println("Expenses Tracker")

  scanner := bufio.NewScanner(os.Stdin)

  fmt.Printf(`
    CMD:
    to add an expense: add --description "Lunch" --amount 20
    to get all expenses: list
    to get summary: summary
    to get summary for a specific month: summary --month 8
    to delete an expense: delete --id 1
  `)

  counter := 0
  data := make([]Expense, 0)

  for scanner.Scan() {
    cmd := scanner.Text()
    split := strings.Split(cmd, " ")

    fmt.Println("Command: ", split)

    switch split[0] {
    case "add":
      counter = addExpense(split, &data, counter)
    case "list":
      listExpenses(data)
    case "summary":
      getSummary(data, split)
    case "delete":
      deleteExpense(split, &data)
    case "exit":
      return
    default:
      fmt.Println("Invalid command")
    }
  }
}
