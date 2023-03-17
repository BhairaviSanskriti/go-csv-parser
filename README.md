# go-csv-parser
This Go program reads a CSV file of people's information, parses its contents into a slice of Person structs, and displays the data in an HTML table using a web app. 
## Usage
This app runs on port 8081 and is deployed on this link: 
```
http://ec2-65-1-223-223.ap-south-1.compute.amazonaws.com:8081/
```
## CSV File Format
The CSV file should have the following format:
```
Name,Age,City
Harry Potter,14,South Hampton
Dumbledore,60,Paris
Hermoine Granger,13,London
Ron,14,Amsterdam
```
The first row should contain the column headers, and each subsequent row should contain the data for each person.

## Conclusion:
In this project, we have developed a Go program that reads data from a CSV file, parses it into a slice of structs, and serves the data as an HTML table through an HTTP server. We have used the encoding/csv and html/template packages to achieve this.
