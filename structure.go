package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type Structure C.GstStructure

func (s *Structure) g() *C.GstStructure {
	return (*C.GstStructure)(s)
}

func (s *Structure) GetName() string {
	return C.GoString((*C.char)(C.gst_structure_get_name(s.g())))
}
