/*
Package stylesheet holds the definitions for a transformation stylesheet.
A transformation stylesheet is a bundle of transformation templates, that
define how the input should be handled
*/
package types

/*
Stylesheet type defines the elements of a stylessheet. The file templates will be unmarshaled
to that type
*/
type Stylesheet struct {
	// array of processing steps
	Templates []Template
}

/*
Template type defines elements that describe actions on one tranformation/
traversation step in the processing of the input
*/
type Template struct {
	// defines on what element the template should be used
	Match   string
	ValueOf ValueOf
}

/*
ValueOf type defines elements that describe actions on one tranformation/
traversation step in the processing of the input
*/
type ValueOf struct {
	// defines on what element/values should be handled
	Select string
}
