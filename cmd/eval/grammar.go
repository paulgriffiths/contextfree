package main

/*
E  : T E'

E' : `\+` T E'
   | `-` T E'
   | e

T  : F T'

T' : `\*` F T'
   | `/` F T'
   | e

F  : N
   | `\(` E `\)`

N  : `-?[[:digit:]]+`
*/

// Grammar for arithmetical expression parser.
var grammarString = `E  : T E'

E' : ` + "`" + `\+` + "`" + ` T E'
   | ` + "`" + `-` + "`" + ` T E'
   | e

T  : F T'

T' : ` + "`" + `\*` + "`" + ` F T'
   | ` + "`" + `/` + "`" + ` F T'
   | e

F  : N
   | ` + "`" + `\(` + "`" + ` E ` + "`" + `\)` + "`" + `

N  : ` + "`" + `-?[[:digit:]]+` + "`" + `
`
