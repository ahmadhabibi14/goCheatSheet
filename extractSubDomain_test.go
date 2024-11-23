package main

import (
	"fmt"
	"regexp"
)

func main() {
  urls := []string{
    "http://admin-2642:1235",
    "http://admin-2642.example.com",
  }

  // Define a regular expression pattern to match the desired part
  re := regexp.MustCompile(`://([^:./]+)`)

  for _, url := range urls {
    match := re.FindStringSubmatch(url)
    if len(match) > 1 {
      fmt.Println("Extracted string:", match[1])
    } else {
      fmt.Println("No match found in URL:", url)
    }
    }
}
