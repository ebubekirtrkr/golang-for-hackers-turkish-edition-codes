package examples

import "fmt"

func Loop_main() {
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	/*
		for { //sonsuz döngü

		}
	*/
	switch i {
	case 1:
		fmt.Println("1")
		fallthrough //devamke
	case 2:
		fmt.Println("2")
	}
	haritalar := map[string]int{}

	haritalar["Ankara"] = 6
	haritalar["Bolu"] = 14
	for k, v := range haritalar {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	if _, ok := haritalar["İstanbul"]; ok {
		fmt.Println("İstanbul var")
	} else {
		fmt.Println("İstanbul yok")
	}

}
