# Bit Manipulation

After failing an interview question about some bit manipulation, I decided to
work through relevant exercises so that I feel less bad about it.

General structure of the repo is that `_test.go` files test the proposed solution.

Solutions to the following problems can be found:

* [5.2 Binary To String][b2s] I rewrote my original solution to use the book's trick of multiplying by 2.
* [5.3 Flip Bit to Win][flip] I save positions of each 0. Then walk through such array.
* [5.4 Next Number][next] Solution closely resembles the one outlined in the book. What's "novel" from is that I test entire range of `uint16` against a naive solution just to make sure I got this one right.


[b2s]: https://github.com/afiodorov/bit_manipulation/tree/master/binary_to_string.go
[flip]: https://github.com/afiodorov/bit_manipulation/tree/master/flip_bit_to_win.go
[next]: https://github.com/afiodorov/bit_manipulation/tree/master/next_number.go
