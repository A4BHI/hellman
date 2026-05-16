package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
)

type hellman struct {
	p           int
	g           int
	privatekeys map[string]int
	publickeys  map[string]int
	sharedkeys  map[string]int
}

func cleanPrefix(key int) int {
	val := strconv.Itoa(key)
	var tmp string
	for i := 0; i < len(val); i++ {
		if val[i] == '-' {
			continue
		}

		tmp += string(val[i])
	}

	key, _ = strconv.Atoi(tmp)
	return key
}

func (hell *hellman) Calculate() {

	hell.calculatePublicKey()
	hell.calcuateSharedKeys()
	fmt.Println("PUBLIC KEY OF ALICE : ", hell.publickeys["alice"])
	fmt.Println("PUBLIC KEY OF BOB : ", hell.publickeys["bob"])
	fmt.Println("SHARED SECRET KEY : ", hell.sharedkeys["alice"])
}

func (hell *hellman) calculatePublicKey() {

	for v, _ := range hell.privatekeys {
		pr := hell.privatekeys[v]
		raiseto := int(math.Pow(float64(hell.g), float64(pr)))

		if raiseto < hell.p {
			hell.publickeys[v] = raiseto
			continue
		}

		div := raiseto / hell.p
		val := strconv.Itoa(div)

		convertedval, _ := strconv.Atoi(val)

		hell.publickeys[v] = cleanPrefix((convertedval * hell.p) - raiseto)

	}

}

func (hell *hellman) calcuateSharedKeys() {

	var public int

	var slice []string
	for v, _ := range hell.privatekeys {

		slice = append(slice, v)
	}
	for _, v := range slice {
		if v == "alice" {
			public = hell.publickeys["bob"]
		} else {
			public = hell.publickeys["alice"]
		}
		pr := hell.privatekeys[v]
		raiseto := int(math.Pow(float64(public), float64(pr)))
		if raiseto < hell.p {
			hell.sharedkeys[v] = raiseto
			continue
		}
		div := raiseto / hell.p

		val := strconv.Itoa(div)

		convertedval, _ := strconv.Atoi(val)

		hell.sharedkeys[v] = cleanPrefix((convertedval * hell.p) - raiseto)

	}
}
func main() {

	hell := hellman{
		privatekeys: make(map[string]int),
		publickeys:  make(map[string]int),
		sharedkeys:  make(map[string]int),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		<-done
		fmt.Println("\nProgram Exited")
		os.Exit(0)
	}()

	var privatekeys int
	for {
		fmt.Print("Enter G: ")
		fmt.Scanln(&hell.g)
		fmt.Print("Enter P: ")
		fmt.Scanln(&hell.p)

		fmt.Print("Enter Private Key Of Alice: ")
		fmt.Scanln(&privatekeys)
		hell.privatekeys["alice"] = privatekeys

		fmt.Print("Enter Private Key Of Bob: ")
		fmt.Scanln(&privatekeys)
		hell.privatekeys["bob"] = privatekeys

		fmt.Println(hell.privatekeys)

		hell.Calculate()

	}

}
