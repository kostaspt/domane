package whois

import "time"

type Message struct {
	Domain string `json:"domain"`
	Info
}

type Info struct {
	Raw         string     `json:"-"`
	IsAvailable *bool      `json:"is_available"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

func (i *Info) SetRaw(raw *string) {
	if raw == nil {
		i.Raw = "N/A"
	} else {
		i.Raw = *raw
	}
}

func (i *Info) SetIsAvailable(val bool) {
	i.IsAvailable = &val
}
