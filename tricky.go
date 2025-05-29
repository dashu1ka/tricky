package main

import (
	"fmt"
	"math/big"
)

func main() {
	val := make(map[string]struct{})
	ten := big.NewInt(10)
	for i := int64(1); i < 1000; i++ {
		if i%10 == 0 {
			continue
		}
		part := big.NewInt(i)
		for j := 0; j < 30; j++ {
			value := new(big.Int).Exp(ten, big.NewInt(int64(j)), nil)
			n := new(big.Int).Mul(part, value)
			fn := f(n)
			ffn := f(fn)
			g := new(big.Rat).SetFrac(ffn, n)
			val[g.RatString()] = struct{}{}
		}
	}

	fmt.Println("count of g(n) dif val: ", len(val))
}

func f(n *big.Int) *big.Int {
	s := n.String()
	runes := []rune(s)
	l := len(runes)
	for i := 0; i < l/2; i++ {
		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	}
	revS := string(runes)
	revInt := new(big.Int)
	revInt.SetString(revS, 10)
	return revInt
}
