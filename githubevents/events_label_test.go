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

func TestOnLabelEventAny(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventAny(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent[LabelEventAnyAction]))
			}
		})
	}
}

func TestSetOnLabelEventAny(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventAny(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent[LabelEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventAny(t *testing.T) {
	ctx := context.Background()

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",

				event: &github.LabelEvent{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",

				event: &github.LabelEvent{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventAny(ctx, tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleLabelEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventCreated(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventCreated(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent[LabelEventCreatedAction]))
			}
		})
	}
}

func TestSetOnLabelEventCreated(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventCreated(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent[LabelEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventCreated(t *testing.T) {
	ctx := context.Background()
	action := LabelEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventCreated(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventEdited(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventEdited(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent[LabelEventEditedAction]))
			}
		})
	}
}

func TestSetOnLabelEventEdited(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventEdited(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent[LabelEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventEdited(t *testing.T) {
	ctx := context.Background()
	action := LabelEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventEdited(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnLabelEventDeleted(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple LabelEventHandleFunc",
			args: args{
				callbacks: []LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventDeleted(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onLabelEvent[LabelEventDeletedAction]))
			}
		})
	}
}

func TestSetOnLabelEventDeleted(t *testing.T) {
	type args struct {
		callbacks []LabelEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single LabelEventHandleFunc",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple LabelEventHandleFuncs",
			args: args{
				[]LabelEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
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
			g.SetOnLabelEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				return nil
			})
			g.SetOnLabelEventDeleted(tt.args.callbacks...)
			if len(g.onLabelEvent[LabelEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onLabelEvent[LabelEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleLabelEventDeleted(t *testing.T) {
	ctx := context.Background()
	action := LabelEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
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
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnLabelEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleLabelEventDeleted(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleLabelEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLabelEvent(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.LabelEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger LabelEventAny with unknown event action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  LabelEvent,

				event: &github.LabelEvent{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger LabelEventCreated",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString(LabelEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail LabelEventCreated with empty action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail LabelEventCreated with nil action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventCreatedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger LabelEventEdited",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString(LabelEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail LabelEventEdited with empty action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail LabelEventEdited with nil action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventEditedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger LabelEventDeleted",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString(LabelEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail LabelEventDeleted with empty action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail LabelEventDeleted with nil action",
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
					onLabelEvent: map[string][]LabelEventHandleFunc{
						LabelEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
						LabelEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.LabelEvent) error {
								t.Logf("%s action called", LabelEventDeletedAction)
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "label",
				event:      &github.LabelEvent{Action: nil},
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
			if err := g.LabelEvent(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("LabelEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
