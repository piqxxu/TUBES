package main

import "fmt"

const NMAX = 20

type karyawan struct {
	nama, posisi, departemen string
	nik, startKerja          int
	cuti                     int
}

type tabKaryawan [NMAX]karyawan //alias

func header() {
	/* I.S. -
	   F.S. menampilkan header dari aplikasi */
	fmt.Println("======================================")
	fmt.Println("||   Sistem Manajemen Data Karyawan ||")
	fmt.Println("||      Algoritma Pemrograman       ||")
	fmt.Println("||               By:                ||")
	fmt.Println("||   Muhammad Azwa Radya Razadisya  ||")
	fmt.Println("||            103022300023          ||")
	fmt.Println("||     Nadya Sekar Rahmawati        ||")
	fmt.Println("||            103022300020          ||")
	fmt.Println("======================================")
}

//dummy
// func dummy(karyawanList *tabKaryawan) {
// 	karyawanList[0] = karyawan{
// 		nama:       "Budi",
// 		posisi:     "Manager",
// 		departemen: "Keuangan",
// 		nik:        101,
// 		startKerja: 2015,
// 		cuti:       10,
// 	}
// 	karyawanList[1] = karyawan{
// 		nama:       "Siti",
// 		posisi:     "Staff",
// 		departemen: "HRD",
// 		nik:        102,
// 		startKerja: 2017,
// 		cuti:       12,
// 	}
// 	karyawanList[2] = karyawan{
// 		nama:       "Joko",
// 		posisi:     "Supervisor",
// 		departemen: "Operasional",
// 		nik:        103,
// 		startKerja: 2018,
// 		cuti:       15,
// 	}
// }

func main() {
	header()
	var karyawanList tabKaryawan
	var n int
	// n = 3
	// dummy(&karyawanList)
	fmt.Println("Masukkan jumlah karyawan:")
	fmt.Scanln(&n)
	inputKaryawan(&karyawanList, n)
	printKaryawan(&karyawanList, n)

	for {
		var choice int
		fmt.Println("Menu:")
		fmt.Println("1. Cari berdasarkan nama")
		fmt.Println("2. Cari berdasarkan NIK")
		fmt.Println("3. Edit data karyawan")
		fmt.Println("4. Hapus data karyawan")
		fmt.Println("5. Tampilkan data karyawan yang terurut berdasarkan tahun mulai kerja")
		fmt.Println("6. Ajukan cuti")
		fmt.Println("7. Tampilkan saldo cuti")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih opsi (1/2/3/4/5/6/7/8): ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println()
			searchByName(&karyawanList, n)
		} else if choice == 2 {
			fmt.Println()
			searchByNIK(&karyawanList, n)
		} else if choice == 3 {
			fmt.Println()
			editKaryawan(&karyawanList, n)
		} else if choice == 4 {
			fmt.Println()
			n = hapusKaryawan(&karyawanList, n)
		} else if choice == 5 {
			fmt.Println()
			urutkanKaryawan(&karyawanList, n)
			fmt.Println("Daftar karyawan terurut berdasarkan tahun mulai kerja:")
			printKaryawan(&karyawanList, n)
		} else if choice == 6 {
			ajukanCuti(&karyawanList, n)
		} else if choice == 7 {
			tampilkanSaldoCuti(&karyawanList, n)
		} else if choice == 8 {
			fmt.Println("Terima kasih telah menggunakan program ini :)")
			return
		} else {
			fmt.Println("Menu tidak ditemukan.")
		}
	}
}

func inputKaryawan(karyawanList *tabKaryawan, n int) {
	var i int
	if n > NMAX {
		n = NMAX
	}
	fmt.Println("Daftar karyawan:")
	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data karyawan ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scanln(&karyawanList[i].nama)
		fmt.Print("Posisi: ")
		fmt.Scanln(&karyawanList[i].posisi)
		fmt.Print("Departemen: ")
		fmt.Scanln(&karyawanList[i].departemen)
		fmt.Print("NIK: ")
		fmt.Scanln(&karyawanList[i].nik)
		fmt.Print("Tahun Mulai Kerja: ")
		fmt.Scanln(&karyawanList[i].startKerja)
		fmt.Print("Saldo Cuti: ")
		fmt.Scanln(&karyawanList[i].cuti)
	}
}

func printKaryawan(karyawanList *tabKaryawan, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("Data karyawan ke-%d:\n", i+1)
		fmt.Printf("Nama: %s, Posisi: %s, Departemen: %s, NIK: %d, Tahun Mulai Kerja: %d, Total Saldo Cuti : %d, \n\n", karyawanList[i].nama, karyawanList[i].posisi, karyawanList[i].departemen, karyawanList[i].nik, karyawanList[i].startKerja, karyawanList[i].cuti)
	}
}

func searchByName(karyawanList *tabKaryawan, n int) {
	var nama string
	var i int
	var found bool
	fmt.Print("Masukkan nama karyawan yang dicari: ")
	fmt.Scan(&nama)
	found = false 
	for i = 0; i < n; i++ { //sequential searching 
		if karyawanList[i].nama == nama { //karyawan 0 : udin, karyawan 1 : asep
			fmt.Printf("Data karyawan ke-%d:\n", i+1)
			fmt.Printf("Nama: %s, Posisi: %s, Departemen: %s, NIK: %d, Tahun Mulai Kerja: %d, Total Saldo Cuti : %d\n\n", karyawanList[i].nama, karyawanList[i].posisi, karyawanList[i].departemen, karyawanList[i].nik, karyawanList[i].startKerja, karyawanList[i].cuti)
			found = true
		}
	}
	if !found {
		fmt.Println("Karyawan dengan nama tersebut tidak ditemukan.")
	}
}

func searchByNIK(karyawanList *tabKaryawan, n int) {
	var nik, i int
	var found bool
	fmt.Print("Masukkan NIK yang dicari: ")
	fmt.Scan(&nik)
	found = false
	for i = 0; i < n; i++ {
		if karyawanList[i].nik == nik {
			fmt.Printf("Data karyawan ke-%d:\n", i+1)
			fmt.Printf("Nama: %s, Posisi: %s, Departemen: %s, NIK: %d, Tahun Mulai Kerja: %d, Total Saldo Cuti : %d\n\n", karyawanList[i].nama, karyawanList[i].posisi, karyawanList[i].departemen, karyawanList[i].nik, karyawanList[i].startKerja, karyawanList[i].cuti)
			found = true
		}
	}
	if !found {
		fmt.Println("Karyawan dengan NIK tersebut tidak ditemukan.")
	}
}

func editKaryawan(karyawanList *tabKaryawan, n int) {
	var i int
	var nama string
	var found bool
	fmt.Print("Masukkan nama karyawan yang ingin diubah: ")
	fmt.Scan(&nama)
	found = false
	for i = 0; i < n; i++ {
		if karyawanList[i].nama == nama {
			fmt.Println("Masukkan data baru untuk karyawan:")
			fmt.Print("Nama: ")
			fmt.Scan(&karyawanList[i].nama)
			fmt.Print("Posisi: ")
			fmt.Scan(&karyawanList[i].posisi)
			fmt.Print("Departemen: ")
			fmt.Scan(&karyawanList[i].departemen)
			fmt.Print("Tahun Mulai Kerja: ")
			fmt.Scan(&karyawanList[i].startKerja)
			fmt.Print("Total Saldo Cuti : ")
			fmt.Scan(&karyawanList[i].cuti)
			fmt.Println("Data karyawan berhasil diubah.")
			found = true
		}
	}
	if !found {
		fmt.Println("Karyawan dengan nama tersebut tidak ditemukan.")
	}
}

func hapusKaryawan(karyawanList *tabKaryawan, n int) int {
	var nama string
	var i, j int
	var found bool
	fmt.Print("Masukkan nama karyawan yang ingin dihapus: ")
	fmt.Scan(&nama)
	found = false
	for i = 0; i < n; i++ {
		if karyawanList[i].nama == nama {
			for j = i; j < n-1; j++ {
				karyawanList[j] = karyawanList[j+1]
			}
			found = true
			n--
			i--
		}
	}
	if found {
		fmt.Println("Data karyawan berhasil dihapus.")
	} else {
		fmt.Println("Karyawan dengan nama tersebut tidak ditemukan.")
	}
	return n
}

func urutkanKaryawan(karyawanList *tabKaryawan, n int) {
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if karyawanList[j].startKerja > karyawanList[j+1].startKerja {
				karyawanList[j], karyawanList[j+1] = karyawanList[j+1], karyawanList[j]
			}
		}
	}
}

func ajukanCuti(karyawanList *tabKaryawan, n int) {
	var nama string
	var hariCuti, i int
	var found bool
	fmt.Print("Masukkan nama karyawan yang mengajukan cuti: ")
	fmt.Scan(&nama)
	found = false
	for i = 0; i < n; i++ {
		if karyawanList[i].nama == nama {
			fmt.Print("Masukkan jumlah hari cuti yang diajukan: ")
			fmt.Scan(&hariCuti)
			if hariCuti <= karyawanList[i].cuti {
				karyawanList[i].cuti = karyawanList[i].cuti - hariCuti
				fmt.Printf("Cuti sebanyak %d hari berhasil diajukan. Sisa saldo cuti: %d\n", hariCuti, karyawanList[i].cuti)
			} else {
				fmt.Println("Jumlah hari cuti melebihi saldo cuti yang tersedia.")
			}
			found = true
		}
	}
	if !found {
		fmt.Println("Karyawan dengan nama tersebut tidak ditemukan.")
	}
}

func tampilkanSaldoCuti(karyawanList *tabKaryawan, n int) {
    var nama string
    var i int
    var found bool
    fmt.Print("Masukkan nama karyawan untuk melihat saldo cuti: ")
    fmt.Scan(&nama)
    found = false
    for i = 0; i < n; i++ {
        if karyawanList[i].nama == nama {
            fmt.Printf("Saldo cuti karyawan dengan nama %s: %d hari\n", nama, karyawanList[i].cuti)
            found = true
        }
    }
    if !found {
        fmt.Println("Karyawan dengan nama tersebut tidak ditemukan.")
    }
}

