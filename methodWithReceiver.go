package main
import "fmt"
//Author struct
type author struct{
	name string
	branch string
	articles int
	salary int
}

//Method with a receiver
func (a author)show(){
	fmt.Println("Author's Name:",a.name)
	fmt.Println("Author's Branch:",a.branch)
	fmt.Println("Author's Articles:",a.articles)
	fmt.Println("Author's Salary:",a.salary)
}

type Mutable struct{
	a int
	b int
}

func(m Mutable) StayTheSame(){
	m.a = 10
	m.b = 20
}
func(m *Mutable)Mutate(){
	m.a = 100
	m.b = 200
}

func main(){
	res := author{
		name: "Sooraj",
		branch: "CSE",
		articles: 10,
		salary: 100000,
	}
	res.show()
	M := &Mutable{9, 0}
	fmt.Println(m)
	M.StayTheSame()
	fmt.Println(m)
	m.Mutate()
	fmt.Println(m)
	}