//GstEvent — Structure describing events that are passed up and down a pipeline
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	//"errors"
	"github.com/lidouf/glib"
	//"time"
	//"unsafe"
)

type Event struct {
	glib.Object
}

func (o *Event) g() *C.GstEvent {
	return (*C.GstEvent)(o.GetPtr())
}

func (o *Event) AsEvent() *Event {
	return o
}

type EventType C.GstEventType

const (
	EVENT_UNKNOWN = EventType(C.GST_EVENT_UNKNOWN)
	/* bidirectional events */
	EVENT_FLUSH_START = EventType(C.GST_EVENT_FLUSH_START)
	EVENT_FLUSH_STOP  = EventType(C.GST_EVENT_FLUSH_STOP)
	/* downstream serialized events */
	EVENT_STREAM_START = EventType(C.GST_EVENT_STREAM_START)
	EVENT_CAPS         = EventType(C.GST_EVENT_CAPS)
	EVENT_SEGMENT      = EventType(C.GST_EVENT_SEGMENT)
	EVENT_TAG          = EventType(C.GST_EVENT_TAG)
	EVENT_BUFFERSIZE   = EventType(C.GST_EVENT_BUFFERSIZE)
	EVENT_SINK_MESSAGE = EventType(C.GST_EVENT_SINK_MESSAGE)
	EVENT_EOS          = EventType(C.GST_EVENT_EOS)
	EVENT_TOC          = EventType(C.GST_EVENT_TOC)
	EVENT_PROTECTION   = EventType(C.GST_EVENT_PROTECTION)
	/* non-sticky downstream serialized */
	EVENT_SEGMENT_DONE = EventType(C.GST_EVENT_SEGMENT_DONE)
	EVENT_GAP          = EventType(C.GST_EVENT_GAP)
	/* upstream events */
	EVENT_QOS         = EventType(C.GST_EVENT_QOS)
	EVENT_SEEK        = EventType(C.GST_EVENT_SEEK)
	EVENT_NAVIGATION  = EventType(C.GST_EVENT_NAVIGATION)
	EVENT_LATENCY     = EventType(C.GST_EVENT_LATENCY)
	EVENT_STEP        = EventType(C.GST_EVENT_STEP)
	EVENT_RECONFIGURE = EventType(C.GST_EVENT_RECONFIGURE)
	EVENT_TOC_SELECT  = EventType(C.GST_EVENT_TOC_SELECT)
	/* custom events start here */
	EVENT_CUSTOM_UPSTREAM          = EventType(C.GST_EVENT_CUSTOM_UPSTREAM)
	EVENT_CUSTOM_DOWNSTREAM        = EventType(C.GST_EVENT_CUSTOM_DOWNSTREAM)
	EVENT_CUSTOM_DOWNSTREAM_OOB    = EventType(C.GST_EVENT_CUSTOM_DOWNSTREAM_OOB)
	EVENT_CUSTOM_DOWNSTREAM_STICKY = EventType(C.GST_EVENT_CUSTOM_DOWNSTREAM_STICKY)
	EVENT_CUSTOM_BOTH              = EventType(C.GST_EVENT_CUSTOM_BOTH)
	EVENT_CUSTOM_BOTH_OOB          = EventType(C.GST_EVENT_CUSTOM_BOTH_OOB)
)
