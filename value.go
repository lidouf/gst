//GstValue — GValue implementations specific to GStreamer
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"fmt"
	"github.com/lidouf/glib"
	"unsafe"
)

func v2g(v *glib.Value) *C.GValue {
	return (*C.GValue)(unsafe.Pointer(v))
}

func g2v(v *C.GValue) *glib.Value {
	return (*glib.Value)(unsafe.Pointer(v))
}

//#define GST_FOURCC_FORMAT "c%c%c%c"
//Can be used together with GST_FOURCC_ARGS to properly output a guint32 fourcc value in a printf()-style text message.
//printf ("fourcc: %" GST_FOURCC_FORMAT "\n", GST_FOURCC_ARGS (fcc));
type Fourcc uint32

func (f Fourcc) Type() glib.Type {
	return glib.TYPE_UINT
}

func (f Fourcc) Value() *glib.Value {
	return glib.ValueOf(f)
}

//#define GST_MAKE_FOURCC(a,b,c,d)        ((guint32)((a)|(b)<<8|(c)<<16|(d)<<24))
//
//Transform four characters into a guint32 fourcc value with host endianness.
//
//guint32 fourcc = GST_MAKE_FOURCC ('M', 'J', 'P', 'G');
//Parameters
//a
//the first character
//b
//the second character
//c
//the third character
//d
//the fourth character
func MakeFourcc(a, b, c, d byte) Fourcc {
	return Fourcc(uint32(a) | uint32(b)<<8 | uint32(c)<<16 | uint32(d)<<24)
}

//#define GST_STR_FOURCC(f)               ((guint32)(((f)[0])|((f)[1]<<8)|((f)[2]<<16)|((f)[3]<<24)))
//
//Transform an input string into a guint32 fourcc value with host endianness.
//Caller is responsible for ensuring the input string consists of at least four characters.
//
//guint32 fourcc = GST_STR_FOURCC ("MJPG");
//Parameters
//f
//a string with at least four characters
func StrFourcc(s string) Fourcc {
	if len(s) != 4 {
		panic("Fourcc string length != 4")
	}
	return MakeFourcc(s[0], s[1], s[2], s[3])
}

//#define             GST_FOURCC_ARGS(fourcc)
//
//Can be used together with GST_FOURCC_FORMAT to properly output a guint32 fourcc value in a printf()-style text message.
//Parameters
//fourcc
//a guint32 fourcc value to output
func (f Fourcc) String() string {
	buf := make([]byte, 4)
	buf[0] = byte(f)
	buf[1] = byte(f >> 8)
	buf[2] = byte(f >> 16)
	buf[3] = byte(f >> 32)
	return string(buf)
}

func ValueFourcc(v *glib.Value) Fourcc {
	return Fourcc(v.GetUint32())
}

//#define GST_TYPE_INT_RANGE               (_gst_int_range_type)
//a GValue type that represents an integer range
//Returns
//the GType of GstIntRange
type IntRange struct {
	Start, End int
}

//#define GST_VALUE_HOLDS_INT_RANGE(x)      ((x) != NULL && G_VALUE_TYPE(x) == _gst_int_range_type)
//
//Checks if the given GValue contains a GST_TYPE_INT_RANGE value.
//Parameters
//x
//the GValue to check
func HoldsIntRange(x *glib.Value) bool {
	if x != nil && x.Type() == TYPE_INT_RANGE {
		return true
	}
	return false
}

//Sets value to the range specified by start and end .
//Parameters
//value
//a GValue initialized to GST_TYPE_INT_RANGE
//start
//the start of the range
//end
//the end of the range
func SetIntRange(value *glib.Value, start, end int) {
	C.gst_value_set_int_range(v2g(value), C.gint(start), C.gint(end))
}

//Gets the minimum of the range specified by value .
//Parameters
//value
//a GValue initialized to GST_TYPE_INT_RANGE
//Returns
//the minimum of the range
func GetIntRangeMin(value *glib.Value) int {
	return int(C.gst_value_get_int_range_min(v2g(value)))
}

//Gets the maximum of the range specified by value .
//Parameters
//value
//a GValue initialized to GST_TYPE_INT_RANGE
//Returns
//the maximum of the range
func GetIntRangeMax(value *glib.Value) int {
	return int(C.gst_value_get_int_range_max(v2g(value)))
}

//Sets value to the range specified by start , end and step .
//Parameters
//value
//a GValue initialized to GST_TYPE_INT_RANGE
//start
//the start of the range
//end
//the end of the range
//step
//the step of the range
func SetIntRangeStep(value *glib.Value, start, end, step int) {
	C.gst_value_set_int_range_step(v2g(value), C.gint(start), C.gint(end), C.gint(step))
}

//Gets the step of the range specified by value .
//Parameters
//value
//a GValue initialized to GST_TYPE_INT_RANGE
//Returns
//the step of the range
func GetIntRangeStep(value *glib.Value) int {
	return int(C.gst_value_get_int_range_step(v2g(value)))
}

func ValueRange(v *glib.Value) *IntRange {
	return &IntRange{
		int(C.gst_value_get_int_range_min(v2g(v))),
		int(C.gst_value_get_int_range_max(v2g(v))),
	}
}

func (r *IntRange) Type() glib.Type {
	return TYPE_INT_RANGE
}

func (r *IntRange) Value() *glib.Value {
	v := glib.NewValue(r.Type())
	C.gst_value_set_int_range(v2g(v), C.gint(r.Start), C.gint(r.End))
	return v
}

func (r *IntRange) String() string {
	return fmt.Sprintf("[%d,%d]", r.Start, r.End)
}

type Fraction struct {
	Numer, Denom int
}

func (f *Fraction) Type() glib.Type {
	return TYPE_FRACTION
}

func (f *Fraction) Value() *glib.Value {
	v := glib.NewValue(f.Type())
	C.gst_value_set_fraction(v2g(v), C.gint(f.Numer), C.gint(f.Denom))
	return v
}

func (r *Fraction) String() string {
	return fmt.Sprintf("%d/%d", r.Numer, r.Denom)
}

func ValueFraction(v *glib.Value) *Fraction {
	return &Fraction{
		int(C.gst_value_get_fraction_numerator(v2g(v))),
		int(C.gst_value_get_fraction_denominator(v2g(v))),
	}
}
