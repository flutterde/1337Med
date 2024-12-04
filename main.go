package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"goreplace/api"
	"goreplace/handlers"
)

func createHeader(name string) {
	log.Println("Creating Header " + name + " file...")
	fileName := name + ".hpp"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	headerContent := fmt.Sprintf(`#ifndef %s_HPP
# define %s_HPP
# include <iostream>

class %s
{
	private:

	public:
		%s(void); // Default constructor
		%s(const %s& obj); // Copy constructor
		%s&	operator=(const %s& obj); // Copy assignment operator
		~%s(); // Destructor

};

#endif
`, 
	strings.ToUpper(name),
	strings.ToUpper(name),
	name,
	name,
	name,
	name,
	name,
	name,
	name)

	_, err = file.WriteString(headerContent)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Header file created successfully:", fileName)
}

func createClass(name string) {

	log.Println("Creating Class " + name + " file...")
	fileName := name + ".cpp"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()
	classContent := fmt.Sprintf(`#include "%s.hpp"

%s::%s(void)
{
	std::cout << "%s Default constructor called" << std::endl;
}

%s::~%s(void)
{
	std::cout << "%s Destructor called" << std::endl;
}

%s::%s(const %s& obj)
{
	std::cout << "%s Copy constructor called" << std::endl;
	*this = obj;
}

%s& %s::operator=(const %s& obj)
{
	std::cout << "%s Copy assignment operator called" << std::endl;
	if (this != &obj)
	{
		// Add any assignment logic here
	}
	return (*this);
}
`, 
	name,
	name,
	name,
	strings.Title(name),
	name,
	name, 
	strings.Title(name),
	name,
	name,
	strings.Title(name),
	name,
	name,
	strings.Title(name),
	name,
	name)

	_, err = file.WriteString(classContent)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Class file created successfully:", fileName)
}

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Println("please set a flag")
		os.Exit(1)
	}
	flags := os.Args[1]
	if (flags == "-lt") {
		var name string
		if (len((os.Args)) > 2) {
			name = os.Args[2]
		} else {name = os.Getenv("USER")}
			res, err := api.GetLogTime(name)
			if err != nil {
				log.Fatalf("Error While getting your logtime: %v\n", err); os.Exit(1)
			} else {
				fmt.Println("Your log time is", res); os.Exit(0)
			}
	} else if (flags == "-mlx") {
		rlt := handlers.MlxHandler()
		if rlt == -1 {
			os.Exit(1)
		} else {os.Exit(0)}
	}
	if flags == "-h" {
		fmt.Println("-ch	Create Class + Header file\n-c	Create Class only\n-lt	Get Log time\n-mlx	install MLX")
		os.Exit(0)
	}
	if (argsLen < 3) {
		log.Fatal("Run the command like that: cppreate -ch MyClass"); os.Exit(1)
	}
	fileName := os.Args[2]
	if (flags == "" || (flags == "-ch" && fileName == "")) {

		log.Fatal("Run the command like that: cppreate -ch MyClass"); os.Exit(1)
	} else if flags == "-ch" {
		createClass(fileName)
		createHeader(fileName)
	} else if flags == "-c" {
		createClass(fileName)
	}

}
