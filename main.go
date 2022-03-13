package main

import "fmt"
import moduleProject "github.com/Muhammad-Iqbal11111/moduleProject"

type Strukpembayaran struct {
	Nama string
}
type DataPelanggan struct {
	Nama         string
	Alamat string
	Umur         int8
	Pekerjaan string
}
func main(){
	
	moduleProject.Welcome("Muhammad Iqbal", "Y" string)

	var pilihanBaju = [] string{"Kemeja Formal", "Kemeja Casual", "Kaos Krah", "Kaos Oblong"}
	moduleProject.ItemPakaian("Muhammad Iqbal", pilihanBaju[4])
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
		x, y := moduleProject.ProfitSale(harga, hargaPokok)
		fmt.Printf("Keuntungan adalah %.2f dengan persentase keuntungan %.2f percen \n", x, y
		modulProject.Pembayaran()

		struk1 := Strukpembayaran(Nama: "Muhammad Iqbal")
		struk1.grandOpening("Berkah")
	}
	defer fmt.Println("terimakasih sudah berbelanja")

	Pelanggan1 := DataPelanggan{
		Nama:         "Muhammad Iqbal",
		Alamat : "Padang",
		Umur:         25,
		Pekerjaan:        "Staf Operasional dan Pembayaran",
	}
	fmt.Println(Pelanggan1.Nama)
	fmt.Println(Pelanggan1.Alamat)

}
func (struk Strukpembayaran) grandOpening(toko string) {
	fmt.Println("terimakasih telah berbelanja di toko ", toko, "semoga senang dengan layanan kami", struk.Nama)
