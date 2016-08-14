// Bindings for GStreamer API
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

char** _gst_init(int* argc, char** argv) {
	gst_init(argc, &argv);
	return argv;
}

typedef struct {
	const char *name;
	const GValue *val;
} Field;

typedef struct {
	Field* tab;
	int    n;
} Fields;

gboolean _parse_field(GQuark id, const GValue* val, gpointer data) {
	Fields *f = (Fields*)(data);
	f->tab[f->n].name = g_quark_to_string(id);
	f->tab[f->n].val = val;
	++f->n;
	return TRUE;
}

Fields _parse_struct(GstStructure *s) {
	int n = gst_structure_n_fields(s);
	Fields f = { malloc(n * sizeof(Field)), 0 };
	gst_structure_foreach(s, _parse_field, (gpointer)(&f));
	return f;
}

#cgo pkg-config: gstreamer-1.0
*/
import "C"

import (
	"os"
	"unsafe"

	"github.com/lidouf/glib"
)

var TYPE_FOURCC, TYPE_INT_RANGE, TYPE_FRACTION glib.Type

func init() {
	alen := C.int(len(os.Args))
	argv := make([]*C.char, alen)
	for i, s := range os.Args {
		argv[i] = C.CString(s)
	}
	ret := C._gst_init(&alen, &argv[0])
	argv = (*[1 << 16]*C.char)(unsafe.Pointer(ret))[:alen]
	os.Args = make([]string, alen)
	for i, s := range argv {
		os.Args[i] = C.GoString(s)
	}

	TYPE_INT_RANGE = glib.Type(C.gst_int_range_get_type())
	TYPE_FRACTION = glib.Type(C.gst_fraction_get_type())
}

func makeGstStructure(name string, fields glib.Params) *C.GstStructure {
	nm := (*C.gchar)(C.CString(name))
	s := C.gst_structure_new_empty(nm)
	C.free(unsafe.Pointer(nm))
	for k, v := range fields {
		n := (*C.gchar)(C.CString(k))
		C.gst_structure_take_value(s, n, v2g(glib.ValueOf(v)))
		C.free(unsafe.Pointer(n))
	}
	return s
}

func parseGstStructure(s *C.GstStructure) (name string, fields glib.Params) {
	name = C.GoString((*C.char)(C.gst_structure_get_name(s)))
	ps := C._parse_struct(s)
	n := (int)(ps.n)
	tab := (*[1 << 16]C.Field)(unsafe.Pointer(ps.tab))[:n]
	fields = make(glib.Params)
	for _, f := range tab {
		fields[C.GoString(f.name)] = g2v(f.val).Get()
	}
	return
}

func serializeGstStructure(s *C.GstStructure) glib.Params {
	ps := C._parse_struct(s)
	n := (int)(ps.n)
	tab := (*[1 << 16]C.Field)(unsafe.Pointer(ps.tab))[:n]
	fields := make(glib.Params)
	for _, f := range tab {
		fields[C.GoString(f.name)] = C.GoString((*C.char)(C.gst_value_serialize(f.val)))
	}
	return fields
}

func convertToGoSlice(ptr **C.gchar, length int) []string {
	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(ptr))[:length:length]
	gostrings := make([]string, 0, length)
	for _, s := range tmpslice {
		if s == nil {
			break
		}
		gostrings = append(gostrings, C.GoString(s))
	}

	return gostrings
}

var CLOCK_TIME_NONE = int64(C.GST_CLOCK_TIME_NONE)
