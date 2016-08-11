//GstEvent — Structure describing events that are passed up and down a pipeline
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

type Event struct {
	glib.Object
}

func (o *Event) g() *C.GstEvent {
	return (*C.GstEvent)(o.GetPtr())
}

func (o *Event) AsEvent() *Event {
	return o
}
