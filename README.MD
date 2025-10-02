# Grade Calculator

This is a library for calculating final grades based on different assignments for different classes. 

## Features 
- Add multiple grades based on student's grades recieved for assignments, exams, and essays. 
- Calculate weighted averages: Assignments 50%, Exams 35%, and Essays 15%
- Match grade percentages into letter grades using this scale: 
  - A >= 90
  - B >= 80
  - C >= 70
  - D >= 60 
  - F < 60 

## Tech Stack
- Language: Go 
- Compatible IDE: GoLang 1.24+ 
- Testing: Built in Go testing package

## Project Structure 
- go.mod: version and module 
- grade_calculator.go: main logic for grade calculation
- grade_calculator_test.go: unit tests for grade calculation
- README.md: instruction for usage and documentation 

## Setup 
After cloning this repo: 

Run unit tests:
```
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade-calculator/
go test -coverprofile=coverage.txt ./...
```
