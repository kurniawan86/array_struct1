package view

import (
	"array_struct1/model"
	"array_struct1/node"
	"bufio"
	"fmt"
	"os"
)

func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai: ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pegawai: ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan Jalan: ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Nomer rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan Nomor Telepon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan Email: ")
	fmt.Scan(&email)

	// create new pegawai
	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email:  email,
	}

	// insert to DaftarPegawai
	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("== Pegawai berhasil ditambahkan ==")
	} else {
		fmt.Println("Pegawai gagal ditambahkan")
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke -", i+1, " ---")
		fmt.Println("ID Pegawai\t: ", emp.ID)
		fmt.Println("Nama Pegawai\t: ", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("No Telepon\t: ", emp.NoTelp)
		fmt.Println("Email\t\t: ", emp.Email)
		fmt.Println()
	}
}
