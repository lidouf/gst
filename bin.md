**progress**
6/21
```
Y   GstElement * 	gst_bin_new ()
Y   gboolean 	gst_bin_add ()
Y   gboolean 	gst_bin_remove ()
Y   GstElement * 	gst_bin_get_by_name ()
    GstElement * 	gst_bin_get_by_name_recurse_up ()
    GstElement * 	gst_bin_get_by_interface ()
    GstIterator * 	gst_bin_iterate_elements ()
    GstIterator * 	gst_bin_iterate_recurse ()
    GstIterator * 	gst_bin_iterate_sinks ()
    GstIterator * 	gst_bin_iterate_sorted ()
    GstIterator * 	gst_bin_iterate_sources ()
    GstIterator * 	gst_bin_iterate_all_by_interface ()
    gboolean 	gst_bin_recalculate_latency ()
*   void 	gst_bin_add_many ()
*   void 	gst_bin_remove_many ()
    GstPad * 	gst_bin_find_unlinked_pad ()
    gboolean 	gst_bin_sync_children_states ()
    #define 	GST_BIN_IS_NO_RESYNC()
    #define 	GST_BIN_CHILDREN()
    #define 	GST_BIN_CHILDREN_COOKIE()
    #define 	GST_BIN_NUMCHILDREN()
```
the item mark * that means can be replaced by other calls