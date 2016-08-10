package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
//static gboolean pad_template_is_fixed(GstPadTemplate *templ) {
//	return GST_PAD_TEMPLATE_IS_FIXED(templ);
//}
*/
import "C"

import (
	"unsafe"
)

type PadTemplate C.GstPadTemplate

func (t *PadTemplate) g() *C.GstPadTemplate {
	return (*C.GstPadTemplate)(t)
}

func (t *PadTemplate) GetCaps() *Caps {
	return (*Caps)(unsafe.Pointer(C.gst_pad_template_get_caps(t.g())))
}

//func (t *PadTemplate) IsFixed() bool {
//	return C.pad_template_is_fixed(t.g()) != 0
//}

func NewPadTemplate(name string, direction PadDirection, presence PadPresence, caps *Caps) *PadTemplate {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	return (*PadTemplate)(unsafe.Pointer(C.gst_pad_template_new(s, direction.g(), presence.g(), caps.g())))
}

type StaticPadTemplate C.GstStaticPadTemplate

func (t *StaticPadTemplate) g() *C.GstStaticPadTemplate {
	return (*C.GstStaticPadTemplate)(t)
}

func (t *StaticPadTemplate) Get() *PadTemplate {
	return (*PadTemplate)(unsafe.Pointer(C.gst_static_pad_template_get(t.g())))
}

func (t *StaticPadTemplate) GetCaps() *Caps {
	return (*Caps)(unsafe.Pointer(C.gst_static_pad_template_get_caps(t.g())))
}

func (t *StaticPadTemplate) GetNameTemplate() string {
	return C.GoString((*C.char)(t.g().name_template))
}

func (t *StaticPadTemplate) GetDirection() PadDirection {
	return PadDirection(t.g().direction)
}

func (t *StaticPadTemplate) GetPresence() PadPresence {
	return PadPresence(t.g().presence)
}

func (t *StaticPadTemplate) GetStaticCaps() *StaticCaps {
	return (*StaticCaps)(&t.g().static_caps)
}
