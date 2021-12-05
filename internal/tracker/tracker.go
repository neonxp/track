package tracker

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/afero"
)

var (
	ErrEntryAlreadyExists = fmt.Errorf("entry with this title already exists")
	ErrActivityNotFound   = fmt.Errorf("there is no activity with given id")
)

type Tracker struct {
	fs       afero.Fs
	document *Document
}

const FileName = "gotrack.json"

func New(fs afero.Fs) (*Tracker, error) {
	t := &Tracker{fs: fs}
	if err := t.load(); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tracker) Add(title string, tags []string, contexts []string) (int, error) {
	for _, activity := range t.document.Activities {
		if strings.ToLower(strings.Trim(activity.Title, " ")) == strings.ToLower(strings.Trim(title, " ")) {
			return 0, ErrEntryAlreadyExists
		}
	}
	activity := Activity{
		ID:      t.document.LastKey + 1,
		Title:   title,
		Tags:    tags,
		Context: contexts,
		Spans:   []*Span{},
	}

	t.document = &Document{
		LastKey:    activity.ID,
		Activities: append(t.document.Activities, &activity),
	}

	return t.document.LastKey, t.save()
}

func (t *Tracker) Start(id int, comment string) error {
	if err := t.Stop(id); err != nil {
		return err
	}
	span := &Span{
		Start:   time.Now(),
		Comment: comment,
	}
	activity := t.Activity(id)
	activity.Spans = append(activity.Spans, span)
	return t.save()
}

func (t *Tracker) Stop(id int) error {
	activity := t.Activity(id)
	if activity == nil {
		return ErrActivityNotFound
	}
	if span := activity.Started(); span != nil {
		t := time.Now()
		span.Stop = &t
	}
	return t.save()
}

func (t *Tracker) List(all bool) []*Activity {
	if all {
		return t.document.Activities
	}
	return filterActivities(t.document.Activities, func(a *Activity) bool {
		return a.Started() != nil
	})
}

func (t *Tracker) Activity(id int) *Activity {
	for _, activity := range t.document.Activities {
		if activity.ID == id {
			return activity
		}
	}
	return nil
}

func (t *Tracker) load() (err error) {
	t.document = new(Document)
	f, err := t.fs.Open(FileName)
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		f, err = t.fs.Create(FileName)
		if err != nil {
			return err
		}
		return nil
	}
	return json.NewDecoder(f).Decode(t.document)
}

func (t *Tracker) save() (err error) {
	f, err := t.fs.Create(FileName)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	return json.NewEncoder(f).Encode(t.document)
}

func filterActivities(list []*Activity, filter func(activity *Activity) bool) []*Activity {
	var filtered []*Activity
	for _, activity := range list {
		if filter(activity) {
			filtered = append(filtered, activity)
		}
	}
	return filtered
}
