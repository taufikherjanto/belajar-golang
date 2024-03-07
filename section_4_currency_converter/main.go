package main

import "fmt"

const (
	// rate konversi dollar ke rupiah
	dollarToRupiah = 15000
	// maksimal dollar yang boleh ditukar
	maxDollarExchange = 100
)

func main() {
	saldoDollar := 120.0

	fmt.Println("Saldo dollar awal: ", saldoDollar)
	fmt.Println("Saldo rupiah: ", exchangeDollarToRupiah(saldoDollar))
}

func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxDollarExchange {
		fmt.Println("Maaf tidak bisa menukar lebih dari ", maxDollarExchange)
		return dollar
	}

	rupiah := dollar * dollarToRupiah
	return rupiah
}
