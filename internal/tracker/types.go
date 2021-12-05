package tracker

import "time"

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

type Span struct {
	Start   time.Time  `json:"start"`
	Stop    *time.Time `json:"stop,omitempty"`
	Comment string     `json:"comment,omitempty"`
}

type Document struct {
	LastKey    int         `json:"last"`
	Activities []*Activity `json:"activities"`
}
