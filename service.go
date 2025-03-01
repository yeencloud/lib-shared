package shared

type ServiceName string

var serviceName = ""

func SetServiceName(name string) {
	serviceName = name
}

func GetServiceName() ServiceName {
	return ServiceName(serviceName)
}
