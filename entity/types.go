package entity

type Person struct {
	Id     uint64 `ini:"id"     json:"id"     db:"id"`
	Name   string `ini:"name"   json:"name"   db:"name"`
	Age    int8   `ini:"age"    json:"age"    db:"age"`
	Gender string `ini:"gender" json:"gender" db:"gender"`
}

func (person *Person) NewPerson(id uint64, name string, age int8, gender string) {

}
