package main

import (
	"fmt"
)

func main()  {
	countryCapitalMap :=map[string]string{"France":"Paris","Italy":"Rome"}
	fmt.Println("old map")
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
	delete(countryCapitalMap,"France")
	fmt.Println("France delete")
	fmt.Println("new map")
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
}
