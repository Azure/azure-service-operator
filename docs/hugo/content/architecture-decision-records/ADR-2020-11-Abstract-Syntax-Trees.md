---
title: '2020-11: AST Library Choice'
---

## Context

When we first started working on the Azure Service Operator Code Generator, we used the standard go [`ast`](https://pkg.go.dev/go/ast) library to create the final Go code.

Unfortunately, we have run into a number of significant limitations with this library.

* **Comments are often emitted in unexpected places.**  
  They would often appear at the end of the previous line, which conveys an incorrect meaning to a casual reader of the code. As a workaround, we found that prefacing some comments with newline characters would force them onto the next line. While hacking the library in this function worked, we are concerned that it's a fragile technique that would break in a future version of Go. We also found situations where this technique did not work.

* **Generated code does not always comply with `go fmt` standards.**  
  For some constructs, the generated code would fail the `go fmt` check included in our continuous integration (CI) builds. We found a workaround that involved using `ast` to read in a generated file and immediately rewrite it back out again, immitating the results of `go fmt`.

* **Generated code does not always compile.**  
  For some constructs, we found the comments were emitted at the start of the line of code (not on the previous line), resulting in the code being commented out. We found no workaround for this.

* **Poor control over whitespace**  
  We were unable to find a reliable technique within `ast` to introduce blank lines betweens stanza of code, or before comment blocks. As a workaround, we would read in the formatted code, scan it for comments and manually inject blank lines before each comment block. 

We have come to the conclusions that the `ast` library is intended as a way to make minor modifications to *existing* well formatted Go files, not for the creation of entirely new Go files from scratch.

A potential alternative is the [`dst`](https://github.com/dave/dst) (Decorated Syntax Tree) package. This package was specifically created to address the kinds of issues described above.

## Decision

After some informative trials, we have decided to adopt the `dst` library.

## Status

Adopted.

Initial change committed in [PR#366](https://github.com/Azure/k8s-infra/pull/336).

## Consequences

With a very high level of API compatibility, we were able to introduce `dst` with a low level of code churn by aliasing the import as `ast` in most files. Future changes should remove this aliasing as files are modified.

The DST library requires that any node be used exactly once, and will panic if a node is reused in mulitiple locations. We've mitigated this by liberally cloning nodes as we build the desired final syntax tree. (See `astbuilder.Expressions()`, introduced in [PR#1613](https://github.com/Azure/azure-service-operator/pull/1613) and `astbuilder.Statements()`, introduced in [PR#427](https://github.com/Azure/k8s-infra/pull/427)).

## Experience Report

TBC

## References

[Go AST](https://pkg.go.dev/go/ast) (Abstract Syntax Tree) library

[DST](https://github.com/dave/dst) (Decorated Syntax Tree) library
