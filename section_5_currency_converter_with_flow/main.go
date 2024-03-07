package main

import (
	"fmt"
	"strings"
)

// maksimal dollar yang boleh ditukar
//maxResultExchange = 1000

func main() {

	var diulangi string
	var jumlahUang float64
	var mataUang string
	var mataUangTujuan string
	var hasilExchange float64

	for {
		fmt.Println("#############################")
		fmt.Println("Input jumlah uang:")

		_, err := fmt.Scanln(&jumlahUang)
		if err != nil {
			fmt.Println("error input tidak valid")
			break
		}

		fmt.Println("Input mata uang:")
		_, err = fmt.Scanln(&mataUang)
		if err != nil {
			fmt.Println("error input tidak valid")
			break
		}

		fmt.Println("Input mata uang tujuan:")
		_, err = fmt.Scanln(&mataUangTujuan)
		if err != nil {
			fmt.Println("error input tidak valid")
			break
		}

		if !checkCurrency(strings.ToUpper(mataUang)) {
			fmt.Println("Mata uang asal tidak tersedia")
			break
		}

		if !checkCurrency(strings.ToUpper(mataUangTujuan)) {
			fmt.Println("Mata uang tujuan tidak tersedia")
			break
		}

		hasilExchange = exchangeMoney(strings.ToUpper(mataUang), strings.ToUpper(mataUangTujuan), jumlahUang)

		fmt.Println("-----------------------------")
		fmt.Println("Jumlah uang : ", jumlahUang)
		fmt.Println("Mata uang: ", mataUang)
		fmt.Println("Mata uang tujuan: ", mataUangTujuan)
		fmt.Println("Hasil konversi: ", hasilExchange)

		if hasilExchange > 1000 {
			fmt.Println("Warning, hasil konversi terlalu besar!!!")
		}

		fmt.Println("-----------------------------")
		fmt.Println("Konversi apakah diulangi lagi, ketik Y jika diulangi: ")

		_, err = fmt.Scanln(&diulangi)
		if err != nil {
			fmt.Println("error input")
			break
		}

		if strings.ToUpper(diulangi) != "Y" {
			break
		}
	}
}

func checkCurrency(mataUang string) bool {
	if mataUang == "USD" || mataUang == "EUR" || mataUang == "GBP" || mataUang == "JPY" {
		return true
	} else {
		return false
	}
}

func exchangeMoney(mataUang string, mataUangTujuan string, uang float64) float64 {
	var hasilExchange float64
	if mataUang == "USD" && mataUangTujuan == "EUR" {
		hasilExchange = uang * 0.91
	} else if mataUang == "USD" && mataUangTujuan == "GBP" {
		hasilExchange = uang * 0.78
	} else if mataUang == "USD" && mataUangTujuan == "JPY" {
		hasilExchange = uang * 148.16
	} else if mataUang == "EUR" && mataUangTujuan == "USD" {
		hasilExchange = uang / 0.91
	} else if mataUang == "EUR" && mataUangTujuan == "GBP" {
		hasilExchange = uang * 0.86
	} else if mataUang == "EUR" && mataUangTujuan == "JPY" {
		hasilExchange = uang * 161.98
	} else if mataUang == "GBP" && mataUangTujuan == "USD" {
		hasilExchange = uang / 0.78
	} else if mataUang == "GBP" && mataUangTujuan == "EUR" {
		hasilExchange = uang / 0.86
	} else if mataUang == "GBP" && mataUangTujuan == "JPY" {
		hasilExchange = uang * 189.43
	} else if mataUang == "JPY" && mataUangTujuan == "USD" {
		hasilExchange = uang / 148.16
	} else if mataUang == "JPY" && mataUangTujuan == "EUR" {
		hasilExchange = uang * 161.98
	} else if mataUang == "JPY" && mataUangTujuan == "GBP" {
		hasilExchange = uang / 189.43
	} else if mataUang == mataUangTujuan {
		hasilExchange = uang
	}

	return hasilExchange
}
