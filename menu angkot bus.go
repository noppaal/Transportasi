package main

import "fmt"

const NMAX int = 100

type transportasi struct {
	Angkutan                            string
	NomorRute                           int
	NamaPerusahaan, AwalRute, AkhirRute string
	PanjangRute                         float64
}

type kendaraan [NMAX]transportasi

func main() {
	var T kendaraan
	var terminal, terminalAwal, terminalAkhir string
	var input, nData, noAngkot, noBus, x int

	fmt.Println("===| Untuk Memulai Silahkan Ketika Angka 1 |===")
	fmt.Scan(&input)
	for input != 1 {
		fmt.Println("===| Untuk Memulai Silahkan Ketika Angka 1 |===")
		fmt.Scan(&input)
	}
	for input != 0 {
		fmt.Println("=========================M E N U==========================")
		fmt.Println("| 1. Input Data                                          |")
		fmt.Println("| 2. Menampilkan Semua Rute Transportasi                 |")
		fmt.Println("| 3. Mencari Rute Bus Dari Terminal Tertentu             |")
		fmt.Println("| 4. Mencari Rute Transportasi Dari 2 Terminal Tertentu  |")
		fmt.Println("| 5. Mencari Angkot Dengan Rute Terpendek                |")
		fmt.Println("| 6. Mencari Angkot Dengan Nomor Tertentu                |")
		fmt.Println("| 7. Mencari Bus Dengan Nomor Tertentu                   |")
		fmt.Println("| 0. Keluar ...                                          |")
		fmt.Println("=========================M E N U==========================")
		fmt.Println("| Silahkan Masukan Pilihan ...")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Println("===| Silahkan Masukan Jumlah Data yang Ingin Dimasukkan |===")
			fmt.Scan(&nData)
			for nData < 0 {
				fmt.Println("===| Jumlah Data Tidak Valid, Silahkan Input Ulang! |===")
				fmt.Scan(&nData)
			}
			fmt.Println("======| Silahkan Isi Data |======")
			IsiData(&T, &nData)
			fmt.Println("========| Data Berhasil Ditambahkan! |========")
			fmt.Println()
		} else if input == 2 {
			fmt.Println("===| List Semua Rute Transportasi |===")
			cetakRute(T, nData)
		} else if input == 3 {
			fmt.Println("===| Silahkan Masukan Terimal Awal Bus yang Ingin Dicari! |===")
			fmt.Scan(&terminal)
			CetakRuteBus(T, nData, terminal)
		} else if input == 4 {
			fmt.Println("===| Silahkan Masukan Terminal Awal dan Akhir yang Ingin Dicari! |===")
			fmt.Scan(&terminalAwal, &terminalAkhir)
			CetakAngkutanAntarTerminal(T, nData, terminalAwal, terminalAkhir)
		} else if input == 5 {
			/*fmt.Println("===| Silahkan Masukan Terminal Awal dan Akhir yang Ingin Dicari Rute Terpendeknya |===")
			fmt.Scan(&terminalAwal, &terminalAkhir)*/
			fmt.Println("========| Angkot Dengan Rute Terpendek! |========")
			AngkotRuteTerpendek(T, nData)
		} else if input == 6 {
			fmt.Println("===| Silahkan Masukan Nomor Angkot yang Ingin Dicari |===")
			fmt.Scan(&noAngkot)
			x = cetaknoangkot(T, nData, noAngkot)
			if x != 0 {
				fmt.Printf("Angkot dengan nomor %v memiliki rute %v-%v\n", T[x].NomorRute, T[x].AwalRute, T[x].AkhirRute)
			} else {
				fmt.Print("Tidak ada angkot dengan nomor tersebut.\n")
			}
		} else if input == 7 {
			fmt.Println("===| Silahkan Masukan Nomor Bus yang Ingin Dicari |===")
			fmt.Scan(&noBus)
			cetaknobus(T, nData, noBus)
		} else if input == 0 {
			fmt.Println("======| Terima Kasih! |======")
		} else {
			fmt.Println("===| Input Invalid Silahkan Ulangi Pilihan! |===")
			fmt.Println()
		}
	}
}

func IsiData(T *kendaraan, n *int) {
	/*
		{I.S. Terdefini array kosong T}
		{F.S. Array T terisi data}
	*/

	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&T[i].Angkutan, &T[i].NomorRute, &T[i].NamaPerusahaan, &T[i].AwalRute, &T[i].AkhirRute, &T[i].PanjangRute)
	}
}

func cetakRute(T kendaraan, n int) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data}
		{F.S. Menampilkan semua rute yang tersedia}
	*/

	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%v %v (%v - %v) %v km\n", T[i].Angkutan, T[i].NomorRute, T[i].AwalRute, T[i].AkhirRute, T[i].PanjangRute)
		}
	}
	fmt.Println()
}

//sequential search
func CetakRuteBus(T kendaraan, n int, x string) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data}
		{F.S. Menampilkan semua rute bus yang tersedia dari terminal tertentu dengan algorima sequential search}
	*/
	var i, found int
	found = 0
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		for i = 0; i < n; i++ {
			if T[i].Angkutan == "bus" && T[i].AwalRute == x {
				fmt.Printf("%v %v (%v - %v) %v km\n", T[i].Angkutan, T[i].NomorRute, T[i].AwalRute, T[i].AkhirRute, T[i].PanjangRute)
				found++
			}
		}
		if found == 0 {
			fmt.Println("Tidak ada bus dengan terminal tersebut.")
		}

		fmt.Println()
	}
}

func CetakAngkutanAntarTerminal(T kendaraan, n int, a, b string) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data, a terminal awal, b terminal akhir}
		{F.S. Menampilkan semua rute kendaraan yang tersedia antar terminal tertentu dengan algorima sequential search}
	*/
	var i, found int
	found = 0
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		for i = 0; i < n; i++ {
			if T[i].AwalRute == a && T[i].AkhirRute == b {
				fmt.Printf("%v %v (%v - %v) %v km\n", T[i].Angkutan, T[i].NomorRute, T[i].AwalRute, T[i].AkhirRute, T[i].PanjangRute)
				found++
			}
		}
		if found == 0 {
			fmt.Println("Tidak ada angkutan dengan rute tersebut.")
		}
		fmt.Println()
	}
}

func AngkotRuteTerpendek(T kendaraan, n int) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data}
		{F.S. Menampilkan angkot dengan rute terpendek dengan memanggil subprogram insertion sorting}
	*/
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		urutPanjangRute(&T, n)
		fmt.Printf("%v %v (%v - %v) %v km\n", T[0].Angkutan, T[0].NomorRute, T[0].AwalRute, T[0].AkhirRute, T[0].PanjangRute)
		fmt.Println()
	}
}

// Sequential Searching
func cetaknoangkot(T kendaraan, n, no int) int {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data dan nomor angkot yang ingin dicari}
		{F.S. Mengembalikan indeks array jika nomor angkot ditemukan dan mengembalikan nilai 0 jika nomor angkot tidak ditemukan}
	*/
	var i, found int
	found = 0
	for i = 0; i < n; i++ {
		if T[i].Angkutan == "angkot" && T[i].NomorRute == no {
			found = i
		}
	}
	return found
}

// Insertion Sort Ascending
func urutPanjangRute(T *kendaraan, n int) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data}
		{F.S. Array T terurut secara ascending (membesar) terhadap panjang rute}
	*/
	var i, j int
	var temp transportasi
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		i = 1
		for i <= n-1 {
			j = i
			temp = T[i]
			for j > 0 && temp.PanjangRute < T[j-1].PanjangRute {
				T[j] = T[j-1]
				j--
			}
			T[j] = temp
			i++
		}
	}
}

//Selection Sort Descending
func urutNomorRute(T *kendaraan, n int) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data}
		{F.S. Array T terurut secara descending (mengecil) terhadap nomor rute}
	*/
	var idx int
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		for i := 1; i < n; i++ {
			idx = i - 1
			for j := i; j < n; j++ {
				if T[idx].NomorRute < T[j].NomorRute {
					idx = j
				}
			}
			T[i-1], T[idx] = T[idx], T[i-1]
		}
	}
}

// Binary Search
func cetaknobus(T kendaraan, n int, no int) {
	/*
		{I.S. Terdefini array T yang berisi n jumlah data dan nomor bus yang ingin dicari}
		{F.S. Menampilkan nomor bus yang dicari JIKA ADA!}
	*/
	var left, right, mid, found int
	if n == 0 {
		fmt.Println("===| Data Kosong, Anda Belum Menginput Array |===")
	} else {
		urutNomorRute(&T, n)
		left = 0
		right = n - 1
		found = -1

		for left <= right && found == -1 {
			mid = (left + right) / 2
			if T[mid].Angkutan == "bus" {
				if T[mid].NomorRute == no {
					found = mid
				} else if no < T[mid].NomorRute {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				if T[mid].NomorRute < no {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}

		if found != -1 {
			fmt.Printf("Bus dengan nomor %v memiliki rute %v-%v\n", T[found].NomorRute, T[found].AwalRute, T[found].AkhirRute)
		} else {
			fmt.Print("Tidak ada bus dengan nomor tersebut.\n")
		}
		fmt.Println()
	}
}

// untuk masukan data
/*
27
angkot 01 PT.MetroPermai BuahBatu Pasteur 12.2
angkot 23 PT.Merdeka Kopo Setiabudi 11.3
angkot 48 PT.SumberAlam Bojongsoang Antapani 10.3
angkot 10 PT.SetiaNegara Cikutra Cibiru 11.5
angkot 43 PT.Primadona BuahBatu Arcamanik 6.8
angkot 30 PT.SetiaNegara Cibiru Pasteur 16.5
angkot 98 PT.MetroPermai Dago Antapani	8.1
angkot 56 PT.SetiaNegara Bojongsoang Setiabudi	31.2
angkot 69 PT.Merdeka BuahBatu PasirKaliki 6.7
angkot 05 PT.MetroPermai Antapani Dago 9.0
angkot 10 PT.SumberAlam Bojongsoang Jatinangor 21.3
angkot 67 PT.Primadona Kopo Cicadas 8.7
angkot 88 PT.RidwanGuacor BuahBatu Pasteur 11.8
bus 99 PT.CahayaPurwanto BuahBatu Pasteur 12.0
bus 07 PT.BandungExpress BuahBatu Jatinangor 20.5
bus 11 PT.BerlianJaya Antapani Pasteur 9.3
bus 77 PT.Nusantara KotaBaruParahyangan	Cibiru 41.7
bus 31 PT.PahalaKencana Bojongsoang KotaBaruParahyangan 30.6
bus 67 PT.CahayaSolo BuahBatu Supratman 7.2
bus 73 PT.BintangUtaraPutra Ciumbuleuit KotaBaruParahyangan 28.8
bus 32 PT.BorneoTransMandiri Kopo Jatinangor 28.1
bus 86 PT.CititransBusline SuryaSumantri Cibiru 20.7
bus 60 PT.DoaIbu Jatinangor Setiabudi 44.9
bus 33 PT.BorneoTransMandiri Bojongsoang Cibiru 15.1
bus 14 PT.PahalaKencana BuahBatu Lembang 19.4
bus 90 PT.BerlianJaya Cibiru BuahBatu 13.3
bus 55 PT.BandungExpress Kopo Cileunyi 27.5
*/
