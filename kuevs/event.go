package kuevs

import (
	"fmt"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Event struct {
	UID       string `gorm:"primarykey"`
	Name      string `gorm:"index"`
	Namespace string `gorm:"index"`
	Object    string
	Reason    string
	Message   string
	Source    string
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"primarykey"`
	Count     int
	Type      string
}

func buildSource(s v1.EventSource) string {
	source := []string{}
	if len(s.Component) > 0 {
		source = append(source, s.Component)
	}
	if len(s.Host) > 0 {
		source = append(source, s.Host)
	}

	return strings.Join(source, ",")
}

func buildObject(o v1.ObjectReference) string {
	return fmt.Sprintf("%s/%s", o.Kind, o.Name)
}
func InitEvent(event watch.Event) *Event {
	evt := event.Object.(*v1.Event)

	return &Event{
		UID:       string(evt.UID),
		Name:      evt.Name,
		Namespace: evt.Namespace,
		Object:    buildObject(evt.InvolvedObject),
		Reason:    evt.Reason,
		Message:   evt.Message,
		Source:    buildSource(evt.Source),
		CreatedAt: evt.FirstTimestamp.Time,
		UpdatedAt: evt.LastTimestamp.Time,
		Count:     int(evt.Count),
		Type:      evt.Type,
	}
}
