package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
//"github.com/lidouf/glib"
//"unsafe"
)

type StaticPadTemplate C.GstStaticPadTemplate

func (t *StaticPadTemplate) g() *C.GstStaticPadTemplate {
	return (*C.GstStaticPadTemplate)(t)
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
