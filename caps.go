package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

int capsRefCount(GstCaps *c) {
	return GST_CAPS_REFCOUNT(c);
}
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type Caps C.GstCaps

func (c *Caps) g() *C.GstCaps {
	return (*C.GstCaps)(c)
}

func (c *Caps) Ref() *Caps {
	return (*Caps)(C.gst_caps_ref(c.g()))
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
	return (*Structure)(C.gst_caps_get_structure(c.g(), C.guint(index)))
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

func NewCapsAny() *Caps {
	return (*Caps)(C.gst_caps_new_any())
}

func NewCapsEmpty() *Caps {
	return (*Caps)(C.gst_caps_new_empty())
}

func NewCapsSimple(media_type string, fields glib.Params) *Caps {
	c := NewCapsEmpty()
	c.AppendStructure(media_type, fields)
	return c
}

func CapsFromString(s string) *Caps {
	cs := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cs))
	return (*Caps)(C.gst_caps_from_string(cs))
}

type StaticCaps C.GstStaticCaps

func (c *StaticCaps) g() *C.GstStaticCaps {
	return (*C.GstStaticCaps)(c)
}

func (c *StaticCaps) Caps() *Caps {
	if c.g().caps == nil {
		return nil
	}
	return (*Caps)(c.g().caps)
}

func (c *StaticCaps) Get() *Caps {
	return (*Caps)(C.gst_static_caps_get(c.g()))
}

func (c *StaticCaps) String() string {
	return C.GoString(c.g().string)
}
