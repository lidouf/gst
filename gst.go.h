#include <glib-object.h>
#include <gst/video/videooverlay.h>

// Type Casting
static GstVideoOverlay *
toGstVideoOverlay(gpointer p) {
    return (GST_VIDEO_OVERLAY(p));
}