package main

import "fmt"
import moduleProject "github.com/Muhammad-Iqbal11111/moduleProject"

type Strukpembayaran struct {
	Name string
}
type DataPelanggan struct {
	Nama         string
	Alamat string
	Umur         int8
	Pekerjaan string
}

func main(){
	moduleProject.AlamatToko()
	moduleProject.Welcome("Muhammad Iqbal")
	var pilihanBaju = []string{"Kemeja Formal", "Kemeja Casual", "Kaos Krah", "Kaos Oblong"}
	moduleProject.ItemPakaian("Muhammad Iqbal", pilihanBaju...)
	fmt.Println("Please choose the clothes you want to buy!")
	var pilihan string
	fmt.Scanln(&pilihan)
	if pilihan == "Kemeja Formal"{
		harga := 200000
		hargaPokok:=150000
		fmt.Println("The Price of this shirt is", harga)
		x, y := moduleProject.ProfitSale(harga, hargaPokok)
		fmt.Printf("The profit for the shop is %d\n", x)
		fmt.Printf("Discount for customers is %d\n", y)
	} else if pilihan == "Kemeja Casual"{
		harga := 180000
		hargaPokok:=130000
		fmt.Println("The Price of this shirt is", harga)
		x, y := moduleProject.ProfitSale(harga, hargaPokok)
		fmt.Printf("The profit for the shop is %d\n", x)
		fmt.Printf("Discount for customers is %d\n", y)
	} else if pilihan == "Kaos Krah"{
		harga := 120000
		hargaPokok:=70000
		fmt.Println("The Price of this shirt is", harga)
		x, y := moduleProject.ProfitSale(harga, hargaPokok)
		fmt.Printf("The profit for the shop is %d\n", x)
		fmt.Printf("Discount for customers is %d\n", y)
	} else {
		harga := 70000
		hargaPokok:=40000
		fmt.Println("The Price of this shirt is", harga)
		x, y := moduleProject.ProfitSale(harga, hargaPokok)
		fmt.Printf("The profit for the shop is %d\n", x)
		fmt.Printf("Discount for customers is %d\n", y)	}
	
		struk1 := Strukpembayaran{Name: "Muhammad Iqbal"}
		struk1.GrandOpening("Berkah")
		moduleProject.CustomerData()
		moduleProject.selesai()

}
