package model

import (
	"fmt"
	"strings"
)

type Action string

const (
	Upsize             Action = "upsize"
	Downsize           Action = "downsize"
	SetFromUnspecified Action = "set-from-unspecified"
	UnsetFromSpecified Action = "unset-from-specified"
)

type ActionPolicy map[Action]bool

type Policy struct {
	Enablement map[Resource]map[Allocation]ActionPolicy `json:"enablement",yaml:"enablement"`
}

type Policies struct {
	AutomationEnabled bool               `json:"automationEnabled",yaml:"automationenabled"`
	RemoteEnablement  bool               `json:"remoteEnablement",yaml:"remoteenablement"`
	DefaultPolicy     string             `json:"defaultPolicy",yaml:"defaultpolicy"`
	PoliciesByName    map[string]*Policy `json:"policiesByName",yaml:"policiesbyname"`
}

func (p *Policy) IsEnabled(r Resource, al Allocation, ac Action) bool {
	return p != nil && p.Enablement[r][al][ac]
}

func (p *Policy) String() string {
	if p == nil {
		return NilString
	}
	var s []string
	for r, als := range p.Enablement {
		for al, acs := range als {
			var ks = []any{r, al}
			var vs []any
			for ac, enabled := range acs {
				if enabled {
					vs = append(vs, ac)
				}
			}
			if len(vs) > 0 {
				s = append(s, fmt.Sprintf("%s%s%s", fmt.Sprintf("%v", ks), Spaces(Colon), fmt.Sprintf("%v", vs)))
			}
		}
	}
	return strings.Join(s, Spaces(Or))
}
