package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const (
	SspService = ServiceName("sspService")
)
