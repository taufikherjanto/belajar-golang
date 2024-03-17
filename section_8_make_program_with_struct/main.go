package main

import (
	"bufio"
	"fmt"
	"os"
)

type Mahasiswa struct {
	Nama    string
	Nim     int
	Jurusan string
}

func tambahMahasiswa(listMahasiswa []Mahasiswa, nama string, nim int, jurusan string) []Mahasiswa {
	dataMahasiswa := Mahasiswa{Nama: nama, Nim: nim, Jurusan: jurusan}
	return append(listMahasiswa, dataMahasiswa)
}

func hapusMahasiswa(listMahasiswa []Mahasiswa, index int) []Mahasiswa {
	// ambil list array dengan index x, listMahasiswa[:index]
	// diappend ke list array listMahasiswa[index + 1]

	/*
		1.	list[:index]
			Ini adalah penggunaan slicing untuk membuat slice baru yang berisi elemen-elemen people dari indeks 0 hingga index-1.
			Dengan kata lain, ini adalah bagian pertama dari slice people sebelum indeks yang ingin dihapus.

		2.	list[index+1:]...
			Ini adalah penggunaan slicing lagi untuk membuat slice baru yang berisi elemen-elemen people mulai dari indeks index+1 hingga elemen terakhir.
			Dengan demikian, ini adalah bagian kedua dari slice people setelah indeks yang ingin dihapus.
			Operator titik tiga kali (...) di sini digunakan untuk membuka slice tersebut sehingga setiap elemen slice people[index+1:] dimasukkan ke dalam fungsi append() secara individual.

		3.	append(list[:index], list[index+1:]...):
			Kode ini menggunakan fungsi append() untuk menggabungkan kedua bagian slice yang dibuat sebelumnya menjadi satu slice baru.
			Dengan cara ini, elemen yang ada di antara people[:index] dan people[index+1:] dihilangkan dari slice people.
	*/

	return append(listMahasiswa[:index], listMahasiswa[index+1:]...) // masih jadi pertanyaan kenapa harus diikut ...
}

func tampilkanData(listMahasiswa []Mahasiswa) {
	fmt.Println("\n\nList Data Mahasiswa: ")
	for index, mahasiswa := range listMahasiswa {
		fmt.Printf("%d. %s \t\t %d \t\t %s \n", index+1, mahasiswa.Nama, mahasiswa.Nim, mahasiswa.Jurusan)
	}
}

func main() {
	var dataMahasiswa []Mahasiswa
	var pilihan int
	var inputNama string
	var inputNim int
	var inputJurusan string
	var inputIndex int

	for {
		fmt.Println("\n\nPilih:")
		fmt.Println("1. Lihat Data Mahasiswa")
		fmt.Println("2. Tambah Mahasiswa")
		fmt.Println("3. Hapus mahasiswa")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Error baca input: ", err)
			return
		}

		switch pilihan {
		case 1:
			tampilkanData(dataMahasiswa)
		case 2:
			fmt.Print("Masukkan Nama: ")
			reader := bufio.NewReader(os.Stdin)
			inputNama, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error baca input: ", err)
				return
			}
			//inputNama = input

			fmt.Print("Masukkan NIM: ")
			_, err = fmt.Scanln(&inputNim)
			if err != nil {
				fmt.Println("Error baca input: ", err)
				return
			}

			fmt.Print("Masukkan jurusan: ")
			reader = bufio.NewReader(os.Stdin)
			inputJurusan, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error baca input: ", err)
				return
			}
			//inputJurusan = input

			dataMahasiswa = tambahMahasiswa(dataMahasiswa, inputNama, inputNim, inputJurusan)
			fmt.Print("Data mahasiswa telah ditambahkan: ", dataMahasiswa)
		case 3:
			fmt.Print("Pilih nomor index yang akan dihapus: ")
			_, err := fmt.Scanln(&inputIndex)
			if err != nil {
				fmt.Println("Error baca input: ", err)
				return
			}

			dataMahasiswa = hapusMahasiswa(dataMahasiswa, inputIndex-1)
			fmt.Println("Data telah dihapus")
		case 4:
			fmt.Print("Keluar...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
