package config

type SourceConfig interface {
	ReadString(key string) (string, error)
}
