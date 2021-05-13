package discussion

type Availability string

const (
	Ready Availability = "READY"
)

func (d Availability) IsReady() bool {
	return d == Ready
}