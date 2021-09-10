package examples

import "fmt"

//küçük harf ile başlayanlar sadece bu package tarafından kullanılır
//büyük harf ile başlayanlar dışarıdan da çağırılabilir
//örneğin
//func read() {} // -> private
//func Read() {} // -> public
//go build hello.go -> build eder ve çalıştırılabilir dosya oluşturur
//go run hello.go -> tmp de build edip runlar

func MainHello() {

	//fmt package'ndaki Println fonksiyonu ile çıktı alabiliriz.
	fmt.Println("Hello World.")

	//int tipinde i'yi tanımlayıp, 42 atadık
	var i int = 42
	println("i =", i)
	//Printf -> C deki gibi extra olarak %T tipini veriyor
	fmt.Printf("i tipi = %T\n\n", i)

	var f float32 = 3.14
	println("f =", f)
	fmt.Printf("f tipi = %T\n\n", f)

	//walrus operatorü :=  --> tip belirtmiyoruz , sağdaki tipi otomatik alıyor
	name := "MDI" // var name string gibi oldu
	println("name =", name)
	fmt.Printf("name tipi = %T\n\n", name)

	g := 3.14 //yine walrus kullandık, float64 alacaktır
	println("g =", g)
	fmt.Printf("g tipi = %T\n\n", g)

	//Pointerlar
	//lastnamePtr adında bir string pointer'ı tanımladık ve bu pointere new ile memoryden alan ayırdık
	var lastnamePtr *string = new(string)
	//dereference edip içine Demir yazdık
	*lastnamePtr = "Demir"
	fmt.Println("lastnamePtr =", lastnamePtr) //memoryde ayırdığımız adresi yazdıracak
	fmt.Printf("lastnamePtr tipi = %T\n", lastnamePtr)
	fmt.Println("*lastnamePtr =", *lastnamePtr) // Demir yazdıracak
	fmt.Printf("*lastnamePtr tipi = %T\n\n", *lastnamePtr)

	//istersek varolan stringin addresini de verebiliriz
	var fNamePtr *string = &name
	fmt.Println("fNamePtr =", fNamePtr)
	fmt.Printf("fNamePtr tipi = %T\n", fNamePtr)
	fmt.Println("*fNamePtr =", *fNamePtr)
	fmt.Printf("*fNamePtr tipi = %T\n\n", *fNamePtr)

	//consts
	const c int = 4
	//c=4
	fmt.Println("c =", c)
	fmt.Printf("c tipi = %T\n\n", c)

	//enum benzeri yapı
	//iota -> kendini ve sonrakileri auto inc yapan keyword gibi düşünebiliriz
	const (
		first  = iota + 1  //iota 0 dan başlatıyor, 1 ekleyerek 1 den başlattık
		second             // kendinden önceki 1 iota olduğundan 2 olacaktır
		third  = 2 << iota // kendini ve sonrakileri hep 2 ile shift left edecektir
		fourth             // shl iotadan dolayı 4 olacak
	)
	println("first =", first,
		"| second =", second,
		"| third =", third,
		"| fourth =", fourth,
	)
	fmt.Printf("first tipi = %T\n\n", first)

	//sabit uzunlukta arrayler
	var arr [3]int
	arr[0] = 1
	//arr[3]  = 1 // hata verir çünkü 3. index yok
	fmt.Println("arr =", arr)
	fmt.Printf("arr tipi = %T\n\n", arr)

	//walrus ile array tanımlayıp süslü parantezler ile ilklendirebiliriz
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println("arr2 =", arr2)
	fmt.Printf("arr2 tipi = %T\n\n", arr2)

	//walrus ile boş array tanımlamak (int default 0 )
	arr3 := [3]int{}
	arr3[0] = 1
	fmt.Println("arr3 =", arr3)
	fmt.Printf("arr3 tipi = %T\n\n", arr3)

	//slicelar
	//sliceları başka arraylerden oluşturabiliriz
	slice := arr2[:] // : -> pythondaki indexleme operatoru gibi
	fmt.Println("slice =", slice)
	fmt.Printf("slice tipi = %T\n\n", slice)

	//ya da walrus ile initilaize
	slice2 := []int{1, 2, 3}
	fmt.Println("slice2 =", slice2)
	fmt.Printf("slice2 tipi = %T\n\n", slice2)

	//maps
	//python'daki dict gibi
	m := map[string]int{"orhun": 42}
	fmt.Println("m =", m)
	fmt.Printf("m tipi = %T\n", m)
	fmt.Println("m[\"orhun\"] =", m["orhun"])
	fmt.Printf("m[\"orhun\"] tipi = %T\n\n", m["orhun"])

	//structs
	//C deki struct gibi
	type person struct {
		Id   int
		Name string
	}
	//walrus ile person struct'u oluşturma
	p := person{
		Id:   1,
		Name: "ali", //buradaki virgül önemli eğer, } i bi satır aşağıda yapacaksak virgül olmalı
		//ya da virgülsüsz Name: "ali"} de yapabilirdik
	}
	fmt.Println("p =", p)
	fmt.Printf("p tipi = %T\n", p)
	fmt.Println("p.Id =", p.Id)
	fmt.Printf("p.Id tipi = %T\n\n", p.Id)

}
