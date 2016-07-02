// Bindings for GStreamer 1.0 API
package gst

/*
#include "gst.go.h"

#cgo pkg-config: gstreamer-1.0
*/
import "C"

import (
	//"fmt"
	"os"
	"unsafe"
	//"github.com/gotk3/gotk3/glib"
)

//var TYPE_FOURCC, TYPE_INT_RANGE, TYPE_FRACTION glib.Type

func Init() {
	alen := C.int(len(os.Args))
	argv := make([]*C.char, alen)
	for i, s := range os.Args {
		argv[i] = C.CString(s)
	}
	ret := C._gst_init(&alen, &argv[0])
	argv = (*[1 << 16]*C.char)(unsafe.Pointer(ret))[:alen]
	os.Args = make([]string, alen)
	for i, s := range argv {
		os.Args[i] = C.GoString(s)
	}
	//
	//TYPE_INT_RANGE = glib.Type(C.gst_int_range_get_type())
	//TYPE_FRACTION = glib.Type(C.gst_fraction_get_type())
}

func Version() (uint, uint, uint, uint) {
	var major, minor, micro, nano C.guint
	C.gst_version(&major, &minor, &micro, &nano)
	return uint(major), uint(minor), uint(micro), uint(nano)
}

//var CLOCK_TIME_NONE = int64(C.GST_CLOCK_TIME_NONE)
