package main

import (
	"encoding/csv"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// Open the CSV file
	file, err := os.Open("people.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Create a slice of Person structs
	var people []Person

	// Iterate over the records and add them to the slice
	for _, record := range records {
		age, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		person := Person{
			Name: record[0],
			Age:  age,
			City: record[2],
		}
		people = append(people, person)
	}

	// Create an HTML template
	tmpl := template.Must(template.New("index").Parse(`
		<!doctype html>
		<html>
			<head>
				<title>List</title>
				<style>
					table {
						border-collapse: collapse;
					}
					th, td {
						padding: 10px;
						border: 1px solid black;
					}
					th {
						background-color: #ddd;
					}
				</style>
			</head>
			<body>
				<h1>People List</h1>
				<table>
					<tr>
						<th>Name</th>
						<th>Age</th>
						<th>City</th>
					</tr>
					{{range .}}
						<tr>
							<td>{{.Name}}</td>
							<td>{{.Age}}</td>
							<td>{{.City}}</td>
						</tr>
					{{end}}
				</table>
			</body>
		</html>
	`))

	// Define an HTTP handler function to serve the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":8081", nil)
}
