#include <gst/gst.h>

char** _gst_init(int* argc, char** argv) {
	gst_init(argc, &argv);
	return argv;
}

//typedef struct {
//	const char *name;
//	const GValue *val;
//} Field;
//
//typedef struct {
//	Field* tab;
//	int    n;
//} Fields;
//
//gboolean _parse_field(GQuark id, const GValue* val, gpointer data) {
//	Fields *f = (Fields*)(data);
//	f->tab[f->n].name = g_quark_to_string(id);
//	f->tab[f->n].val = val;
//	++f->n;
//	return TRUE;
//}
//
//Fields _parse_struct(GstStructure *s) {
//	int n = gst_structure_n_fields(s);
//	Fields f = { malloc(n * sizeof(Field)), 0 };
//	gst_structure_foreach(s, _parse_field, (gpointer)(&f));
//	return f;
//}