package model

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// SetAnalyzedOnTime is a utility method to set the AnalyzedOn field from a time.Time.
func (x *Container) SetAnalyzedOnTime(t *time.Time) {
	if x == nil {
		return
	}
	if t == nil {
		x.AnalyzedOn = nil
	} else {
		x.AnalyzedOn = timestamppb.New(*t)
	}
}

// GetAnalyzedOnTime is a utility method to get the AnalyzedOn field as a time.Time.
// It returns nil if AnalyzedOn is not set or invalid.
func (x *Container) GetAnalyzedOnTime() *time.Time {
	if x == nil || !x.AnalyzedOn.IsValid() {
		return nil
	}
	t := x.AnalyzedOn.AsTime()
	return &t
}