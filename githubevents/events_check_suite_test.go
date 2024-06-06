// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/google/go-github/v62/github"
)

func TestOnCheckSuiteEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventAny(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent[CheckSuiteEventAnyAction]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventAny(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventAny(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent[CheckSuiteEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventAny(t *testing.T) {
	ctx := context.Background()

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",

				event: &github.CheckSuiteEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",

				event: &github.CheckSuiteEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventAny(ctx, tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleCheckSuiteEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventCompleted(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventCompletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent[CheckSuiteEventCompletedAction]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventCompleted(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventCompleted(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventCompleted(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventCompletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent[CheckSuiteEventCompletedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventCompleted(t *testing.T) {
	ctx := context.Background()
	action := CheckSuiteEventCompletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventCompleted(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventCompleted(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent[CheckSuiteEventRequestedAction]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventRequested(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent[CheckSuiteEventRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventRequested(t *testing.T) {
	ctx := context.Background()
	action := CheckSuiteEventRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventRequested(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventRequested(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCheckSuiteEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple CheckSuiteEventHandleFunc",
			args: args{
				callbacks: []CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventReRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction]))
			}
		})
	}
}

func TestSetOnCheckSuiteEventReRequested(t *testing.T) {
	type args struct {
		callbacks []CheckSuiteEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single CheckSuiteEventHandleFunc",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple CheckSuiteEventHandleFuncs",
			args: args{
				[]CheckSuiteEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
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
			g.SetOnCheckSuiteEventReRequested(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				return nil
			})
			g.SetOnCheckSuiteEventReRequested(tt.args.callbacks...)
			if len(g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onCheckSuiteEvent[CheckSuiteEventReRequestedAction]), tt.want)
			}
		})
	}
}

func TestHandleCheckSuiteEventReRequested(t *testing.T) {
	ctx := context.Background()
	action := CheckSuiteEventReRequestedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
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
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnCheckSuiteEventReRequested(func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleCheckSuiteEventReRequested(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleCheckSuiteEventReRequested() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckSuiteEvent(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.CheckSuiteEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger CheckSuiteEventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  CheckSuiteEvent,

				event: &github.CheckSuiteEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger CheckSuiteEventCompleted",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventCompletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString(CheckSuiteEventCompletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckSuiteEventCompleted with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventCompletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckSuiteEventCompleted with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventCompletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventCompletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger CheckSuiteEventRequested",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString(CheckSuiteEventRequestedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckSuiteEventRequested with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckSuiteEventRequested with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger CheckSuiteEventReRequested",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventReRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString(CheckSuiteEventReRequestedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail CheckSuiteEventReRequested with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventReRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail CheckSuiteEventReRequested with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onCheckSuiteEvent: map[string][]CheckSuiteEventHandleFunc{
						CheckSuiteEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						CheckSuiteEventReRequestedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.CheckSuiteEvent) error {
								t.Logf("%s action called", CheckSuiteEventReRequestedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "check_suite",
				event:      &github.CheckSuiteEvent{Action: nil},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.CheckSuiteEvent(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("CheckSuiteEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
