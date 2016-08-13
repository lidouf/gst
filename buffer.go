//GstBuffer — Data-passing buffer type
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

type Buffer struct {
	glib.Object
}

func (b *Buffer) g() *C.GstBuffer {
	return (*C.GstBuffer)(b.GetPtr())
}

func (b *Buffer) AsBuffer() *Buffer {
	return b
}
