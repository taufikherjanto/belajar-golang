package main

import (
	"fmt"
	"strings"
)

const (
	// maksimal dollar yang boleh ditukar
	maxResultExchange = 1000
)

// list mata uang yang tersedia
var listMataUang = [4]string{"USD", "EUR", "GBP", "JPY"}

// rate mata uang
var rateMataUang = []float64{0.91, 0.78, 148.16, 0.86, 161.98, 189.43}

// mapping convert currency
var mappingNilaiTukar = map[string]float64{
	"USDEUR": rateMataUang[0],
	"USDGBP": rateMataUang[1],
	"USDJPY": rateMataUang[2],
	"EURGBP": rateMataUang[3],
	"EURJPY": rateMataUang[4],
	"GBPJPY": rateMataUang[5],
	"EURUSD": 1 / rateMataUang[0],
	"GBPUSD": 1 / rateMataUang[1],
	"JPYUSD": 1 / rateMataUang[2],
	"GBPEUR": 1 / rateMataUang[3],
	"JPYEUR": 1 / rateMataUang[4],
	"JPYGBP": 1 / rateMataUang[5],
}

func main() {

	var diulangi string
	var jumlahUang float64
	var mataUang string
	var mataUangTujuan string
	var hasilExchange float64
	var persentasePerubahan float64

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

		if !checkCurrency(listMataUang, strings.ToUpper(mataUang)) {
			fmt.Println("Mata uang asal tidak tersedia")
			break
		}

		fmt.Println("Input mata uang tujuan:")
		_, err = fmt.Scanln(&mataUangTujuan)
		if err != nil {
			fmt.Println("error input tidak valid")
			break
		}

		if !checkCurrency(listMataUang, strings.ToUpper(mataUangTujuan)) {
			fmt.Println("Mata uang tujuan tidak tersedia")
			break
		}

		fmt.Println("Input persentase perubahan:")
		_, err = fmt.Scanln(&persentasePerubahan)
		if err != nil {
			fmt.Println("error input tidak valid")
			break
		}

		hasilExchange = convertCurrency(strings.ToUpper(mataUang), strings.ToUpper(mataUangTujuan), jumlahUang, persentasePerubahan)

		fmt.Println("-----------------------------")
		fmt.Println("Jumlah uang : ", jumlahUang)
		fmt.Println("Mata uang: ", mataUang)
		fmt.Println("Mata uang tujuan: ", mataUangTujuan)
		fmt.Printf("Persentase Perubahan: %.2f%%\n", persentasePerubahan)
		fmt.Println("Hasil konversi: ", hasilExchange)

		if hasilExchange > maxResultExchange {
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

func checkCurrency(listCurrency [4]string, mataUang string) bool {
	// iterate array dengan looping
	for i := 0; i < len(listCurrency); i++ {
		if listCurrency[i] == mataUang {
			return true
		}
	}
	return false
}

func convertCurrency(mataUang string, mataUangTujuan string, uang float64, persentase float64) float64 {
	var hasilExchange float64
	var nilaiTukarBaru float64
	if mataUang == mataUangTujuan {
		hasilExchange = uang
	} else {
		hasilExchange = uang * mappingNilaiTukar[mataUang+mataUangTujuan]
	}

	// hitung konversi mata uang berdasarkan persentase perubahan nilai tukar yang diberikan
	nilaiTukarBaru = hasilExchange + (hasilExchange * persentase / 100)

	return nilaiTukarBaru
}
