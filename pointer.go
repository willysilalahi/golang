package main
import "fmt"

type Address struct{
	Kecamatan string
	Kabupaten string
	Provinsi string
}

func main()  {
	address1 := Address{"Bintang Bayu", "Serdang Bedagai", "Sumatera Utara"}
	address2 := &address1
	address3 := &address1
	//yang sebenarnya
	var address_A Address = Address{"Bintang Bayu", "Serdang Bedagai", "Sumatera Utara"}
	var address_B *Address = &address_A

	address2.Kecamatan = "Dolok Masihul"
	//tanpa *
	//address2 = &Address{"Raya Kahean", "Simalungun", "Sumatera Utara"} // akan membentuk memori baru
	
	//dengan *
	*address2 = Address{"Raya Kahean", "Simalungun", "Sumatera Utara"} // akan membentuk memori baru

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address_A)
	fmt.Println(address_B)

	var alamat *Address = new(Address)
	fmt.Println(alamat)
}