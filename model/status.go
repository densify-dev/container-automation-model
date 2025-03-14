package model

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var version string

func SetVersion(v string) {
	version = v
}

type Subject string

const (
	Overall            Subject = "Overall"
	TLS                Subject = "TLS"
	DensifyCredentials Subject = "DensifyCredentials"
	Config             Subject = "ConfigMap"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusOK
	StatusWarning
	StatusError
)

const (
	// StatusUnknownStr is the string representation of StatusUnknown
	StatusUnknownStr = "unknown"
	// StatusOKStr is the string representation of StatusOK
	StatusOKStr = "ok"
	// StatusWarningStr is the string representation of StatusWarning
	StatusWarningStr = "warning"
	// StatusErrorStr is the string representation of StatusError
	StatusErrorStr = "error"
)

func (s Status) String() string {
	switch s {
	case StatusOK:
		return StatusOKStr
	case StatusWarning:
		return StatusWarningStr
	case StatusError:
		return StatusErrorStr
	default:
		return StatusUnknownStr
	}
}

func ParseStatus(s string) (st Status, err error) {
	switch Unquote(strings.ToLower(s)) {
	case StatusOKStr:
		st = StatusOK
	case StatusWarningStr:
		st = StatusWarning
	case StatusErrorStr:
		st = StatusError
	case StatusUnknownStr:
	default:
		err = fmt.Errorf("invalid status: %s", s)
	}
	return
}

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(Quote(s.String())), nil
}

func (s *Status) UnmarshalJSON(data []byte) (err error) {
	*s, err = ParseStatus(string(data))
	return
}

type SubjectStatus struct {
	Version string      `json:"macVersion"`
	Time    RFC3339Time `json:"time"`
	Status  Status      `json:"status"`
	PodName string      `json:"podName"`
	Details string      `json:"statusDetails,omitempty"`
}

type MacStatus struct {
	m        map[Subject]*SubjectStatus
	mu       *sync.Mutex
	modified bool
}

var MacStat = NewMacStatus()

func NewMacStatus() *MacStatus {
	return &MacStatus{
		m:  make(map[Subject]*SubjectStatus),
		mu: &sync.Mutex{},
	}
}

func NewStatus(status Status, details string) (st *SubjectStatus) {
	st = newStatus()
	st.Status = status
	st.Details = details
	return
}

func (ms MacStatus) SetStatus(subject Subject, st *SubjectStatus) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.m[subject] = st
	ms.modified = true
}

// GetOverallStatus returns the overall status of the MacStatus and
// a boolean indicating whether the status has been modified since the last call.
// It sets the modified flag to false after returning the status.
func (ms MacStatus) GetOverallStatus() (overallStatus *SubjectStatus, modified bool) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	modified = ms.modified
	ms.modified = false
	overallStatus = newStatus()
	for s, v := range ms.m {
		aggregate(s, overallStatus, v)
	}
	return
}

func newStatus() *SubjectStatus {
	return &SubjectStatus{
		Version: version,
		Time:    RFC3339Time(time.Now()),
		PodName: GetPodName(),
	}
}

func aggregate(subject Subject, overallStatus, subjectStatus *SubjectStatus) {
	if overallStatus == nil || subjectStatus == nil {
		return
	}
	if subjectStatus.Status > overallStatus.Status {
		overallStatus.Status = subjectStatus.Status
		overallStatus.Details = formatDetails(subject, subjectStatus.Details)
	} else if subjectStatus.Status == overallStatus.Status {
		overallStatus.Details = fmt.Sprintf("%s; %s", overallStatus.Details, formatDetails(subject, subjectStatus.Details))
	}
}

func formatDetails(subject Subject, details string) string {
	return fmt.Sprintf("%v: %s", subject, details)
}
