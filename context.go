//GstContext — Lightweight objects to represent element contexts
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

type Context struct {
	glib.Object
}

func (o *Context) g() *C.GstContext {
	return (*C.GstContext)(o.GetPtr())
}

func (o *Context) AsContext() *Context {
	return o
}
