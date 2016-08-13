package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
static inline
GstPadProbeType CALL_MACRO_GST_PAD_PROBE_INFO_TYPE(GstPadProbeInfo *o) {
	return GST_PAD_PROBE_INFO_TYPE(o);
}
static inline
GstFlowReturn CALL_MACRO_GST_PAD_PROBE_INFO_FLOW_RETURN(GstPadProbeInfo *o) {
	return GST_PAD_PROBE_INFO_FLOW_RETURN(o);
}
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type PadLinkReturn C.GstPadLinkReturn

const (
	PAD_LINK_OK              = PadLinkReturn(C.GST_PAD_LINK_OK)
	PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
	PAD_LINK_WAS_LINKED      = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
	PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
	PAD_LINK_NOFORMAT        = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
	PAD_LINK_NOSCHED         = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
	PAD_LINK_REFUSED         = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

func (p PadLinkReturn) String() string {
	switch p {
	case PAD_LINK_OK:
		return "PAD_LINK_OK"
	case PAD_LINK_WRONG_HIERARCHY:
		return "PAD_LINK_WRONG_HIERARCHY"
	case PAD_LINK_WAS_LINKED:
		return "PAD_LINK_WAS_LINKED"
	case PAD_LINK_WRONG_DIRECTION:
		return "PAD_LINK_WRONG_DIRECTION"
	case PAD_LINK_NOFORMAT:
		return "PAD_LINK_NOFORMAT"
	case PAD_LINK_NOSCHED:
		return "PAD_LINK_NOSCHED"
	case PAD_LINK_REFUSED:
		return "PAD_LINK_REFUSED"
	}
	panic("Wrong value of PadLinkReturn variable")
}

//Gets a string representing the given pad-link return.
//Returns
//a static string with the name of the pad-link return.
//Since: 1.4
func (p PadLinkReturn) GetName() string {
	return C.GoString((*C.char)(C.gst_pad_link_get_name(C.GstPadLinkReturn(p))))
}

//#define GST_PAD_LINK_FAILED(ret) ((ret) < GST_PAD_LINK_OK)
//Macro to test if the given GstPadLinkReturn value indicates a failed link step.
func (p PadLinkReturn) LinkFailed() bool {
	return p < PAD_LINK_OK
}

//#define GST_PAD_LINK_SUCCESSFUL(ret) ((ret) >= GST_PAD_LINK_OK)
//Macro to test if the given GstPadLinkReturn value indicates a successful link step.
func (p PadLinkReturn) LinkSuccessful() bool {
	return p >= PAD_LINK_OK
}

type PadDirection C.GstPadDirection

const (
	PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
	PAD_SRC     = PadDirection(C.GST_PAD_SRC)
	PAD_SINK    = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) g() C.GstPadDirection {
	return C.GstPadDirection(p)
}

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type PadPresence C.GstPadPresence

const (
	PAD_ALWAYS    = PadPresence(C.GST_PAD_ALWAYS)
	PAD_SOMETIMES = PadPresence(C.GST_PAD_SOMETIMES)
	PAD_REQUEST   = PadPresence(C.GST_PAD_REQUEST)
)

func (p PadPresence) g() C.GstPadPresence {
	return C.GstPadPresence(p)
}

func (p PadPresence) String() string {
	switch p {
	case PAD_ALWAYS:
		return "PAD_ALWAYS"
	case PAD_SOMETIMES:
		return "PAD_SOMETIMES"
	case PAD_REQUEST:
		return "PAD_REQUEST"
	}
	panic("Wrong value of PadPresence variable")
}

type PadFlags C.GstPadFlags

const (
	PAD_FLAG_BLOCKED          = PadFlags(C.GST_PAD_FLAG_BLOCKED)
	PAD_FLAG_FLUSHING         = PadFlags(C.GST_PAD_FLAG_FLUSHING)
	PAD_FLAG_EOS              = PadFlags(C.GST_PAD_FLAG_EOS)
	PAD_FLAG_BLOCKING         = PadFlags(C.GST_PAD_FLAG_BLOCKING)
	PAD_FLAG_NEED_PARENT      = PadFlags(C.GST_PAD_FLAG_NEED_PARENT)
	PAD_FLAG_NEED_RECONFIGURE = PadFlags(C.GST_PAD_FLAG_NEED_RECONFIGURE)
	PAD_FLAG_PENDING_EVENTS   = PadFlags(C.GST_PAD_FLAG_PENDING_EVENTS)
	PAD_FLAG_FIXED_CAPS       = PadFlags(C.GST_PAD_FLAG_FIXED_CAPS)
	PAD_FLAG_PROXY_CAPS       = PadFlags(C.GST_PAD_FLAG_PROXY_CAPS)
	PAD_FLAG_PROXY_ALLOCATION = PadFlags(C.GST_PAD_FLAG_PROXY_ALLOCATION)
	PAD_FLAG_PROXY_SCHEDULING = PadFlags(C.GST_PAD_FLAG_PROXY_SCHEDULING)
	PAD_FLAG_ACCEPT_INTERSECT = PadFlags(C.GST_PAD_FLAG_ACCEPT_INTERSECT)
	PAD_FLAG_ACCEPT_TEMPLATE  = PadFlags(C.GST_PAD_FLAG_ACCEPT_TEMPLATE)
	/* padding */
	PAD_FLAG_LAST = PadFlags(C.GST_PAD_FLAG_LAST)
)

type PadLinkCheck C.GstPadLinkCheck

const (
	PAD_LINK_CHECK_NOTHING       = PadLinkCheck(C.GST_PAD_LINK_CHECK_NOTHING)
	PAD_LINK_CHECK_HIERARCHY     = PadLinkCheck(C.GST_PAD_LINK_CHECK_HIERARCHY)
	PAD_LINK_CHECK_TEMPLATE_CAPS = PadLinkCheck(C.GST_PAD_LINK_CHECK_TEMPLATE_CAPS)
	PAD_LINK_CHECK_CAPS          = PadLinkCheck(C.GST_PAD_LINK_CHECK_CAPS)
	PAD_LINK_CHECK_DEFAULT       = PadLinkCheck(C.GST_PAD_LINK_CHECK_DEFAULT)
)

type FlowReturn C.GstFlowReturn

const (
	FLOW_CUSTOM_SUCCESS_2 = FlowReturn(C.GST_FLOW_CUSTOM_SUCCESS_2)
	FLOW_CUSTOM_SUCCESS_1 = FlowReturn(C.GST_FLOW_CUSTOM_SUCCESS_1)
	FLOW_CUSTOM_SUCCESS   = FlowReturn(C.GST_FLOW_CUSTOM_SUCCESS)
	FLOW_OK               = FlowReturn(C.GST_FLOW_OK)
	FLOW_NOT_LINKED       = FlowReturn(C.GST_FLOW_NOT_LINKED)
	FLOW_FLUSHING         = FlowReturn(C.GST_FLOW_FLUSHING)
	FLOW_EOS              = FlowReturn(C.GST_FLOW_EOS)
	FLOW_NOT_NEGOTIATED   = FlowReturn(C.GST_FLOW_NOT_NEGOTIATED)
	FLOW_ERROR            = FlowReturn(C.GST_FLOW_ERROR)
	FLOW_NOT_SUPPORTED    = FlowReturn(C.GST_FLOW_NOT_SUPPORTED)
	FLOW_CUSTOM_ERROR     = FlowReturn(C.GST_FLOW_CUSTOM_ERROR)
	FLOW_CUSTOM_ERROR_1   = FlowReturn(C.GST_FLOW_CUSTOM_ERROR_1)
	FLOW_CUSTOM_ERROR_2   = FlowReturn(C.GST_FLOW_CUSTOM_ERROR_2)
)

//Gets a string representing the given flow return.
//Returns
//a static string with the name of the flow return.
func (p FlowReturn) GetName() string {
	return C.GoString((*C.char)(C.gst_flow_get_name(C.GstFlowReturn(p))))
}

//Get the unique quark for the given GstFlowReturn.
//Returns
//the quark associated with the flow return or 0 if an invalid return was specified.
func (p FlowReturn) ToQuark() glib.Quark {
	return glib.Quark(C.gst_flow_to_quark(C.GstFlowReturn(p)))
}

type PadMode C.GstPadMode

const (
	GST_PAD_MODE_NONE = PadMode(C.GST_PAD_MODE_NONE)
	GST_PAD_MODE_PUSH = PadMode(C.GST_PAD_MODE_PUSH)
	GST_PAD_MODE_PULL = PadMode(C.GST_PAD_MODE_PULL)
)

//Return the name of a pad mode, for use in debug messages mostly.
//Returns
//short mnemonic for pad mode mode
func (p PadMode) GetName() string {
	return C.GoString((*C.char)(C.gst_pad_mode_get_name(C.GstPadMode(p))))
}

type PadProbeType C.GstPadProbeType

const (
	PAD_PROBE_TYPE_INVALID = PadProbeType(C.GST_PAD_PROBE_TYPE_INVALID)
	/* flags to control blocking */
	PAD_PROBE_TYPE_IDLE  = PadProbeType(C.GST_PAD_PROBE_TYPE_IDLE)
	PAD_PROBE_TYPE_BLOCK = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK)
	/* flags to select datatypes */
	PAD_PROBE_TYPE_BUFFER           = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER)
	PAD_PROBE_TYPE_BUFFER_LIST      = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER_LIST)
	PAD_PROBE_TYPE_EVENT_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM)
	PAD_PROBE_TYPE_EVENT_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_FLUSH      = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_FLUSH)
	PAD_PROBE_TYPE_QUERY_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_DOWNSTREAM)
	PAD_PROBE_TYPE_QUERY_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_UPSTREAM)
	/* flags to select scheduling mode */
	PAD_PROBE_TYPE_PUSH = PadProbeType(C.GST_PAD_PROBE_TYPE_PUSH)
	PAD_PROBE_TYPE_PULL = PadProbeType(C.GST_PAD_PROBE_TYPE_PULL)
	/* flag combinations */
	PAD_PROBE_TYPE_BLOCKING         = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCKING)
	PAD_PROBE_TYPE_DATA_DOWNSTREAM  = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_DOWNSTREAM)
	PAD_PROBE_TYPE_DATA_UPSTREAM    = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_UPSTREAM)
	PAD_PROBE_TYPE_DATA_BOTH        = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_BOTH)
	PAD_PROBE_TYPE_BLOCK_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_DOWNSTREAM)
	PAD_PROBE_TYPE_BLOCK_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_BOTH       = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_BOTH)
	PAD_PROBE_TYPE_QUERY_BOTH       = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_BOTH)
	PAD_PROBE_TYPE_ALL_BOTH         = PadProbeType(C.GST_PAD_PROBE_TYPE_ALL_BOTH)
	PAD_PROBE_TYPE_SCHEDULING       = PadProbeType(C.GST_PAD_PROBE_TYPE_SCHEDULING)
)

type PadProbeReturn C.GstPadProbeReturn

const (
	GST_PAD_PROBE_DROP    = PadProbeReturn(C.GST_PAD_PROBE_DROP)
	GST_PAD_PROBE_OK      = PadProbeReturn(C.GST_PAD_PROBE_OK)
	GST_PAD_PROBE_REMOVE  = PadProbeReturn(C.GST_PAD_PROBE_REMOVE)
	GST_PAD_PROBE_PASS    = PadProbeReturn(C.GST_PAD_PROBE_PASS)
	GST_PAD_PROBE_HANDLED = PadProbeReturn(C.GST_PAD_PROBE_HANDLED)
)

type PadProbeInfo C.GstPadProbeInfo

func (p *PadProbeInfo) g() *C.GstPadProbeInfo {
	return (*C.GstPadProbeInfo)(p)
}

func (p *PadProbeInfo) AsPadProbeInfo() *PadProbeInfo {
	return p
}

func (p *PadProbeInfo) Type() PadProbeType {
	return PadProbeType(C.CALL_MACRO_GST_PAD_PROBE_INFO_TYPE(p.g()))
}

func (p *PadProbeInfo) Id() uint64 {
	return uint64(p.g().id)
}

func (p *PadProbeInfo) Data() unsafe.Pointer {
	return unsafe.Pointer(p.g().data)
}

//Returns
//The GstBuffer from the probe.
func (p *PadProbeInfo) Buffer() *Buffer {
	b := p.Data()
	if b == nil {
		return nil
	}
	r := new(Buffer)
	r.SetPtr(glib.Pointer(b))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstBufferList from the probe.
func (p *PadProbeInfo) GetBufferList() *BufferList {
	c := C.gst_pad_probe_info_get_buffer_list(p.g())
	if c == nil {
		return nil
	}
	r := new(BufferList)
	r.SetPtr(glib.Pointer(c))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstEvent from the probe.
func (p *PadProbeInfo) GetEvent() *Event {
	c := C.gst_pad_probe_info_get_event(p.g())
	if c == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(c))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstQuery from the probe.
func (p *PadProbeInfo) GetQuery() *Query {
	c := C.gst_pad_probe_info_get_query(p.g())
	if c == nil {
		return nil
	}
	r := new(Query)
	r.SetPtr(glib.Pointer(c))
	return r
}

func (p *PadProbeInfo) Offset() uint64 {
	return uint64(p.g().offset)
}

func (p *PadProbeInfo) Size() uint {
	return uint(p.g().size)
}

func (p *PadProbeInfo) FlowRet() FlowReturn {
	return FlowReturn(C.CALL_MACRO_GST_PAD_PROBE_INFO_FLOW_RETURN(p.g()))
}

type Pad struct {
	GstObj
}

//LiD: add GstPad Type() interface
func (p *Pad) Type() glib.Type {
	return glib.TypeFromName("GstPad")
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

//Store the sticky event on pad
//Parameters
//event
//a GstEvent
//Returns
//GST_FLOW_OK on success, GST_FLOW_FLUSHING when the pad was flushing or GST_FLOW_EOS when the pad was EOS.
//Since: 1.2
func (p *Pad) StoreStickyEvent(event *Event) FlowReturn {
	return FlowReturn(C.gst_pad_store_sticky_event(p.g(), event.g()))
}

//Gets the direction of the pad. The direction of the pad is decided at construction time so this function does not take the LOCK.
//Returns
//the GstPadDirection of the pad.
//MT safe.
func (p *Pad) GetDirection() PadDirection {
	return PadDirection(C.gst_pad_get_direction(p.g()))
}

//Gets the parent of pad , cast to a GstElement. If a pad has no parent or its parent is not an element, return NULL.
//Returns
//the parent of the pad. The caller has a reference on the parent, so unref when you're finished with it.
//MT safe.
func (p *Pad) GetParentElement() *Element {
	ce := C.gst_pad_get_parent_element(p.g())
	if ce == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(glib.Pointer(ce))
	return e
}

//Gets the template for pad .
//Returns
//the GstPadTemplate from which this pad was instantiated, or NULL if this pad has no template. Unref after usage.
func (p *Pad) GetPadTemplate() *PadTemplate {
	ce := C.gst_pad_get_pad_template(p.g())
	if ce == nil {
		return nil
	}
	e := new(PadTemplate)
	e.SetPtr(glib.Pointer(ce))
	return e
}

//Links the source pad and the sink pad.
//Parameters
//sinkpad
//the sink GstPad to link.
//Returns
//A result code indicating if the connection worked or what went wrong.
//MT Safe.
func (p *Pad) Link(sinkpad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sinkpad.g()))
}

//Links the source pad and the sink pad.
//This variant of gst_pad_link provides a more granular control on the checks being done when linking.
//While providing some considerable speedups the caller of this method must be aware that wrong usage of those flags can cause severe issues.
//Refer to the documentation of GstPadLinkCheck for more information.
//MT Safe.
//Parameters
//sinkpad
//the sink GstPad to link.
//flags
//the checks to validate when linking
//Returns
//A result code indicating if the connection worked or what went wrong.
func (p *Pad) LinkFull(sinkpad *Pad, flags PadLinkCheck) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link_full(p.g(), sinkpad.g(), C.GstPadLinkCheck(flags)))
}

//Unlinks the source pad from the sink pad. Will emit the “unlinked” signal on both pads.
//Parameters
//sinkpad
//the sink GstPad to unlink.
//Returns
//TRUE if the pads were unlinked. This function returns FALSE if the pads were not linked together.
//MT safe.
func (p *Pad) Unlink(sinkpad *Pad) bool {
	return C.gst_pad_unlink(p.g(), sinkpad.g()) != 0
}

//Checks if a pad is linked to another pad or not.
//Returns
//TRUE if the pad is linked, FALSE otherwise.
//MT safe.
func (p *Pad) IsLinked() bool {
	return C.gst_pad_is_linked(p.g()) != 0
}

//Checks if the source pad and the sink pad are compatible so they can be linked.
//Parameters
//sinkpad
//the sink GstPad.
//Returns
//TRUE if the pads can be linked.
func (p *Pad) CanLink(sinkpad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sinkpad.g()) != 0
}

//Gets the capabilities of the allowed media types that can flow through pad and its peer.
//The allowed capabilities is calculated as the intersection of the results of calling gst_pad_query_caps() on pad and its peer.
//The caller owns a reference on the resulting caps.
//Returns
//the allowed GstCaps of the pad link. Unref the caps when you no longer need it. This function returns NULL when pad has no peer.
//MT safe.
func (p *Pad) GetAllowedCaps() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_pad_get_allowed_caps(p.g())))
	return r
}

//Gets the capabilities currently configured on pad with the last GST_EVENT_CAPS event.
//Returns
//the current caps of the pad with incremented ref-count or NULL when pad has no caps. Unref after usage.
func (p *Pad) GetCurrentCaps() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_pad_get_current_caps(p.g())))
	return r
}

//Gets the capabilities for pad 's template.
//Returns
//the GstCaps of this pad template. Unref after usage.
func (p *Pad) GetPadTemplateCaps() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_pad_get_pad_template_caps(p.g())))
	return r
}

//Gets the peer of pad . This function refs the peer pad so you need to unref it after use.
//Returns
//the peer GstPad. Unref after usage.
//MT safe.
func (p *Pad) GetPeer() *Pad {
	ce := C.gst_pad_get_peer(p.g())
	if ce == nil {
		return nil
	}
	e := new(Pad)
	e.SetPtr(glib.Pointer(ce))
	return e
}

//A helper function you can use that sets the FIXED_CAPS flag
//This way the default CAPS query will always return the negotiated caps or in case the pad is not negotiated, the padtemplate caps.
//The negotiated caps are the caps of the last CAPS event that passed on the pad. Use this function on a pad that,
//once it negotiated to a CAPS, cannot be renegotiated to something else.
func (p *Pad) UseFixedCaps() {
	C.gst_pad_use_fixed_caps(p.g())
}

//Check if pad has caps set on it with a GST_EVENT_CAPS event.
//Returns
//TRUE when pad has caps associated with it.
func (p *Pad) HasCurrentCaps() bool {
	return C.gst_pad_has_current_caps(p.g()) != 0
}

//Returns a new reference of the sticky event of type event_type from the event.
//Parameters
//event_type
//the GstEventType that should be retrieved.
//idx
//the index of the event
//Returns
//a GstEvent of type event_type or NULL when no event of event_type was on pad . Unref after usage.
func (p *Pad) GetStickyEvent(event_type EventType, idx uint) *Event {
	e := C.gst_pad_get_sticky_event(p.g(), C.GstEventType(event_type), C.guint(idx))
	if e == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(e))
	return r
}

//gboolean
//(*GstPadStickyEventsForeachFunction) (GstPad *pad,
//GstEvent **event,
//gpointer user_data);
//Callback used by gst_pad_sticky_events_foreach().
//When this function returns TRUE, the next event will be returned. When FALSE is returned, gst_pad_sticky_events_foreach() will return.
//When event is set to NULL, the item will be removed from the list of sticky events.
//event can be replaced by assigning a new reference to it. This function is responsible for unreffing the old event when removing or modifying.
//Parameters
//pad
//the GstPad.
//event
//a sticky GstEvent.
//user_data
//the gpointer to optional user data.
//Returns
//TRUE if the iteration should continue

//void
//gst_pad_sticky_events_foreach (GstPad *pad,
//GstPadStickyEventsForeachFunction foreach_func,
//gpointer user_data);
//Iterates all sticky events on pad and calls foreach_func for every event. If foreach_func returns FALSE the iteration is immediately stopped.
//Parameters
//foreach_func
//the GstPadStickyEventsForeachFunction that should be called for every event.
//user_data
//the optional user data.
//[closure]
//func (p *Pad) StickyEventsForeach(foreach_func func(userdata interface{})) {
//	C.gst_pad_sticky_events_foreach(p.g(), foreach_func)
//}

//Query if a pad is active
//Parameters
//pad
//the GstPad to query
//Returns
//TRUE if the pad is active.
//MT safe.
func (p *Pad) IsActive() bool {
	return C.gst_pad_is_active(p.g()) != 0
}

//Gets the GstFlowReturn return from the last data passed by this pad.
//Parameters
//pad
//the GstPad
//Since: 1.4
func (p *Pad) GetLastFlowReturn(sinkpad *Pad, flags PadLinkCheck) FlowReturn {
	return FlowReturn(C.gst_pad_get_last_flow_return(p.g()))
}

//Remove the probe with id from pad .
//MT safe.
//Parameters
//pad
//the GstPad with the probe
//id
//the probe id to remove
func (p *Pad) RemoveProbe(event_type EventType, id uint64) {
	C.gst_pad_remove_probe(p.g(), C.gulong(id))
}

//Checks if the pad is blocked or not. This function returns the last requested state of the pad. It is not certain that the pad is actually blocking at this point (see gst_pad_is_blocking()).
//Parameters
//pad
//the GstPad to query
//Returns
//TRUE if the pad is blocked.
//MT safe.
func (p *Pad) IsBlocked() bool {
	return C.gst_pad_is_blocked(p.g()) != 0
}

//Checks if the pad is blocking or not. This is a guaranteed state of whether the pad is actually blocking on a GstBuffer or a GstEvent.
//Parameters
//pad
//the GstPad to query
//Returns
//TRUE if the pad is blocking.
//MT safe.
func (p *Pad) IsBlocking() bool {
	return C.gst_pad_is_blocking(p.g()) != 0
}

//Get the offset applied to the running time of pad . pad has to be a source pad.
//Parameters
//pad
//a GstPad
//Returns
//the offset.
func (p *Pad) GetOffset() int64 {
	return int64(C.gst_pad_get_offset(p.g()))
}

//Set the offset that will be applied to the running time of pad .
//Parameters
//pad
//a GstPad
//offset
//the offset
func (p *Pad) SetOffset(offset int64) {
	C.gst_pad_set_offset(p.g(), C.gint64(offset))
}

//Creates a new pad with the given name in the given direction. If name is NULL, a guaranteed unique name (across all pads) will be assigned.
//This function makes a copy of the name so you can safely free the name.
//Parameters
//name
//the name of the new pad.
//direction
//the GstPadDirection of the pad.
//Returns
//a new GstPad, or NULL in case of an error.
//MT safe.
func NewPad(name string, direction PadDirection) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(Pad)
	p.SetPtr(glib.Pointer(C.gst_pad_new(s, C.GstPadDirection(direction))))
	return p
}

//Creates a new pad with the given name from the given template. If name is NULL, a guaranteed unique name (across all pads) will be assigned.
//This function makes a copy of the name so you can safely free the name.
//Parameters
//templ
//the pad template to use
//name
//the name of the pad.
//Returns
//a new GstPad, or NULL in case of an error.
func NewPadFromTemplate(templ *PadTemplate, name string) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(Pad)
	p.SetPtr(glib.Pointer(C.gst_pad_new_from_template(templ.g(), s)))
	return p
}

//Creates a new pad with the given name from the given static template. If name is NULL, a guaranteed unique name (across all pads) will be assigned.
//This function makes a copy of the name so you can safely free the name.
//Parameters
//templ
//the GstStaticPadTemplate to use
//name
//the name of the pad
//Returns
//a new GstPad, or NULL in case of an error.
func NewPadFromStaticTemplate(templ *StaticPadTemplate, name string) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(Pad)
	p.SetPtr(glib.Pointer(C.gst_pad_new_from_static_template(templ.g(), s)))
	return p
}

//Calls gst_pad_query_caps() for all internally linked pads of pad and returns the intersection of the results.
//This function is useful as a default caps query function for an element that can handle any stream format, but requires all its pads to have the same caps.
//Two such elements are tee and adder.
//Parameters
//pad
//a GstPad to proxy.
//query
//a CAPS GstQuery.
//Returns
//TRUE if query could be executed
func (p *Pad) ProxyQueryCaps(query *Query) bool {
	return C.gst_pad_proxy_query_caps(p.g(), query.g()) != 0
}

//Checks if all internally linked pads of pad accepts the caps in query and returns the intersection of the results.
//This function is useful as a default accept caps query function for an element that can handle any stream format,
//but requires caps that are acceptable for all opposite pads.
//Parameters
//pad
//a GstPad to proxy.
//query
//an ACCEPT_CAPS GstQuery.
//Returns
//TRUE if query could be executed
func (p *Pad) ProxyQueryAcceptCaps(query *Query) bool {
	return C.gst_pad_proxy_query_accept_caps(p.g(), query.g()) != 0
}

//Check the GST_PAD_FLAG_NEED_RECONFIGURE flag on pad and return TRUE if the flag was set.
//Parameters
//pad
//the GstPad to check
//Returns
//TRUE is the GST_PAD_FLAG_NEED_RECONFIGURE flag is set on pad .
func (p *Pad) NeedsReconfigure() bool {
	return C.gst_pad_needs_reconfigure(p.g()) != 0
}

//Check and clear the GST_PAD_FLAG_NEED_RECONFIGURE flag on pad and return TRUE if the flag was set.
//Parameters
//pad
//the GstPad to check
//Returns
//TRUE is the GST_PAD_FLAG_NEED_RECONFIGURE flag was set on pad .
func (p *Pad) CheckReconfigure() bool {
	return C.gst_pad_check_reconfigure(p.g()) != 0
}

//Mark a pad for needing reconfiguration. The next call to gst_pad_check_reconfigure() will return TRUE after this call.
//Parameters
//pad
//the GstPad to mark
func (p *Pad) MarkReconfigure() {
	C.gst_pad_mark_reconfigure(p.g())
}

//Pushes a buffer to the peer of pad .
//This function will call installed block probes before triggering any installed data probes.
//The function proceeds calling gst_pad_chain() on the peer pad and returns the value from that function.
//If pad has no peer, GST_FLOW_NOT_LINKED will be returned.
//In all cases, success or failure, the caller loses its reference to buffer after calling this function.
//Parameters
//pad
//a source GstPad, returns GST_FLOW_ERROR if not.
//buffer
//the GstBuffer to push returns GST_FLOW_ERROR if not.
//Returns
//a GstFlowReturn from the peer pad.
//MT safe.
func (p *Pad) Push(buffer *Buffer) FlowReturn {
	return FlowReturn(C.gst_pad_push(p.g(), buffer.g()))
}

//Sends the event to the peer of the given pad. This function is mainly used by elements to send events to their peer elements.
//This function takes ownership of the provided event so you should gst_event_ref() it if you want to reuse the event after this call.
//Parameters
//pad
//a GstPad to push the event to.
//event
//the GstEvent to send to the pad.
//Returns
//TRUE if the event was handled.
//MT safe.
func (p *Pad) PushEvent(event *Event) bool {
	return C.gst_pad_push_event(p.g(), event.g()) != 0
}

//Pushes a buffer list to the peer of pad .
//This function will call installed block probes before triggering any installed data probes.
//The function proceeds calling the chain function on the peer pad and returns the value from that function. If pad has no peer, GST_FLOW_NOT_LINKED will be returned.
//If the peer pad does not have any installed chainlist function every group buffer of the list will be merged into a normal GstBuffer and chained via gst_pad_chain().
//In all cases, success or failure, the caller loses its reference to list after calling this function.
//Parameters
//pad
//a source GstPad, returns GST_FLOW_ERROR if not.
//list
//the GstBufferList to push returns GST_FLOW_ERROR if not.
//Returns
//a GstFlowReturn from the peer pad.
//MT safe.
func (p *Pad) PushList(list *BufferList) FlowReturn {
	return FlowReturn(C.gst_pad_push_list(p.g(), list.g()))
}

//Pulls a buffer from the peer pad or fills up a provided buffer.
//This function will first trigger the pad block signal if it was installed.
//When pad is not linked GST_FLOW_NOT_LINKED is returned else this function returns the result of gst_pad_get_range() on the peer pad.
//See gst_pad_get_range() for a list of return values and for the semantics of the arguments of this function.
//
//If buffer points to a variable holding NULL, a valid new GstBuffer will be placed in buffer when this function returns GST_FLOW_OK.
//The new buffer must be freed with gst_buffer_unref() after usage. When this function returns any other result value, buffer will still point to NULL.
//
//When buffer points to a variable that points to a valid GstBuffer, the buffer will be filled with the result data when this function returns GST_FLOW_OK.
//When this function returns any other result value, buffer will be unchanged.
//If the provided buffer is larger than size , only size bytes will be filled in the result buffer and its size will be updated accordingly.
//
//Note that less than size bytes can be returned in buffer when, for example, an EOS condition is near or when buffer is not large enough to hold size bytes.
//The caller should check the result buffer size to get the result size.
//Parameters
//pad
//a sink GstPad, returns GST_FLOW_ERROR if not.
//offset
//The start offset of the buffer
//size
//The length of the buffer
//Returns
//a GstFlowReturn from the peer pad.
//buffer
//a pointer to hold the GstBuffer, returns GST_FLOW_ERROR if NULL.
//MT safe.
func (p *Pad) PullRange(offset uint64, size uint) (FlowReturn, *Buffer) {
	var b *C.GstBuffer
	ret := C.gst_pad_pull_range(p.g(), C.guint64(offset), C.guint(size), &b)
	r := new(Buffer)
	r.SetPtr(glib.Pointer(b))
	return FlowReturn(ret), r
}

//Activates or deactivates the given pad in mode via dispatching to the pad's activatemodefunc. For use from within pad activation functions only.
//If you don't know what this is, you probably don't want to call it.
//Parameters
//pad
//the GstPad to activate or deactivate.
//mode
//the requested activation mode
//active
//whether or not the pad should be active.
//Returns
//TRUE if the operation was successful.
//MT safe.
func (p *Pad) ActiveMode(mode PadMode, active bool) bool {
	return C.gst_pad_activate_mode(p.g(), C.GstPadMode(mode), gBoolean(active)) != 0
}

//Sends the event to the pad. This function can be used by applications to send events in the pipeline.
//
//If pad is a source pad, event should be an upstream event.
//If pad is a sink pad, event should be a downstream event. For example, you would not send a GST_EVENT_EOS on a src pad; EOS events only propagate downstream.
//Furthermore, some downstream events have to be serialized with data flow, like EOS, while some can travel out-of-band, like GST_EVENT_FLUSH_START.
//If the event needs to be serialized with data flow, this function will take the pad's stream lock while calling its event function.
//
//To find out whether an event type is upstream, downstream, or downstream and serialized,
//see GstEventTypeFlags, gst_event_type_get_flags(), GST_EVENT_IS_UPSTREAM, GST_EVENT_IS_DOWNSTREAM, and GST_EVENT_IS_SERIALIZED.
//Note that in practice that an application or plugin doesn't need to bother itself with this information; the core handles all necessary locks and checks.
//
//This function takes ownership of the provided event so you should gst_event_ref() it if you want to reuse the event after this call.
//Parameters
//pad
//a GstPad to send the event to.
//event
//the GstEvent to send to the pad.
//Returns
//TRUE if the event was handled.
func (p *Pad) SendEvent(event *Event) bool {
	return C.gst_pad_send_event(p.g(), event.g()) != 0
}

//Invokes the default event handler for the given pad.
//
//The EOS event will pause the task associated with pad before it is forwarded to all internally linked pads,
//
//The event is sent to all pads internally linked to pad . This function takes ownership of event .
//Parameters
//pad
//a GstPad to call the default event handler on.
//parent
//the parent of pad or NULL.
//event
//the GstEvent to handle.
//Returns
//TRUE if the event was sent successfully.
func (p *Pad) EventDefault(parent *GstObj, event *Event) bool {
	return C.gst_pad_event_default(p.g(), parent.g(), event.g()) != 0
}

//Dispatches a query to a pad. The query should have been allocated by the caller via one of the type-specific allocation functions.
//The element that the pad belongs to is responsible for filling the query with an appropriate response,
//which should then be parsed with a type-specific query parsing function.
//
//Again, the caller is responsible for both the allocation and deallocation of the query structure.
//
//Please also note that some queries might need a running pipeline to work.
//Parameters
//pad
//a GstPad to invoke the default query on.
//query
//the GstQuery to perform.
//Returns
//TRUE if the query could be performed.
func (p *Pad) Query(query *Query) bool {
	return C.gst_pad_query(p.g(), query.g()) != 0
}

//Performs gst_pad_query() on the peer of pad .
//
//The caller is responsible for both the allocation and deallocation of the query structure.
//Parameters
//pad
//a GstPad to invoke the peer query on.
//query
//the GstQuery to perform.
//Returns
//TRUE if the query could be performed. This function returns FALSE if pad has no peer.
func (p *Pad) PeerQuery(query *Query) bool {
	return C.gst_pad_peer_query(p.g(), query.g()) != 0
}

//Invokes the default query handler for the given pad. The query is sent to all pads internally linked to pad .
//Note that if there are many possible sink pads that are internally linked to pad , only one will be sent the query.
//Multi-sinkpad elements should implement custom query handlers.
//Parameters
//pad
//a GstPad to call the default query handler on.
//parent
//the parent of pad or NULL.
//query
//the GstQuery to handle.
//Returns
//TRUE if the query was performed successfully.
func (p *Pad) QueryDefault(parent *GstObj, query *Query) bool {
	return C.gst_pad_query_default(p.g(), parent.g(), query.g()) != 0
}

//Queries a pad for the stream position.
//Parameters
//pad
//a GstPad to invoke the position query on.
//format
//the GstFormat requested
//Returns
//TRUE if the query could be performed.
//cur
//A location in which to store the current position, or NULL.
func (p *Pad) QueryPosition(format Format) (bool, int64) {
	var cur C.gint64
	ret := C.gst_pad_query_position(p.g(), C.GstFormat(format), &cur)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(cur)
	}
}

//Queries a pad for the total stream duration.
//Parameters
//pad
//a GstPad to invoke the duration query on.
//format
//the GstFormat requested
//Returns
//TRUE if the query could be performed.
//duration
//a location in which to store the total duration, or NULL.
func (p *Pad) QueryDuration(format Format) (bool, int64) {
	var duration C.gint64
	ret := C.gst_pad_query_duration(p.g(), C.GstFormat(format), &duration)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(duration)
	}
}

//Queries a pad to convert src_val in src_format to dest_format .
//Parameters
//pad
//a GstPad to invoke the convert query on.
//src_format
//a GstFormat to convert from.
//src_val
//a value to convert.
//dest_format
//the GstFormat to convert to.
//Returns
//TRUE if the query could be performed.
//dest_val
//a pointer to the result.
func (p *Pad) QueryConvert(src_format Format, src_val int64, dest_format Format) (bool, int64) {
	var dest_val C.gint64
	ret := C.gst_pad_query_convert(p.g(), C.GstFormat(src_format), C.gint64(src_val), C.GstFormat(dest_format), &dest_val)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(dest_val)
	}
}

//Check if the given pad accepts the caps.
//Parameters
//pad
//a GstPad to check
//caps
//a GstCaps to check on the pad
//Returns
//TRUE if the pad can accept the caps.
func (p *Pad) QueryAcceptCaps(caps *Caps) bool {
	return C.gst_pad_query_accept_caps(p.g(), caps.g()) != 0
}

//Gets the capabilities this pad can produce or consume.
//Note that this method doesn't necessarily return the caps set by sending a gst_event_new_caps() - use gst_pad_get_current_caps() for that instead.
//gst_pad_query_caps returns all possible caps a pad can operate with, using the pad's CAPS query function,
//If the query fails, this function will return filter , if not NULL, otherwise ANY.
//
//When called on sinkpads filter contains the caps that upstream could produce in the order preferred by upstream.
//When called on srcpads filter contains the caps accepted by downstream in the preferred order.
//filter might be NULL but if it is not NULL the returned caps will be a subset of filter .
//
//Note that this function does not return writable GstCaps, use gst_caps_make_writable() before modifying the caps.
//Parameters
//pad
//a GstPad to get the capabilities of.
//filter
//suggested GstCaps, or NULL.
//Returns
//the caps of the pad with incremented ref-count.
func (p *Pad) QueryCaps(filter *Caps) *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_pad_query_caps(p.g(), filter.g())))
	return r
}

//Queries the peer of a given sink pad for the stream position.
//Parameters
//pad
//a GstPad on whose peer to invoke the position query on. Must be a sink pad.
//format
//the GstFormat requested
//Returns
//TRUE if the query could be performed.
//cur
//a location in which to store the current position, or NULL.
func (p *Pad) PeerQueryPosition(format Format) (bool, int64) {
	var cur C.gint64
	ret := C.gst_pad_peer_query_position(p.g(), C.GstFormat(format), &cur)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(cur)
	}
}

//Queries the peer pad of a given sink pad for the total stream duration.
//Parameters
//pad
//a GstPad on whose peer pad to invoke the duration query on. Must be a sink pad.
//format
//the GstFormat requested
//Returns
//TRUE if the query could be performed.
//duration
//a location in which to store the total duration, or NULL.
func (p *Pad) PeerQueryDuration(format Format) (bool, int64) {
	var duration C.gint64
	ret := C.gst_pad_peer_query_duration(p.g(), C.GstFormat(format), &duration)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(duration)
	}
}

//Queries the peer pad of a given sink pad to convert src_val in src_format to dest_format .
//Parameters
//pad
//a GstPad, on whose peer pad to invoke the convert query on. Must be a sink pad.
//src_format
//a GstFormat to convert from.
//src_val
//a value to convert.
//dest_format
//the GstFormat to convert to.
//Returns
//TRUE if the query could be performed.
//dest_val
//a pointer to the result.
func (p *Pad) PeerQueryConvert(src_format Format, src_val int64, dest_format Format) (bool, int64) {
	var dest_val C.gint64
	ret := C.gst_pad_peer_query_convert(p.g(), C.GstFormat(src_format), C.gint64(src_val), C.GstFormat(dest_format), &dest_val)
	if ret == 0 {
		return false, 0
	} else {
		return true, int64(dest_val)
	}
}

//Check if the peer of pad accepts caps . If pad has no peer, this function returns TRUE.
//Parameters
//pad
//a GstPad to check the peer of
//caps
//a GstCaps to check on the pad
//Returns
//TRUE if the peer of pad can accept the caps or pad has no peer.
func (p *Pad) PeerQueryAcceptCaps(caps *Caps) bool {
	return C.gst_pad_peer_query_accept_caps(p.g(), caps.g()) != 0
}

//Gets the capabilities of the peer connected to this pad. Similar to gst_pad_query_caps().
//
//When called on srcpads filter contains the caps that upstream could produce in the order preferred by upstream.
//When called on sinkpads filter contains the caps accepted by downstream in the preferred order.
//filter might be NULL but if it is not NULL the returned caps will be a subset of filter .
//Parameters
//pad
//a GstPad to get the capabilities of.
//filter
//a GstCaps filter, or NULL.
//Returns
//the caps of the peer pad with incremented ref-count. When there is no peer pad, this function returns filter or, when filter is NULL, ANY caps.
func (p *Pad) PeerQueryCaps(filter *Caps) *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_pad_peer_query_caps(p.g(), filter.g())))
	return r
}

//Set the given private data gpointer on the pad. This function can only be used by the element that owns the pad. No locking is performed in this function.
//Parameters
//pad
//the GstPad to set the private data of.
//priv
//The private data to attach to the pad.
func (p *Pad) SetElementPrivate(priv glib.Pointer) {
	C.gst_pad_set_element_private(p.g(), C.gpointer(unsafe.Pointer(priv)))
}

//Gets the private data of a pad. No locking is performed in this function.
//Parameters
//pad
//the GstPad to get the private data of.
//Returns
//a gpointer to the private data.
func (p *Pad) GetElementPrivate() glib.Pointer {
	return glib.Pointer(C.gst_pad_get_element_private(p.g()))
}

//Creates a stream-id for the source GstPad pad by combining the upstream information with the optional stream_id of the stream of pad .
//pad must have a parent GstElement and which must have zero or one sinkpad.
//stream_id can only be NULL if the parent element of pad has only a single source pad.
//
//This function generates an unique stream-id by getting the upstream stream-start event stream ID and appending stream_id to it.
//If the element has no sinkpad it will generate an upstream stream-id by doing an URI query on the element and in the worst case just uses a random number.
//Source elements that don't implement the URI handler interface should ideally generate a unique, deterministic stream-id manually instead.
//
//Since stream IDs are sorted alphabetically, any numbers in the stream ID should be printed with a fixed number of characters, preceded by 0's,
//such as by using the format %03u instead of %u.
//Parameters
//pad
//A source GstPad
//parent
//Parent GstElement of pad
//stream_id
//The stream-id.
//Returns
//A stream-id for pad . g_free() after usage.
func (p *Pad) CreateStreamId(parent *Element, stream_id string) string {
	s := (*C.gchar)(C.CString(stream_id))
	defer C.free(unsafe.Pointer(s))
	return C.GoString((*C.char)(C.gst_pad_create_stream_id(p.g(), parent.g(), s)))
}

//Returns the current stream-id for the pad , or NULL if none has been set yet, i.e. the pad has not received a stream-start event yet.
//
//This is a convenience wrapper around gst_pad_get_sticky_event() and gst_event_parse_stream_start().
//
//The returned stream-id string should be treated as an opaque string, its contents should not be interpreted.
//Parameters
//pad
//A source GstPad
//Returns
//a newly-allocated copy of the stream-id for pad , or NULL. g_free() the returned string when no longer needed.
//Since: 1.2
func (p *Pad) GetStreamId(parent *Element, stream_id string) string {
	return C.GoString((*C.char)(C.gst_pad_get_stream_id(p.g())))
}

//Chain a buffer to pad .
//
//The function returns GST_FLOW_FLUSHING if the pad was flushing.
//
//If the buffer type is not acceptable for pad (as negotiated with a preceding GST_EVENT_CAPS event), this function returns GST_FLOW_NOT_NEGOTIATED.
//
//The function proceeds calling the chain function installed on pad (see gst_pad_set_chain_function())
//and the return value of that function is returned to the caller.
//GST_FLOW_NOT_SUPPORTED is returned if pad has no chain function.
//
//In all cases, success or failure, the caller loses its reference to buffer after calling this function.
//Parameters
//pad
//a sink GstPad, returns GST_FLOW_ERROR if not.
//buffer
//the GstBuffer to send, return GST_FLOW_ERROR if not.
//Returns
//a GstFlowReturn from the pad.
//MT safe.
func (p *Pad) Chain(buffer *Buffer) FlowReturn {
	return FlowReturn(C.gst_pad_chain(p.g(), buffer.g()))
}

//Chain a bufferlist to pad .
//
//The function returns GST_FLOW_FLUSHING if the pad was flushing.
//
//If pad was not negotiated properly with a CAPS event, this function returns GST_FLOW_NOT_NEGOTIATED.
//
//The function proceeds calling the chainlist function installed on pad (see gst_pad_set_chain_list_function())
//and the return value of that function is returned to the caller.
//GST_FLOW_NOT_SUPPORTED is returned if pad has no chainlist function.
//
//In all cases, success or failure, the caller loses its reference to list after calling this function.
//
//MT safe.
//Parameters
//pad
//a sink GstPad, returns GST_FLOW_ERROR if not.
//list
//the GstBufferList to send, return GST_FLOW_ERROR if not.
//Returns
//a GstFlowReturn from the pad.
func (p *Pad) ChainList(list *BufferList) FlowReturn {
	return FlowReturn(C.gst_pad_chain_list(p.g(), list.g()))
}

//Pause the task of pad .
//This function will also wait until the function executed by the task is finished if this function is not called from the task function.
//Parameters
//pad
//the GstPad to pause the task of
//Returns
//a TRUE if the task could be paused or FALSE when the pad has no task.
func (p *Pad) PauseTask() bool {
	return C.gst_pad_pause_task(p.g()) != 0
}

//Stop the task of pad .
//This function will also make sure that the function executed by the task will effectively stop if not called from the GstTaskFunction.
//
//This function will deadlock if called from the GstTaskFunction of the task. Use gst_task_pause() instead.
//
//Regardless of whether the pad has a task, the stream lock is acquired and released so as to ensure that streaming through this pad has finished.
//Parameters
//pad
//the GstPad to stop the task of
//Returns
//a TRUE if the task could be stopped or FALSE on error.
func (p *Pad) StopTask() bool {
	return C.gst_pad_stop_task(p.g()) != 0
}

//Activates or deactivates the given pad. Normally called from within core state change functions.
//
//If active , makes sure the pad is active. If it is already active, either in push or pull mode, just return.
//Otherwise dispatches to the pad's activate function to perform the actual activation.
//
//If not active , calls gst_pad_activate_mode() with the pad's current mode and a FALSE argument.
//Parameters
//pad
//the GstPad to activate or deactivate.
//active
//whether or not the pad should be active.
//Returns
//TRUE if the operation was successful.
//MT safe.
func (p *Pad) SetActive(active bool) bool {
	return C.gst_pad_set_active(p.g(), gBoolean(active)) != 0
}

//Get the GstPadMode of pad, which will be GST_PAD_MODE_NONE if the pad has not been activated yet,
//and otherwise either GST_PAD_MODE_PUSH or GST_PAD_MODE_PULL depending on which mode the pad was activated in.
//Parameters
//pad
//a GstPad
func (p *Pad) Mode() PadMode {
	return PadMode(p.g().mode)
}

//#define GST_PAD_IS_SRC(pad)		(GST_PAD_DIRECTION(pad) == GST_PAD_SRC)
//Parameters
//pad
//a GstPad
//Returns
//TRUE if the pad is a source pad (i.e. produces data).
func (p *Pad) IsSrc() bool {
	return p.GetDirection() == PAD_SRC
}

//#define GST_PAD_IS_SINK(pad)		(GST_PAD_DIRECTION(pad) == GST_PAD_SINK)
//Parameters
//pad
//a GstPad
//Returns
//TRUE if the pad is a sink pad (i.e. consumes data).
func (p *Pad) IsSink() bool {
	return p.GetDirection() == PAD_SINK
}

//#define GST_PAD_IS_FLUSHING(pad) (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_FLUSHING))
//Check if the given pad is flushing.
//Parameters
//pad
//a GstPad
func (p *Pad) IsFlushing() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_FLUSHING))
}

//#define GST_PAD_SET_FLUSHING(pad) (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_FLUSHING))
//Set the given pad to flushing state, which means it will not accept any more events, queries or buffers,
//and return GST_FLOW_FLUSHING if any buffers are pushed on it. This usually happens when the pad is shut down or when a flushing seek happens.
//This is used inside GStreamer when flush start/stop events pass through pads, or when an element state is changed and pads are activated or deactivated.
//Parameters
//pad
//a GstPad
func (p *Pad) SetFlushing() {
	p.FlagSet(uint32(PAD_FLAG_FLUSHING))
}

//#define GST_PAD_UNSET_FLUSHING(pad) (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_FLUSHING))
//Unset the flushing flag.
//Parameters
//pad
//a GstPad
func (p *Pad) UnsetFlushing() {
	p.FlagUnset(uint32(PAD_FLAG_FLUSHING))
}

//#define GST_PAD_IS_EOS(pad)	        (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_EOS))
//Check if the pad is in EOS state.
//Parameters
//pad
//a GstPad
func (p *Pad) IsEOS() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_EOS))
}

//#define GST_PAD_HAS_PENDING_EVENTS(pad) (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_PENDING_EVENTS))
//Check if the given pad has pending events. This is used internally by GStreamer.
//Parameters
//pad
//a GstPad
func (p *Pad) HasPendingEvents() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_PENDING_EVENTS))
}

//#define GST_PAD_IS_FIXED_CAPS(pad) (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_FIXED_CAPS))
//Check if the given pad is using fixed caps, which means that once the caps are set on the pad ,
//the caps query function will only return those caps. See gst_pad_use_fixed_caps().
func (p *Pad) IsFixedCaps() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_FIXED_CAPS))
}

//#define GST_PAD_NEEDS_PARENT(pad)       (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_NEED_PARENT))
//Check if there is a parent object before calling into the pad callbacks. This is used internally by GStreamer.
func (p *Pad) NeedsParent() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_NEED_PARENT))
}

//#define GST_PAD_IS_PROXY_CAPS(pad)      (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_PROXY_CAPS))
//Check if the given pad is set to proxy caps.
//This means that the default event and query handler will forward all events and queries to the internally linked pads instead of discarding them.
func (p *Pad) IsProxyCaps() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_PROXY_CAPS))
}

//#define GST_PAD_SET_PROXY_CAPS(pad)     (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_PROXY_CAPS))
//Set pad to proxy caps, so that all caps-related events and queries are proxied down- or upstream to the other side of the element automatically.
//Set this if the element always outputs data in the exact same format as it receives as input.
//This is just for convenience to avoid implementing some standard event and query handling code in an element.
func (p *Pad) SetProxyCaps() {
	p.FlagSet(uint32(PAD_FLAG_PROXY_CAPS))
}

//#define GST_PAD_UNSET_PROXY_CAPS(pad)   (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_PROXY_CAPS))
//Unset proxy caps flag.
func (p *Pad) UnsetProxyCaps() {
	p.FlagUnset(uint32(PAD_FLAG_PROXY_CAPS))
}

//#define GST_PAD_IS_PROXY_ALLOCATION(pad)    (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_PROXY_ALLOCATION))
//Check if the given pad is set as proxy allocation which means
//that the default query handler will forward allocation queries to the internally linked pads instead of discarding them.
//Parameters
//pad
//a GstPad
func (p *Pad) IsProxyAllocation() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_PROXY_ALLOCATION))
}

//#define GST_PAD_SET_PROXY_ALLOCATION(pad)   (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_PROXY_ALLOCATION))
//Set pad to proxy allocation queries, which means
//that the default query handler will forward allocation queries to the internally linked pads instead of discarding them.
//Set this if the element always outputs data in the exact same format as it receives as input.
//This is just for convenience to avoid implementing some standard query handling code in an element.
//Parameters
//pad
//a GstPad
func (p *Pad) SetProxyAllocation() {
	p.FlagSet(uint32(PAD_FLAG_PROXY_ALLOCATION))
}

//#define GST_PAD_UNSET_PROXY_ALLOCATION(pad) (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_PROXY_ALLOCATION))
//Unset proxy allocation flag.
func (p *Pad) UnsetProxyAllocation() {
	p.FlagUnset(uint32(PAD_FLAG_PROXY_ALLOCATION))
}

//#define GST_PAD_IS_PROXY_SCHEDULING(pad)    (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_PROXY_SCHEDULING))
//Check if the given pad is set to proxy scheduling queries,
//which means that the default query handler will forward scheduling queries to the internally linked pads instead of discarding them.
func (p *Pad) IsProxyScheduling() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_PROXY_SCHEDULING))
}

//#define GST_PAD_SET_PROXY_SCHEDULING(pad)   (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_PROXY_SCHEDULING))
//Set pad to proxy scheduling queries, which means that the default query handler will forward scheduling queries to the internally linked pads
//instead of discarding them. You will usually want to handle scheduling queries explicitly if your element supports multiple scheduling modes.
func (p *Pad) SetProxyScheduling() {
	p.FlagSet(uint32(PAD_FLAG_PROXY_SCHEDULING))
}

//#define GST_PAD_UNSET_PROXY_SCHEDULING(pad) (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_PROXY_SCHEDULING))
//Unset proxy scheduling flag.
func (p *Pad) UnsetProxyScheduling() {
	p.FlagUnset(uint32(PAD_FLAG_PROXY_SCHEDULING))
}

//#define GST_PAD_IS_ACCEPT_INTERSECT(pad)    (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_ACCEPT_INTERSECT))
//Check if the pad's accept intersect flag is set.
//The default accept-caps handler will check if the caps intersect the query-caps result instead of checking for a subset.
//This is interesting for parser elements that can accept incompletely specified caps.
func (p *Pad) IsAcceptIntersect() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_ACCEPT_INTERSECT))
}

//#define GST_PAD_SET_ACCEPT_INTERSECT(pad)   (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_ACCEPT_INTERSECT))
//Set pad to by default accept caps by intersecting the result instead of checking for a subset.
//This is interesting for parser elements that can accept incompletely specified caps.
func (p *Pad) SetAcceptIntersect() {
	p.FlagSet(uint32(PAD_FLAG_ACCEPT_INTERSECT))
}

//#define GST_PAD_UNSET_ACCEPT_INTERSECT(pad) (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_ACCEPT_INTERSECT))
//Unset accept intersect flag.
func (p *Pad) UnsetAcceptIntersect() {
	p.FlagUnset(uint32(PAD_FLAG_ACCEPT_INTERSECT))
}

//#define GST_PAD_IS_ACCEPT_TEMPLATE(pad)    (GST_OBJECT_FLAG_IS_SET (pad, GST_PAD_FLAG_ACCEPT_TEMPLATE))
//Check if the pad's accept caps operation will use the pad template caps.
//The default accept-caps will do a query caps to get the caps, which might be querying downstream causing unnecessary overhead.
//It is recommended to implement a proper accept-caps query handler or to use this flag to prevent recursive accept-caps handling.
//Since: 1.6
func (p *Pad) IsAcceptTemplate() bool {
	return p.FlagIsSet(uint32(PAD_FLAG_ACCEPT_TEMPLATE))
}

//#define GST_PAD_SET_ACCEPT_TEMPLATE(pad)   (GST_OBJECT_FLAG_SET (pad, GST_PAD_FLAG_ACCEPT_TEMPLATE))
//Set pad to by default use the pad template caps to compare with the accept caps instead of using a caps query result.
//Since: 1.6
func (p *Pad) SetAcceptTemplate() {
	p.FlagSet(uint32(PAD_FLAG_ACCEPT_TEMPLATE))
}

//#define GST_PAD_UNSET_ACCEPT_TEMPLATE(pad) (GST_OBJECT_FLAG_UNSET (pad, GST_PAD_FLAG_ACCEPT_TEMPLATE))
//Unset accept template flag.
//Since: 1.6
func (p *Pad) UnsetAcceptTemplate() {
	p.FlagUnset(uint32(PAD_FLAG_ACCEPT_TEMPLATE))
}

type GhostPad struct {
	Pad
}

func (p *GhostPad) g() *C.GstGhostPad {
	return (*C.GstGhostPad)(p.GetPtr())
}

func (p *GhostPad) AsGhostPad() *GhostPad {
	return p
}

func (p *GhostPad) SetTarget(new_target *Pad) bool {
	return C.gst_ghost_pad_set_target(p.g(), new_target.g()) != 0
}

func (p *GhostPad) GetTarget() *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_ghost_pad_get_target(p.g())))
	return r
}

func (p *GhostPad) Construct() bool {
	return C.gst_ghost_pad_construct(p.g()) != 0
}

func NewGhostPad(name string, target *Pad) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new(s, target.g())))
	return p
}

func NewGhostPadNoTarget(name string, dir PadDirection) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new_no_target(s, dir.g())))
	return p
}
