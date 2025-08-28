package model

import (
	"time"
)

// SetAnalyzedOnTime is a utility method to set the AnalyzedOn field from a time.Time.
func (x *Container) SetAnalyzedOnTime(t *time.Time) {
	if x != nil {
		x.AnalyzedOn = NewTimestamp(t)
	}
}

// GetAnalyzedOnTime is a utility method to get the AnalyzedOn field as a time.Time.
// It returns nil if AnalyzedOn is not set or invalid.
func (x *Container) GetAnalyzedOnTime() (t *time.Time) {
	if x != nil {
		t = AsTime(x.AnalyzedOn)
	}
	return
}

// SetSuspendedOnTime is a utility method to set the SuspendedOn field from a time.Time.
func (x *Container) SetSuspendedOnTime(t *time.Time) {
	if x != nil {
		x.SuspendedOn = NewTimestamp(t)
	}
}

// GetSuspendedOnTime is a utility method to get the SuspendedOn field as a time.Time.
// It returns nil if SuspendedOn is not set or invalid.
func (x *Container) GetSuspendedOnTime() (t *time.Time) {
	if x != nil {
		t = AsTime(x.SuspendedOn)
	}
	return
}

// SetSpecChangeDateTime is a utility method to set the SpecChangeDate field from a time.Time.
func (x *OwnerSpecs) SetSpecChangeDateTime(t *time.Time) {
	if x != nil {
		x.SpecChangeDate = NewTimestamp(t)
	}
}

// GetSpecChangeDateTime is a utility method to get the SpecChangeDate field as a time.Time.
// It returns nil if SpecChangeDate is not set or invalid.
func (x *OwnerSpecs) GetSpecChangeDateTime() (t *time.Time) {
	if x != nil {
		t = AsTime(x.SpecChangeDate)
	}
	return
}
