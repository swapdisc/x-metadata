# QueenAttack

Write a program that positions two queens on a chess board and indicates whether or not they are positioned so that they can attack each other.

In the game of chess, a queen can attack pieces which are on the same row,
column, or diagonal.

A chessboard can be represented by an 8 by 8 array.

So if you're told the white queen is at (2, 3) and the black queen at (5, 6),
then you'd know you've got a set-up like so:

```plain
O O O O O O O O
O O O O O O O O
O O O W O O O O
O O O O O O O O
O O O O O O O O
O O O O O O B O
O O O O O O O O
O O O O O O O O
```

You'd also be able to answer whether the queens can attack each other.
In this case, that answer would be yes, they can,
because both pieces share a diagonal.


## Source

J Dalbey's Programming Practice problems [view source](http://users.csc.calpoly.edu/~jdalbey/103/Projects/ProgrammingPractice.html)