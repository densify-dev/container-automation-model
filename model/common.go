package model

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Resource string

const (
	Cpu              Resource = "cpu"
	Memory           Resource = "memory"
	EphemeralStorage Resource = "ephemeral-storage"
	HugePages2Mi     Resource = "hugepages-2Mi"
	HugePages1Gi     Resource = "hugepages-1Gi"
	NvidiaGpu        Resource = "nvidia.com/gpu"
)

type Allocation string

const (
	Request Allocation = "request"
	Limit   Allocation = "limit"
)

func Quote(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func Unquote(s string) string {
	return strings.Trim(s, `"`)
}

const (
	RFC3339Micro = "2006-01-02T15:04:05.999999Z07:00"
)

type RFC3339Time time.Time

func (rt RFC3339Time) MarshalJSON() ([]byte, error) {
	return []byte(Quote(time.Time(rt).Format(RFC3339Micro))), nil
}

func (rt *RFC3339Time) UnmarshalJSON(data []byte) (err error) {
	var t time.Time
	if t, err = time.Parse(RFC3339Micro, Unquote(string(data))); err == nil {
		*rt = RFC3339Time(t)
	}
	return
}

const (
	hostnameEnvVar = "HOSTNAME"
)

func GetPodName() string {
	var hostname string
	var err error
	if hostname, err = os.Hostname(); err != nil {
		hostname = os.Getenv(hostnameEnvVar)
	}
	return strings.ToLower(hostname)
}
