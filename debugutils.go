package gst

/*
#include <gst/gst.h>
*/
import "C"

type DebugGraphDetails C.GstDebugGraphDetails

const (
	DEBUG_GRAPH_SHOW_MEDIA_TYPE         = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_MEDIA_TYPE)
	DEBUG_GRAPH_SHOW_CAPS_DETAILS       = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_CAPS_DETAILS)
	DEBUG_GRAPH_SHOW_NON_DEFAULT_PARAMS = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_NON_DEFAULT_PARAMS)
	DEBUG_GRAPH_SHOW_STATES             = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_STATES)
	DEBUG_GRAPH_SHOW_FULL_PARAMS        = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_FULL_PARAMS)
	DEBUG_GRAPH_SHOW_ALL                = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_ALL)
	DEBUG_GRAPH_SHOW_VERBOSE            = DebugGraphDetails(C.GST_DEBUG_GRAPH_SHOW_VERBOSE)
)