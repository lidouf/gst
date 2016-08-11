**progress**
10/10
```
-   #define 	GST_STATIC_PAD_TEMPLATE()
Y   GstPadTemplate * 	gst_static_pad_template_get ()
Y   GstCaps * 	gst_static_pad_template_get_caps ()
Y   #define 	GST_PAD_TEMPLATE_NAME_TEMPLATE()
Y   #define 	GST_PAD_TEMPLATE_DIRECTION()
Y   #define 	GST_PAD_TEMPLATE_PRESENCE()
Y   #define 	GST_PAD_TEMPLATE_CAPS()
!   #define 	GST_PAD_TEMPLATE_IS_FIXED()
Y   GstPadTemplate * 	gst_pad_template_new ()
Y   GstCaps * 	gst_pad_template_get_caps ()
```
**Note: It looks like `GST_PAD_TEMPLATE_FIXED` constant is removed in gstreamer-1.0, so the `GST_PAD_TEMPLATE_IS_FIXED` is unavailable now.**
in gstreamer 0.10:
```
typedef enum {
  /* FIXME0.11: this is not used and the purpose is unclear */
  GST_PAD_TEMPLATE_FIXED        = (GST_OBJECT_FLAG_LAST << 0),
  /* padding */
  GST_PAD_TEMPLATE_FLAG_LAST    = (GST_OBJECT_FLAG_LAST << 4)
} GstPadTemplateFlags;
```
in gstreamer 1.0:
```
typedef enum {
  /* padding */
  GST_PAD_TEMPLATE_FLAG_LAST    = (GST_OBJECT_FLAG_LAST << 4)
} GstPadTemplateFlags;
```