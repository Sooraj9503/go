package main

import "fmt"
import "App/Database"

func main(){
	_, err := database.DatabaseConnection()
	if err != nil{
		fmt.Println("Database conn err:", err.Error())
	}
}