package main

import "fmt"

type sepeda struct {
	jumlahBan   int
	jumlahGigi  int
	warnaSepeda string
}

func (durasi *sepeda) waktuTempuh(jarak float32) float32 {
	return jarak * 2.5
}

func main() {
	var arraySepeda [5]sepeda
	var inputJarak float32

	// for i := range arraySepeda {
	// 	fmt.Println(i)
	// 	fmt.Printf("Masukkan Jarak yang ditempuh untuk sepeda ke %d: ", i+1)
	// 	fmt.Scanln(&inputJarak)
	// }

	arraySepeda[0] = sepeda{jumlahBan: 2, jumlahGigi: 1, warnaSepeda: "merah"}
	arraySepeda[1] = sepeda{jumlahBan: 2, jumlahGigi: 6, warnaSepeda: "biru"}
	arraySepeda[2] = sepeda{jumlahBan: 2, jumlahGigi: 11, warnaSepeda: "hijau"}
	arraySepeda[3] = sepeda{jumlahBan: 2, jumlahGigi: 7, warnaSepeda: "kuning"}
	arraySepeda[4] = sepeda{jumlahBan: 2, jumlahGigi: 10, warnaSepeda: "abu"}

	for i, basikal := range arraySepeda {
		fmt.Printf("Masukkan jarak tempuh sepeda ke %d : ", i+1)
		fmt.Scanln(&inputJarak)
		fmt.Printf("Sepeda ke %d dengan jumlah gigi %d dapat menempuh jarak %.f km dalam waktu %.f menit \n", i+1, basikal.jumlahGigi, inputJarak, basikal.waktuTempuh(inputJarak))
	}
}
