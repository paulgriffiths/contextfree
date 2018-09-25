# eval

**eval** is a demonstration program illustrating the use of a context-free
grammar and a predictive parser to implement a simple arithmetical
expression evaluation engine.

Addition, subtraction, multiplication and division operations are supported,
as well as parenthesized expressions. Numbers must be integers, although they
may be positive or negative.

## Example

Sample session:

	paul@horus:eval$ ./eval
	Enter simple arithmetical expressions ('q' to quit).
	> 3+4*5
	 = 23
	> (3+4)*5
	 = 35
	> 13 / -5
	 = -2.6
	> -4 * -8
	 = 32
	> (6 + 7) * (8 + 9)
	 = 221
	> 2 ^ 3
	Illegal expression.
	> 10/0
	 = +Inf
	> 42
	 = 42
	> q
	paul@horus:eval$ 

