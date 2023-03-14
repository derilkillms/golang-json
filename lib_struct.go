package golang_json

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	Firstname, Lastname string
	Age                 int
	Hobbies             []string
	Address             []Address
}
