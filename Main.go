/* All code must belong to a package
The first statement in GO file must be a package*/
//A package is a collection of source files// 
package main

/* All the packages need to be defined for all the commands in the code */
import "fmt" //Stands for format [Built-in package]
/* Define an entry point, to let compiler know where to start executing */
func main() {
	fmt.Println("Hello World")
}
