package plugins

import (
	"context"
	"fmt"

	"github.com/google/go-github/v62/github"

	"github.com/tylergannon/githubevents/githubevents"
)

func NewResponder(msg string) githubevents.IssueCommentEventHandleFunc {
	// do some configuration here
	// ...
	return func(_ context.Context, deliveryID string, eventName string, event *github.IssueCommentEvent) error {
		fmt.Printf("commenting %s", msg)
		return nil
	}
}
