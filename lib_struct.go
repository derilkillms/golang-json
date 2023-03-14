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

type Product struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}
