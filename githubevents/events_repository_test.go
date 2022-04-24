package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

func TestOnRepositoryEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventAny(tt.args.callbacks...)
			if len(g.onRepositoryEvent["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["*"]))
			}
		})
	}
}

func TestSetOnRepositoryEventAny(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventAny(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventAny(tt.args.callbacks...)
			if len(g.onRepositoryEvent["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["*"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",

				event: &github.RepositoryEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",

				event: &github.RepositoryEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventAny(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventAny(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleRepositoryEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEvenCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEvenCreated(tt.args.callbacks...)
			if len(g.onRepositoryEvent["created"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["created"]))
			}
		})
	}
}

func TestSetOnRepositoryEvenCreated(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEvenCreated(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEvenCreated(tt.args.callbacks...)
			if len(g.onRepositoryEvent["created"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["created"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEvenCreated(t *testing.T) {
	action := "created"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEvenCreated(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEvenCreated(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEvenCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryEvent["deleted"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["deleted"]))
			}
		})
	}
}

func TestSetOnRepositoryEventDeleted(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventDeleted(tt.args.callbacks...)
			if len(g.onRepositoryEvent["deleted"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["deleted"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventDeleted(t *testing.T) {
	action := "deleted"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventDeleted(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventDeleted(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventArchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventArchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent["archived"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["archived"]))
			}
		})
	}
}

func TestSetOnRepositoryEventArchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventArchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventArchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent["archived"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["archived"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventArchived(t *testing.T) {
	action := "archived"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventArchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventArchived(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventArchived() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventUnarchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventUnarchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent["unarchived"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["unarchived"]))
			}
		})
	}
}

func TestSetOnRepositoryEventUnarchived(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventUnarchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventUnarchived(tt.args.callbacks...)
			if len(g.onRepositoryEvent["unarchived"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["unarchived"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventUnarchived(t *testing.T) {
	action := "unarchived"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventUnarchived(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventUnarchived(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventUnarchived() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryEvent["edited"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["edited"]))
			}
		})
	}
}

func TestSetOnRepositoryEventEdited(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventEdited(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventEdited(tt.args.callbacks...)
			if len(g.onRepositoryEvent["edited"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["edited"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventEdited(t *testing.T) {
	action := "edited"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventEdited(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventEdited(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventRenamed(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventRenamed(tt.args.callbacks...)
			if len(g.onRepositoryEvent["renamed"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["renamed"]))
			}
		})
	}
}

func TestSetOnRepositoryEventRenamed(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventRenamed(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventRenamed(tt.args.callbacks...)
			if len(g.onRepositoryEvent["renamed"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["renamed"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventRenamed(t *testing.T) {
	action := "renamed"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventRenamed(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventRenamed(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventRenamed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventTransferred(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventTransferred(tt.args.callbacks...)
			if len(g.onRepositoryEvent["transferred"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["transferred"]))
			}
		})
	}
}

func TestSetOnRepositoryEventTransferred(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventTransferred(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventTransferred(tt.args.callbacks...)
			if len(g.onRepositoryEvent["transferred"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["transferred"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventTransferred(t *testing.T) {
	action := "transferred"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventTransferred(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventTransferred(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventTransferred() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventPublicized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPublicized(tt.args.callbacks...)
			if len(g.onRepositoryEvent["publicized"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["publicized"]))
			}
		})
	}
}

func TestSetOnRepositoryEventPublicized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventPublicized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventPublicized(tt.args.callbacks...)
			if len(g.onRepositoryEvent["publicized"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["publicized"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventPublicized(t *testing.T) {
	action := "publicized"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPublicized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventPublicized(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventPublicized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnRepositoryEventPrivatized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple RepositoryEventHandleFunc",
			args: args{
				callbacks: []RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPrivatized(tt.args.callbacks...)
			if len(g.onRepositoryEvent["privatized"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onRepositoryEvent["privatized"]))
			}
		})
	}
}

func TestSetOnRepositoryEventPrivatized(t *testing.T) {
	type args struct {
		callbacks []RepositoryEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single RepositoryEventHandleFunc",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple RepositoryEventHandleFuncs",
			args: args{
				[]RepositoryEventHandleFunc{
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
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
			g.SetOnRepositoryEventPrivatized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				return nil
			})
			g.SetOnRepositoryEventPrivatized(tt.args.callbacks...)
			if len(g.onRepositoryEvent["privatized"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onRepositoryEvent["privatized"]), tt.want)
			}
		})
	}
}

func TestHandleRepositoryEventPrivatized(t *testing.T) {
	action := "privatized"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.RepositoryEvent
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
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "repository",
				event:      &github.RepositoryEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnRepositoryEventPrivatized(func(deliveryID string, eventName string, event *github.RepositoryEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleRepositoryEventPrivatized(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleRepositoryEventPrivatized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}