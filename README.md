# Advent of Code 2021

These are my solutions for the [Advent of Code 2021](https://adventofcode.com/) challenge.

## Approach

I will be tackling the problems with [Go](https://go.dev/).  This is because I am working to improve my proficiency with the language.  My goal is _not_ to quickly hack my way to the answer.  My goal is explore Go and see what it takes to write quality code that solves the problems.

## Questions and Observations

### Day 1

I have a couple of test cases that use the same test data.  How can I set this shared code up as a constant?  Currently, only basic data types (int, float, bool, string, ...) can be set up as constants.  In order to reuse something like an `[]int` wrap it in a function.

### Day 2

How can I implement dependency injection (DI)?  I want to be able to inject either version of sub into the script runner code.  

DONE.  Injected script runner function.

### Day 3

Using `type` and `struct` to accomplish what I would do with a class in other languages.

## Reference

- [Examples](https://github.com/k0emt/go-ws/tree/main/examples) from [my own go learning spike](https://github.com/k0emt/go-ws)
