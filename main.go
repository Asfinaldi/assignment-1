package main

import (
	"fmt"
	"os"
	"strconv"
)

type pelajar struct {
	name string
	address string
	profession string
	argument string
}

func main() {
	var pelajar []pelajar = []pelajar{
		{
			name: "Asfinaldi",
			address: "Padang Pariaman",
			profession: "Frelance",
			argument: "Karena saya tertarik menggunakan golang",
		},
		{
			name: "Sammi Aldhi Yanto",
			address: "Depok",
			profession: "Mahasiswa",
			argument: "Ingin menambah pengalaman dalam bidang backend",
		},
		{
			name: "Igram Mullah",
			address: "Jakarta",
			profession: "IT",
			argument: "golang mulai banyak digunakan",
		},
		{
			name: "Krisna Madani",
			address: "Medan",
			profession: "Mahasiswa",
			argument: "Karena tertarik menggunakan goalang di backend",
		},
		{
			name: "Nugraha",
			address: "Bandung",
			profession: "Frond End",
			argument: "Saya tertarik menggunakan microservis di golang",
		},	
	}

	if len(os.Args) > 1 {
		absen, err := strconv.Atoi(os.Args[1])
		_ = err
		checkAbsen(absen, pelajar)
	} else {
		fmt.Println("Anda tidak menginput nomor absen!")
	}
}

func checkAbsen(absen int, slc []pelajar) {
	if absen <= len(slc) && absen > 0 {
		printBiodata(absen ,slc)
	} else {
		printError(len(slc))
	}
}

func printBiodata(absen int, slc []pelajar) {
	fmt.Printf("Absen ke-%d\n\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih Golang: %s", 
	absen, slc[absen - 1].name, slc[absen - 1].address, slc[absen - 1].profession, slc[absen - 1].argument);
}

func printError(lengthValid int) {
	fmt.Printf("Maaf, absen hanya dari 1 hingga %d", lengthValid)
}