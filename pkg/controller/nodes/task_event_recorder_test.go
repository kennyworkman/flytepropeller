package nodes

import (
	"context"
	"fmt"
	"github.com/flyteorg/flyteidl/clients/go/events"
	eventsErr "github.com/flyteorg/flyteidl/clients/go/events/errors"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/event"
	"testing"
)

type fakeTaskEventsRecorder struct {
	err error
}

func (f fakeTaskEventsRecorder) RecordTaskEvent(ctx context.Context, event *event.TaskExecutionEvent) error {
	if f.err != nil {
		return f.err
	}
	return nil
}

func Test_taskEventRecorder_RecordTaskEvent(t1 *testing.T) {
	noErrRecorder := fakeTaskEventsRecorder{}
	alreadyExistsError := fakeTaskEventsRecorder{&eventsErr.EventError{Code: eventsErr.AlreadyExists, Cause: fmt.Errorf("err")}}
	inTerminalError := fakeTaskEventsRecorder{&eventsErr.EventError{Code: eventsErr.EventAlreadyInTerminalStateError, Cause: fmt.Errorf("err")}}
	otherError := fakeTaskEventsRecorder{&eventsErr.EventError{Code: eventsErr.ResourceExhausted, Cause: fmt.Errorf("err")}}

	tests := []struct {
		name    string
		rec  events.TaskEventRecorder
		p core.TaskExecution_Phase
		wantErr bool
	}{
		{"aborted-success", noErrRecorder, core.TaskExecution_ABORTED, false},
		{"aborted-failure", otherError, core.TaskExecution_ABORTED, true},
		{"aborted-already", alreadyExistsError, core.TaskExecution_ABORTED, false},
		{"aborted-terminal", inTerminalError, core.TaskExecution_ABORTED, false},
		{"running-terminal", inTerminalError, core.TaskExecution_RUNNING, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := taskEventRecorder{
				TaskEventRecorder: tt.rec,
			}
			ev := &event.TaskExecutionEvent{
				Phase: tt.p,
			}
			if err := t.RecordTaskEvent(context.TODO(), ev); (err != nil) != tt.wantErr {
				t1.Errorf("RecordTaskEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}