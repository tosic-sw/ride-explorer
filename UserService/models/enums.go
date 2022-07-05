package models

type Role string

const (
	ADMIN     Role = "ADMIN"
	DRIVER    Role = "DRIVER"
	PASSENGER Role = "PASSENGER"
)

func (s Role) String() string {
	switch s {
	case ADMIN:
		return "ADMIN"
	case DRIVER:
		return "DRIVER"
	case PASSENGER:
		return "PASSENGER"
	}
	return "unknown"
}
