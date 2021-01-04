package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
)

type publicKey struct {
	N, E *big.Int
}

func generatePair() (*big.Int, *big.Int) {
	p, err := rand.Prime(rand.Reader, 1024)
	for err != nil {
		p, err = rand.Prime(rand.Reader, 1024)
	}

	q, err := rand.Prime(rand.Reader, 1024)
	for err != nil {
		q, err = rand.Prime(rand.Reader, 1024)
	}

	if p == q {
		return generatePair()
	} else {
		return p, q
	}
}

func generateKeys(n int, keys chan publicKey, quit chan bool) {
	for i := 0; i < n; i++ {
		p, q := generatePair()

		n := &big.Int{}
		n.Mul(p, q)

		keys <- publicKey{
			N: n,
			E: big.NewInt(65537),
		}
	}

	quit <- true
}

func main() {
	keysChan := make(chan publicKey)
	quit := make(chan bool)

	n := 8
	for i := 0; i < n; i++ {
		go generateKeys(1000/n, keysChan, quit)
	}

	i := 0
	keys := make([]publicKey, 1002)

	nrClosed := 0
	for nrClosed < n {
		fmt.Printf("\rProgress : %.1f%%", float64(i)/1000*100)
		select {
		case <-quit:
			nrClosed += 1
		case key := <-keysChan:
			keys[i] = key
			i++
		}
	}
	fmt.Println()

	p, q := generatePair()
	r, err := rand.Prime(rand.Reader, 1024)
	for err != nil {
		r, err = rand.Prime(rand.Reader, 1024)
	}

	fmt.Println(i)

	n2 := &big.Int{}
	n2.Mul(p, q)

	n3 := &big.Int{}
	n3.Mul(p, r)

	keys[i] = publicKey{
		N: n2,
		E: big.NewInt(65537),
	}
	keys[i+1] = publicKey{
		N: n3,
		E: big.NewInt(65537),
	}

	f, _ := os.Create("keys.csv")
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"n", "e"})
	for _, k := range keys {
		w.Write([]string{k.N.String(), k.E.String()})
	}
	w.Flush()

	fmt.Printf("%s\n", p)
	fmt.Printf("%s\n", q)
}
