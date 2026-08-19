# CLASP.Go - TODO <!-- omit in toc -->


## Table of Contents <!-- omit in toc -->

- [Functional improvements](#functional-improvements)
- [Performance improvements](#performance-improvements)
- [Packaging improvements](#packaging-improvements)


## Functional improvements

* [ ] consider whether handling of `Version` should be implemented in terms of **ver2go**;
* [x] Standardise format of all (public) constants;
* [ ] Flags with bit-flags receiver variables should be marked used during `Parse()` (except when suppressed);


## Performance improvements

* \<none>


## Packaging improvements

* [ ] Before the next official release: confirm **`go.mod`** (`go 1.21`) and the CI Go-version matrix, bump Synesis `require`s to newly published tags, then run **`go mod tidy`** (not against currently published tags). Prior Synesis Go releases, in order:
  * **ver2go**;
  * **STEGoL**;
  * **ANGoLS**;
* [ ] Ensure all documentation markup is adequate;
* [ ] Flesh out all documentation;


<!-- ########################### end of file ########################### -->
