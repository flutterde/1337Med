package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func createHeader(name string) {
	log.Println("Creating Header " + name + " file...")
	fileName := strings.ToLower(name) + ".hpp"
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
		%s();
		%s(const %s& inst);
		%s&	operator=(const %s& inst);
		~%s();

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
	fileName := strings.ToLower(name) + ".cpp"
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

%s::%s(const %s& inst)
{
	std::cout << "%s Copy constructor called" << std::endl;
}

%s& %s::operator=(const %s& inst)
{
	std::cout << "%s Copy assignment operator called" << std::endl;
	if (this != &inst)
	{
		// Add any assignment logic here
	}
	return (*this);
}
`, 
	strings.ToUpper(name),
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
	name)

	_, err = file.WriteString(classContent)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Class file created successfully:", fileName)
}

func main() {
	if len(os.Args) < 2 {
		log.Println("please set a flag")
		os.Exit(1)
	}
	flags := os.Args[1]

	if flags == "-h" {
		fmt.Println("-ch	Create Class + Header file\n-c	Create Class only")
		os.Exit(0)
	}
	if (len(os.Args) < 3) {
		log.Fatal("Run the command like that: cppreate -ch MyClass"); os.Exit(1)
	}
	fileName := os.Args[2]
	if flags == "" || fileName == "" {
		log.Fatal("Run the command like that: cppreate -ch MyClass"); os.Exit(1)
	} else if flags == "-ch" {
		createClass(fileName)
		createHeader(fileName)
	} else if flags == "-c" {
		createClass(fileName)
	}
}
