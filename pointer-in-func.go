package main 
import "fmt"

type Address struct{
	Kecamatan string
	Kabupaten string
	Provinsi string
}

func changeAddress(address Address){
	address.Kecamatan = "Medan Timur"
}

func UbahAddress(address *Address){
	address.Kecamatan = "Medan Timur"
}


func  main()  {
	address := Address{"Medan Maimun", "Medan", "Sumatera Utara"}
	changeAddress(address)
	fmt.Println(address) // ini takkan merubah isi dari address
	
	
	alamat := Address{"Medan Maimun", "Medan", "Sumatera Utara"}
	UbahAddress(&alamat)
	fmt.Println(alamat) // ini merubah isi dari address
}