package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

typedef struct {
	const char *name;
	const char *val;
} TagStruct;

typedef struct {
	TagStruct* tab;
	int        n;
} TagsStruct;

void _parse_tag(const GstTagList *tags, const gchar *tag, gpointer data) {
    GValue val = { 0, };
    gchar *str;
    TagsStruct *f = (TagsStruct*)(data);

    gst_tag_list_copy_value(&val, tags, tag);

    if (G_VALUE_HOLDS_STRING(&val)) {
        str = g_value_dup_string(&val);
    } else {
        str = gst_value_serialize(&val);
    }

    f->tab[f->n].name = tag;
    f->tab[f->n].val = str;

    ++f->n;

    g_value_unset(&val);
}

TagsStruct _parse_tag_list(GstTagList *t) {
	int n = gst_tag_list_n_tags(t);;
	TagsStruct f = { malloc(n * sizeof(TagStruct)), 0 };

	gst_tag_list_foreach(t, _parse_tag, (gpointer)(&f));
	return f;
}
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type TagMergeMode C.GstTagMergeMode

const (
	TAG_MERGE_UNDEFINED   = TagMergeMode(C.GST_TAG_MERGE_UNDEFINED)
	TAG_MERGE_REPLACE_ALL = TagMergeMode(C.GST_TAG_MERGE_REPLACE_ALL)
	TAG_MERGE_REPLACE     = TagMergeMode(C.GST_TAG_MERGE_REPLACE)
	TAG_MERGE_APPEND      = TagMergeMode(C.GST_TAG_MERGE_APPEND)
	TAG_MERGE_PREPEND     = TagMergeMode(C.GST_TAG_MERGE_PREPEND)
	TAG_MERGE_KEEP        = TagMergeMode(C.GST_TAG_MERGE_KEEP)
	TAG_MERGE_KEEP_ALL    = TagMergeMode(C.GST_TAG_MERGE_KEEP_ALL)
	/* add more */
	TAG_MERGE_COUNT = TagMergeMode(C.GST_TAG_MERGE_COUNT)
)

type TagFlag C.GstTagFlag

const (
	TAG_FLAG_UNDEFINED = TagFlag(C.GST_TAG_FLAG_UNDEFINED)
	TAG_FLAG_META      = TagFlag(C.GST_TAG_FLAG_META)
	TAG_FLAG_ENCODED   = TagFlag(C.GST_TAG_FLAG_ENCODED)
	TAG_FLAG_DECODED   = TagFlag(C.GST_TAG_FLAG_DECODED)
	TAG_FLAG_COUNT     = TagFlag(C.GST_TAG_FLAG_COUNT)
)

func (t *TagFlag) g() *C.GstTagFlag {
	return (*C.GstTagFlag)(t)
}

type TagScope C.GstTagScope

const (
	TAG_SCOPE_STREAM = TagScope(C.GST_TAG_SCOPE_GLOBAL)
	TAG_SCOPE_GLOBAL = TagScope(C.GST_TAG_SCOPE_GLOBAL)
)

func (t *TagScope) g() *C.GstTagScope {
	return (*C.GstTagScope)(t)
}

//type TagList C.GstTagList
type TagList struct {
	glib.Object
}

func (t *TagList) g() *C.GstTagList {
	return (*C.GstTagList)(t.GetPtr())
}

func (t *TagList) Type() glib.Type {
	return glib.TypeFromName("GstTagList")
}

func (t *TagList) TagsNumber() int {
	return int(C.gst_tag_list_n_tags(t.g()))
}

func (t *TagList) Serialize() glib.Params {
	ps := C._parse_tag_list(t.g())
	n := (int)(ps.n)
	tab := (*[1 << 16]C.TagStruct)(unsafe.Pointer(ps.tab))[:n]
	fields := make(glib.Params)
	for _, f := range tab {
		fields[C.GoString(f.name)] = C.GoString((*C.char)(f.val))
	}
	return fields
}

func GetTagNick(tag string) string {
	s := (*C.gchar)(C.CString(tag))
	defer C.free(unsafe.Pointer(s))
	return C.GoString((*C.char)(C.gst_tag_get_nick(s)))
}
