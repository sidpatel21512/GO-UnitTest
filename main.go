package main

import "fmt"

func main(){
	var a int
	var b int
	fmt.Printf("Enter two numbers: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil{		
		fmt.Println("Error while scanning numbers")
	}
	v := lcm(a,b)
	fmt.Printf("LCM of %d and %d is: %v \n", a, b, v)
}

// Function to return LCM of two numbers
func lcm(a int, b int)int {
	if a == 0 || b == 0 {
		return 0;
	}
	if a == 1 {
		return b;
	}
	if b == 1 {
		return a;
	}
    return (a / gcd(a, b)) * b;
}

// Recursive function to return gcd of a and b
func gcd(a int, b int) int {
    if a == 0{
        return b;
	}
    return gcd(b % a, a);
}