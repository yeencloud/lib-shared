package config

type Secret struct {
	Value string
}

func (s Secret) String() string {
	return "***"
}

func NewSecret(value string) Secret {
	return Secret{Value: value}
}
