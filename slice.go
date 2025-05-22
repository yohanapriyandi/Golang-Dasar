package main

import "fmt"

func main() {

	names := [...]string{"Yohan", "Nabilah", "Fathar", "Fariq", "Tsabita", "Dhabit", "Taqi", "Uswah"}

	slice1 := names[4:6]
	fmt.Println(slice1)
	fmt.Println("==========================================")
	slice2 := names[2:]
	fmt.Println(slice2)
	fmt.Println("==========================================")
	slice3 := names[:5]
	fmt.Println(slice3)
	fmt.Println("==========================================")
	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println("==========================================")
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	dayslice1 := days[5:]
	fmt.Println(dayslice1)
	fmt.Println("==========================================")
	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"
	fmt.Println(dayslice1)
	fmt.Println(days)
	fmt.Println("==========================================")
	dayslice2 := append(dayslice1, "Libur Baru")
	dayslice2[0] = "Sabtu Lama"
	fmt.Println(dayslice2)
	fmt.Println(dayslice1)
	fmt.Println(days)

	fmt.Println("==================NEW SLICE========================")
	var newslice []string = make([]string, 2, 5)
	newslice[0] = "Yohan"
	newslice[1] = "Nabilah"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))
	fmt.Println("==========================================")
	//tambah data slice masih di array yang sama karena masih ada kapasitas untuk menambah data
	newslice2 := append(newslice, "Fathar")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	newslice2[0] = "Yohan Apriyandi S.Kom"
	fmt.Println(newslice)
	fmt.Println(newslice2)

	fmt.Println("==================COPY DATA SLICE========================")
	fromslice := days[:]
	toslice := make([]string, len(days), cap(days))

	copy(toslice, fromslice)
	fmt.Println(fromslice)
	fmt.Println(toslice)
}
