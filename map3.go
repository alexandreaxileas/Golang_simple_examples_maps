package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	funcMApPorLetra := map[string]map[string]float64{
		"A": {"Ana": 1500.20,
			"Alexandre": 6540.08},
		"B": {"Benedito": 6400.80,
			"Baltazar": 3500.00,
			"Breno":    3500.00},
		"C": {"Carlos": 4860.77,
			"Carina": 7202.14}}

	//fmt.Println(funcMApPorLetra["B"]["Benedito"]) //fmt.Println("-------------")

	fmt.Println("-------------")
	fmt.Println("-------------")
	// Atribuindo um Slice para cada letra do Map
	s := make([]string, 0)
	for letra, _ := range funcMApPorLetra {
		for key, valor2 := range funcMApPorLetra[letra] {
			fmt.Println(letra, key, valor2)
			s = append(s, key)
		}
	}

	sort.Strings(s)
	//fmt.Println(sort.SearchStrings(s, "C"))
	fmt.Println("-------------")
	fmt.Println("-------------")
	fmt.Println(strings.Join(s, "-"))
	fmt.Println("-------------")
	fmt.Println("-------------")
	keys := make([]string, 0, len(funcMApPorLetra)) // criou um slice
	for k := range funcMApPorLetra {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	fmt.Println(strings.Join(keys, "-"))

}
