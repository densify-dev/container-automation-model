package model

type Action string

const (
	Upsize                 Action = "upsize"
	Downsize               Action = "downsize"
	SetUninitializedValues Action = "set-uninitialized-values"
)

type ActionPolicy map[Action]bool

type Policy struct {
	Enablement map[Resource]map[Allocation]ActionPolicy `json:"enablement",yaml:"enablement"`
}

type Policies struct {
	PoliciesByName map[string]*Policy `json:"policiesByName",yaml:"policiesbyname"`
}

func (p *Policy) IsEnabled(r Resource, al Allocation, ac Action) bool {
	return p != nil && p.Enablement[r][al][ac]
}
