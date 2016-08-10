package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type Pipeline struct {
	Bin
}

func (p *Pipeline) g() *C.GstPipeline {
	return (*C.GstPipeline)(p.GetPtr())
}

func (p *Pipeline) AsPipeline() *Pipeline {
	return p
}

func (p *Pipeline) GetBus() *Bus {
	bus := C.gst_pipeline_get_bus(p.g())
	if bus == nil {
		return nil
	}
	b := new(Bus)
	b.SetPtr(glib.Pointer(bus))
	return b
}

func (p *Pipeline) SetClock(clock *Clock) bool {
	return C.gst_pipeline_set_clock(p.g(), clock.g()) != 0
}

func (p *Pipeline) GetPipelineClock() *Clock {
	return (*Clock)(unsafe.Pointer(C.gst_pipeline_get_pipeline_clock(p.g())))
}

func (p *Pipeline) GetClock() *Clock {
	return (*Clock)(unsafe.Pointer(C.gst_pipeline_get_clock(p.g())))
}

func (p *Pipeline) UseClock(clock *Clock) {
	C.gst_pipeline_use_clock(p.g(), clock.g())
}

func (p *Pipeline) AutoClock() {
	C.gst_pipeline_auto_clock(p.g())
}

func (p *Pipeline) SetAutoFlushBus(autoFlush bool) {
	var af int
	if autoFlush {
		af = 1
	} else {
		af = 0
	}
	C.gst_pipeline_set_auto_flush_bus(p.g(), C.gboolean(af))
}

func (p *Pipeline) GetAutoFlushBus() bool {
	return C.gst_pipeline_get_auto_flush_bus(p.g()) != 0
}

func (p *Pipeline) SetDelay(delay ClockTime) {
	C.gst_pipeline_set_delay(p.g(), C.GstClockTime(delay))
}

func (p *Pipeline) GetDelay() ClockTime {
	return ClockTime(uint64(C.gst_pipeline_get_delay(p.g())))
}

func (p *Pipeline) SetLatency(latency ClockTime) {
	C.gst_pipeline_set_latency(p.g(), C.GstClockTime(latency))
}

func (p *Pipeline) GetLatency() ClockTime {
	return ClockTime(uint64(C.gst_pipeline_get_delay(p.g())))
}

func NewPipeline(name string) *Pipeline {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(Pipeline)
	p.SetPtr(glib.Pointer(C.gst_pipeline_new(s)))
	return p
}

func ParseLaunch(pipeline_description string) (*Pipeline, error) {
	pd := (*C.gchar)(C.CString(pipeline_description))
	defer C.free(unsafe.Pointer(pd))
	p := new(Pipeline)
	var Cerr *C.GError
	p.SetPtr(glib.Pointer(C.gst_parse_launch(pd, &Cerr)))
	if Cerr != nil {
		err := *(*glib.Error)(unsafe.Pointer(Cerr))
		C.g_error_free(Cerr)
		return nil, &err
	}
	return p, nil
}
