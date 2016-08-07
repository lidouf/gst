package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
)

type Structure C.GstStructure

func (s *Structure) g() *C.GstStructure {
	return (*C.GstStructure)(s)
}

func (s *Structure) GetName() string {
	return C.GoString((*C.char)(C.gst_structure_get_name(s.g())))
}

func (s *Structure) Serialize() glib.Params {
	return serializeGstStructure(s.g())
}

func (s *Structure) ToString() string {
	return C.GoString((*C.char)(C.gst_structure_to_string(s.g())))
}

func MakeStructure(name string, fields *glib.Params) *Structure {
	var f glib.Params
	if fields == nil {
		f = glib.Params{}
	} else {
		f = *fields
	}
	return (*Structure)(makeGstStructure(name, f))
}
