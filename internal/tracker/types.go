package tracker

import (
	"fmt"
	"time"
)

type Activity struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Context []string `json:"context"`
	Spans   []*Span  `json:"spans"`
}

func (a *Activity) Started() *Span {
	for _, span := range a.Spans {
		if span.Stop == nil {
			return span
		}
	}
	return nil
}

func (a *Activity) Duration() time.Duration {
	t := time.Duration(0)
	for _, span := range a.Spans {
		stop := time.Now()
		if span.Stop != nil {
			stop = *span.Stop
		}
		t += stop.Sub(span.Start)
	}
	return t
}

func (a *Activity) LastDuration() time.Duration {
	if len(a.Spans) == 0 {
		return 0
	}
	span := a.Spans[len(a.Spans)-1]
	stop := time.Now()
	if span.Stop != nil {
		stop = *span.Stop
	}
	return stop.Sub(span.Start)
}

type Span struct {
	Start   time.Time  `json:"start"`
	Stop    *time.Time `json:"stop,omitempty"`
	Comment string     `json:"comment,omitempty"`
}

type Document struct {
	LastKey    int         `json:"last"`
	Activities []*Activity `json:"activities"`
}

type Timespan time.Duration

func (t Timespan) Format() string {
	z := time.Unix(0, 0).UTC()
	tt := z.Add(time.Duration(t))
	result := fmt.Sprintf("%d minutes", tt.Minute())
	if tt.Hour() > 0 {
		result = fmt.Sprintf("%d hours, ", tt.Hour()) + result
	}
	if tt.Day() > 1 {
		result = fmt.Sprintf("%d days, ", tt.Day()-1) + result
	}
	if tt.Month() > 1 {
		result = fmt.Sprintf("%d months, ", tt.Month()-1) + result
	}
	if tt.Year() > 1970 {
		result = fmt.Sprintf("%d years, ", tt.Year()-1970) + result
	}
	return result
}
