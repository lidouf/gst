//type converter
package gst

/*
#include <gst/gst.h>
*/
import "C"

import (
//"github.com/lidouf/glib"
)

func gBoolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
