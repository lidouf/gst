package gst

/*
#include <gst/gst.h>
static gboolean clock_time_is_valid(gint64 t) {
	return GST_CLOCK_TIME_IS_VALID(t);
}
*/
import "C"

type Clock struct {
	GstObj
}

func (c *Clock) g() *C.GstClock {
	return (*C.GstClock)(c.GetPtr())
}

func (c *Clock) AsClock() *Clock {
	return c
}

func ClockTimeIsValid(t int64) bool {
	return C.clock_time_is_valid((C.gint64)(t)) == 1
}
