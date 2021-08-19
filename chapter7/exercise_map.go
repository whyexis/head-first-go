package main

import "fmt"

func main() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Printf("%9s %.2f\n", "Earrings:",jewelry["earrings"])
	fmt.Printf("%9s %.2f\n", "Necklace:", jewelry["necklace"])
	fmt.Printf("%9s %.2f\n", "Shirt:", clothing["shirt"])
	fmt.Printf("%9s %.2f\n", "Pants:", clothing["pants"])
}