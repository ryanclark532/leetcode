package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, _ := http.Get("https://au.indeed.com/jobs?q=software+developer&l=Sydney+NSW&oq=software+developer+%24100%2C000&ol=sydney+nsw&ts=1727227189215&pts=1727061385273&salaryType=%24100%2C000&sc=0kf%3Ajt%28fulltime%29%3B&from=searchOnHP&rq=1&rsIdx=0&fromage=last&vjk=8de4664121038748")

	fmt.Println(response)
}
