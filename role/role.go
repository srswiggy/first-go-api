package role

type Role int

const (
	Admin Role = iota
	User
	SpecialUser
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case User:
		return "User"
	case SpecialUser:
		return "Special User"
	default:
		return "Unknown"
	}
}

func GetRole(name string) Role {
	switch name {
	case "admin":
		return Admin
	case "user":
		return User
	case "special":
		return SpecialUser
	default:
		return User
	}
}
