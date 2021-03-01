package main

import "fmt"

// type vehicles interface{}

// type vehicle struct {
// 	Seats    int
// 	MaxSpeed int
// 	Color    string
// }

// type car struct {
// 	vehicle
// 	Wheels int
// 	Doors  int
// }

// type plane struct {
// 	vehicle
// 	Jet bool
// }

// type boat struct {
// 	vehicle
// 	Length int
// }

// func main(){
// 	prius := car{}
// 	tacoma := car{}
// 	bmw258 := car{}
// 	cars := []car{prius,tacoma,bmw258}

// 	boeing747 := plane{}
// 	boeing757 := plane{}
// 	boeing767 := plane{}

// 	planes := []plane{boeing747,boeing757,boeing767}

// 	// sanger := boat{}
// 	// nautique := boat{}

// 	for key,value := range cars {
// 		fmt.Println(key," - ",value)
// 	}

// }

// func main() {
// 	prius := car{}
// 	tacoma := car{}
// 	bmw258 := car{}
// 	// cars := []car{prius,tacoma,bmw258}

// 	boeing747 := plane{}
// 	boeing757 := plane{}
// 	boeing767 := plane{}

// 	// planes := []plane{boeing747,boeing757,boeing767}
// 	rides := []vehicles{prius, tacoma, bmw258, boeing747, boeing757, boeing767}

// 	// sanger := boat{}
// 	// nautique := boat{}

// 	for key, value := range rides {
// 		fmt.Println(key, " - ", value)
// 	}

// }

// params accpets any types

func spaceres(a interface{}) {
	fmt.Println(a)
}

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"miao"}, true}

	// spaceres(fido)
	// spaceres(fifi)
	shadow := dog{animal{"woof"}, true}
	// interfae{}  like java Object
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)

}
