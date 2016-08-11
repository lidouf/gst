package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
#include "gst.go.h"
static GstPad* get_active_switch_pad(GstElement *switcher) {
	GstPad* active_pad;
	g_object_get(G_OBJECT(switcher), "active-pad", &active_pad, NULL);
	return active_pad;
}
GstElementClass* get_class_by_element(GstElement *element) {
	return GST_ELEMENT_GET_CLASS(element);
}
*/
import "C"

import (
	"errors"
	"github.com/lidouf/glib"
	"time"
	"unsafe"
)

type State C.GstState

const (
	STATE_VOID_PENDING = State(C.GST_STATE_VOID_PENDING)
	STATE_NULL         = State(C.GST_STATE_NULL)
	STATE_READY        = State(C.GST_STATE_READY)
	STATE_PAUSED       = State(C.GST_STATE_PAUSED)
	STATE_PLAYING      = State(C.GST_STATE_PLAYING)
)

func (s *State) g() *C.GstState {
	return (*C.GstState)(s)
}

func (s State) String() string {
	switch s {
	case STATE_VOID_PENDING:
		return "STATE_VOID_PENDING"
	case STATE_NULL:
		return "STATE_NULL"
	case STATE_READY:
		return "STATE_READY"
	case STATE_PAUSED:
		return "STATE_PAUSED"
	case STATE_PLAYING:
		return "STATE_PLAYING"
	}
	panic("Unknown state")
}

//Gets a string representing the given state.
//Returns
//a string with the name of the state.
func (s State) GetName() string {
	return C.GoString((*C.char)(C.gst_element_state_get_name(C.GstState(s))))
}

type StateChangeReturn C.GstStateChangeReturn

const (
	STATE_CHANGE_FAILURE    = StateChangeReturn(C.GST_STATE_CHANGE_FAILURE)
	STATE_CHANGE_SUCCESS    = StateChangeReturn(C.GST_STATE_CHANGE_SUCCESS)
	STATE_CHANGE_ASYNC      = StateChangeReturn(C.GST_STATE_CHANGE_ASYNC)
	STATE_CHANGE_NO_PREROLL = StateChangeReturn(C.GST_STATE_CHANGE_NO_PREROLL)
)

//Gets a string representing the given state change result.
//Returns
//a string with the name of the state result.
func (s StateChangeReturn) GetName() string {
	return C.GoString((*C.char)(C.gst_element_state_change_return_get_name(C.GstStateChangeReturn(s))))
}

type StateChange C.GstStateChange

const (
	STATE_CHANGE_NULL_TO_READY     = StateChange(C.GST_STATE_CHANGE_NULL_TO_READY)
	STATE_CHANGE_READY_TO_PAUSED   = StateChange(C.GST_STATE_CHANGE_READY_TO_PAUSED)
	STATE_CHANGE_PAUSED_TO_PLAYING = StateChange(C.GST_STATE_CHANGE_PAUSED_TO_PLAYING)
	STATE_CHANGE_PLAYING_TO_PAUSED = StateChange(C.GST_STATE_CHANGE_PLAYING_TO_PAUSED)
	STATE_CHANGE_PAUSED_TO_READY   = StateChange(C.GST_STATE_CHANGE_PAUSED_TO_READY)
	STATE_CHANGE_READY_TO_NULL     = StateChange(C.GST_STATE_CHANGE_READY_TO_NULL)
)

type ElementFlags C.GstElementFlags

const (
	ELEMENT_FLAG_LOCKED_STATE  = ElementFlags(C.GST_ELEMENT_FLAG_LOCKED_STATE)
	ELEMENT_FLAG_SINK          = ElementFlags(C.GST_ELEMENT_FLAG_SINK)
	ELEMENT_FLAG_SOURCE        = ElementFlags(C.GST_ELEMENT_FLAG_SOURCE)
	ELEMENT_FLAG_PROVIDE_CLOCK = ElementFlags(C.GST_ELEMENT_FLAG_PROVIDE_CLOCK)
	ELEMENT_FLAG_REQUIRE_CLOCK = ElementFlags(C.GST_ELEMENT_FLAG_REQUIRE_CLOCK)
	ELEMENT_FLAG_INDEXABLE     = ElementFlags(C.GST_ELEMENT_FLAG_INDEXABLE)
	/* padding */
	ELEMENT_FLAG_LAST = ElementFlags(C.GST_ELEMENT_FLAG_LAST)
)

type Element struct {
	GstObj
}

func (e *Element) g() *C.GstElement {
	return (*C.GstElement)(e.GetPtr())
}

func (e *Element) AsElement() *Element {
	return e
}

func (e *Element) AsVideoOverlay() *VideoOverlay {
	cVideoOverlay := C.toGstVideoOverlay(e.GetPtr())
	//defer C.free(unsafe.Pointer(cVideoOverlay))
	goVideoOverlay := new(VideoOverlay)
	goVideoOverlay.SetPtr(glib.Pointer(cVideoOverlay))

	return goVideoOverlay
}

//LiD: add GetSwitchActivePad for input-selector active-pad property
func (e *Element) GetSwitchActivePad() *Pad {
	s := (*C.gchar)(C.CString("active-pad"))
	defer C.free(unsafe.Pointer(s))

	var r1 *C.GstPad
	r1 = C.get_active_switch_pad(e.g())
	if r1 == nil {
		return nil
	}
	r := new(Pad)
	r.SetPtr(glib.Pointer(r1))

	return r
}

func (e *Element) GetClass() *ElementClass {
	c := C.get_class_by_element(e.g())
	ec := new(ElementClass)
	ec.SetPtr(glib.Pointer(c))
	return ec
}

type ElementClass struct {
	GstObjClass
}

func (c *ElementClass) g() *C.GstElementClass {
	return (*C.GstElementClass)(c.GetPtr())
}

func (c *ElementClass) AsElementClass() *ElementClass {
	return c
}

//Adds a padtemplate to an element class. This is mainly used in the _class_init functions of classes.
//If a pad template with the same name as an already existing one is added the old one is replaced by the new one.
func (c *ElementClass) AddPadTemplate(pt *PadTemplate) {
	C.gst_element_class_add_pad_template(c.g(), pt.g())
}

//Adds a pad template to an element class based on the static pad template templ .
//This is mainly used in the _class_init functions of element implementations.
//If a pad template with the same name already exists, the old one is replaced by the new one.
func (c *ElementClass) AddStaticPadTemplate(pt *StaticPadTemplate) {
	C.gst_element_class_add_static_pad_template(c.g(), pt.g())
}

//Retrieves a padtemplate from element_class with the given name.
//If you use this function in the GInstanceInitFunc of an object class that has subclasses, make sure to pass the g_class parameter of the GInstanceInitFunc here.
//Returns
//the GstPadTemplate with the given name, or NULL if none was found. No unreferencing is necessary.
func (c *ElementClass) GetPadTemplate(name string) *PadTemplate {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	pt := new(PadTemplate)
	pt.SetPtr(glib.Pointer(C.gst_element_class_get_pad_template(c.g(), s)))
	return pt
}

//Retrieves a list of the pad templates associated with element_class . The list must not be modified by the calling code.
//If you use this function in the GInstanceInitFunc of an object class that has subclasses, make sure to pass the g_class parameter of the GInstanceInitFunc here.
//Returns
//the GList of pad templates.
func (c *ElementClass) GetPadTemplateList() *glib.List {
	return glib.WrapList(uintptr(unsafe.Pointer(C.gst_element_class_get_pad_template_list(c.g()))))
}

//Sets the detailed information for a GstElementClass.
//This function is for use in _class_init functions only.
//Parameters
//longname
//The long English name of the element. E.g. "File Sink"
//classification
//String describing the type of element, as an unordered list separated with slashes ('/').
//See draft-klass.txt of the design docs for more details and common types. E.g: "Sink/File"
//description
//Sentence describing the purpose of the element. E.g: "Write stream to a file"
//author
//Name and contact details of the author(s). Use \n to separate multiple author metadata. E.g: "Joe Bloggs <joe.blogs at foo.com>"
func (c *ElementClass) SetMetadata(longName, classification, description, author string) {
	s1 := (*C.gchar)(C.CString(longName))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(classification))
	defer C.free(unsafe.Pointer(s2))
	s3 := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(s3))
	s4 := (*C.gchar)(C.CString(author))
	defer C.free(unsafe.Pointer(s4))
	C.gst_element_class_set_metadata(c.g(), s1, s2, s3, s4)
}

//Sets the detailed information for a GstElementClass.
//This function is for use in _class_init functions only.
//Same as gst_element_class_set_metadata(), but longname , classification , description , and author must be static strings or inlined strings,
//as they will not be copied. (GStreamer plugins will be made resident once loaded, so this function can be used even from dynamically loaded plugins.)
//Parameters
//longname
//The long English name of the element. E.g. "File Sink"
//classification
//String describing the type of element, as an unordered list separated with slashes ('/').
//See draft-klass.txt of the design docs for more details and common types. E.g: "Sink/File"
//description
//Sentence describing the purpose of the element. E.g: "Write stream to a file"
//author
//Name and contact details of the author(s). Use \n to separate multiple author metadata. E.g: "Joe Bloggs <joe.blogs at foo.com>"
func (c *ElementClass) SetStaticMetadata(longName, classification, description, author string) {
	s1 := (*C.gchar)(C.CString(longName))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(classification))
	defer C.free(unsafe.Pointer(s2))
	s3 := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(s3))
	s4 := (*C.gchar)(C.CString(author))
	defer C.free(unsafe.Pointer(s4))
	C.gst_element_class_set_static_metadata(c.g(), s1, s2, s3, s4)
}

//Set key with value as metadata in klass .
//Parameters
//key
//the key to set
//value
//the value to set
func (c *ElementClass) AddMetadata(key, value string) {
	s1 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(s2))
	C.gst_element_class_add_metadata(c.g(), s1, s2)
}

//Set key with value as metadata in klass .
//Same as gst_element_class_add_metadata(), but value must be a static string or an inlined string, as it will not be copied. (GStreamer plugins will be made resident once loaded, so this function can be used even from dynamically loaded plugins.)
//Parameters
//key
//the key to set
//value
//the value to set
func (c *ElementClass) AddStaticMetadata(key, value string) {
	s1 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(s2))
	C.gst_element_class_add_static_metadata(c.g(), s1, s2)
}

//Get metadata with key in klass .
//Parameters
//key
//the key to get
//Returns
//the metadata for key .
func (c *ElementClass) GetMetadata(key string) string {
	s := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(s))
	return C.GoString((*C.char)(C.gst_element_class_get_metadata(c.g(), s)))
}

//#define GST_ELEMENT_IS_LOCKED_STATE(elem)        (GST_OBJECT_FLAG_IS_SET(elem,GST_ELEMENT_FLAG_LOCKED_STATE))
//Check if the element is in the locked state and therefore will ignore state changes from its parent object.
func (e *Element) IS_LOCKED_STATE() bool {
	return e.FlagIsSet(uint32(ELEMENT_FLAG_LOCKED_STATE))
}

//Adds a pad (link point) to element . pad 's parent will be set to element ; see gst_object_set_parent() for refcounting information.
//Pads are not automatically activated so elements should perform the needed steps to activate the pad in case this pad is added in the PAUSED or PLAYING state.
//See gst_pad_set_active() for more information about activating pads.
//The pad and the element should be unlocked when calling this function.
//This function will emit the “pad-added” signal on the element.
//Parameters
//pad
//the GstPad to add to the element.
//Returns
//TRUE if the pad could be added. This function can fail when a pad with the same name already existed or the pad already had another parent.
//MT safe.
func (e *Element) AddPad(p *Pad) bool {
	return C.gst_element_add_pad(e.g(), p.g()) != 0
}

//Creates a pad for each pad template that is always available. This function is only useful during object initialization of subclasses of GstElement.
func (e *Element) CreateAllPads() {
	C.gst_element_create_all_pads(e.g())
}

//Looks for an unlinked pad to which the given pad can link. It is not guaranteed that linking the pads will work, though it should work in most cases.
//This function will first attempt to find a compatible unlinked ALWAYS pad,
//and if none can be found, it will request a compatible REQUEST pad by looking at the templates of element .
//Parameters
//pad
//the GstPad to find a compatible one for.
//caps
//the GstCaps to use as a filter.
//Returns
//the GstPad to which a link can be made, or NULL if one cannot be found. gst_object_unref() after usage.
func (e *Element) GetCompatiblePad(p *Pad, c *Caps) *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_element_get_compatible_pad(e.g(), p.g(), c.g())))
	return p
}

//Retrieves a pad template from element that is compatible with compattempl . Pads from compatible templates can be linked together.
//Parameters
//compattempl
//the GstPadTemplate to find a compatible template for.
//Returns
//a compatible GstPadTemplate, or NULL if none was found. No unreferencing is necessary.
func (e *Element) GetCompatiblePadTemplate(t *PadTemplate) *PadTemplate {
	p := new(PadTemplate)
	p.SetPtr(glib.Pointer(C.gst_element_get_compatible_pad_template(e.g(), t.g())))
	return p
}

//Retrieves a pad from the element by name (e.g. "src_%d"). This version only retrieves request pads.
//The pad should be released with gst_element_release_request_pad().
//This method is slower than manually getting the pad template and calling gst_element_request_pad()
//if the pads should have a specific name (e.g. name is "src_1" instead of "src_%u").
//Parameters
//name
//the name of the request GstPad to retrieve.
//Returns
//requested GstPad if found, otherwise NULL. Release after usage.
func (e *Element) GetRequestPad(name string) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	cp := C.gst_element_get_request_pad(e.g(), s)
	if cp == nil {
		return nil
	}
	p := new(Pad)
	p.SetPtr(glib.Pointer(cp))
	return p
}

//Retrieves a pad from element by name. This version only retrieves already-existing (i.e. 'static') pads.
//Parameters
//name
//the name of the static GstPad to retrieve.
//Returns
//the requested GstPad if found, otherwise NULL. unref after usage.
//MT safe.
func (e *Element) GetStaticPad(name string) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	cp := C.gst_element_get_static_pad(e.g(), s)
	if cp == nil {
		return nil
	}
	p := new(Pad)
	p.SetPtr(glib.Pointer(cp))
	return p
}

//Retrieves a request pad from the element according to the provided template.
//Pad templates can be looked up using gst_element_factory_get_static_pad_templates().
//The pad should be released with gst_element_release_request_pad().
//Parameters
//templ
//a GstPadTemplate of which we want a pad of.
//name
//the name of the request GstPad to retrieve. Can be NULL.
//caps
//the caps of the pad we want to request. Can be NULL.
//Returns
//requested GstPad if found, otherwise NULL. Release after usage.
func (e *Element) RequestPad(templ *PadTemplate, name string, caps *Caps) *Pad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	if name == "" {
		s = nil
	}
	var cp *C.GstPad
	if caps == nil {
		cp = C.gst_element_request_pad(e.g(), templ.g(), s, nil)
	} else {
		cp = C.gst_element_request_pad(e.g(), templ.g(), s, caps.g())
	}
	if cp == nil {
		return nil
	}
	p := new(Pad)
	p.SetPtr(glib.Pointer(cp))
	return p
}

//Use this function to signal that the element does not expect any more pads to show up in the current pipeline.
//This function should be called whenever pads have been added by the element itself.
//Elements with GST_PAD_SOMETIMES pad templates use this in combination with autopluggers to figure out that the element is done initializing its pads.
//This function emits the “no-more-pads” signal.
//MT safe.
func (e *Element) NoMorePads() {
	C.gst_element_no_more_pads(e.g())
}

//Makes the element free the previously requested pad as obtained with gst_element_request_pad().
//This does not unref the pad.
//If the pad was created by using gst_element_request_pad(), gst_element_release_request_pad() needs to be followed by gst_object_unref() to free the pad .
//MT safe.
//Parameters
//pad
//the GstPad to release.
func (e *Element) ReleaseRequestPad(pad *Pad) {
	C.gst_element_release_request_pad(e.g(), pad.g())
}

//Removes pad from element . pad will be destroyed if it has not been referenced elsewhere using gst_object_unparent().
//This function is used by plugin developers and should not be used by applications.
//Pads that were dynamically requested from elements with gst_element_request_pad() should be released with the gst_element_release_request_pad() function instead.
//Pads are not automatically deactivated so elements should perform the needed steps to deactivate the pad in case this pad is removed in the PAUSED or PLAYING state.
//See gst_pad_set_active() for more information about deactivating pads.
//The pad and the element should be unlocked when calling this function.
//This function will emit the “pad-removed” signal on the element.
//Parameters
//pad
//the GstPad to remove from the element.
//Returns
//TRUE if the pad could be removed. Can return FALSE if the pad does not belong to the provided element.
//MT safe.
func (e *Element) RemovePad(pad *Pad) bool {
	return C.gst_element_remove_pad(e.g(), pad.g()) != 0
}

//Links src to dest . The link must be from source to destination; the other direction will not be tried.
//The function looks for existing pads that aren't linked yet. It will request new pads if necessary.
//Such pads need to be released manually when unlinking. If multiple links are possible, only one is established.
//Make sure you have added your elements to a bin or pipeline with gst_bin_add() before trying to link them.
//Parameters
//element_1
//the first GstElement in the link chain.
//element_2
//the second GstElement in the link chain.
//...
//Returns
//TRUE if the elements could be linked, FALSE otherwise.
func (e *Element) Link(next ...*Element) bool {
	for _, dst := range next {
		if C.gst_element_link(e.g(), dst.g()) == 0 {
			return false
		}
		e = dst
	}
	return true
}

//Unlinks all source pads of the source element with all sink pads of the sink element to which they are linked.
//If the link has been made using gst_element_link(), it could have created an requestpad, which has to be released using gst_element_release_request_pad().
//Parameters
//element_1
//the first GstElement in the link chain.
//element_2
//the second GstElement in the link chain.
//...
//the sink GstElement to unlink.
func (e *Element) Unlink(next ...*Element) {
	for _, dst := range next {
		C.gst_element_unlink(e.g(), dst.g())
		e = dst
	}
}

//Links the two named pads of the source and destination elements.
//Side effect is that if one of the pads has no parent, it becomes a child of the parent of the other element.
//If they have different parents, the link fails.
//Parameters
//srcpadname
//the name of the GstPad in source element or NULL for any pad.
//dest
//the GstElement containing the destination pad.
//destpadname
//the name of the GstPad in destination element, or NULL for any pad.
//Returns
//TRUE if the pads could be linked, FALSE otherwise.
func (e *Element) LinkPads(pad_name string, dst *Element, dst_pad_name string) bool {
	src_pname := (*C.gchar)(C.CString(pad_name))
	defer C.free(unsafe.Pointer(src_pname))
	dst_pname := (*C.gchar)(C.CString(dst_pad_name))
	defer C.free(unsafe.Pointer(dst_pname))
	return C.gst_element_link_pads(e.g(), src_pname, dst.g(), dst_pname) != 0
}

//Links the two named pads of the source and destination elements.
//Side effect is that if one of the pads has no parent, it becomes a child of the parent of the other element.
//If they have different parents, the link fails.
//Calling gst_element_link_pads_full() with flags == GST_PAD_LINK_CHECK_DEFAULT is the same as
//calling gst_element_link_pads() and the recommended way of linking pads with safety checks applied.
//This is a convenience function for gst_pad_link_full().
//Parameters
//srcpadname
//the name of the GstPad in source element or NULL for any pad.
//dest
//the GstElement containing the destination pad.
//destpadname
//the name of the GstPad in destination element, or NULL for any pad.
//flags
//the GstPadLinkCheck to be performed when linking pads.
//Returns
//TRUE if the pads could be linked, FALSE otherwise.
func (e *Element) LinkPadsFull(pad_name string, dst *Element, dst_pad_name string, flags PadLinkCheck) bool {
	src_pname := (*C.gchar)(C.CString(pad_name))
	defer C.free(unsafe.Pointer(src_pname))
	dst_pname := (*C.gchar)(C.CString(dst_pad_name))
	defer C.free(unsafe.Pointer(dst_pname))
	return C.gst_element_link_pads_full(e.g(), src_pname, dst.g(), dst_pname, C.GstPadLinkCheck(flags)) != 0
}

//Unlinks the two named pads of the source and destination elements.
//This is a convenience function for gst_pad_unlink().
//Parameters
//srcpadname
//the name of the GstPad in source element.
//dest
//a GstElement containing the destination pad.
//destpadname
//the name of the GstPad in destination element.
func (e *Element) UnlinkPads(pad_name string, dst *Element, dst_pad_name string) {
	src_pname := (*C.gchar)(C.CString(pad_name))
	defer C.free(unsafe.Pointer(src_pname))
	dst_pname := (*C.gchar)(C.CString(dst_pad_name))
	defer C.free(unsafe.Pointer(dst_pname))
	C.gst_element_unlink_pads(e.g(), src_pname, dst.g(), dst_pname)
}

//Links the two named pads of the source and destination elements.
//Side effect is that if one of the pads has no parent, it becomes a child of the parent of the other element.
//If they have different parents, the link fails.
//If caps is not NULL, makes sure that the caps of the link is a subset of caps .
//Parameters
//srcpadname
//the name of the GstPad in source element or NULL for any pad.
//dest
//the GstElement containing the destination pad.
//destpadname
//the name of the GstPad in destination element or NULL for any pad.
//filter
//the GstCaps to filter the link, or NULL for no filter.
//Returns
//TRUE if the pads could be linked, FALSE otherwise.
func (e *Element) LinkPadsFiltered(pad_name string, dst *Element, dst_pad_name string, filter *Caps) bool {
	src_pname := (*C.gchar)(C.CString(pad_name))
	defer C.free(unsafe.Pointer(src_pname))
	dst_pname := (*C.gchar)(C.CString(dst_pad_name))
	defer C.free(unsafe.Pointer(dst_pname))
	return C.gst_element_link_pads_filtered(e.g(), src_pname, dst.g(), dst_pname, filter.g()) != 0
}

//Links src to dest using the given caps as filtercaps. The link must be from source to destination; the other direction will not be tried.
//The function looks for existing pads that aren't linked yet. It will request new pads if necessary.
//If multiple links are possible, only one is established.
//Make sure you have added your elements to a bin or pipeline with gst_bin_add() before trying to link them.
//Parameters
//dest
//the GstElement containing the destination pad.
//filter
//the GstCaps to filter the link, or NULL for no filter.
//Returns
//TRUE if the pads could be linked, FALSE otherwise.
func (e *Element) LinkFiltered(dst *Element, filter *Caps) bool {
	return C.gst_element_link_filtered(e.g(), dst.g(), filter.g()) != 0
}

//Set the base time of an element. See gst_element_get_base_time().
//MT safe.
//Parameters
//time
//the base time to set.
func (e *Element) SetBaseTime(t ClockTime) {
	C.gst_element_set_base_time(e.g(), C.GstClockTime(t))
}

//Returns the base time of the element. The base time is the absolute time of the clock when this element was last put to PLAYING.
//Subtracting the base time from the clock time gives the running time of the element.
//Returns
//the base time of the element.
//MT safe.
func (e *Element) GetBaseTime() ClockTime {
	return ClockTime(C.gst_element_get_base_time(e.g()))
}

//Set the start time of an element. The start time of the element is the running time of the element when it last went to the PAUSED state.
//In READY or after a flushing seek, it is set to 0.
//Toplevel elements like GstPipeline will manage the start_time and base_time on its children.
//Setting the start_time to GST_CLOCK_TIME_NONE on such a toplevel element will disable the distribution of the base_time to the children
//and can be useful if the application manages the base_time itself,
//for example if you want to synchronize capture from multiple pipelines, and you can also ensure that the pipelines have the same clock.
//MT safe.
//Parameters
//time
//the base time to set.
func (e *Element) SetStartTime(t ClockTime) {
	C.gst_element_set_start_time(e.g(), C.GstClockTime(t))
}

//Returns the start time of the element. The start time is the running time of the clock when this element was last put to PAUSED.
//Usually the start_time is managed by a toplevel element such as GstPipeline.
//MT safe.
//Returns
//the start time of the element.
func (e *Element) GetStartTime() ClockTime {
	return ClockTime(C.gst_element_get_start_time(e.g()))
}

//Sets the bus of the element. Increases the refcount on the bus. For internal use only, unless you're testing elements.
//MT safe.
//Parameters
//bus
//the GstBus to set.
func (e *Element) SetBus(bus *Bus) {
	C.gst_element_set_bus(e.g(), bus.g())
}

//Returns the bus of the element. Note that only a GstPipeline will provide a bus for the application.
//Parameters
//Returns
//the element's GstBus. unref after usage.
//MT safe.
func (e *Element) GetBus() *Bus {
	bus := C.gst_element_get_bus(e.g())
	if bus == nil {
		return nil
	}
	b := new(Bus)
	b.SetPtr(glib.Pointer(bus))
	return b
}

//Sets the context of the element. Increases the refcount of the context.
//MT safe.
//Parameters
//context
//the GstContext to set.
func (e *Element) SetContext(context *Context) {
	C.gst_element_set_context(e.g(), context.g())
}

//Gets the context with context_type set on the element or NULL.
//MT safe.
//Parameters
//context_type
//a name of a context to retrieve
//Returns
//A GstContext or NULL.
//Since: 1.8
func (e *Element) GetContext(name string) *Context {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	context := C.gst_element_get_context(e.g(), s)
	if context == nil {
		return nil
	}
	r := new(Context)
	r.SetPtr(glib.Pointer(context))
	return r
}

//Gets the context with context_type set on the element or NULL.
//Parameters
//context_type
//a name of a context to retrieve
//Returns
//A GstContext or NULL.
//Since: 1.8
func (e *Element) GetContextUnlocked(name string) *Context {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	context := C.gst_element_get_context_unlocked(e.g(), s)
	if context == nil {
		return nil
	}
	r := new(Context)
	r.SetPtr(glib.Pointer(context))
	return r
}

//Gets the contexts set on the element.
//MT safe.
//Returns
//List of GstContext.
//Since: 1.8
func (e *Element) GetContexts() *glib.List {
	return glib.WrapList(uintptr(unsafe.Pointer(C.gst_element_get_contexts(e.g()))))
}

//Retrieves the factory that was used to create this element.
//Returns
//the GstElementFactory used for creating this element. no refcounting is needed.
func (e *Element) GetFactory() *ElementFactory {
	factory := C.gst_element_get_factory(e.g())
	if factory == nil {
		return nil
	}
	f := new(ElementFactory)
	f.SetPtr(glib.Pointer(factory))
	return f
}

//Sets the clock for the element. This function increases the refcount on the clock. Any previously set clock on the object is unreffed.
//Parameters
//clock
//the GstClock to set for the element.
//Returns
//TRUE if the element accepted the clock.
//An element can refuse a clock when it, for example, is not able to slave its internal clock to the clock or when it requires a specific clock to operate.
//MT safe.
func (e *Element) SetClock(t *Clock) {
	C.gst_element_set_clock(e.g(), t.g())
}

//Gets the currently configured clock of the element. This is the clock as was last set with gst_element_set_clock().
//Elements in a pipeline will only have their clock set when the pipeline is in the PLAYING state.
//Returns
//the GstClock of the element. unref after usage.
//MT safe.
func (e *Element) GetClock() *Clock {
	c := new(Clock)
	c.SetPtr(glib.Pointer(C.gst_element_get_clock(e.g())))
	return c
}

//Get the clock provided by the given element.
//An element is only required to provide a clock in the PAUSED state. Some elements can provide a clock in other states.
//Returns
//the GstClock provided by the element or NULL if no clock could be provided. Unref after usage.
//MT safe.
func (e *Element) ProvideClock() *Clock {
	c := new(Clock)
	c.SetPtr(glib.Pointer(C.gst_element_provide_clock(e.g())))
	return c
}

//Sets the state of the element.
//This function will try to set the requested state by going through all the intermediary states and calling the class's state change function for each.
//This function can return GST_STATE_CHANGE_ASYNC, in which case the element will perform the remainder of the state change asynchronously in another thread.
//An application can use gst_element_get_state() to wait for the completion of the state change
//or it can wait for a GST_MESSAGE_ASYNC_DONE or GST_MESSAGE_STATE_CHANGED on the bus.
//State changes to GST_STATE_READY or GST_STATE_NULL never return GST_STATE_CHANGE_ASYNC.
//Parameters
//state
//the element's new GstState.
//Returns
//Result of the state change using GstStateChangeReturn.
//MT safe.
func (e *Element) SetState(state State) StateChangeReturn {
	return StateChangeReturn(C.gst_element_set_state(e.g(), C.GstState(state)))
}

//Gets the state of the element.
//For elements that performed an ASYNC state change, as reported by gst_element_set_state(),
//this function will block up to the specified timeout value for the state change to complete.
//If the element completes the state change or goes into an error,
//this function returns immediately with a return value of GST_STATE_CHANGE_SUCCESS or GST_STATE_CHANGE_FAILURE respectively.
//For elements that did not return GST_STATE_CHANGE_ASYNC, this function returns the current and pending state immediately.
//This function returns GST_STATE_CHANGE_NO_PREROLL if the element successfully changed its state but is not able to provide data yet.
//This mostly happens for live sources that only produce data in GST_STATE_PLAYING.
//While the state change return is equivalent to GST_STATE_CHANGE_SUCCESS,
//it is returned to the application to signal that some sink elements might not be able to complete their state change
//because an element is not producing data to complete the preroll. When setting the element to playing, the preroll will complete and playback will start.
//Parameters
//timeout
//a GstClockTime to specify the timeout for an async state change or GST_CLOCK_TIME_NONE for infinite timeout.
//Returns
//state
//a pointer to GstState to hold the state. Can be NULL.
//pending
//a pointer to GstState to hold the pending state. Can be NULL.
//GST_STATE_CHANGE_SUCCESS if the element has no more pending state and the last state change succeeded,
//GST_STATE_CHANGE_ASYNC if the element is still performing a state change or GST_STATE_CHANGE_FAILURE if the last state change failed.
//MT safe.
func (e *Element) GetState(timeout_ns int64) (state, pending State, ret StateChangeReturn) {
	ret = StateChangeReturn(C.gst_element_get_state(
		e.g(), state.g(), pending.g(), C.GstClockTime(timeout_ns),
	))
	return
}

//Locks the state of an element, so state changes of the parent don't affect this element anymore.
//MT safe.
//Parameters
//locked_state
//TRUE to lock the element's state
//Returns
//TRUE if the state was changed, FALSE if bad parameters were given or the elements state-locking needed no change.
func (e *Element) SetLockedState(lockedState bool) bool {
	var s int
	if lockedState {
		s = 1
	} else {
		s = 0
	}
	return C.gst_element_set_locked_state(e.g(), C.gboolean(s)) != 0
}

//Checks if the state of an element is locked. If the state of an element is locked, state changes of the parent don't affect the element.
//This way you can leave currently unused elements inside bins. Just lock their state before changing the state from GST_STATE_NULL.
//MT safe.
//Returns
//TRUE, if the element's state is locked.
func (e *Element) IsLockedState() bool {
	return C.gst_element_is_locked_state(e.g()) != 0
}

//Abort the state change of the element. This function is used by elements that do asynchronous state changes and find out something is wrong.
//This function should be called with the STATE_LOCK held.
//MT safe.
func (e *Element) AbortState() {
	C.gst_element_abort_state(e.g())
}

//Commit the state change of the element and proceed to the next pending state if any.
//This function is used by elements that do asynchronous state changes.
//The core will normally call this method automatically when an element returned GST_STATE_CHANGE_SUCCESS from the state change function.
//If after calling this method the element still has not reached the pending state, the next state change is performed.
//This method is used internally and should normally not be called by plugins or applications.
//Parameters
//ret
//The previous state return value
//Returns
//The result of the commit state change.
//MT safe.
func (e *Element) ContinueState(ret StateChangeReturn) StateChangeReturn {
	return StateChangeReturn(C.gst_element_continue_state(e.g(), C.GstStateChangeReturn(ret)))
}

//Brings the element to the lost state.
//The current state of the element is copied to the pending state so that any call to gst_element_get_state() will return GST_STATE_CHANGE_ASYNC.
//An ASYNC_START message is posted. If the element was PLAYING, it will go to PAUSED.
//The element will be restored to its PLAYING state by the parent pipeline when it prerolls again.
//This is mostly used for elements that lost their preroll buffer in the GST_STATE_PAUSED or GST_STATE_PLAYING state after a flush,
//they will go to their pending state again when a new preroll buffer is queued.
//This function can only be called when the element is currently not in error or an async state change.
//This function is used internally and should normally not be called from plugins or applications.
func (e *Element) LostState() {
	C.gst_element_lost_state(e.g())
}

//Tries to change the state of the element to the same as its parent. If this function returns FALSE, the state of element is undefined.
//Returns
//TRUE, if the element's state could be synced to the parent's state.
//MT safe.
func (e *Element) SyncStateWithParent() bool {
	return C.gst_element_sync_state_with_parent(e.g()) != 0
}

//Perform transition on element .
//This function must be called with STATE_LOCK held and is mainly used internally.
//Parameters
//transition
//the requested transition
//Returns
//the GstStateChangeReturn of the state transition.
func (e *Element) ChangeState(transition StateChange) StateChangeReturn {
	return StateChangeReturn(C.gst_element_change_state(e.g(), C.GstStateChange(transition)))
}

//Post an error, warning or info message on the bus from inside an element.
//type must be of GST_MESSAGE_ERROR, GST_MESSAGE_WARNING or GST_MESSAGE_INFO.
//MT safe.
//Parameters
//type
//the GstMessageType
//domain
//the GStreamer GError domain this message belongs to
//code
//the GError code belonging to the domain
//text
//an allocated text string to be used as a replacement for the default message connected to code, or NULL.
//debug
//an allocated debug message to be used as a replacement for the default debugging information, or NULL.
//file
//the source code file where the error was generated
//function
//the source code function where the error was generated
//line
//the source code line where the error was generated
func (e *Element) MessageFull(tp MessageType, domain glib.Quark, code int, text, debug, file, function string, line int) {
	s1 := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(s1))
	s2 := (*C.gchar)(C.CString(debug))
	defer C.free(unsafe.Pointer(s2))
	s3 := (*C.gchar)(C.CString(file))
	defer C.free(unsafe.Pointer(s3))
	s4 := (*C.gchar)(C.CString(function))
	defer C.free(unsafe.Pointer(s4))
	C.gst_element_message_full(e.g(), C.GstMessageType(tp), C.GQuark(domain), C.gint(code), s1, s2, s3, s4, C.gint(line))
}

//Post a message on the element's GstBus.
//This function takes ownership of the message; if you want to access the message after this call, you should add an additional reference before calling.
//Parameters
//message
//a GstMessage to post.
//Returns
//TRUE if the message was successfully posted. The function returns FALSE if the element did not have a bus.
//MT safe.
func (e *Element) PostMessage(msg *Message) bool {
	return C.gst_element_post_message(e.g(), msg.g()) == 1
}

//Performs a query on the given element.
//For elements that don't implement a query handler, this function forwards the query to a random srcpad or to the peer of a random linked sinkpad of this element.
//Please note that some queries might need a running pipeline to work.
//Parameters
//query
//the GstQuery.
//Returns
//TRUE if the query could be performed.
//MT safe.
func (e *Element) Query(q *Query) bool {
	return C.gst_element_query(e.g(), q.g()) == 1
}

//Queries an element to convert src_val in src_format to dest_format .
//src_format
//a GstFormat to convert from.
//src_val
//a value to convert.
//dest_format
//the GstFormat to convert to.
//Returns
//dest_val
//a pointer to the result.
//error is nil if the query could be performed.
func (e *Element) QueryConvert(srcFormat Format, srcVal int64, destFormat Format) (int64, error) {
	var destVal C.gint64
	result := C.gst_element_query_convert(e.g(), C.GstFormat(srcFormat), C.gint64(srcVal), C.GstFormat(destFormat), &destVal)
	if result == 0 {
		return int64(0), errors.New("Query could not be performed")
	} else {
		return int64(destVal), nil
	}
}

//Queries an element (usually top-level pipeline or playbin element) for the stream position in nanoseconds.
//This will be a value between 0 and the stream duration (if the stream duration is known).
//This query will usually only work once the pipeline is prerolled (i.e. reached PAUSED or PLAYING state).
//The application will receive an ASYNC_DONE message on the pipeline bus when that is the case.
//If one repeatedly calls this function one can also create a query and reuse it in gst_element_query().
//Parameters
//format
//the GstFormat requested
//Returns
//a location in which to store the current position, or NULL.
//error is nil if the query could be performed.
func (e *Element) QueryPosition(format Format) (time.Duration, error) {
	var pos C.gint64
	ret := C.gst_element_query_position(e.g(), *(format.g()), &pos)
	if ret == 0 {
		return -1, errors.New("Query position from element failed")
	} else {
		return time.Duration(pos), nil
	}
}

//Queries an element (usually top-level pipeline or playbin element) for the total stream duration in nanoseconds.
//This query will only work once the pipeline is prerolled (i.e. reached PAUSED or PLAYING state).
//The application will receive an ASYNC_DONE message on the pipeline bus when that is the case.
//If the duration changes for some reason, you will get a DURATION_CHANGED message on the pipeline bus,
//in which case you should re-query the duration using this function.
//Parameters
//format
//the GstFormat requested
//Returns
//duration
//A location in which to store the total duration, or NULL.
//error is nil if the query could be performed.
func (e *Element) QueryDuration(format Format) (time.Duration, error) {
	var duration C.gint64
	ret := C.gst_element_query_duration(e.g(), *(format.g()), &duration)
	if ret == 0 {
		return -1, errors.New("Query duration from element failed")
	} else {
		return time.Duration(duration), nil
	}
}

//Sends an event to an element. If the element doesn't implement an event handler,
//the event will be pushed on a random linked sink pad for downstream events or a random linked source pad for upstream events.
//This function takes ownership of the provided event so you should gst_event_ref() it if you want to reuse the event after this call.
//MT safe.
//Parameters
//event
//the GstEvent to send to the element.
//Returns
//TRUE if the event was handled. Events that trigger a preroll (such as flushing seeks and steps) will emit GST_MESSAGE_ASYNC_DONE.
func (e *Element) SendEvent(event Event) bool {
	return C.gst_element_send_event(e.g(), event.g()) == 1
}

//Simple API to perform a seek on the given element, meaning it just seeks to the given position relative to the start of the stream.
//For more complex operations like segment seeks (e.g. for looping) or changing the playback rate
//or seeking relative to the last configured playback segment you should use gst_element_seek().
//In a completely prerolled PAUSED or PLAYING pipeline, seeking is always guaranteed to return TRUE on a seekable media type or FALSE
//when the media type is certainly not seekable (such as a live stream).
//Some elements allow for seeking in the READY state, in this case they will store the seek event and execute it when they are put to PAUSED.
//If the element supports seek in READY, it will always return TRUE when it receives the event in the READY state.
//Parameters
//format
//a GstFormat to execute the seek in, such as GST_FORMAT_TIME
//seek_flags
//seek options; playback applications will usually want to use GST_SEEK_FLAG_FLUSH | GST_SEEK_FLAG_KEY_UNIT here
//seek_pos
//position to seek to (relative to the start); if you are doing a seek in GST_FORMAT_TIME this value is in nanoseconds - multiply with GST_SECOND to convert seconds to nanoseconds or with GST_MSECOND to convert milliseconds to nanoseconds.
//Returns
//TRUE if the seek operation succeeded. Flushing seeks will trigger a preroll, which will emit GST_MESSAGE_ASYNC_DONE.
func (e *Element) SeekSimple(format Format, flags SeekFlags, pos int64) bool {
	return C.gst_element_seek_simple(e.g(), *(format.g()), flags.g(), (C.gint64)(pos)) == 1
}

//Sends a seek event to an element. See gst_event_new_seek() for the details of the parameters.
//The seek event is sent to the element using gst_element_send_event().
//MT safe.
//Parameters
//rate
//The new playback rate
//format
//The format of the seek values
//flags
//The optional seek flags.
//start_type
//The type and flags for the new start position
//start
//The value of the new start position
//stop_type
//The type and flags for the new stop position
//stop
//The value of the new stop position
//Returns
//TRUE if the event was handled. Flushing seeks will trigger a preroll, which will emit GST_MESSAGE_ASYNC_DONE.
func (e *Element) Seek(rate float64, format Format, flags SeekFlags, startType SeekType, start int64, stopType SeekType, stop int64) bool {
	return C.gst_element_seek(e.g(), C.gdouble(rate), C.GstFormat(format), flags.g(), C.GstSeekType(startType), C.gint64(start), C.GstSeekType(stopType), C.gint64(stop)) == 1
}
