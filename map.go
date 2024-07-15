package main
import "fmt"

func main(){
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Println(a)
	fmt.Println(b)

var c = make(map[string]string) //The map is empty now
c["brand"] = "Ford"
c["model"] = "Mustang"
c["year"] = "1964"

d := make(map[string]int)
d["Oslo"] = 1
d["Bergen"] = 2
d["Trondheim"] = 3
d["Stavanger"] = 4

fmt.Println(c)
fmt.Println(d)

var e = make(map[string]string)
e["brand"] = "Renault"
e["model"] = "Duster"
e["year"] = "2018"

fmt.Println("Value of Brand in Map a:", a["brand"])
a["brand"] = "Kia"
fmt.Println("Value of Brand in Map a:", a["brand"])

val1, ok1 :=a["brand"] //Checking for existing key and its value
val2, ok2 :=a["color"] //Checking for non-existing key and its value
val3, ok3 :=a["day"] //Checking for existing key and its value
_, ok4 :=a["year"] //Only checking for existing key and not its value

fmt.Println(val1, ok1)
fmt.Println(val2, ok2)
fmt.Println(val3, ok3)
fmt.Println(ok4)
}