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

func TestOnPingEventAny(t *testing.T) {
	type args struct {
		callbacks []PingEventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single PingEventHandleFunc",
			args: args{
				[]PingEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple PingEventHandleFuncs",
			args: args{
				[]PingEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPingEventAny(tt.args.callbacks...)
			if len(g.onPingEvent[PingEventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onPingEvent[PingEventAnyAction]))
			}
		})
	}
}

func TestSetOnPingEventAny(t *testing.T) {
	type args struct {
		callbacks []PingEventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single PingEventHandleFunc",
			args: args{
				[]PingEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple PingEventHandleFuncs",
			args: args{
				[]PingEventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
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
			g.SetOnPingEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
				return nil
			})
			g.SetOnPingEventAny(tt.args.callbacks...)
			if len(g.onPingEvent[PingEventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onPingEvent[PingEventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandlePingEventAny(t *testing.T) {
	ctx := context.Background()

	type args struct {
		deliveryID string
		eventName  string
		event      *github.PingEvent
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
				eventName:  "ping",

				event: &github.PingEvent{},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "ping",

				event: &github.PingEvent{},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "ping",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnPingEventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handlePingEventAny(ctx, tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandlePingEventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPingEvent(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.PingEvent
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger PingEventAny with unknown event action",
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
					onPingEvent: map[string][]PingEventHandleFunc{
						PingEventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.PingEvent) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  PingEvent,

				event: &github.PingEvent{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
			}
			if err := g.PingEvent(ctx, tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("PingEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
