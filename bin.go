//GstBin — Base class and element that can contain other elements
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

type BinFlags C.GstBinFlags

const (
	BIN_FLAG_NO_RESYNC = BinFlags(C.GST_BIN_FLAG_NO_RESYNC)
	/* padding */
	BIN_FLAG_LAST = BinFlags(C.GST_BIN_FLAG_LAST)
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

//Creates a new bin with the given name.
//Parameters
//name
//the name of the new bin.
//Returns
//a new GstBin.
func NewBin(name string) *Bin {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(Bin)
	b.SetPtr(glib.Pointer(C.gst_bin_new(s)))
	return b
}

//Adds the given element to the bin. Sets the element's parent, and thus takes ownership of the element. An element can only be added to one bin.
//
//If the element's pads are linked to other pads, the pads will be unlinked before the element is added to the bin.
//When you add an element to an already-running pipeline, you will have to take care to set the state of the newly-added element to the desired state
//(usually PLAYING or PAUSED, same you set the pipeline to originally) with gst_element_set_state(), or use gst_element_sync_state_with_parent().
//The bin or pipeline will not take care of this for you.
//
//MT safe.
//Parameters
//bin
//a GstBin
//element
//the GstElement to add.
//Returns
//TRUE if the element could be added, FALSE if the bin does not want to accept the element.
func (b *Bin) Add(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_add(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

//Removes the element from the bin, unparenting it as well.
//Unparenting the element means that the element will be dereferenced, so if the bin holds the only reference to the element,
//the element will be freed in the process of removing it from the bin.
//If you want the element to still exist after removing, you need to call gst_object_ref() before removing it from the bin.
//
//If the element's pads are linked to other pads, the pads will be unlinked before the element is removed from the bin.
//
//MT safe.
//Parameters
//bin
//a GstBin
//element
//the GstElement to remove.
//Returns
//TRUE if the element could be removed, FALSE if the bin does not want to remove the element.
func (b *Bin) Remove(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_remove(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

//GetByName returns the element with the given name from a bin. Returns nil
//if no element with the given name is found in the bin.
//MT safe. Caller owns returned reference.
//Parameters
//bin
//a GstBin
//name
//the element name to search for
//Returns
//the GstElement with the given name, or NULL.
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

//Gets the element with the given name from this bin. If the element is not found, a recursion is performed on the parent bin.
//
//Returns NULL if:
//
//no element with the given name is found in the bin
//
//MT safe. Caller owns returned reference.
//Parameters
//bin
//a GstBin
//name
//the element name to search for
//Returns
//the GstElement with the given name, or NULL.
func (b *Bin) GetByNameRecurseUp(name string) *Element {
	en := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(en))
	p := glib.Pointer(C.gst_bin_get_by_name_recurse_up(b.g(), en))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(p)
	return e
}

//Looks for an element inside the bin that implements the given interface.
//If such an element is found, it returns the element. You can cast this element to the given interface afterwards.
//If you want all elements that implement the interface, use gst_bin_iterate_all_by_interface().
//This function recurses into child bins.
//
//MT safe. Caller owns returned reference.
//Parameters
//bin
//a GstBin
//iface
//the GType of an interface
//Returns
//A GstElement inside the bin implementing the interface.
func (b *Bin) GetByInterface(tp glib.Type) *Element {
	p := C.gst_bin_get_by_interface(b.g(), C.GType(tp))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(glib.Pointer(p))
	return e
}

//Gets an iterator for the elements in this bin.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//Returns
//a GstIterator of GstElement, or NULL.
func (b *Bin) IterateElements() *Iterator {
	p := C.gst_bin_iterate_elements(b.g())
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Gets an iterator for the elements in this bin. This iterator recurses into GstBin children.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//Returns
//a GstIterator of GstElement, or NULL.
func (b *Bin) IterateRecurse() *Iterator {
	p := C.gst_bin_iterate_recurse(b.g())
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Gets an iterator for all elements in the bin that have the GST_ELEMENT_FLAG_SINK flag set.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//Returns
//a GstIterator of GstElement, or NULL.
func (b *Bin) IterateSinks() *Iterator {
	p := C.gst_bin_iterate_sinks(b.g())
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Gets an iterator for the elements in this bin in topologically sorted order.
//This means that the elements are returned from the most downstream elements (sinks) to the sources.
//
//This function is used internally to perform the state changes of the bin elements and for clock selection.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//Returns
//a GstIterator of GstElement, or NULL.
func (b *Bin) IterateSorted() *Iterator {
	p := C.gst_bin_iterate_sorted(b.g())
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Gets an iterator for all elements in the bin that have the GST_ELEMENT_FLAG_SOURCE flag set.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//Returns
//a GstIterator of GstElement, or NULL.
func (b *Bin) IterateSources() *Iterator {
	p := C.gst_bin_iterate_sources(b.g())
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Looks for all elements inside the bin that implements the given interface. You can safely cast all returned elements to the given interface.
//The function recurses inside child bins. The iterator will yield a series of GstElement that should be unreffed after use.
//
//MT safe. Caller owns returned value.
//Parameters
//bin
//a GstBin
//iface
//the GType of an interface
//Returns
//a GstIterator of GstElement for all elements in the bin implementing the given interface, or NULL.
func (b *Bin) IterateAllByInterface(tp glib.Type) *Iterator {
	p := C.gst_bin_iterate_all_by_interface(b.g(), C.GType(tp))
	if p == nil {
		return nil
	}
	return (*Iterator)(unsafe.Pointer(p))
}

//Query bin for the current latency using and reconfigures this latency to all the elements with a LATENCY event.
//
//This method is typically called on the pipeline when a GST_MESSAGE_LATENCY is posted on the bus.
//
//This function simply emits the 'do-latency' signal so any custom latency calculations will be performed.
//Parameters
//bin
//a GstBin
//Returns
//TRUE if the latency could be queried and reconfigured.
func (b *Bin) RecalculateLatency() bool {
	return C.gst_bin_recalculate_latency(b.g()) != 0
}

//Recursively looks for elements with an unlinked pad of the given direction within the specified bin and returns an unlinked pad if one is found, or NULL otherwise. If a pad is found, the caller owns a reference to it and should use gst_object_unref() on the pad when it is not needed any longer.
//Parameters
//bin
//bin in which to look for elements with unlinked pads
//direction
//whether to look for an unlinked source or sink pad
//Returns
//unlinked pad of the given direction, NULL.
func (b *Bin) FindUnlinkedPad(direction PadDirection) *Pad {
	p := C.gst_bin_find_unlinked_pad(b.g(), direction.g())
	if p == nil {
		return nil
	}
	r := new(Pad)
	r.SetPtr(glib.Pointer(p))
	return r
}

//Synchronizes the state of every child of bin with the state of bin . See also gst_element_sync_state_with_parent().
//Parameters
//bin
//a GstBin
//Returns
//TRUE if syncing the state was successful for all children, otherwise FALSE.
//Since: 1.6
func (b *Bin) SyncChildrenStates() bool {
	return C.gst_bin_sync_children_states(b.g()) != 0
}

//#define GST_BIN_IS_NO_RESYNC(bin)        (GST_OBJECT_FLAG_IS_SET(bin,GST_BIN_FLAG_NO_RESYNC))
//Check if bin will resync its state change when elements are added and removed.
//Parameters
//bin
//A GstBin
//Since: 1.0.5
func (b *Bin) IsNoResync() bool {
	return b.FlagIsSet(uint32(BIN_FLAG_NO_RESYNC))
}

//#define GST_BIN_CHILDREN(bin)		(GST_BIN_CAST(bin)->children)
//Gets the list with children in a bin.
//Parameters
//bin
//a GstBin
func (b *Bin) Children() *glib.List {
	return glib.WrapList(uintptr(unsafe.Pointer(b.g().children)))
}

//#define GST_BIN_CHILDREN_COOKIE(bin) (GST_BIN_CAST(bin)->children_cookie)
//Gets the children cookie that watches the children list.
//Parameters
//bin
//a GstBin
func (b *Bin) ChildrenCookie() uint32 {
	return uint32(b.g().children_cookie)
}

//#define GST_BIN_NUMCHILDREN(bin) (GST_BIN_CAST(bin)->numchildren)
//Gets the number of children in a bin.
//Parameters
//bin
//a GstBin
func (b *Bin) NumChildren() int {
	return int(b.g().numchildren)
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
