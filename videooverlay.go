package gst

/*
#cgo pkg-config: gstreamer-video-1.0
#include <gst/gst.h>
#include "gst.go.h"
*/
import "C"

type VideoOverlay struct {
	GstObj
}

func (o *VideoOverlay) g() *C.GstVideoOverlay {
	return (*C.GstVideoOverlay)(o.GetPtr())
}

func (o *VideoOverlay) SetWindowHandle(handle uint32) {
	C.gst_video_overlay_set_window_handle(o.g(), (C.guintptr)(handle))
}
