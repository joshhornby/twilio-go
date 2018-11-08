package twilio

import (
	"context"
	"net/url"
	"testing"
)

func TestGetTaskRouterQueue(t *testing.T) {
	t.Parallel()
	client, server := getServer(taskRouterTaskQueueResponse)
	defer server.Close()

	sid := "WQ63868a235fc1cf3987e6a2b67346273f"
	name := "English"

	queue, err := client.TaskRouter.TaskRouterQueues.Get(context.Background(), sid)
	if err != nil {
		t.Fatal(err)
	}
	if queue.Sid != sid {
		t.Errorf("task router queue: got sid %q, want %q", queue.Sid, sid)
	}

	if queue.FriendlyName != name {
		t.Errorf("ask router queue: got sid %q, want %q", queue.FriendlyName, name)
	}
}

func TestCreateTaskRouterQueue(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()

	data := url.Values{}
	newname := "Some Activity"
	data.Set("FriendlyName", newname)
	data.Set("AssignmentActivitySid", "WA086126500699bba752b0485c185013d1")
	data.Set("ReservationActivitySid", "WAd9ad5e7b1cb9c8327cf7eb14a8a31131")

	queue, err := envClient.TaskRouter.TaskRouterQueues.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if queue.FriendlyName != newname {
		t.Errorf("FriendlyName don't match")
	}
}
