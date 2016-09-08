//GstQuery — Provide functions to create queries, and to set and parse values in them.
package gst

/*
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
	"time"
)

type QueryTypeFlags C.GstQueryTypeFlags

const (
	QUERY_TYPE_UPSTREAM = 1 << iota
	QUERY_TYPE_DOWNSTREAM
	QUERY_TYPE_SERIALIZED
)

//type Query C.GstQuery
type Query struct {
	glib.Object
}

func (q *Query) g() *C.GstQuery {
	return (*C.GstQuery)(q.GetPtr())
}

func (q *Query) AsQuery() *Query {
	return q
}

func (q *Query) ParseSeeking(format *Format) (bool, time.Duration, time.Duration) {
	var seekable C.gboolean
	var start, end C.gint64
	if format == nil {
		C.gst_query_parse_seeking(q.g(), nil, &seekable, &start, &end)
	} else {
		C.gst_query_parse_seeking(q.g(), format.g(), &seekable, &start, &end)
	}

	return seekable == 1, (time.Duration)(start), (time.Duration)(end)
}

func (q *Query) GetNBufferingRanges() int {
	return int(C.gst_query_get_n_buffering_ranges(q.g()))
}

func (q *Query) ParseNthBufferingRange(r uint) (int64, int64) {
	var start, stop C.gint64
	C.gst_query_parse_nth_buffering_range(q.g(), C.guint(r), &start, &stop)
	return int64(start), int64(stop)
}

func (q *Query) Unref() {
	C.gst_query_unref(q.g())
}

func NewQuerySeeking(format Format) *Query {
	seeking := C.gst_query_new_seeking(*(format.g()))
	if seeking == nil {
		return nil
	}
	q := new(Query)
	q.SetPtr(glib.Pointer(seeking))
	return q
}

func NewQueryBuffering(format Format) *Query {
	buffering := C.gst_query_new_buffering(*(format.g()))
	if buffering == nil {
		return nil
	}
	q := new(Query)
	q.SetPtr(glib.Pointer(buffering))
	return q
}
