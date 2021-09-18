package examples

import "fmt"

type Insan struct {
	boy int
}

func (i *Insan) Boy() int {
	return i.boy
}

type Hayvan struct {
	boy int
}

func (h *Hayvan) Boy() int {
	return h.boy
}

type Canli interface {
	Boy() int
}
func banaboyunugetir(c Canli) int{
	return c.Boy()
}

func Insan_main() {
	h := Hayvan{boy: 3}
	i := Insan{boy: 4}
	fmt.Println(banaboyunugetir(&i))
	fmt.Println(banaboyunugetir(&h))
}
