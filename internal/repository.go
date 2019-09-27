package internal

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	tracingLog "github.com/opentracing/opentracing-go/log"
)

// UserData ...
type UserData struct {
	ID    int
	Name  string
	Email string
}

// RepoGetUserData ...
func RepoGetUserData(ctx context.Context, userID int) (UserData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository.RepoGetUserData")
	defer span.Finish()

	var tracingFields = []tracingLog.Field{}
	tracingFields = append(tracingFields, tracingLog.Int("user_id", userID))

	// FAKE PROCESS get user request to api
	time.Sleep(78 * time.Millisecond)

	response := UserData{
		ID:    userID,
		Name:  "Andrew Smith",
		Email: "andrew.smith@123.com",
	}

	tracingFields = append(tracingFields, tracingLog.Object("user data response", response))
	span.LogFields(tracingFields...)
	span.SetTag("get_user_data.http_status", http.StatusOK)

	return response, nil

}

// RepoUpdateUserData ...
func RepoUpdateUserData(ctx context.Context, data UserData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository.RepoUpdateUserData")
	defer span.Finish()

	var tracingFields = []tracingLog.Field{}
	tracingFields = append(tracingFields, tracingLog.Object("user_data", data))

	// FAKE PROCESS get user request to api
	// ERROR response
	time.Sleep(3 * time.Second)
	err := errors.New("cannot update user data")

	tracingFields = append(tracingFields, tracingLog.Object("error update user", err))
	span.LogFields(tracingFields...)
	span.SetTag("update_user_data.http_status", http.StatusInternalServerError)

	return err
}
