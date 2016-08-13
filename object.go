package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
static inline
guint32 CALL_MACRO_GST_OBJECT_FLAGS(GstObject* obj) {
	return GST_OBJECT_FLAGS(obj);
}
static inline
gboolean CALL_MACRO_GST_OBJECT_FLAG_IS_SET(GstObject* obj, guint32 flag) {
	return GST_OBJECT_FLAG_IS_SET(obj, flag);
}
static inline
void CALL_MACRO_GST_OBJECT_FLAG_SET(GstObject* obj, guint32 flag) {
	GST_OBJECT_FLAG_SET(obj, flag);
}
static inline
void CALL_MACRO_GST_OBJECT_FLAG_UNSET(GstObject* obj, guint32 flag) {
	GST_OBJECT_FLAG_UNSET(obj, flag);
}
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type GstObj struct {
	glib.Object
}

func (o *GstObj) g() *C.GstObject {
	return (*C.GstObject)(o.GetPtr())
}

func (o *GstObj) AsGstObj() *GstObj {
	return o
}

// This macro returns the entire set of flags for the object.
func (o *GstObj) Flags() uint32 {
	return uint32(C.CALL_MACRO_GST_OBJECT_FLAGS(o.g()))
}

//This macro checks to see if the given flag is set.
func (o *GstObj) FlagIsSet(flag uint32) bool {
	return C.CALL_MACRO_GST_OBJECT_FLAG_IS_SET(o.g(), C.guint32(flag)) != 0
}

//This macro sets the given bits.
func (o *GstObj) FlagSet(flag uint32) {
	C.CALL_MACRO_GST_OBJECT_FLAG_SET(o.g(), C.guint32(flag))
}

//This macro unsets the given bits.
func (o *GstObj) FlagUnset(flag uint32) {
	C.CALL_MACRO_GST_OBJECT_FLAG_UNSET(o.g(), C.guint32(flag))
}

// Sets the name of object.
// Returns true if the name could be set. MT safe.
func (o *GstObj) SetName(name string) bool {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	return C.gst_object_set_name(o.g(), (*C.gchar)(s)) != 0
}

// MT safe.
func (o *GstObj) GetName() string {
	s := C.gst_object_get_name(o.g())
	if s == nil {
		return ""
	}
	defer C.g_free(C.gpointer(s))
	return C.GoString((*C.char)(s))
}

// Sets the parent of o to p. This function causes the parent-set signal to be
// emitted when the parent was successfully set.
func (o *GstObj) SetParent(p *GstObj) bool {
	return C.gst_object_set_parent(o.g(), p.g()) != 0
}

// Returns the parent of o. Increases the refcount of the parent object so you
// should Unref it after usage.
func (o *GstObj) GetParent() *GstObj {
	p := new(GstObj)
	p.SetPtr(glib.Pointer(C.gst_object_get_parent(o.g())))
	return p
}

/**
Check if parent is the parent of object . E.g. a GstElement can check if it owns a given GstPad.
Returns
FALSE if either object or parent is NULL. TRUE if parent is the parent of object . Otherwise FALSE.
MT safe. Grabs and releases object 's locks.
Since: 1.6
*/
func (o *GstObj) HasAsParent(parent *GstObj) bool {
	return C.gst_object_has_as_parent(o.g(), parent.g()) != 0
}

// Clear the parent of object, removing the associated reference. This function
// decreases the refcount of o. MT safe. Grabs and releases object's lock.
func (o *GstObj) Unparent() {
	C.gst_object_unparent(o.g())
}

/**
Checks to see if there is any object named name in list .
This function does not do any locking of any kind. You might want to protect the provided list with the lock of the owner of the list.
This function will lock each GstObject in the list to compare the name, so be careful when passing a list with a locked object.
Returns
TRUE if a GstObject named name does not appear in list , FALSE if it does.
MT safe. Grabs and releases the LOCK of each object in the list.
*/
func CheckUniqueness(list *glib.List, name string) bool {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	return C.gst_object_check_uniqueness((*C.struct__GList)(unsafe.Pointer(list.Native())), s) != 0
}

/**
Check if object has an ancestor ancestor somewhere up in the hierarchy. One can e.g. check if a GstElement is inside a GstPipeline.
Returns
TRUE if ancestor is an ancestor of object .
MT safe. Grabs and releases object 's locks.
*/
func (o *GstObj) HasAsAncestor(ancestor *GstObj) bool {
	return C.gst_object_has_as_ancestor(o.g(), ancestor.g()) != 0
}

/**
@deprecated
gst_object_has_ancestor is deprecated and should not be used in newly-written code.
Use gst_object_has_as_ancestor() instead.
MT safe. Grabs and releases object 's locks.
Check if object has an ancestor ancestor somewhere up in the hierarchy. One can e.g. check if a GstElement is inside a GstPipeline.
Returns
TRUE if ancestor is an ancestor of object .
*/
func (o *GstObj) HasAncestor(ancestor *GstObj) bool {
	return C.gst_object_has_ancestor(o.g(), ancestor.g()) != 0
}

/**
Increments the reference count on object . This function does not take the lock on object because it relies on atomic refcounting.
This object returns the input parameter to ease writing constructs like : result = gst_object_ref (object->parent);
Returns
A pointer to object .
*/
func (o *GstObj) Ref(obj *GstObj) *GstObj {
	r := new(GstObj)
	r.SetPtr(glib.Pointer(C.gst_object_ref(o.g())))
	return r
}

/**
Decrements the reference count on object . If reference count hits zero, destroy object .
This function does not take the lock on object as it relies on atomic refcounting.
The unref method should never be called with the LOCK held since this might deadlock the dispose function.
*/
func (o *GstObj) Unref() {
	C.gst_object_unref(o.g())
}

/**
Increase the reference count of object , and possibly remove the floating reference, if object has a floating reference.
In other words, if the object is floating, then this call "assumes ownership" of the floating reference,
converting it to a normal reference by clearing the floating flag while leaving the reference count unchanged.
If the object is not floating, then this call adds a new normal reference increasing the reference count by one.
*/
func (o *GstObj) RefSink(obj *GstObj) *GstObj {
	r := new(GstObj)
	r.SetPtr(glib.Pointer(C.gst_object_ref_sink(o.g())))
	return r
}

/**
Atomically modifies a pointer to point to a new object. The reference count of oldobj is decreased and the reference count of newobj is increased.
Either newobj and the value pointed to by oldobj may be NULL.
Returns
TRUE if newobj was different from oldobj
*/
func (o *GstObj) Replace(newObj *GstObj) bool {
	oAddr := o.g()
	return C.gst_object_replace(&oAddr, newObj.g()) != 0
}

// Generates a string describing the path of object in the object hierarchy.
// Only useful (or used) for debugging.
//Returns
//a string describing the path of object . You must g_free() the string after usage.
//MT safe. Grabs and releases the GstObject's LOCK for all objects in the hierarchy.
func (o *GstObj) GetPathString() string {
	s := C.gst_object_get_path_string(o.g())
	defer C.g_free(C.gpointer(s))
	return C.GoString((*C.char)(s))
}

/**
Returns a suggestion for timestamps where buffers should be split to get best controller results.
Returns
Returns the suggested timestamp or GST_CLOCK_TIME_NONE if no control-rate was set.
*/
func (o *GstObj) SuggestNextSync() ClockTime {
	return ClockTime(C.gst_object_suggest_next_sync(o.g()))
}

/**
Sets the properties of the object, according to the GstControlSources that (maybe) handle them and for the given timestamp.
If this function fails, it is most likely the application developers fault. Most probably the control sources are not setup correctly.
Parameters
Returns
TRUE if the controller values could be applied to the object properties, FALSE otherwise
*/
func (o *GstObj) SyncValues(t ClockTime) bool {
	return C.gst_object_sync_values(o.g(), C.GstClockTime(t)) != 0
}

/**
Check if the object has an active controlled properties.
Returns
TRUE if the object has active controlled properties
*/
func (o *GstObj) HasActiveControlBindings() bool {
	return C.gst_object_has_active_control_bindings(o.g()) != 0
}

/**
This function is used to disable all controlled properties of the object for some time, i.e. gst_object_sync_values() will do nothing.
*/
func (o *GstObj) SetControlBindingsDisabled(disabled bool) {
	C.gst_object_set_control_bindings_disabled(o.g(), gBoolean(disabled))
}

/**
This function is used to disable the control bindings on a property for some time, i.e. gst_object_sync_values() will do nothing for the property.
*/
func (o *GstObj) SetControlBindingDisabled(propertyName string, disabled bool) {
	s := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(s))
	C.gst_object_set_control_binding_disabled(o.g(), s, gBoolean(disabled))
}

/**
Obtain the control-rate for this object .
Audio processing GstElement objects will use this rate to sub-divide their processing loop and call gst_object_sync_values() inbetween.
The length of the processing segment should be up to control -rate nanoseconds.
If the object is not under property control, this will return GST_CLOCK_TIME_NONE. This allows the element to avoid the sub-dividing.
The control-rate is not expected to change if the element is in GST_STATE_PAUSED or GST_STATE_PLAYING.
Returns
the control rate in nanoseconds
*/
func (o *GstObj) GetControlRate() ClockTime {
	return ClockTime(C.gst_object_get_control_rate(o.g()))
}

/**
Change the control-rate for this object .
Audio processing GstElement objects will use this rate to sub-divide their processing loop and call gst_object_sync_values() inbetween.
The length of the processing segment should be up to control -rate nanoseconds.
The control-rate should not change if the element is in GST_STATE_PAUSED or GST_STATE_PLAYING.
*/
func (o *GstObj) SetControlRate(t ClockTime) {
	C.gst_object_set_control_rate(o.g(), C.GstClockTime(t))
}

//func (o *GstObj) GetGstProperty(name string, value interface{}) {
//	s := C.CString(name)
//	defer C.free(unsafe.Pointer(s))
//	glib.TypeOf(obj)
//	C.g_object_get_property(o.g(), (*C.gchar)(s), v.g())
//	return v.Get()
//}

/*func (o *GstObj) ImplementsInterfaceCheck(typ glib.Type) bool {
	return C.gst_implements_interface_check(C.gpointer(o.GetPtr()),
		C.GType(typ)) != 0
}

func (o *GstObj) ImplementsInterfaceCast(typ glib.Type) glib.Pointer {
	return glib.Pointer(C.gst_implements_interface_cast(C.gpointer(o.GetPtr()),
		C.GType(typ)))
}*/

type GstObjClass struct {
	glib.Object
}

func (o *GstObjClass) g() *C.GstObjectClass {
	return (*C.GstObjectClass)(o.GetPtr())
}

func (o *GstObjClass) AsGstObjClass() *GstObjClass {
	return o
}
