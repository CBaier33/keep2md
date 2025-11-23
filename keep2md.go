package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)


type Label struct {
	Name string `json:"name"`
}

type Input struct {
	Title string `json:"title"`
	Text  string `json:"textContent"`
	Timestamp int64 `json:"createdTimestampUsec"`
	Labels []Label `json:"labels"`
}

type Page struct {
	Title string
	Text string
	Date string
	Tags string
}


func getDate(t int64) string {

	sec := t / 1_000_000
	nsec := (t % 1_000_000) * 1000

	date := time.Unix(sec, nsec).Format("January 2, 2006");

	return date

}

func getTags(labels []Label) string {

  if len(labels) == 0 {
      return "[[Keep]]"
  }
  
  var parts []string
  for _, label := range labels {
      parts = append(parts, "[["+label.Name+"]]")
  }
  
  return strings.Join(parts, ", ")
}

var defaultTemplate string = `##### {{.Date}}

{{.Text}}

Tags: {{.Tags}}`

func main() {

	// Load File
	filePath := flag.String("f", "", "path to JSON file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("missing -f (file)")
		os.Exit(1)
	} else {
		fmt.Printf("Converting: %s | ", *filePath)
	}

	file, err := os.ReadFile(*filePath)

	if err != nil {
		panic(err)
	}

	
	// Extract JSON Input
	var input Input
	err = json.Unmarshal(file, &input)

	if err != nil {
		panic(err)
	}

	// Load template
	userTemplate, err := template.New("default").Parse(defaultTemplate)

	if err != nil {
		panic(err)
	}

	// Create New Page
	title := input.Title
	date := getDate(input.Timestamp)
	tags := getTags(input.Labels)
	text := input.Text

	p := Page{
		Title: title,
		Date: date,
		Tags: tags,
		Text: text,
	}

	newFileName := fmt.Sprintf("%s.md", title)
	outFile, err := os.Create(newFileName)

	if err != nil {
			panic(err)
	}
	defer outFile.Close()

	// Build from template
	if err := userTemplate.Execute(outFile, p); err != nil {
			panic(err)
	}

	fmt.Printf("Saved: %s", newFileName)

}
