package gst

/*
#include <gst/gst.h>
static gboolean clock_time_is_valid(gint64 t) {
	return GST_CLOCK_TIME_IS_VALID(t);
}
static gchar* clock_duration_output(GstClockTime t) {
	return g_strdup_printf("%" GST_TIME_FORMAT, GST_TIME_ARGS(t));
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

type ClockTime C.GstClockTime

func (t *ClockTime) g() *C.GstClockTime {
	return (*C.GstClockTime)(t)
}

func (t ClockTime) AsUint64() uint64 {
	return uint64(t)
}

func (t ClockTime) DurationOutput() string {
	return C.GoString((*C.char)(C.clock_duration_output(C.GstClockTime(t))))
}

func ClockTimeIsValid(t int64) bool {
	return C.clock_time_is_valid((C.gint64)(t)) == 1
}
