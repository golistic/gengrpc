package gengrpc

type sourcer interface {
	Contracts() ([]string, error)
	ContractPath() string
}
