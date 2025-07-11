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

type Policies struct {
	AutomationEnabled bool               `json:"automationEnabled" yaml:"automationEnabled"`
	RemoteEnablement  bool               `json:"remoteEnablement" yaml:"remoteEnablement"`
	DefaultPolicy     string             `json:"defaultPolicy" yaml:"defaultPolicy"`
	Policies          map[string]*Policy `json:"policies" yaml:"policies"`
}

func (p *Policy) IsEnabled(r Resource, al Allocation, ac Action) bool {
	return p != nil && p.Enablement[r][al][ac]
}
