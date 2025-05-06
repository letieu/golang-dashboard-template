package user

type User struct {
	ID    int
	Name  string
	Email string
}

func GetMockUsers() []User {
	return []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}
}
