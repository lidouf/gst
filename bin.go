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

type Bin struct {
	Element
}

func (b *Bin) g() *C.GstBin {
	return (*C.GstBin)(b.GetPtr())
}

func (b *Bin) AsBin() *Bin {
	return b
}

func NewBin(name string) *Bin {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(Bin)
	b.SetPtr(glib.Pointer(C.gst_bin_new(s)))
	return b
}

func (b *Bin) Add(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_add(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

func (b *Bin) Remove(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_remove(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

// GetByName returns the element with the given name from a bin. Returns nil
// if no element with the given name is found in the bin.
func (b *Bin) GetByName(name string) *Element {
	en := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(en))
	p := glib.Pointer(C.gst_bin_get_by_name(b.g(), en))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(p)
	return e
}

func (b *Bin) DebugToDotData(details DebugGraphDetails) string {
	return (C.GoString)((*C.char)(C.gst_debug_bin_to_dot_data(b.g(), (C.GstDebugGraphDetails)(details))))
}

func (b *Bin) DebugToDotFile(filename string, details DebugGraphDetails) {
	s := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(s))
	C.gst_debug_bin_to_dot_file(b.g(), (C.GstDebugGraphDetails)(details), s)
}

func (b *Bin) DebugToDotFileWithTs(filename string, details DebugGraphDetails) {
	s := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(s))
	C.gst_debug_bin_to_dot_file_with_ts(b.g(), (C.GstDebugGraphDetails)(details), s)
}
