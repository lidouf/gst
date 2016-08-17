//GstEvent — Structure describing events that are passed up and down a pipeline
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
static inline
GstEventType CALL_MACRO_GST_EVENT_TYPE(GstEvent *event) {
	return GST_EVENT_TYPE(event);
}
static inline
gboolean CALL_MACRO_gst_event_is_writable(GstEvent *event) {
	return gst_event_is_writable(event);
}
static inline
GstEvent* CALL_MACRO_gst_event_make_writable(GstEvent *event) {
	return gst_event_make_writable(event);
}
*/
import "C"

import (
	//"errors"
	"github.com/lidouf/glib"
	//"time"
	"unsafe"
)

type EventTypeFlags C.GstEventTypeFlags

const (
	EVENT_TYPE_UPSTREAM     = EventTypeFlags(C.GST_EVENT_TYPE_UPSTREAM)
	EVENT_TYPE_DOWNSTREAM   = EventTypeFlags(C.GST_EVENT_TYPE_DOWNSTREAM)
	EVENT_TYPE_SERIALIZED   = EventTypeFlags(C.GST_EVENT_TYPE_SERIALIZED)
	EVENT_TYPE_STICKY       = EventTypeFlags(C.GST_EVENT_TYPE_STICKY)
	EVENT_TYPE_STICKY_MULTI = EventTypeFlags(C.GST_EVENT_TYPE_STICKY_MULTI)
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

//Gets the GstEventTypeFlags associated with type .
//Parameters
//type
//a GstEventType
//Returns
//a GstEventTypeFlags.
func (t EventType) Flags() EventTypeFlags {
	return EventTypeFlags(C.gst_event_type_get_flags(C.GstEventType(t)))
}

//Get a printable name for the given event type. Do not modify or free.
//Parameters
//type
//the event type
//Returns
//a reference to the static name of the event.
func (t EventType) GetName() string {
	return C.GoString((*C.char)(C.gst_event_type_get_name(C.GstEventType(t))))
}

//Get the unique quark for the given event type.
//Parameters
//type
//the event type
//Returns
//the quark associated with the event type
func (t EventType) ToQuark() glib.Quark {
	return glib.Quark(C.gst_event_type_to_quark(C.GstEventType(t)))
}

const EVENT_NUM_SHIFT = 8

//#define             GST_EVENT_MAKE_TYPE(num,flags)
//when making custom event types, use this macro with the num and the given flags
//Parameters
//num
//the event number to create
//flags
//the event flags
//#define GST_EVENT_MAKE_TYPE(num,flags) \
//(((num) << GST_EVENT_NUM_SHIFT) | (flags))
func MakeEventType(num int, flags EventTypeFlags) EventType {
	return EventType(num<<EVENT_NUM_SHIFT | int(flags))
}

type QOSType C.GstQOSType

const (
	QOS_TYPE_OVERFLOW  = QOSType(C.GST_QOS_TYPE_OVERFLOW)
	QOS_TYPE_UNDERFLOW = QOSType(C.GST_QOS_TYPE_UNDERFLOW)
	QOS_TYPE_THROTTLE  = QOSType(C.GST_QOS_TYPE_THROTTLE)
)

type StreamFlags C.GstStreamFlags

const (
	STREAM_FLAG_NONE     = StreamFlags(C.GST_STREAM_FLAG_NONE)
	STREAM_FLAG_SPARSE   = StreamFlags(C.GST_STREAM_FLAG_SPARSE)
	STREAM_FLAG_SELECT   = StreamFlags(C.GST_STREAM_FLAG_SELECT)
	STREAM_FLAG_UNSELECT = StreamFlags(C.GST_STREAM_FLAG_UNSELECT)
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

//#define GST_EVENT_TYPE(event)           (GST_EVENT_CAST(event)->type)
//Get the GstEventType of the event.
//Parameters
//event
//the event to query
func (e *Event) GetType() EventType {
	return EventType(C.CALL_MACRO_GST_EVENT_TYPE(e.g()))
}

//#define GST_EVENT_TYPE_NAME(event)      (gst_event_type_get_name(GST_EVENT_TYPE(event)))
//Get a constant string representation of the GstEventType of the event.
//Parameters
//event
//the event to query
func (e *Event) GetTypeName() string {
	return e.GetType().GetName()
}

//#define GST_EVENT_TIMESTAMP(event)      (GST_EVENT_CAST(event)->timestamp)
//Get the GstClockTime timestamp of the event. This is the time when the event was created.
//Parameters
//event
//the event to query
func (e *Event) Timestamp() uint64 {
	return uint64(e.g().timestamp)
}

//#define GST_EVENT_SEQNUM(event)         (GST_EVENT_CAST(event)->seqnum)
//The sequence number of event .
//Parameters
//event
//the event to query
func (e *Event) SeqNum() uint32 {
	return uint32(e.g().seqnum)
}

//#define GST_EVENT_IS_UPSTREAM(ev)       !!(GST_EVENT_TYPE (ev) & GST_EVENT_TYPE_UPSTREAM)
//Check if an event can travel upstream.
//Parameters
//ev
//the event to query
func (e *Event) IsUpstream() bool {
	return int32(e.GetType())&int32(EVENT_TYPE_UPSTREAM) != 0
}

//#define GST_EVENT_IS_DOWNSTREAM(ev)     !!(GST_EVENT_TYPE (ev) & GST_EVENT_TYPE_DOWNSTREAM)
//Check if an event can travel downstream.
//Parameters
//ev
//the event to query
func (e *Event) IsDownstream() bool {
	return int32(e.GetType())&int32(EVENT_TYPE_DOWNSTREAM) != 0
}

//#define GST_EVENT_IS_SERIALIZED(ev)     !!(GST_EVENT_TYPE (ev) & GST_EVENT_TYPE_SERIALIZED)
//Check if an event is serialized with the data stream.
//Parameters
//ev
//the event to query
func (e *Event) IsSerialized() bool {
	return int32(e.GetType())&int32(EVENT_TYPE_SERIALIZED) != 0
}

//#define GST_EVENT_IS_STICKY(ev)     !!(GST_EVENT_TYPE (ev) & GST_EVENT_TYPE_STICKY)
//Check if an event is sticky on the pads.
//Parameters
//ev
//the event to query
func (e *Event) IsSticky() bool {
	return int32(e.GetType())&int32(EVENT_TYPE_STICKY) != 0
}

//Increase the refcount of this event.
//Parameters
//event
//The event to refcount
//Returns
//event (for convenience when doing assignments).
func (e *Event) Ref() *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_ref(e.g())))
	return r
}

//Decrease the refcount of an event, freeing it if the refcount reaches 0.
//Parameters
//event
//the event to refcount.
func (e *Event) Unref() {
	C.gst_event_unref(e.g())
}

//Modifies a pointer to a GstEvent to point to a different GstEvent.
//The modification is done atomically (so this is useful for ensuring thread safety in some cases),
//and the reference counts are updated appropriately (the old event is unreffed, the new one is reffed).
//
//Either new_event or the GstEvent pointed to by old_event may be NULL.
//Parameters
//old_event
//pointer to a pointer to a GstEvent to be replaced.
//new_event
//pointer to a GstEvent that will replace the event pointed to by old_event .
//Returns
//TRUE if new_event was different from old_event
//func (e *Event) Replace(old_event *Event) bool {
//	return C.gst_event_replace(&old_event.g(), e.g())
//}

//Copy the event using the event specific copy function.
//Parameters
//event
//The event to copy
//Returns
//the new event.
func (e *Event) Copy() *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_copy(e.g())))
	return r
}

//GstEvent *
//gst_event_steal (GstEvent **old_event);

//Atomically replace the GstEvent pointed to by old_event with NULL and return the original event.
//Parameters
//old_event
//pointer to a pointer to a GstEvent to be stolen.
//Returns
//the GstEvent that was in old_event

//gboolean
//gst_event_take (GstEvent **old_event,
//GstEvent *new_event);
//
//Modifies a pointer to a GstEvent to point to a different GstEvent.
//This function is similar to gst_event_replace() except that it takes ownership of new_event .
//
//Either new_event or the GstEvent pointed to by old_event may be NULL.
//Parameters
//old_event
//pointer to a pointer to a GstEvent to be stolen.
//new_event
//pointer to a GstEvent that will replace the event pointed to by old_event .
//Returns
//TRUE if new_event was different from old_event

//#define         gst_event_is_writable(ev)     gst_mini_object_is_writable (GST_MINI_OBJECT_CAST (ev))
//Tests if you can safely write data into a event's structure or validly modify the seqnum and timestamp field.
//Parameters
//ev
//a GstEvent
func (e *Event) IsWritable() bool {
	return C.CALL_MACRO_gst_event_is_writable(e.g()) != 0
}

//#define         gst_event_make_writable(ev)   GST_EVENT_CAST (gst_mini_object_make_writable (GST_MINI_OBJECT_CAST (ev)))
//Makes a writable event from the given event. If the source event is already writable, this will simply return the same event.
//A copy will otherwise be made using gst_event_copy().
//Parameters
//ev
//a GstEvent.
//Returns
//a writable event which may or may not be the same as ev .
func (e *Event) MakeWritable() *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.CALL_MACRO_gst_event_make_writable(e.g())))
	return r
}

//Get a writable version of the structure.
//Parameters
//event
//The GstEvent.
//Returns
//The structure of the event.
//The structure is still owned by the event, which means that you should not free it and that the pointer becomes invalid when you free the event.
//This function checks if event is writable and will never return NULL.
//MT safe.
func (e *Event) WritableStructure() *Structure {
	r := new(Structure)
	r.SetPtr(glib.Pointer(C.gst_event_writable_structure(e.g())))
	return r
}

//Create a new custom-typed event. This can be used for anything not handled by other event-specific functions to pass an event to another element.
//Make sure to allocate an event type with the GST_EVENT_MAKE_TYPE macro, assigning a free number and filling in the correct direction and serialization flags.
//New custom events can also be created by subclassing the event type if needed.
//Parameters
//type
//The type of the new event
//structure
//the structure for the event. The event will take ownership of the structure.
//Returns
//the new custom event.
func NewCustomEvent(tp EventType, structure *Structure) *Event {
	e := C.gst_event_new_custom(C.GstEventType(tp), structure.g())
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Access the structure of the event.
//Parameters
//event
//The GstEvent.
//Returns
//The structure of the event.
//The structure is still owned by the event, which means that you should not free it and that the pointer becomes invalid when you free the event.
//MT safe.
func (e *Event) GetStructure() *Structure {
	r := new(Structure)
	r.SetPtr(glib.Pointer(C.gst_event_get_structure(e.g())))
	return r
}

//Checks if event has the given name . This function is usually used to check the name of a custom event.
//Parameters
//event
//The GstEvent.
//name
//name to check
//Returns
//TRUE if name matches the name of the event structure.
func (e *Event) HasName(name string) bool {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	return C.gst_event_has_name(e.g(), s) != 0
}

//Retrieve the sequence number of a event.
//
//Events have ever-incrementing sequence numbers, which may also be set explicitly via gst_event_set_seqnum().
//Sequence numbers are typically used to indicate that a event corresponds to some other set of events or messages,
//for example an EOS event corresponding to a SEEK event.
//It is considered good practice to make this correspondence when possible, though it is not required.
//
//Note that events and messages share the same sequence number incrementor;
//two events or messages will never have the same sequence number unless that correspondence was made explicitly.
//Parameters
//event
//A GstEvent.
//Returns
//The event's sequence number.
//MT safe.
func (e *Event) GetSeqNum() uint32 {
	return uint32(C.gst_event_get_seqnum(e.g()))
}

//Set the sequence number of a event.
//
//This function might be called by the creator of a event to indicate that the event relates to other events or messages.
//See gst_event_get_seqnum() for more information.
//MT safe.
//Parameters
//event
//A GstEvent.
//seqnum
//A sequence number.
func (e *Event) SetSeqNum(seqnum uint32) {
	C.gst_event_set_seqnum(e.g(), C.guint32(seqnum))
}

//Retrieve the accumulated running time offset of the event.
//
//Events passing through GstPads that have a running time offset set via gst_pad_set_offset() will get their offset adjusted according to the pad's offset.
//
//If the event contains any information that related to the running time, this information will need to be updated before usage with this offset.
//Parameters
//event
//A GstEvent.
//Returns
//The event's running time offset
//MT safe.
//Since: 1.4
func (e *Event) GetRunningTimeOffset() int64 {
	return int64(C.gst_event_get_running_time_offset(e.g()))
}

//Set the running time offset of a event. See gst_event_get_running_time_offset() for more information.
//MT safe.
//Parameters
//event
//A GstEvent.
//offset
//A the new running time offset
//Since: 1.4
func (e *Event) SetRunningTimeOffset(offset int64) {
	C.gst_event_set_running_time_offset(e.g(), C.gint64(offset))
}

//Allocate a new flush start event. The flush start event can be sent upstream and downstream and travels out-of-bounds with the dataflow.
//
//It marks pads as being flushing and will make them return GST_FLOW_FLUSHING when used for data flow with
//gst_pad_push(), gst_pad_chain(), gst_pad_get_range() and gst_pad_pull_range().
//Any event (except a GST_EVENT_FLUSH_STOP) received on a flushing pad will return FALSE immediately.
//
//Elements should unlock any blocking functions and exit their streaming functions as fast as possible when this event is received.
//
//This event is typically generated after a seek to flush out all queued data in the pipeline so that the new media is played as soon as possible.
//Returns
//a new flush start event.
func NewFlushStartEvent() *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_flush_start()))
	return r
}

//Allocate a new flush stop event. The flush stop event can be sent upstream and downstream and travels serialized with the dataflow.
//It is typically sent after sending a FLUSH_START event to make the pads accept data again.
//
//Elements can process this event synchronized with the dataflow since the preceding FLUSH_START event stopped the dataflow.
//
//This event is typically generated to complete a seek and to resume dataflow.
//Parameters
//reset_time
//if time should be reset
//Returns
//a new flush stop event.
func NewFlushStopEvent(reset_time bool) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_flush_stop(gBoolean(reset_time))))
	return r
}

//Parse the FLUSH_STOP event and retrieve the reset_time member.
//Parameters
//event
//The event to parse
//Return
//reset_time
//if time should be reset.
func (e *Event) ParseFlushStop() bool {
	var reset_time C.gboolean
	C.gst_event_parse_flush_stop(e.g(), &reset_time)
	return reset_time != 0
}

//Create a new EOS event. The eos event can only travel downstream synchronized with the buffer flow.
//Elements that receive the EOS event on a pad can return GST_FLOW_EOS as a GstFlowReturn when data after the EOS event arrives.
//
//The EOS event will travel down to the sink elements in the pipeline which will then post the GST_MESSAGE_EOS on the bus
//after they have finished playing any buffered data.
//
//When all sinks have posted an EOS message, an EOS message is forwarded to the application.
//
//The EOS event itself will not cause any state transitions of the pipeline.
//Returns
//the new EOS event.
func NewEosEvent() *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_eos()))
	return r
}

//Create a new GAP event.
//A gap event can be thought of as conceptually equivalent to a buffer to signal that there is no data for a certain amount of time.
//This is useful to signal a gap to downstream elements which may wait for data, such as muxers or mixers or overlays,
//especially for sparse streams such as subtitle streams.
//Parameters
//timestamp
//the start time (pts) of the gap
//duration
//the duration of the gap
//Returns
//the new GAP event.
func NewGapEvent(timestamp, duration ClockTime) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_gap(C.GstClockTime(timestamp), C.GstClockTime(duration))))
	return r
}

//Extract timestamp and duration from a new GAP event.
//Parameters
//event
//a GstEvent of type GST_EVENT_GAP
//Return
//timestamp
//location where to store the start time (pts) of the gap, or NULL.
//duration
//location where to store the duration of the gap, or NULL.
func (e *Event) ParseGap() (ClockTime, ClockTime) {
	var t, d C.GstClockTime
	C.gst_event_parse_gap(e.g(), &t, &d)
	return ClockTime(t), ClockTime(d)
}

//Create a new STREAM_START event. The stream start event can only travel downstream synchronized with the buffer flow.
//It is expected to be the first event that is sent for a new stream.
//
//Source elements, demuxers and other elements that create new streams are supposed to send this event as the first event of a new stream.
//It should not be sent after a flushing seek or in similar situations and is used to mark the beginning of a new logical stream.
//Elements combining multiple streams must ensure that this event is only forwarded downstream once and not for every single input stream.
//
//The stream_id should be a unique string that consists of the upstream stream-id, / as separator and a unique stream-id for this specific stream.
//A new stream-id should only be created for a stream if the upstream stream is split into (potentially) multiple new streams, e.g.
//in a demuxer, but not for every single element in the pipeline.
//gst_pad_create_stream_id() or gst_pad_create_stream_id_printf() can be used to create a stream-id.
//There are no particular semantics for the stream-id, though it should be deterministic (to support stream matching)
//and it might be used to order streams (besides any information conveyed by stream flags).
//
//Parameters
//stream_id
//Identifier for this stream
//Returns
//the new STREAM_START event.
func NewStreamStartEvent(stream_id string) *Event {
	s := (*C.gchar)(C.CString(stream_id))
	defer C.free(unsafe.Pointer(s))
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_stream_start(s)))
	return r
}

//Parse a stream-id event and store the result in the given stream_id location.
//The string stored in stream_id must not be modified and will remain valid only until event gets freed.
//Make a copy if you want to modify it or store it for later use.
//Parameters
//event
//a stream-start event.
//Return
//stream_id
//pointer to store the stream-id.
func (e *Event) ParseStreamStart() string {
	var stream_id *C.gchar
	defer C.free(unsafe.Pointer(stream_id))
	C.gst_event_parse_stream_start(e.g(), &stream_id)
	return C.GoString((*C.char)(stream_id))
}

//Parameters
//event
//a stream-start event
//flags
//the stream flags to set
//Since: 1.2
func (e *Event) SetStreamFlags(flags StreamFlags) {
	C.gst_event_set_stream_flags(e.g(), C.GstStreamFlags(flags))
}

//Parameters
//event
//a stream-start event
//Return
//flags
//address of variable where to store the stream flags.
//Since: 1.2
func (e *Event) ParseStreamFlags() StreamFlags {
	var f C.GstStreamFlags
	C.gst_event_parse_stream_flags(e.g(), &f)
	return StreamFlags(f)
}

//All streams that have the same group id are supposed to be played together,
//i.e. all streams inside a container file should have the same group id but different stream ids.
//The group id should change each time the stream is started, resulting in different group ids each time a file is played for example.
//
//Use gst_util_group_id_next() to get a new group id.
//Parameters
//event
//a stream-start event
//group_id
//the group id to set
//Since: 1.2
func (e *Event) SetGroupId(group_id uint) {
	C.gst_event_set_group_id(e.g(), C.guint(group_id))
}

//Parameters
//event
//a stream-start event
//Returns
//group_id
//address of variable where to store the group id.
//TRUE if a group id was set on the event and could be parsed, FALSE otherwise.
//Since: 1.2
func (e *Event) ParseGroupId() (uint, bool) {
	var g C.guint
	ret := C.gst_event_parse_group_id(e.g(), &g)
	return uint(g), ret != 0
}

//Create a new SEGMENT event for segment .
//The segment event can only travel downstream synchronized with the buffer flow and contains timing information
//and playback properties for the buffers that will follow.
//
//The segment event marks the range of buffers to be processed. All data not within the segment range is not to be processed.
//This can be used intelligently by plugins to apply more efficient methods of skipping unneeded data.
//The valid range is expressed with the start and stop values.
//
//The time value of the segment is used in conjunction with the start value to convert the buffer timestamps into the stream time.
//This is usually done in sinks to report the current stream_time.
//time represents the stream_time of a buffer carrying a timestamp of start . time cannot be -1.
//
//start cannot be -1, stop can be -1. If there is a valid stop given, it must be greater or equal the start ,
//including when the indicated playback rate is < 0.
//
//The applied_rate value provides information about any rate adjustment that has already been made to the timestamps and content on the buffers of the stream.
//(rate * applied_rate ) should always equal the rate that has been requested for playback.
//For example, if an element has an input segment with intended playback rate of 2.0 and applied_rate of 1.0,
//it can adjust incoming timestamps and buffer content by half and output a segment event with rate of 1.0 and applied_rate of 2.0
//
//After a segment event, the buffer stream time is calculated with:
//
//time + (TIMESTAMP(buf) - start) * ABS (rate * applied_rate)
//Parameters
//segment
//a GstSegment.
//
//Returns
//the new SEGMENT event.
func NewSegmentEvent(segment *Segment) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_segment(segment.g())))
	return r
}

//Parses a segment event and stores the result in the given segment location. segment remains valid only until the event is freed.
//Don't modify the segment and make a copy if you want to modify it or store it for later use.
//Parameters
//event
//The event to parse
//Returns
//segment
//a pointer to a GstSegment.
//Need to unref
func (e *Event) ParseSegment() *Segment {
	var s *C.GstSegment
	C.gst_event_parse_segment(e.g(), &s)
	r := new(Segment)
	r.SetPtr(glib.Pointer(s))
	return r
}

//Parses a segment event and copies the GstSegment into the location given by segment .
//Parameters
//event
//The event to parse
//Returns
//segment
//a pointer to a GstSegment
func (e *Event) CopySegment() *Segment {
	var s *C.GstSegment
	C.gst_event_copy_segment(e.g(), s)
	r := new(Segment)
	r.SetPtr(glib.Pointer(&s))
	return r
}

//Generates a metadata tag event from the given taglist .
//
//The scope of the taglist specifies if the taglist applies to the complete medium or only to this specific stream.
//As the tag event is a sticky event, elements should merge tags received from upstream with a given scope with their own tags with the same scope
//and create a new tag event from it.
//Parameters
//taglist
//metadata list. The event will take ownership of the taglist.
//Returns
//a new GstEvent.
func NewTagEvent(taglist *TagList) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_tag(taglist.g())))
	return r
}

//Parses a tag event and stores the results in the given taglist location.
//No reference to the taglist will be returned, it remains valid only until the event is freed. Don't modify or free the taglist,
//make a copy if you want to modify it or store it for later use.
//Parameters
//event
//a tag event
//Returns
//taglist
//pointer to metadata list.
func (e *Event) ParseTag() *TagList {
	var s *C.GstTagList
	C.gst_event_parse_tag(e.g(), &s)
	r := new(TagList)
	r.SetPtr(glib.Pointer(s))
	return r
}

//Create a new buffersize event.
//The event is sent downstream and notifies elements that they should provide a buffer of the specified dimensions.
//When the async flag is set, a thread boundary is preferred.
//Parameters
//format
//buffer format
//minsize
//minimum buffer size
//maxsize
//maximum buffer size
//async
//thread behavior
//Returns
//a new GstEvent.
func NewBufferSizeEvent(format Format, minsize, maxsize int64, async bool) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_buffer_size(C.GstFormat(format), C.gint64(minsize), C.gint64(maxsize), gBoolean(async))))
	return r
}

//Get the format, minsize, maxsize and async-flag in the buffersize event.
//Parameters
//event
//The event to query
//Returns
//format
//A pointer to store the format in.
//minsize
//A pointer to store the minsize in.
//maxsize
//A pointer to store the maxsize in.
//async
//A pointer to store the async-flag in.
func (e *Event) ParseBufferSize() (Format, int64, int64, bool) {
	var format C.GstFormat
	var minsize, maxsize C.gint64
	var async C.gboolean
	C.gst_event_parse_buffer_size(e.g(), &format, &minsize, &maxsize, &async)
	return Format(format), int64(minsize), int64(maxsize), async != 0
}

//Allocate a new qos event with the given values.
//The QOS event is generated in an element that wants an upstream element to either reduce or increase its rate
//because of high/low CPU load or other resource usage such as network performance or throttling.
//Typically sinks generate these events for each buffer they receive.
//
//type indicates the reason for the QoS event.
//GST_QOS_TYPE_OVERFLOW is used when a buffer arrived in time or when the sink cannot keep up with the upstream datarate.
//GST_QOS_TYPE_UNDERFLOW is when the sink is not receiving buffers fast enough and thus has to drop late buffers.
//GST_QOS_TYPE_THROTTLE is used when the datarate is artificially limited by the application, for example to reduce power consumption.
//
//proportion indicates the real-time performance of the streaming in the element that generated the QoS event (usually the sink).
//The value is generally computed based on more long term statistics about the streams timestamps compared to the clock.
//A value < 1.0 indicates that the upstream element is producing data faster than real-time.
//A value > 1.0 indicates that the upstream element is not producing data fast enough. 1.0 is the ideal proportion value.
//The proportion value can safely be used to lower or increase the quality of the element.
//
//diff is the difference against the clock in running time of the last buffer that caused the element to generate the QOS event.
//A negative value means that the buffer with timestamp arrived in time.
//A positive value indicates how late the buffer with timestamp was.
//When throttling is enabled, diff will be set to the requested throttling interval.
//
//timestamp is the timestamp of the last buffer that cause the element to generate the QOS event.
//It is expressed in running time and thus an ever increasing value.
//
//The upstream element can use the diff and timestamp values to decide whether to process more buffers.
//For positive diff , all buffers with timestamp <= timestamp + diff will certainly arrive late in the sink as well.
//A (negative) diff value so that timestamp + diff would yield a result smaller than 0 is not allowed.
//
//The application can use general event probes to intercept the QoS event and implement custom application specific QoS handling.
//Parameters
//type
//the QoS type
//proportion
//the proportion of the qos message
//diff
//The time difference of the last Clock sync
//timestamp
//The timestamp of the buffer
//Returns
//a new QOS event.
func NewQOSEvent(tp QOSType, proportion float64, diff int64, timestamp ClockTime) *Event {
	r := new(Event)
	r.SetPtr(glib.Pointer(C.gst_event_new_qos(C.GstQOSType(tp), C.gdouble(proportion), C.GstClockTimeDiff(diff), C.GstClockTime(timestamp))))
	return r
}

//Get the type, proportion, diff and timestamp in the qos event. See gst_event_new_qos() for more information about the different QoS values.
//timestamp will be adjusted for any pad offsets of pads it was passing through.
//Parameters
//event
//The event to query
//Return
//type
//A pointer to store the QoS type in.
//proportion
//A pointer to store the proportion in.
//diff
//A pointer to store the diff in.
//timestamp
//A pointer to store the timestamp in.
func (e *Event) ParseQOS() (QOSType, float64, int64, ClockTime) {
	var tp C.GstQOSType
	var proportion C.gdouble
	var diff C.GstClockTimeDiff
	var timestamp C.GstClockTime
	C.gst_event_parse_qos(e.g(), &tp, &proportion, &diff, &timestamp)
	return QOSType(tp), float64(proportion), int64(diff), ClockTime(timestamp)
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

//Parses a seek event and stores the results in the given result locations.
//Parameters
//event
//a seek event
//Returns
//rate
//result location for the rate.
//format
//result location for the stream format.
//flags
//result location for the GstSeekFlags.
//start_type
//result location for the GstSeekType of the start position.
//start
//result location for the start position expressed in format .
//stop_type
//result location for the GstSeekType of the stop position.
//stop
//result location for the stop position expressed in format .
func (e *Event) ParseSeek() (float64, Format, SeekFlags, SeekType, int64, SeekType, int64) {
	var rate C.gdouble
	var format C.GstFormat
	var flags C.GstSeekFlags
	var start_type, stop_type C.GstSeekType
	var start, stop C.gint64
	C.gst_event_parse_seek(e.g(), &rate, &format, &flags, &start_type, &start, &stop_type, &stop)
	return float64(rate), Format(format), SeekFlags(flags), SeekType(start_type), int64(start), SeekType(stop_type), int64(stop)
}

//Create a new navigation event from the given description.
//Parameters
//structure
//description of the event. The event will take ownership of the structure.
//
//Returns
//a new GstEvent.
func NewNavigationEvent(structure *Structure) *Event {
	e := C.gst_event_new_navigation(structure.g())
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Create a new latency event.
//The event is sent upstream from the sinks and notifies elements that they should add an additional latency
//to the running time before synchronising against the clock.
//
//The latency is mostly used in live sinks and is always expressed in the time format.
//Parameters
//latency
//the new latency value
//Returns
//a new GstEvent.
func NewLatencyEvent(latency ClockTime) *Event {
	e := C.gst_event_new_latency(C.GstClockTime(latency))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Get the latency in the latency event.
//Parameters
//event
//The event to query
//Returns
//latency
//A pointer to store the latency in.
func (e *Event) ParseLatency() ClockTime {
	var latency C.GstClockTime
	C.gst_event_parse_latency(e.g(), &latency)
	return ClockTime(latency)
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

//Parse the step event.
//Parameters
//event
//The event to query
//Returns
//format
//a pointer to store the format in.
//amount
//a pointer to store the amount in.
//rate
//a pointer to store the rate in.
//flush
//a pointer to store the flush boolean in.
//intermediate
//a pointer to store the intermediate boolean in.
func (e *Event) ParseStep() (Format, uint64, float64, bool, bool) {
	var format C.GstFormat
	var amount C.guint64
	var rate C.gdouble
	var flush, intermediate C.gboolean
	C.gst_event_parse_step(e.g(), &format, &amount, &rate, &flush, &intermediate)
	return Format(format), uint64(amount), float64(rate), flush != 0, intermediate != 0
}

//Create a new sink-message event.
//The purpose of the sink-message event is to instruct a sink to post the message contained in the event synchronized with the stream.
//
//name is used to store multiple sticky events on one pad.
//Parameters
//name
//a name for the event
//msg
//the GstMessage to be posted.
//Returns
//a new GstEvent.
func NewSinkMessageEvent(name string, msg *Message) *Event {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	e := C.gst_event_new_sink_message(s, msg.g())
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Parse the sink-message event. Unref msg after usage.
//Parameters
//event
//The event to query
//Returns
//msg
//a pointer to store the GstMessage in.
func (e *Event) ParseSinkMessage() *Message {
	var msg *C.GstMessage
	C.gst_event_parse_sink_message(e.g(), &msg)
	r := new(Message)
	r.SetPtr(glib.Pointer(msg))
	return r
}

//Create a new reconfigure event.
//The purpose of the reconfigure event is to travel upstream and make elements renegotiate their caps or reconfigure their buffer pools.
//This is useful when changing properties on elements or changing the topology of the pipeline.
//Returns
//a new GstEvent.
func NewReconfigureEvent() *Event {
	e := C.gst_event_new_reconfigure()
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Create a new CAPS event for caps .
//The caps event can only travel downstream synchronized with the buffer flow and contains the format of the buffers that will follow after the event.
//Parameters
//caps
//a GstCaps.
//Returns
//the new CAPS event.
func NewCapsEvent(caps *Caps) *Event {
	e := C.gst_event_new_caps(caps.g())
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Get the caps from event . The caps remains valid as long as event remains valid.
//Parameters
//event
//The event to parse
//Returns
//caps
//A pointer to the caps.
func (e *Event) ParseCaps() *Caps {
	var caps *C.GstCaps
	C.gst_event_parse_caps(e.g(), &caps)
	r := new(Caps)
	r.SetPtr(glib.Pointer(caps))
	return r
}

//Generate a TOC event from the given toc .
//The purpose of the TOC event is to inform elements that some kind of the TOC was found.
//Parameters
//toc
//GstToc structure.
//updated
//whether toc was updated or not.
//Returns
//a new GstEvent.
func NewTOCEvent(toc *Toc, updated bool) *Event {
	e := C.gst_event_new_toc(toc.g(), gBoolean(updated))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Parse a TOC event and store the results in the given toc and updated locations.
//Parameters
//event
//a TOC event.
//Return
//toc
//pointer to GstToc structure.
//updated
//pointer to store TOC updated flag.
func (e *Event) ParseTOC() (*Toc, bool) {
	var toc *C.GstToc
	var updated C.gboolean
	C.gst_event_parse_toc(e.g(), &toc, &updated)
	r := new(Toc)
	r.SetPtr(glib.Pointer(toc))
	return r, updated != 0
}

//Generate a TOC select event with the given uid . T
//he purpose of the TOC select event is to start playback based on the TOC's entry with the given uid .
//Parameters
//uid
//UID in the TOC to start playback from.
//Returns
//a new GstEvent.
func NewTOCSelectEvent(uid string) *Event {
	s := (*C.gchar)(C.CString(uid))
	defer C.free(unsafe.Pointer(s))
	e := C.gst_event_new_toc_select(s)
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Parse a TOC select event and store the results in the given uid location.
//Parameters
//event
//a TOC select event.
//Returns
//uid
//storage for the selection UID.
func (e *Event) ParseTOCSelect() string {
	var uid *C.gchar
	defer C.free(unsafe.Pointer(uid))
	C.gst_event_parse_toc_select(e.g(), &uid)
	return C.GoString((*C.char)(uid))
}

//Create a new segment-done event.
//This event is sent by elements that finish playback of a segment as a result of a segment seek.
//Parameters
//format
//The format of the position being done
//position
//The position of the segment being done
//Returns
//a new GstEvent.
func NewSegmentDoneEvent(format Format, position int64) *Event {
	e := C.gst_event_new_segment_done(C.GstFormat(format), C.gint64(position))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Extracts the position and format from the segment done message.
//Parameters
//event
//A valid GstEvent of type GST_EVENT_SEGMENT_DONE.
//Returns
//format
//Result location for the format, or NULL.
//position
//Result location for the position, or NULL.
func (e *Event) ParseSegmentDone() (Format, int64) {
	var format C.GstFormat
	var position C.gint64
	C.gst_event_parse_segment_done(e.g(), &format, &position)
	return Format(format), int64(position)
}

//Creates a new event containing information specific to a particular protection system (uniquely identified by system_id ),
//by which that protection system can acquire key(s) to decrypt a protected stream.
//
//In order for a decryption element to decrypt media protected using a specific system,
//it first needs all the protection system specific information necessary to acquire the decryption key(s) for that stream.
//The functions defined here enable this information to be passed in events from elements
//that extract it (e.g., ISOBMFF demuxers, MPEG DASH demuxers) to protection decrypter elements that use it.
//
//Events containing protection system specific information are created using gst_event_new_protection,
//and they can be parsed by downstream elements using gst_event_parse_protection.
//
//In Common Encryption, protection system specific information may be located within ISOBMFF files,
//both in movie (moov) boxes and movie fragment (moof) boxes; it may also be contained in ContentProtection elements within MPEG DASH MPDs.
//The events created by gst_event_new_protection contain data identifying from which of these locations
//the encapsulated protection system specific information originated.
//This origin information is required as some protection systems use different encodings depending upon where the information originates.
//
//The events returned by gst_event_new_protection() are implemented in such a way as to ensure
//that the most recently-pushed protection info event of a particular origin and system_id will be stuck to the output pad of the sending element.
//Parameters
//system_id
//a string holding a UUID that uniquely identifies a protection system.
//data
//a GstBuffer holding protection system specific information. The reference count of the buffer will be incremented by one.
//origin
//a string indicating where the protection information carried in the event was extracted from.
//The allowed values of this string will depend upon the protection scheme.
//Returns
//a GST_EVENT_PROTECTION event, if successful; NULL if unsuccessful.
//Since: 1.6
func NewProtectionEvent(system_id string, data *Buffer, origin string) *Event {
	s1 := (*C.gchar)(C.CString(system_id))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(s2))
	e := C.gst_event_new_protection(s1, data.g(), s2)
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//Parses an event containing protection system specific information and stores the results in system_id , data and origin .
//The data stored in system_id , origin and data are valid until event is released.
//Parameters
//event
//a GST_EVENT_PROTECTION event.
//Returns
//system_id
//pointer to store the UUID string uniquely identifying a content protection system.
//data
//pointer to store a GstBuffer holding protection system specific information.
//origin
//pointer to store a value that indicates where the protection information carried by event was extracted from.
//Since: 1.6
func (e *Event) ParseProtection() (string, *Buffer, string) {
	var system_id, origin *C.gchar
	defer C.free(unsafe.Pointer(system_id))
	defer C.free(unsafe.Pointer(origin))
	var data *C.GstBuffer
	C.gst_event_parse_protection(e.g(), &system_id, &data, &origin)
	r := new(Buffer)
	r.SetPtr(glib.Pointer(data))
	return C.GoString((*C.char)(system_id)), r, C.GoString((*C.char)(origin))
}
