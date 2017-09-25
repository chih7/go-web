package data

var users = []User{
	{
		Name: "chih",
		Email: "i@chih.me",
		Password: "test123",
	},
	{
		Name: "chih2",
		Email: "i2@chih.me",
		Password: "test1234",
	},
}

func setup()  {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}
