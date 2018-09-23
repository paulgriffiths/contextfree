/*
Package lar provides a single character lookahead reader.

When performing lexical analysis, it often greatly simplifies matters to
have available one character of lookahead, i.e. to be able to peek at what
the next character to be read would be without actually reading it. With
this ability you can construct a series of tests to check if the next
character matches the start of a pattern, but only read and consume that
character from the input if it does, in fact, match.

The single character lookahead reader implemented by this package uses that
functionality to provide a set of matching functions which will extract
a character or set of characters from the input if and only if they match
a specified category.
*/
package lar
