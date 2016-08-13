//GstBufferList — Lists of buffers for data-passing
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	//"errors"
	"github.com/lidouf/glib"
	//"time"
	//"unsafe"
)

type BufferList struct {
	glib.Object
}

func (b *BufferList) g() *C.GstBufferList {
	return (*C.GstBufferList)(b.GetPtr())
}

func (b *BufferList) AsBuffer() *BufferList {
	return b
}
