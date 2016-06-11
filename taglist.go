package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
)

type TagFlag C.GstTagFlag

const (
	GST_TAG_FLAG_UNDEFINED = TagFlag(C.GST_TAG_FLAG_UNDEFINED)
	GST_TAG_FLAG_META      = TagFlag(C.GST_TAG_FLAG_META)
	GST_TAG_FLAG_ENCODED   = TagFlag(C.GST_TAG_FLAG_ENCODED)
	GST_TAG_FLAG_DECODED   = TagFlag(C.GST_TAG_FLAG_DECODED)
	GST_TAG_FLAG_COUNT     = TagFlag(C.GST_TAG_FLAG_COUNT)
)

func (t *TagFlag) g() *C.GstTagFlag {
	return (*C.GstTagFlag)(t)
}

type TagList C.GstTagList

func (t *TagList) g() *C.GstTagList {
	return (*C.GstTagList)(t)
}

func (t *TagList) Type() glib.Type {
	return glib.TypeFromName("GstTagList")
}
