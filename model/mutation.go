package model

type Mutation struct {
	Original int64 `json:"original"`
	New      int64 `json:"new"`
}

type ContainerMutation struct {
	Time        RFC3339Time                           `json:"time"`
	EntityId    string                                `json:"entityId"`
	ContainerId string                                `json:"containerId"`
	PolicyName  string                                `json:"policyName"`
	Operation   string                                `json:"operation,omitempty"`
	Source      string                                `json:"source,omitempty"`
	Mutations   map[Resource]map[Allocation]*Mutation `json:"mutations"`
}

type ContainerMutations []ContainerMutation
