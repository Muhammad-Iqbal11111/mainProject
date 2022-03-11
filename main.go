package main

import "fmt"
import modul "github.com/Muhammad-Iqbal11111/moduleProject"
func main(){
	
	modul.Welcome("Muhammad Iqbal", "Y" string)

	var pilihanBaju = [] string{"Kemeja Formal", "Kemeja Casual", "Kaos Krah", "Kaos Oblong"}
	modul.ItemPakaian("Muhammad Iqbal", pilihanBaju[4])
	for _ , baju := range pilihanBaju {
		if baju ==  "Kemeja Formal"{
			harga := 200000
			hargaPokok:= 120000
			fmt.Println("Harga kemeja formal ini adalah", harga)
		} else if baju == "Kemeja Casual"{
			harga := 180000
			hargaPokok:= 100000
			fmt.Println("Harga kemeja casual ini adalah", harga)
		} else if baju == "Kaos Krah"{
			harga := 150000
			hargaPokok:=70000
			fmt.Println("Harga kaos krah ini adalah", harga)
		} else {
			harga := 50000
			hargaPokok:=20000
			fmt.Println("Harga kaos oblong ini adalah", harga)
		}
		x, y := modul.ProfitSale(harga, hargaPokok)
	fmt.Printf("Keuntungan adalah %.2f dengan persentase keuntungan %.2f percen \n", x, y
	modul.Pembayaran()
	}

	defer fmt.Println("terimakasih sudah berbelanja")
}