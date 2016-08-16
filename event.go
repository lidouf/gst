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

type Event struct {
	glib.Object
}

func (e *Event) g() *C.GstEvent {
	return (*C.GstEvent)(e.GetPtr())
}

func (e *Event) AsEvent() *Event {
	return e
}

//Allocate a new seek event with the given parameters.
//
//The seek event configures playback of the pipeline between start to stop at the speed given in rate , also called a playback segment.
//The start and stop values are expressed in format .
//
//A rate of 1.0 means normal playback rate, 2.0 means double speed. Negatives values means backwards playback.
//A value of 0.0 for the rate is not allowed and should be accomplished instead by PAUSING the pipeline.
//
//A pipeline has a default playback segment configured with a start position of 0, a stop position of -1 and a rate of 1.0.
//The currently configured playback segment can be queried with GST_QUERY_SEGMENT.
//
//start_type and stop_type specify how to adjust the currently configured start and stop fields in playback segment.
//Adjustments can be made relative or absolute to the last configured values. A type of GST_SEEK_TYPE_NONE means that the position should not be updated.
//
//When the rate is positive and start has been updated, playback will start from the newly configured start position.
//
//For negative rates, playback will start from the newly configured stop position (if any).
//If the stop position is updated, it must be different from -1 (GST_CLOCK_TIME_NONE) for negative rates.
//
//It is not possible to seek relative to the current playback position, to do this, PAUSE the pipeline,
//query the current playback position with GST_QUERY_POSITION and update the playback segment current position with a GST_SEEK_TYPE_SET to the desired position.
//Parameters
//rate
//The new playback rate
//format
//The format of the seek values
//flags
//The optional seek flags
//start_type
//The type and flags for the new start position
//start
//The value of the new start position
//stop_type
//The type and flags for the new stop position
//stop
//The value of the new stop position
//Returns
//a new seek event.
func NewSeekEvent(rate float64, format Format, flags SeekFlags, start_type SeekType, start int64, stop_type SeekType, stop int64) *Event {
	e := C.gst_event_new_seek(C.gdouble(rate), C.GstFormat(format), flags.g(), C.GstSeekType(start_type), C.gint64(start), C.GstSeekType(stop_type), C.gint64(stop))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Create a new step event. The purpose of the step event is to instruct a sink to skip amount (expressed in format ) of media.
//It can be used to implement stepping through the video frame by frame or for doing fast trick modes.
//
//A rate of <= 0.0 is not allowed.
//Pause the pipeline, for the effect of rate = 0.0 or first reverse the direction of playback using a seek event to get the same effect as rate < 0.0.
//
//The flush flag will clear any pending data in the pipeline before starting the step operation.
//
//The intermediate flag instructs the pipeline that this step operation is part of a larger step operation.
//Parameters
//format
//the format of amount
//amount
//the amount of data to step
//rate
//the step rate
//flush
//flushing steps
//intermediate
//intermediate steps
//Returns
//a new GstEvent.
func NewStepEvent(format Format, amount uint64, rate float64, flush bool, intermediate bool) *Event {
	e := C.gst_event_new_step(C.GstFormat(format), C.guint64(amount), C.gdouble(rate), gBoolean(flush), gBoolean(intermediate))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}
