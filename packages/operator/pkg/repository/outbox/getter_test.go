package outbox_test

import (

	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/odahu/odahu-flow/packages/operator/api/v1alpha1"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/deployment"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/event"
	"github.com/odahu/odahu-flow/packages/operator/pkg/repository/outbox"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestModelRouteGet(t *testing.T) {
	eventPublisher := &outbox.EventPublisher{DB: db}
	routeEventGetter := &outbox.RouteEventGetter{DB: db}


	payload1 := deployment.ModelRoute{Spec: v1alpha1.ModelRouteSpec{URLPrefix: "Test"}}
	event1 := event.Event{
		EntityID:   "event1", EventType: event.ModelRouteCreatedEventType,
		EventGroup: event.ModelRouteEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    payload1,
	}
	err := eventPublisher.PublishEvent(context.Background(), nil, event1)
	assert.NoError(t, err)

	routes, newC, err := routeEventGetter.Get(context.TODO(), 0)
	assert.NoError(t, err)
	assert.Len(t, routes, 1)
	r := routes[0]
	assert.Equal(t, event1.EntityID, r.EntityID)
	assert.Equal(t, event1.EventType, r.EventType)
	assert.Equal(t, payload1, r.Payload)
	assert.True(t, event1.Datetime.Equal(r.Datetime))

	payload2 := deployment.ModelRoute{Spec: v1alpha1.ModelRouteSpec{URLPrefix: "Test"}}
	event2 := event.Event{
		EntityID:   "event2", EventType: event.ModelRouteCreatedEventType,
		EventGroup: event.ModelRouteEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    payload2,
	}
	err = eventPublisher.PublishEvent(context.Background(), nil, event2)
	assert.NoError(t, err)

	payload3 := deployment.ModelRoute{Spec: v1alpha1.ModelRouteSpec{URLPrefix: "Test"}}
	event3 := event.Event{
		EntityID:   "event3", EventType: event.ModelRouteCreatedEventType,
		EventGroup: event.ModelRouteEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    payload3,
	}
	err = eventPublisher.PublishEvent(context.Background(), nil, event3)
	assert.NoError(t, err)

	// Fetch next route updates
	routes, _, err = routeEventGetter.Get(context.TODO(), newC)
	assert.NoError(t, err)
	assert.Len(t, routes, 2)
	r = routes[0]
	assert.Equal(t, event2.EntityID, r.EntityID)
	assert.Equal(t, event2.EventType, r.EventType)
	assert.Equal(t, payload2, r.Payload)
	assert.True(t, event2.Datetime.Equal(r.Datetime))
	r = routes[1]
	assert.Equal(t, event3.EntityID, r.EntityID)
	assert.Equal(t, event3.EventType, r.EventType)
	assert.Equal(t, payload3, r.Payload)
	assert.True(t, event3.Datetime.Equal(r.Datetime))

	stmt, _, _ := sq.Delete(outbox.Table).ToSql()
	_, _ = db.Exec(stmt)

}

func TestModelDeploymentGet(t *testing.T) {
	eventPublisher := &outbox.EventPublisher{DB: db}
	deploymentEventGetter := &outbox.DeploymentEventGetter{DB: db}


	payload1 := deployment.ModelDeployment{Spec: v1alpha1.ModelDeploymentSpec{Image: "Image"}}
	event1 := event.Event{
		EntityID:   "event1", EventType: event.ModelDeploymentCreatedEventType,
		EventGroup: event.ModelDeploymentEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    payload1,
	}
	err := eventPublisher.PublishEvent(context.Background(), nil, event1)
	assert.NoError(t, err)

	routes, newC, err := deploymentEventGetter.Get(context.TODO(), 0)
	assert.NoError(t, err)
	assert.Len(t, routes, 1)
	r := routes[0]
	assert.Equal(t, event1.EntityID, r.EntityID)
	assert.Equal(t, event1.EventType, r.EventType)
	assert.Equal(t, payload1, r.Payload)
	assert.True(t, event1.Datetime.Equal(r.Datetime))

	payload2 := deployment.ModelDeployment{Spec: v1alpha1.ModelDeploymentSpec{Image: "Image"}}
	event2 := event.Event{
		EntityID:   "event2", EventType: event.ModelDeploymentCreatedEventType,
		EventGroup: event.ModelDeploymentEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    payload2,
	}
	err = eventPublisher.PublishEvent(context.Background(), nil, event2)
	assert.NoError(t, err)

	event3 := event.Event{
		EntityID:   "event3", EventType: event.ModelDeploymentDeletedEventType,
		EventGroup: event.ModelDeploymentEventGroup, Datetime:   time.Now().Round(time.Microsecond).UTC(),
		Payload:    nil,
	}
	err = eventPublisher.PublishEvent(context.Background(), nil, event3)
	assert.NoError(t, err)

	// Fetch next route updates
	routes, _, err = deploymentEventGetter.Get(context.TODO(), newC)
	assert.NoError(t, err)
	assert.Len(t, routes, 2)
	r = routes[0]
	assert.Equal(t, event2.EntityID, r.EntityID)
	assert.Equal(t, event2.EventType, r.EventType)
	assert.Equal(t, payload2, r.Payload)
	assert.True(t, event2.Datetime.Equal(r.Datetime.UTC()))
	r = routes[1]
	assert.Equal(t, event3.EntityID, r.EntityID)
	assert.Equal(t, event3.EventType, r.EventType)
	assert.True(t, event3.Datetime.Equal(r.Datetime.UTC()))

	stmt, _, _ := sq.Delete(outbox.Table).ToSql()
	_, _ = db.Exec(stmt)

}