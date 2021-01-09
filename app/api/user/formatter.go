package user

type UserFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func FormatUser(user User) UserFormatter {
	return UserFormatter{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	}
}

func FormatUsers(users []User) []UserFormatter {
	var formatters []UserFormatter

	if len(users) > 0 {
		for _, user := range users {
			formatter := FormatUser(user)
			formatters = append(formatters, formatter)
		}
	} else {
		formatters = []UserFormatter{}
	}

	return formatters
}
