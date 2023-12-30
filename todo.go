package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	
	
	if len(os.Args) != 1 {
		if os.Args[1] == "del"{
			if len(os.Args) != 3 {
				fmt.Println("del requires one argument")
				os.Exit(0)
			} else {
				dat, err := os.ReadFile("todo.txt")
				if err != nil {
					panic(err)
				}
				n, err := strconv.Atoi(os.Args[2])
				if err != nil {
					panic(err)
				}
				d := strings.Split(string(dat), "\n")
				os.Truncate("todo.txt", 0)	
				f, err := os.OpenFile("todo.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //move
				if err != nil {
					panic(err)
				}
				for i := 0; i < len(d); i++ {
					if i != n {
						if i != 0 && n != 0{
							f.WriteString("\n")	
						}
						f.WriteString(d[i])
					}
				}
				f.Close() 
			}
		} else {
			f, err := os.OpenFile("todo.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //move
			if err != nil {
				panic(err)
			}
			if o, _ := os.Stat("todo.txt"); o.Size() != 0 {	
				f.WriteString("\n")
			}
		
			for i := 1; i < len(os.Args); i++ {
				//fmt.Println(os.Args[i])
	
				f.WriteString(os.Args[i])
				f.WriteString(" ")
			} 
			f.Close()
		}
	}
	dat, err := os.ReadFile("todo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
}