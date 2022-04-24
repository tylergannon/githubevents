package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnPullRequestReviewEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFuncs",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewEvent["*"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewEventAny(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFuncs",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnPullRequestReviewEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewEventAny(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",

				event: &github.PullRequestReviewEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",

				event: &github.PullRequestReviewEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventAny(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePullRequestReviewEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewEventSubmitted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventSubmitted(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["submitted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewEvent["submitted"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewEventSubmitted(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFuncs",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnPullRequestReviewEventSubmitted(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewEventSubmitted(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["submitted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewEvent["submitted"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewEventSubmitted(t *testing.T) {
	action := "submitted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventSubmitted(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewEventSubmitted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewEventSubmitted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewEvent["edited"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewEventEdited(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFuncs",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnPullRequestReviewEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewEventEdited(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventEdited(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnPullRequestReviewEventDismissed(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFunc",
			args: args{
				callbacks: []PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventDismissed(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["dismissed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPullRequestReviewEvent["dismissed"]))
			}
		})
	}
}

func TestSetOnPullRequestReviewEventDismissed(t *testing.T) {
	type args struct {
		callbacks []PullRequestReviewEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PullRequestReviewEventHandleFunc",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PullRequestReviewEventHandleFuncs",
			args: args{
				[]PullRequestReviewEventHandleFunc{
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnPullRequestReviewEventDismissed(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				return nil
			})
			g.SetOnPullRequestReviewEventDismissed(tt.args.callbacks...)
			if len(g.onPullRequestReviewEvent["dismissed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPullRequestReviewEvent["dismissed"]), tt.want)
			}
		})
	}
}

func TestHandlePullRequestReviewEventDismissed(t *testing.T) {
	action := "dismissed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PullRequestReviewEvent
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "pull_request_review",
				event:      &github.PullRequestReviewEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPullRequestReviewEventDismissed(func(deliveryID string, eventName string, event *github.PullRequestReviewEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePullRequestReviewEventDismissed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handlePullRequestReviewEventDismissed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}