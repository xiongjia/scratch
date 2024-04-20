package pedestal

type pedestalStatus int

const (
	stopped pedestalStatus = iota
	started
)

type NetworkService struct {
	IP   string `json:"ip"`
	Port uint16 `json:"port"`
}
