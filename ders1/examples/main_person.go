package examples

import (
	"fmt"

	"ders1/models" // go mod init "test"
)

func MainPerson() {
	// yeni bir Person objesi oluşturup, idsine 100 atadık
	p1 := models.NewPerson("Mehmet", "İnce")
	err := p1.SetId(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)

	// yeni bir Person objesi oluşturup, idsine -1 atamaya çalışıyoruz, hata verecek.
	p2 := models.NewPerson("Ege", "Balcı")
	err = p2.SetId(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
