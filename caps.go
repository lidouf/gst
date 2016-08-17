package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
#include <gst/pbutils/pbutils.h>

int capsRefCount(GstCaps *c) {
	return GST_CAPS_REFCOUNT(c);
}
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

//type Caps C.GstCaps
type Caps struct {
	glib.Object
}

func (c *Caps) Type() glib.Type {
	return glib.TypeFromName("GstCaps")
}

func (c *Caps) Value() *glib.Value {
	v := glib.NewValue(c.Type())
	C.gst_value_set_caps(v2g(v), c.g())
	return v
}

func (c *Caps) g() *C.GstCaps {
	return (*C.GstCaps)(c.GetPtr())
}

func (c *Caps) Ref() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_caps_ref(c.g())))
	return r
}

func (c *Caps) Unref() {
	C.gst_caps_unref(c.g())
}

func (c *Caps) RefCount() int {
	return int(C.capsRefCount(c.g()))
}

func (c *Caps) AppendStructure(media_type string, fields glib.Params) {
	C.gst_caps_append_structure(c.g(), makeGstStructure(media_type, fields))
}

func (c *Caps) GetSize() int {
	return int(C.gst_caps_get_size(c.g()))
}

func (c *Caps) GetStructure(index uint) *Structure {
	r := new(Structure)
	r.SetPtr(glib.Pointer(C.gst_caps_get_structure(c.g(), C.guint(index))))
	return r
}

func (c *Caps) String() string {
	s := (*C.char)(C.gst_caps_to_string(c.g()))
	defer C.free(unsafe.Pointer(s))
	return C.GoString(s)
}

func (c *Caps) IsAny() bool {
	return C.gst_caps_is_any(c.g()) != 0
}

func (c *Caps) IsEmpty() bool {
	return C.gst_caps_is_empty(c.g()) != 0
}

func (c *Caps) IsFixed() bool {
	return C.gst_caps_is_fixed(c.g()) != 0
}

func (c *Caps) GetCodecDescription() string {
	return C.GoString((*C.char)(C.gst_pb_utils_get_codec_description(c.g())))
}

func NewCapsAny() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_caps_new_any()))
	return r
}

func NewCapsEmpty() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_caps_new_empty()))
	return r
}

func NewCapsSimple(media_type string, fields glib.Params) *Caps {
	c := NewCapsEmpty()
	c.AppendStructure(media_type, fields)
	return c
}

func CapsFromString(s string) *Caps {
	cs := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cs))
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_caps_from_string(cs)))
	return r
}

//type StaticCaps C.GstStaticCaps
type StaticCaps struct {
	glib.Object
}

func (c *StaticCaps) g() *C.GstStaticCaps {
	return (*C.GstStaticCaps)(c.GetPtr())
}

func (c *StaticCaps) Caps() *Caps {
	if c.g().caps == nil {
		return nil
	}
	r := new(Caps)
	r.SetPtr(glib.Pointer(c.g().caps))
	return r
}

func (c *StaticCaps) Get() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_static_caps_get(c.g())))
	return r
}

func (c *StaticCaps) String() string {
	return C.GoString(c.g().string)
}
