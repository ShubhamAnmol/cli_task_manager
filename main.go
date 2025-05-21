package main 

import (
	"log"
	"flag"
)

func main() {
	var firstname string
	var secondname string

	flag.StringVar(&firstname,"firstname","Anmol","Provide your first name")
	flag.StringVar(&secondname,"secondname","Shubham","Provide your second name")
	flag.Parse()
	log.Println("First name is ",firstname)
	log.Println("Second name is",secondname)
	
}