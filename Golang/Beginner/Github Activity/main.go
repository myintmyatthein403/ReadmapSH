package main;

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "net/http"
  "io"
  "encoding/json"

);

func parseActivities(body io.Reader) []string {
  var activities []map[string]interface{} // Define a slice of maps
  err := json.NewDecoder(body).Decode(&activities) // Decode the JSON data into the slice
  if err != nil {
    panic(err)
  }
  var result []string
  for _, activity := range activities {
    actor := activity["actor"].(map[string]interface{})
    repo := activity["repo"].(map[string]interface{})
    result = append(result, fmt.Sprintf("%s %s %s", actor["login"], activity["type"], repo["name"]))
  } // Loop through the activities and append the formatted string to the result slice
  return result
}

func getActivities(username string) []string {
  url := "https://api.github.com/users/" + username + "/events"
  response, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer response.Body.Close()
  return parseActivities(response.Body)
}

func main() {
  fmt.Println("Github Activities")
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("CMD:\nTo search Activities - github-activity <username>\nTo exit - exit\n")

  loop:
  for scanner.Scan() {
    cmd := scanner.Text() 
    split := strings.Split(cmd, " ")
    switch split[0] {
      case "github-activity":
        if len(split) != 2 {
          fmt.Print("Invalid command\n")
          continue
        }
        username := split[1]
        activities := getActivities(username)
        for _, activity := range activities {
          fmt.Println(activity)
        }
      case "exit":
        break loop
      default:
        fmt.Print("Invalid command\n")
    }
  }
}
