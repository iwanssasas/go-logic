package entity

type (
	Student struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Class   string `json:"class"`
		Contact int    `json:"contact"`
		Email   int    `json:"email"`
	}

	Students []Student
)
