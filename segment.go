package gst

/*
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
)

type SeekFlags C.GstSeekFlags

const (
	SEEK_FLAG_NONE      = SeekFlags(C.GST_SEEK_FLAG_NONE)
	SEEK_FLAG_FLUSH     = SeekFlags(C.GST_SEEK_FLAG_FLUSH)
	SEEK_FLAG_ACCURATE  = SeekFlags(C.GST_SEEK_FLAG_ACCURATE)
	SEEK_FLAG_KEY_UNIT  = SeekFlags(C.GST_SEEK_FLAG_KEY_UNIT)
	SEEK_FLAG_SEGMENT   = SeekFlags(C.GST_SEEK_FLAG_SEGMENT)
	SEEK_FLAG_TRICKMODE = SeekFlags(C.GST_SEEK_FLAG_TRICKMODE)
	/* FIXME 2.0: Remove _SKIP flag,
	 * which was kept for backward compat when _TRICKMODE was added */
	SEEK_FLAG_SKIP         = SeekFlags(C.GST_SEEK_FLAG_SKIP)
	SEEK_FLAG_SNAP_BEFORE  = SeekFlags(C.GST_SEEK_FLAG_SNAP_BEFORE)
	SEEK_FLAG_SNAP_AFTER   = SeekFlags(C.GST_SEEK_FLAG_SNAP_AFTER)
	SEEK_FLAG_SNAP_NEAREST = SeekFlags(C.GST_SEEK_FLAG_SNAP_NEAREST)
	/* Careful to restart next flag with 1<<7 here */
	SEEK_FLAG_TRICKMODE_KEY_UNITS = SeekFlags(C.GST_SEEK_FLAG_TRICKMODE_KEY_UNITS)
	SEEK_FLAG_TRICKMODE_NO_AUDIO  = SeekFlags(C.GST_SEEK_FLAG_TRICKMODE_NO_AUDIO)
)

func (s SeekFlags) g() C.GstSeekFlags {
	return (C.GstSeekFlags)(s)
}

type SeekType C.GstSeekType

const (
	SEEK_TYPE_NONE = SeekType(C.GST_SEEK_TYPE_NONE)
	SEEK_TYPE_SET  = SeekType(C.GST_SEEK_TYPE_SET)
	SEEK_TYPE_END  = SeekType(C.GST_SEEK_TYPE_END)
)

type Segment struct {
	glib.Object
}

func (s *Segment) Type() glib.Type {
	return glib.TypeFromName("GstSegment")
}

func (s *Segment) g() *C.GstSegment {
	return (*C.GstSegment)(s.GetPtr())
}
