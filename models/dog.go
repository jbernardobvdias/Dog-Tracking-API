package models

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  int    `json:"age"`
}

var Dogs = []Dog{
	{Name: "Doggy", Race: "German Shepherd", Age: 10},
	{Name: "Dogger", Race: "German Shepherd", Age: 5},
	{Name: "Doggo", Race: "German Shepherd", Age: 21},
}
