package gst

/*
#include <gst/gst.h>
*/
import "C"

type Format C.GstFormat

const (
	FORMAT_UNDEFINED = Format(C.GST_FORMAT_UNDEFINED)
	FORMAT_DEFAULT   = Format(C.GST_FORMAT_DEFAULT)
	FORMAT_BYTES     = Format(C.GST_FORMAT_BYTES)
	FORMAT_TIME      = Format(C.GST_FORMAT_TIME)
	FORMAT_BUFFERS   = Format(C.GST_FORMAT_BUFFERS)
	FORMAT_PERCENT   = Format(C.GST_FORMAT_PERCENT)
)

func (f Format) String() string {
	switch f {
	case FORMAT_UNDEFINED:
		return "FORMAT_UNDEFINED"
	case FORMAT_DEFAULT:
		return "FORMAT_DEFAULT"
	case FORMAT_BYTES:
		return "FORMAT_BYTES"
	case FORMAT_TIME:
		return "FORMAT_TIME"
	case FORMAT_BUFFERS:
		return "FORMAT_BUFFERS"
	case FORMAT_PERCENT:
		return "FORMAT_PERCENT"
	}
	panic("Unknown value of gst.Format")
}

func (f *Format) g() *C.GstFormat {
	return (*C.GstFormat)(f)
}
