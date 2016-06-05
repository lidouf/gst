package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"
)

type QueryTypeFlags C.GstQueryTypeFlags

const (
	QUERY_TYPE_UPSTREAM = 1 << iota
	QUERY_TYPE_DOWNSTREAM
	QUERY_TYPE_SERIALIZED
)

type Query C.GstQuery

func (q *Query) g() *C.GstQuery {
	return (*C.GstQuery)(q)
}

func (q *Query) AsQuery() *Query {
	return q
}

func (q *Query) ParseSeeking(format *Format) (bool, int64, int64) {
	seekable := new(C.gboolean)
	defer C.free(unsafe.Pointer(seekable))
	start := new(C.gint64)
	defer C.free(unsafe.Pointer(start))
	end := new(C.gint64)
	defer C.free(unsafe.Pointer(end))
	C.gst_query_parse_seeking(q.g(), format.g(), seekable, start, end)
	return *seekable == 1, (int64)(*start), (int64)(*end)
}

func (q *Query) Unref() {
	C.gst_query_unref((*C.GstQuery)(q))
}

func NewQuerySeeking(format Format) *Query {
	seeking := C.gst_query_new_seeking(*(format.g()))
	if seeking == nil {
		return nil
	}
	return (*Query)(seeking)
}
