// Bindings for GStreamer API
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

static char** _gst_init(int* argc, char** argv) {
	gst_init(argc, &argv);
	return argv;
}

static gboolean _gst_init_check(int* argc, char** argv, GError** err) {
	return gst_init_check(argc, &argv, err);
}
*/
import "C"

import (
	"errors"
	"github.com/lidouf/glib"
	"os"
	"unsafe"
)

func Init() {
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
}

func InitCheck() error {
	alen := C.int(len(os.Args))
	argv := make([]*C.char, alen)
	for i, s := range os.Args {
		argv[i] = C.CString(s)
	}
	var err *C.GError
	defer C.free(unsafe.Pointer(err))
	ret := C._gst_init_check(&alen, &argv[0], &err)
	if ret == 0 {
		return errors.New((*glib.Error)(unsafe.Pointer(err)).Error())
	}
	return nil
}

func IsInitialized() bool {
	return C.gst_is_initialized() != 0
}

func Deinit() {
	C.gst_deinit()
}

func Version() (uint, uint, uint, uint) {
	var major, minor, micro, nano C.guint
	C.gst_version(&major, &minor, &micro, &nano)
	return uint(major), uint(minor), uint(micro), uint(nano)
}

func VersionString() string {
	return C.GoString((*C.char)(C.gst_version_string()))
}

func SegtrapIsEnabled() bool {
	return C.gst_segtrap_is_enabled() != 0
}

func SegtrapSetEnabled(enabled bool) {
	var e int
	if enabled {
		e = 1
	} else {
		e = 0
	}
	C.gst_segtrap_set_enabled(C.gboolean(e))
}

func RegistryForkIsEnabled() bool {
	return C.gst_registry_fork_is_enabled() != 0
}

func RegistryForkSetEnabled(enabled bool) {
	var e int
	if enabled {
		e = 1
	} else {
		e = 0
	}
	C.gst_registry_fork_set_enabled(C.gboolean(e))
}

func UpdateRegistry() bool {
	return C.gst_update_registry() != 0
}
