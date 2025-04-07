package model

import (
	"fmt"
	"os"
	"strconv"
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

const (
	Empty       = ""
	DoubleQuote = `"`
	Dash        = "-"
	Colon       = ":"
	Comma       = ","
	Space       = " "
	Or          = "|"
	NilString   = "<nil>"
)

func Quote(s string) string {
	return fmt.Sprintf("%s%s%s", DoubleQuote, s, DoubleQuote)
}

func Unquote(s string) string {
	return strings.Trim(s, DoubleQuote)
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

type StringBool bool

func (sb StringBool) MarshalJSON() ([]byte, error) {
	return []byte(Quote(strconv.FormatBool(bool(sb)))), nil
}

func (sb *StringBool) UnmarshalJSON(data []byte) (err error) {
	var b bool
	if b, err = strconv.ParseBool(Unquote(string(data))); err == nil {
		*sb = StringBool(b)
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
