//GstIterator — Object to retrieve multiple elements in a threadsafe way.
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
//"errors"
//"github.com/lidouf/glib"
//"time"
//"unsafe"
)

type IteratorResult C.GstIteratorResult

const (
	ITERATOR_DONE   = IteratorResult(C.GST_ITERATOR_DONE)
	ITERATOR_OK     = IteratorResult(C.GST_ITERATOR_OK)
	ITERATOR_RESYNC = IteratorResult(C.GST_ITERATOR_RESYNC)
	ITERATOR_ERROR  = IteratorResult(C.GST_ITERATOR_ERROR)
)

type IteratorItem C.GstIteratorItem

const (
	ITERATOR_ITEM_SKIP = IteratorItem(C.GST_ITERATOR_ITEM_SKIP)
	ITERATOR_ITEM_PASS = IteratorItem(C.GST_ITERATOR_ITEM_PASS)
	ITERATOR_ITEM_END  = IteratorItem(C.GST_ITERATOR_ITEM_END)
)

type Iterator C.GstIterator
