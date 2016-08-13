package gst

/*
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
)

type ControlBinding struct {
	GstObj
}

func (d *ControlBinding) Type() glib.Type {
	return glib.TypeFromName("GstControlBinding")
}

func (c *ControlBinding) g() *C.GstControlBinding {
	return (*C.GstControlBinding)(c.GetPtr())
}

func (c *ControlBinding) AsControlBinding() *ControlBinding {
	return c
}

//Sets the property of the object , according to the GstControlSources that handle them and for the given timestamp.
//If this function fails, it is most likely the application developers fault. Most probably the control sources are not setup correctly.
//Returns
//TRUE if the controller value could be applied to the object property, FALSE otherwise
func (c *ControlBinding) SyncValues(o *GstObj, timestamp, lastSync ClockTime) bool {
	return C.gst_control_binding_sync_values(c.g(), o.g(), C.GstClockTime(timestamp), C.GstClockTime(lastSync)) != 0
}

//This function is used to disable a control binding for some time, i.e. gst_object_sync_values() will do nothing.
func (c *ControlBinding) SetDisabled(disabled bool) {
	C.gst_control_binding_set_disabled(c.g(), gBoolean(disabled))
}

//Check if the control binding is disabled.
//Returns
//TRUE if the binding is inactive
func (o *ControlBinding) IsDisabled() bool {
	return C.gst_control_binding_is_disabled(o.g()) != 0
}
