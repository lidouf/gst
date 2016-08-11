package gst

/*
#include <gst/gst.h>
*/
import "C"

type SeekFlags C.GstSeekFlags

const (
	SEEK_FLAG_NONE      = 0
	SEEK_FLAG_FLUSH     = 1 << 0
	SEEK_FLAG_ACCURATE  = 1 << 1
	SEEK_FLAG_KEY_UNIT  = 1 << 2
	SEEK_FLAG_SEGMENT   = 1 << 3
	SEEK_FLAG_TRICKMODE = 1 << 4
	/* FIXME 2.0: Remove _SKIP flag,
	 * which was kept for backward compat when _TRICKMODE was added */
	SEEK_FLAG_SKIP         = 1 << 4
	SEEK_FLAG_SNAP_BEFORE  = 1 << 5
	SEEK_FLAG_SNAP_AFTER   = 1 << 6
	SEEK_FLAG_SNAP_NEAREST = SEEK_FLAG_SNAP_BEFORE | SEEK_FLAG_SNAP_AFTER
	/* Careful to restart next flag with 1<<7 here */
	SEEK_FLAG_TRICKMODE_KEY_UNITS = 1 << 7
	SEEK_FLAG_TRICKMODE_NO_AUDIO  = 1 << 8
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
