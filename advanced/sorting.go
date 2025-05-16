package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type PersonSorter struct {
	people []Person
	by     By
}

func (s *PersonSorter) Len() int { 
	return len(s.people)
}
func (s *PersonSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}	
func (s *PersonSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}
func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func sorting() {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(age).Sort(people)
	fmt.Println("Sorting by age:", people)
	By(name).Sort(people)
	fmt.Println("Sorting by name:", people)

	// numbers := []int{5, 2, 9, 1, 5, 6}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)	
	
	// strSlice := []string{"banana", "apple", "cherry"}
	// sort.Strings(strSlice)
	// fmt.Println("Sorted strings:", strSlice)


}