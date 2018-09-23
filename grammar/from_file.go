package grammar

import "os"

// FromFile constructs a context-free grammar object from
// a representation in a text file.
func FromFile(filename string) (*Grammar, error) {
	infile, fileErr := os.Open(filename)
	if fileErr != nil {
		return nil, fileErr
	}
	defer infile.Close()

	g, perr := parse(infile)
	if perr != nil {
		return nil, perr
	}

	return g, nil
}
