/*
Twofer outputs a simple text string to the command line
*/
package twofer

// ShareWith returns a string literal including name.
// The returned string either includes the name variable or generic "you" if the name is not passed in.
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
