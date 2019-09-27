package internal

import (
	"context"

	tracingLog "github.com/opentracing/opentracing-go/log"

	"github.com/opentracing/opentracing-go"
)

// ServiceDummy ...
func ServiceDummy(ctx context.Context, userID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Service.DummySvc")
	defer span.Finish()

	var tracingFields = []tracingLog.Field{}
	tracingFields = append(tracingFields, tracingLog.Int("user_id", userID))

	// go to repo dummy 1
	userData, err := RepoGetUserData(ctx, userID)
	if err != nil {
		tracingFields = append(tracingFields, tracingLog.Object("error RepoGetUserData", err))
		span.LogFields(tracingFields...)
		return err
	}

	// go to repo dummy 2
	userData.Name = "New Andrew Smith"
	userData.Email = "new.andrew.smith@123.com"
	err = RepoUpdateUserData(ctx, userData)
	if err != nil {
		tracingFields = append(tracingFields, tracingLog.Object("error RepoGetUserData", err))
		span.LogFields(tracingFields...)
		return err
	}

	tracingFields = append(tracingFields, tracingLog.Object("response", nil))
	span.LogFields(tracingFields...)
	return nil

}
