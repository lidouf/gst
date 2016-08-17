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

type Toc struct {
	glib.Object
}

func (t *Toc) g() *C.GstToc {
	return (*C.GstToc)(t.GetPtr())
}

func (t *Toc) AsTOC() *Toc {
	return t
}
