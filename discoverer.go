package gst

/*
#cgo pkg-config: gstreamer-pbutils-1.0
#include <stdlib.h>
#include <gst/gst.h>
#include <gst/pbutils/pbutils.h>

static gboolean is_discoverer_container_info(GstDiscovererStreamInfo* info) {
	return GST_IS_DISCOVERER_CONTAINER_INFO(info);
}
static GstDiscovererContainerInfo* to_discoverer_container_info(GstDiscovererStreamInfo* info) {
	return GST_DISCOVERER_CONTAINER_INFO(info);
}
*/
import "C"

import (
	"github.com/lidouf/glib"

	"errors"
	"time"
	"unsafe"
)

type DiscovererResult C.GstDiscovererResult

const (
	DISCOVERER_OK              = DiscovererResult(C.GST_DISCOVERER_OK)
	DISCOVERER_URI_INVALID     = DiscovererResult(C.GST_DISCOVERER_URI_INVALID)
	DISCOVERER_ERROR           = DiscovererResult(C.GST_DISCOVERER_ERROR)
	DISCOVERER_TIMEOUT         = DiscovererResult(C.GST_DISCOVERER_TIMEOUT)
	DISCOVERER_BUSY            = DiscovererResult(C.GST_DISCOVERER_BUSY)
	DISCOVERER_MISSING_PLUGINS = DiscovererResult(C.GST_DISCOVERER_MISSING_PLUGINS)
)

func (d *DiscovererResult) Type() glib.Type {
	return glib.TypeFromName("GstDiscovererResult")
}

func (d *DiscovererResult) g() *C.GstDiscovererResult {
	return (*C.GstDiscovererResult)(d)
}

type DiscovererStreamInfo C.GstDiscovererStreamInfo

func (d *DiscovererStreamInfo) Type() glib.Type {
	return glib.TypeFromName("DiscovererStreamInfo")
}

func (d *DiscovererStreamInfo) g() *C.GstDiscovererStreamInfo {
	return (*C.GstDiscovererStreamInfo)(d)
}

func (d *DiscovererStreamInfo) GetCaps() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_discoverer_stream_info_get_caps(d.g())))
	return r
}

func (d *DiscovererStreamInfo) GetTags() *TagList {
	return (*TagList)(unsafe.Pointer(C.gst_discoverer_stream_info_get_tags(d.g())))
}

func (d *DiscovererStreamInfo) GetNext() *DiscovererStreamInfo {
	return (*DiscovererStreamInfo)(C.gst_discoverer_stream_info_get_next(d.g()))
}

func (d *DiscovererStreamInfo) GetPrevious() *DiscovererStreamInfo {
	return (*DiscovererStreamInfo)(C.gst_discoverer_stream_info_get_previous(d.g()))
}

func (d *DiscovererStreamInfo) GetStreamId() string {
	return C.GoString((*C.char)(C.gst_discoverer_stream_info_get_stream_id(d.g())))
}

func (d *DiscovererStreamInfo) GetStreamTypeNick() string {
	return C.GoString((*C.char)(C.gst_discoverer_stream_info_get_stream_type_nick(d.g())))
}

func (d *DiscovererStreamInfo) IsContainerInfo() bool {
	return C.is_discoverer_container_info(d.g()) != 0
}

func (d *DiscovererStreamInfo) AsContainerInfo() *DiscovererContainerInfo {
	return (*DiscovererContainerInfo)(unsafe.Pointer(C.to_discoverer_container_info(d.g())))
}

type DiscovererContainerInfo C.GstDiscovererContainerInfo

func (d *DiscovererContainerInfo) g() *C.GstDiscovererContainerInfo {
	return (*C.GstDiscovererContainerInfo)(d)
}

func DisCovererContainerInfoGetType() glib.Type {
	return glib.Type(C.gst_discoverer_container_info_get_type())
}

func (d *DiscovererContainerInfo) GetStreams() *glib.List {
	return glib.WrapList(uintptr(unsafe.Pointer(C.gst_discoverer_container_info_get_streams(d.g()))))
}

type DiscovererInfo C.GstDiscovererInfo

func (d *DiscovererInfo) Type() glib.Type {
	return glib.TypeFromName("GstDiscovererInfo")
}

func (d *DiscovererInfo) g() *C.GstDiscovererInfo {
	return (*C.GstDiscovererInfo)(d)
}

func (d *DiscovererInfo) GetUri() string {
	return C.GoString((*C.char)(C.gst_discoverer_info_get_uri(d.g())))
}

func (d *DiscovererInfo) GetResult() DiscovererResult {
	return DiscovererResult(C.gst_discoverer_info_get_result(d.g()))
}

func (d *DiscovererInfo) GetMisc() *Structure {
	r := new(Structure)
	r.SetPtr(glib.Pointer(C.gst_discoverer_info_get_misc(d.g())))
	return r
}

func (d *DiscovererInfo) GetDuration() ClockTime {
	return ClockTime(C.gst_discoverer_info_get_duration(d.g()))
}

func (d *DiscovererInfo) GetTags() *TagList {
	r := new(TagList)
	r.SetPtr(glib.Pointer(C.gst_discoverer_info_get_tags(d.g())))
	return r
}

func (d *DiscovererInfo) GetSeekable() bool {
	return C.gst_discoverer_info_get_seekable(d.g()) != 0
}

func (d *DiscovererInfo) GetStreamInfo() *DiscovererStreamInfo {
	return (*DiscovererStreamInfo)(C.gst_discoverer_info_get_stream_info(d.g()))
}

func (d *DiscovererInfo) GetStreamList() *glib.List {
	return glib.WrapList(uintptr(unsafe.Pointer(C.gst_discoverer_info_get_stream_list(d.g()))))
}

type Discoverer struct {
	GstObj
}

func (d *Discoverer) g() *C.GstDiscoverer {
	return (*C.GstDiscoverer)(d.GetPtr())
}

/* Asynchronous API */
func (d *Discoverer) Start() {
	C.gst_discoverer_start(d.g())
}

func (d *Discoverer) Stop() {
	C.gst_discoverer_stop(d.g())
}

func (d *Discoverer) DiscoverUriAsync(uri string) bool {
	s := (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(s))
	return C.gst_discoverer_discover_uri_async(d.g(), s) != 0
}

/* Synchronous API */
func (d *Discoverer) DiscoverUri(uri string) (*DiscovererInfo, error) {
	s := (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(s))
	var err *C.GError = nil

	res := C.gst_discoverer_discover_uri(d.g(), s, &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return (*DiscovererInfo)(res), nil
}

func NewDiscoverer(timeout time.Duration) (discoverer *Discoverer, err error) {
	var e, ret_e *C.GError
	cDiscoverer := C.gst_discoverer_new((C.GstClockTime)(uint64(timeout)), &e)
	discoverer = new(Discoverer)
	ret_e = new(C.GError)
	if e == nil && cDiscoverer != nil {
		discoverer.SetPtr(glib.Pointer(cDiscoverer))
		return
	}

	err = (*glib.Error)(unsafe.Pointer(ret_e))
	return
}
