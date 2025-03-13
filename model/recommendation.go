package model

import (
	"net/http"
	"time"
)

type Recommendation struct {
	Cluster                      string    `json:"cluster"`
	Namespace                    string    `json:"namespace"`
	PodOwnerName                 string    `json:"podOwnerName"`
	PodOwnerKind                 string    `json:"podOwnerKind"`
	Container                    string    `json:"container"`
	EntityId                     string    `json:"entityId"`
	ContainerId                  string    `json:"containerId"`
	AvgContainerCount            int       `json:"avgContainerCount"`
	CurrentCpuRequestMCores      int64     `json:"currentCpuRequestmCores"`
	CurrentCpuLimitMCores        int64     `json:"currentCpuLimitmCores"`
	CurrentMemRequestBytes       int64     `json:"currentMemRequestBytes"`
	CurrentMemLimitBytes         int64     `json:"currentMemLimitBytes"`
	RecommendedCpuRequestMCores  int64     `json:"recommendedCpuRequestmCores"`
	RecommendedCpuLimitMCores    int64     `json:"recommendedCpuLimitmCores"`
	RecommendedMemRequestBytes   int64     `json:"recommendedMemRequestBytes"`
	RecommendedMemLimitBytes     int64     `json:"recommendedMemLimitBytes"`
	EstimatedSavingsPerContainer float64   `json:"estimatedSavingsPerContainer"`
	AnalyzedOn                   time.Time `json:"analyzedOn"`
	AutomationPolicy             string    `json:"automationPolicy"`
	HpaMetricThreshold           float64   `json:"hpaMetricThreshold"`
}
type Recommendations struct {
	Checksum string            `json:"checksum"`
	Results  []*Recommendation `json:"results"`
}

const (
	StatusNotModified = http.StatusNotModified
)
