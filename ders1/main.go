package main

//go mod init "ders1" diyerek go.mod dosyasını ders1 package'i olarak oluşturduk
//daha sonra examples klasöründe bulunan dosyalara aşağıdaki gibi import ederek ulaşabiliriz
import "ders1/examples"

func main() {
	println("-----------------------------------------------------")
	println("Main Hello")
	println("-----------------------------------------------------")
	examples.MainHello()
	println()
	println()
	println("-----------------------------------------------------")
	println("Main Person")
	println("-----------------------------------------------------")
	examples.MainPerson()
}
