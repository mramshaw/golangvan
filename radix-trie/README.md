# Radix Trie

This is mainly based on the [Wikipedia article](https://en.wikipedia.org/wiki/Radix_tree).

Note that a __radix tree__ is a space-optimized [trie](https://en.wikipedia.org/wiki/Trie).

More particularly, this seems to be a __Patricia tree__ - which I believe is the binary form.

![Patricia_trie](https://upload.wikimedia.org/wikipedia/commons/a/ae/Patricia_trie.svg)

## Motivation

I've been doing a lot of high-level programming lately (RESTful APIS, Scala/Akka) so
something a little more low-level sounded like a nice change of pace.

Also, it's a nice opportunity to use
[Table-driven testing](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go).

## Prerequisites

1. A recent version of Go

2. `make` installed

[Or can simply type `go test -v`.]

## Expected Usage

First four inserts:

```
   1          2             3               4
   |          |             |               |
   r          r             r               r
   |          |             |              / \
 omane      oman            om            om  ubens
             /\            / \           / \
            e  us         an  ulus      an  ulus
                         / \           / \
                        e  us         e   us
```

## To Run

Type the following command:

    make
