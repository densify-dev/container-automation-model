package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Recommendation struct {
	Cluster                      string     `json:"cluster" validate:"required"`
	Namespace                    string     `json:"namespace" validate:"required"`
	PodOwnerName                 string     `json:"podOwnerName" validate:"required"`
	PodOwnerKind                 string     `json:"podOwnerKind" validate:"required"`
	Container                    string     `json:"container" validate:"required"`
	EntityId                     string     `json:"entityId" validate:"required"`
	ContainerId                  string     `json:"containerId" validate:"required"`
	AvgContainerCount            int        `json:"avgContainerCount"`
	CurrentCpuRequestMCores      int64      `json:"currentCpuRequestmCores"`
	CurrentCpuLimitMCores        int64      `json:"currentCpuLimitmCores"`
	CurrentMemRequestBytes       int64      `json:"currentMemRequestBytes"`
	CurrentMemLimitBytes         int64      `json:"currentMemLimitBytes"`
	RecommendedCpuRequestMCores  int64      `json:"recommendedCpuRequestmCores"`
	RecommendedCpuLimitMCores    int64      `json:"recommendedCpuLimitmCores"`
	RecommendedMemRequestBytes   int64      `json:"recommendedMemRequestBytes"`
	RecommendedMemLimitBytes     int64      `json:"recommendedMemLimitBytes"`
	EstimatedSavingsPerContainer float64    `json:"estimatedSavingsPerContainer"`
	AnalyzedOn                   time.Time  `json:"analyzedOn" validate:"required"`
	HpaMetricName                string     `json:"hpaMetricName"`
	KubexAutomation              StringBool `json:"kubexAutomation"`
}

func (r *Recommendation) Validate() error {
	return validate.Struct(r)
}

func FilterValidRecommendations(recs []*Recommendation) ([]*Recommendation, error) {
	validRecs := make([]*Recommendation, 0, len(recs))
	var err error
	for _, rec := range recs {
		if rec == nil {
			continue
		}
		if err1 := rec.Validate(); err1 == nil {
			validRecs = append(validRecs, rec)
		} else {
			if err == nil {
				err = err1
			} else {
				var verrs validator.ValidationErrors
				if errors.As(err, &verrs) {
					var verrs1 validator.ValidationErrors
					if errors.As(err1, &verrs1) {
						verrs = append(verrs, verrs1...)
						err = verrs
					}
				}
			}
		}
	}
	return validRecs, err
}

type Recommendations struct {
	Checksum string            `json:"checksum"`
	Results  []*Recommendation `json:"results"`
}

const (
	StatusNotModified = http.StatusNotModified
)
