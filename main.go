package main

import (
	"fmt"

	"rsc.io/quote"
)

func main(){
	fmt.Print("Hello world");
	fmt.Println("I");//prints and moves execution to next line
	fmt.Println("am");
	fmt.Println("Groot");
	fmt.Print("I am ");
	fmt.Println("Ironman");
	var villain="Thanos";
	fmt.Println("----------------------------");
	fmt.Println("i am ---- ",villain);
	fmt.Println("----------------------------");
	const hero="Tony Stark";
	fmt.Printf("%v said i am inevitable\n",villain)//%v represents it is special placeholder for variable which is declared later after comma
	fmt.Printf("%v said i am Iron man\n",hero)
	    fmt.Println(quote.Go())
}