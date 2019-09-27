package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	tracingLog "github.com/opentracing/opentracing-go/log"

	"github.com/opentracing/opentracing-go"
)

type payload struct {
	ID    int
	Name  string
	Email string
}

// HandlerDummy ...
func HandlerDummy(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "Handler.DummyHdlr")
	defer span.Finish()

	// dummy payload
	userID, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	var tracingFields = []tracingLog.Field{}
	tracingFields = append(tracingFields,
		tracingLog.String("url path", r.URL.Path),
		tracingLog.Int("path param: user_id", userID),
	)

	// ADDITIONAL PROCESS ...
	time.Sleep(230 * time.Millisecond)

	// go to service
	err := ServiceDummy(ctx, userID)
	if err != nil {
		tracingFields = append(tracingFields, tracingLog.Object("error ServiceDummy()", err))
		span.LogFields(tracingFields...)
		return
	}

	tracingFields = append(tracingFields, tracingLog.Object("response", nil))
	span.LogFields(tracingFields...)
	return
}
