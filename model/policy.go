package model

type Action string

const (
	Upsize             Action = "upsize"
	Downsize           Action = "downsize"
	SetFromUnspecified Action = "setFromUnspecified"
	UnsetFromSpecified Action = "unsetFromSpecified"
)

type ActionPolicy map[Action]bool

type Strategy struct {
	Enabled  bool   `json:"enabled" yaml:"enabled"`
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`
}

type SafetyChecks struct {
	CooldownAfterSpecChangeDays int `json:"cooldownAfterSpecChangeDays" yaml:"cooldownAfterSpecChangeDays"`
	MaxAnalysisAgeDays          int `json:"maxAnalysisAgeDays" yaml:"maxAnalysisAgeDays"`
	PodPendingRollbackPeriod    int `json:"podPendingRollbackPeriod" yaml:"podPendingRollbackPeriod"`
}

type ResourceQuotaChecks struct {
	MinQuotaHeadroomPercentForUpsize int `json:"minQuotaHeadroomPercentForUpsize" yaml:"minQuotaHeadroomPercentForUpsize"`
}

type Policy struct {
	AllowedPodOwners              string                                   `json:"allowedPodOwners,omitempty" yaml:"allowedPodOwners,omitempty"`
	Enablement                    map[Resource]map[Allocation]ActionPolicy `json:"enablement" yaml:"enablement"`
	InPlaceResize                 Strategy                                 `json:"inPlaceResize" yaml:"inPlaceResize"`
	InPlaceResizeContainerRestart Strategy                                 `json:"inPlaceResizeContainerRestart" yaml:"inPlaceResizeContainerRestart"`
	PodEviction                   Strategy                                 `json:"podEviction" yaml:"podEviction"`
	SafetyChecks                  SafetyChecks                             `json:"safetyChecks" yaml:"safetyChecks"`
	ResourceQuotaChecks           ResourceQuotaChecks                      `json:"resourceQuotaChecks" yaml:"resourceQuotaChecks"`
}

func (p *Policy) IsEnabled(r Resource, al Allocation, ac Action) bool {
	return p != nil && p.Enablement[r][al][ac]
}

type Policies struct {
	AutomationEnabled bool               `json:"automationEnabled" yaml:"automationEnabled"`
	RemoteEnablement  bool               `json:"remoteEnablement" yaml:"remoteEnablement"`
	DefaultPolicy     string             `json:"defaultPolicy" yaml:"defaultPolicy"`
	Policies          map[string]*Policy `json:"policies" yaml:"policies"`
}

type Operator string

const (
	OperatorIn    Operator = "In"
	OperatorNotIn Operator = "NotIn"
)

type Selector struct {
	Operator Operator `json:"operator" yaml:"operator"`
	Values   []string `json:"values" yaml:"values"`
}

type KeySelector struct {
	Selector `json:",inline" yaml:",inline"`
	Key      string `json:"key" yaml:"key"`
}

type Scope struct {
	PolicyName string      `json:"policyName" yaml:"policyName"`
	Namespaces Selector    `json:"namespaces" yaml:"namespaces"`
	PodLabels  KeySelector `json:"podLabels" yaml:"podLabels"`
}

type Scopes map[string]Scope
