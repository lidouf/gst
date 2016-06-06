package gst

/*
#include <gst/gst.h>
*/
import "C"

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
	var seekable C.gboolean
	var start, end C.gint64
	if format == nil {
		C.gst_query_parse_seeking(q.g(), nil, &seekable, &start, &end)
	} else {
		C.gst_query_parse_seeking(q.g(), format.g(), &seekable, &start, &end)
	}

	return seekable == 1, (int64)(start), (int64)(end)
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
