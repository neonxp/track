package tracker

import (
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestTracker(t *testing.T) {
	fs := afero.NewMemMapFs()
	tracker, err := New(fs)
	if err != nil {
		t.Errorf("Must no err, got %v", err)
	}
	tid1, err := tracker.Add("activity 1", []string{}, []string{})
	if err != nil {
		t.Errorf("Must no err, got %v", err)
	}
	if tid1 != 1 {
		t.Errorf("Expected task id = 1, got %d", tid1)
	}
	tid2, err := tracker.Add("activity 2", []string{"tag1", "tag2"}, []string{"context1"})
	if err != nil {
		t.Errorf("Must no err, got %v", err)
	}
	if tid2 != 2 {
		t.Errorf("Expected task id = 2, got %d", tid2)
	}
	if err = tracker.Start(tid1, "work 1"); err != nil {
		t.Errorf("Must no err, got %v", err)
	}
	list := tracker.List(false)
	if len(list) != 1 {
		t.Errorf("List %v expected to be from 1 element", list)
	}
	list2 := tracker.List(true)
	if len(list2) != 2 {
		t.Errorf("List %v expected to be from 2 elements", list2)
	}
	<- time.After(2 * time.Second)
	if err := tracker.Stop(tid1); err != nil {
		t.Errorf("Must no err, got %v", err)
	}
	list3 := tracker.List(false)
	if len(list3) != 0 {
		t.Errorf("List %v expected to be from 0 element", list3)
	}
	list4 := tracker.List(true)
	for _, activity := range list4 {
		if activity.ID != tid1 {
			continue
		}
		if len(activity.Spans) != 1 {
			t.Errorf("List %v expected to be from 1 element", activity.Spans)
		}
		sp := activity.Spans[0]
		if sp.Stop == nil {
			t.Errorf("Span end time must be not empty")
		}
		if !sp.Stop.After(sp.Start) {
			t.Errorf("End span must be after start time")
		}
		if int(sp.Stop.Sub(sp.Start).Seconds()) != 2 {
			t.Errorf("difference between %v and %v must be 2 seconds, got %f", sp.Start, sp.Stop, sp.Stop.Sub(sp.Start).Seconds())
		}
	}
}
