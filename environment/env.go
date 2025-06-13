package environment

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	Production  Type = "production"
	Staging     Type = "staging"
	Development Type = "development"
)
