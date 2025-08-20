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

// SetSuspendedOnTime is a utility method to set the SuspendedOn field from a time.Time.
func (x *Container) SetSuspendedOnTime(t *time.Time) {
	if x == nil {
		return
	}
	if t == nil {
		x.SuspendedOn = nil
	} else {
		x.SuspendedOn = timestamppb.New(*t)
	}
}

// GetSuspendedOnTime is a utility method to get the SuspendedOn field as a time.Time.
// It returns nil if SuspendedOn is not set or invalid.
func (x *Container) GetSuspendedOnTime() *time.Time {
	if x == nil || !x.SuspendedOn.IsValid() {
		return nil
	}
	t := x.SuspendedOn.AsTime()
	return &t
}

// SetSpecChangeDateTime is a utility method to set the SpecChangeDate field from a time.Time.
func (x *OwnerSpecs) SetSpecChangeDateTime(t *time.Time) {
	if x == nil {
		return
	}
	if t == nil {
		x.SpecChangeDate = nil
	} else {
		x.SpecChangeDate = timestamppb.New(*t)
	}
}

// GetSpecChangeDateTime is a utility method to get the SpecChangeDate field as a time.Time.
// It returns nil if SpecChangeDate is not set or invalid.
func (x *OwnerSpecs) GetSpecChangeDateTime() *time.Time {
	if x == nil || !x.SpecChangeDate.IsValid() {
		return nil
	}
	t := x.SpecChangeDate.AsTime()
	return &t
}
