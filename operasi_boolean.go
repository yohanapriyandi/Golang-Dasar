package main

import "fmt"

func main() {
	var nilaiAkhir = 95
	var absensi = 80

	var lulusNiaiAKhir bool = nilaiAkhir > 90
	var lulusAbsensi bool = absensi >= 80

	var result bool = lulusNiaiAKhir && lulusAbsensi

	fmt.Println(result)
}
