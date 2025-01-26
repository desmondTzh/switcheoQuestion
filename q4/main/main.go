package main

import "fmt"

func main() {

    n := 10

    fmt.Printf("Results for method A: %d\n", sum_to_n_a(n))
    fmt.Printf("Results for method B: %d\n", sum_to_n_b(n))
    fmt.Printf("Results for method C: %d\n", sum_to_n_c(n))
}



/*
This method is the most efficient way to calculate the sum of the all digit from 1 to n.
But only any modification to the requirement, would require the rework of this function.
    Complexity of O(1)

*/
func sum_to_n_a(n int) int {
	return  (n * (n + 1)) / 2
}

/*
This method is the easiest to understand way of calculating the sum of the all digit from 1 to n, as it most humanly readable.
The efficiency is lacking but allow easier modification for other similar cases (Example: adding all odd number from 1 to n).
    Complexity of O(n)
*/
func sum_to_n_b(n int) int {
    sum := 0
    for i := 1; i <= n; i++{
        sum += i
    }
    return sum
}

/*
This function is the uses recursive to loop each number and provide the summation of all digit from 1 to n.
The only benefit this method brings is the ability to customize each digit handling logic in the recursive function.
    Complexity of O(n)
*/
func sum_to_n_c(n int) int {
	return sum_to_n_c_recursive(n)
}
func sum_to_n_c_recursive(n int) int {
    if n == 1 {
        return 1
    }
    return n + sum_to_n_c_recursive(n-1)
}