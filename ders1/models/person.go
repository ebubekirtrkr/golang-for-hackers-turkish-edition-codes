package models
import "errors"

//STRUCTS
//Burada dikkat etmemiz gereken nokta, "Person" büyük harf ile başladığından başka package'lardan erişilebilir ancak
//	Person'un içerdiği üyeler, id,firstname, lastname küçük harfle başladığından başka package'den erişilemez.
//	Bu özelliği kullanarak class gibi kullanıyoruz.
type Person struct {
	id        int
	firstname string
	lastname  string
}

//fName ve lName alıyor, Person objesi döndürüyor.
func NewPerson(fName, lName string) Person {
	return Person{
		firstname: fName, // aynı package'de olduğumuzdan
		lastname:  lName, // firstname ve lastname'e ulaşabiliyoruz.
	}

}
//NewPerson fonksiyonunda id'yi atamamıştık, onun için ayrı bi member func tanımlıyoruz.
//SetId fonksiyonu id alıyor ve error döndürüyor
//"(p *Person)" kısmı önemli çünkü Person struct'una bu fonksiyonu bind etmiş oluyoruz.
func (p *Person) SetId(id int) error {
	if id < 0 {
		return errors.New("id must be positive")
	}
	p.id = id
	return nil

}
