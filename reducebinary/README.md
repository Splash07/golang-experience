# Tricks and hacks to reduce size of app's binary

- Making structs incomparable (which is intended to be) by adding an incomprable field
  - use _ [0][]byte
  - check out **incomprable-struct** package for more details
  - reference: https://dave.cheney.net/2020/05/09/ensmallening-go-binaries-by-prohibiting-comparisons
  - explanation:
    - due to the lack of support for generics, go compiler has to generate equality and hash functions for every date type in the program's binary 
    - the bigger the size and complexity the type is, the bigger the overhead is
    - making structs incomprable prevents hash and equality functions from being generated for those types, resulting in a small reduction in the size of the binary