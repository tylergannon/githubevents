package main

var webhookTestsTemplate = `
package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"errors"
	"github.com/google/go-github/v43/github"
	"testing"
)

{{ range $_, $webhook := .Webhooks }}

func TestOn{{ $webhook.Event }}Any(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $webhook.Event }}Any(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}["*"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.on{{ $webhook.Event }}["*"]))
			}
		})
	}
}

func TestSetOn{{ $webhook.Event }}Any(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
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
			g.SetOn{{ $webhook.Event }}Any(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				return nil
			})
			g.SetOn{{ $webhook.Event }}Any(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}["*"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.on{{ $webhook.Event }}["*"]), tt.want)
			}
		})
	}
}

func TestHandle{{ $webhook.Event }}Any(t *testing.T) {
	{{ if $webhook.HasActions }}
	action := "*"
	{{ end }}

	type args struct {
		deliveryID string
		eventName  string
		event      *github.{{ $webhook.Event }}
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
				eventName:  "{{ $webhook.Name }}",
				{{ if $webhook.HasActions }}
				event:    &github.{{ $webhook.Event }}{Action: &action},
				{{ else }}
				event:    &github.{{ $webhook.Event }}{},
				{{ end }}
				fail:     false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				{{ if $webhook.HasActions }}
				event:    &github.{{ $webhook.Event }}{Action: &action},
				{{ else }}
				event:    &github.{{ $webhook.Event }}{},
				{{ end }}
				fail:     true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:    nil,
				fail:     false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $webhook.Event }}Any(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handle{{ $webhook.Event }}Any(tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandle{{ $webhook.Event }}Any() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

{{ range $_, $action := $webhook.Actions }}

func TestOn{{ $action.Handler }}(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				callbacks: []{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFunc",
			args: args{
				callbacks: []{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $action.Handler }}(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}["{{ $action.Action }}"]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.on{{ $webhook.Event }}["{{ $action.Action }}"]))
			}
		})
	}
}

func TestSetOn{{ $action.Handler }}(t *testing.T) {
	type args struct {
		callbacks []{{ $webhook.Event }}HandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single {{ $webhook.Event }}HandleFunc",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple {{ $webhook.Event }}HandleFuncs",
			args: args{
				[]{{ $webhook.Event }}HandleFunc{
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
						return nil
					},
					func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
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
			g.SetOn{{ $action.Handler }}(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				return nil
			})
			g.SetOn{{ $action.Handler }}(tt.args.callbacks...)
			if len(g.on{{ $webhook.Event }}["{{ $action.Action }}"]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.on{{ $webhook.Event }}["{{ $action.Action }}"]), tt.want)
			}
		})
	}
}

func TestHandle{{ $action.Handler }}(t *testing.T) {
	action := "{{ $action.Action }}"

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.{{ $webhook.Event }}
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
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "{{ $webhook.Name }}",
				event:      &github.{{ $webhook.Event }}{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.On{{ $action.Handler }}(func(deliveryID string, eventName string, event *github.{{ $webhook.Event }}) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handle{{ $action.Handler }}(tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handle{{ $action.Handler }}() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

{{ end }}

{{ end }}
`