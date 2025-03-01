# **CLASP.Go** Changes <!-- omit in toc -->


## 0.17.0-alpha2 - 1st March 2025

* - removed `Arguments#CheckAllFlagBits()` and `Arguments#CheckUnusedFlagBits()`;
* ~ fixed defect in handling of `Parse_DontMergeBitFlagsIntoBitFlags64`;
* ~ flags with bit-flags receiver variables should be marked used during `Parse()` (except when suppressed);
* + more testing of `Parse_DontMergeBitFlagsIntoBitFlags64`;
* ~ standardised names of `ParseFlag` constants;
* ~ refactored `#String()` form of `Specification`;
* ~ updated dependencies;
* ~ applied STEGoL to various test cases;
* ~ simplification;
* ~ preparatory refactoring;
* + documentation markup;


## 0.17.0-alpha1 - 28th February 2025

* + added `Specification#SetBitFlags()` / `Specification#SetBitFlags64()`, for associating a given flag specification with a bit-flags value and, optionally, a bitmask variable, to be applied automatically when detected from the command-line during parsing;
* ~ improved (internal) field names
* ~ formatting;
* ~ simplification of markdown;


## 0.16.3-alpha1 - 28th February 2025

* - removed unncessary use of `iota`;
* + added more test cases;
* ~ tidying / doc-markup;
* ~ formatting;


## 0.16.2 - 28th February 2025

* ~ `Version` additionally accepts type `[]uint16`;
* ~ changed some tests to exercise the ability of `Version` to be array of numbers;


## 0.16.2-alpha2 - 28th February 2025

* ~ boilerplate;
* ~ formatting;


## 0.16.2-alpha1 - 23rd February 2025

* ~ updated for use of Go modules;


## 0.16.1 - 22nd July 2019

* ~ improvement in display of usage to ensure sections are always surrounded by blank lines, and to deal with flags/options without help descriptions


## 0.16.0 - 22nd July 2019

* + added support for sections (via ``Section()`` function)


## 0.15.0 - 20th July 2019

* ~ minor improvements to README.md


## 0.15.0 - 10th April 2019

* ~ changed *Alias to *Specification


## 0.14.4 - 8th April 2019

* ~ option-value aliases are now listed correctly in ShowUsage()


## 0.14.3 - 8th April 2019

* ~ fixed failure to properly assign alias for uncombined option-value aliases


## 0.14.2 - 30th March 2019

* ~ fixed failure to call specified exiter in ShowUsage()
* ~ now uses features from ANGoLS (https://github.com/synesissoftware/ANGoLS/) for testing support
* ~ fixed HelpFlag() help message (which said "Shows this helps and exits"; discovered in unit-testing libCLImate.Go)


## 0.14.1 - 29th March 2019

* ~ fix for FlagsAndOptionsString handling


## 0.14.0 - 29th March 2019

* + added End() builder method


## 0.13.1 - 28th March 2019

* ~ handling omitted alias resolution for options


## 0.13.0 - 25th March 2019

 * ~ changed ``Flag``, ``Option``, and ``Value`` constants to ``FlagType``, ``OptionType``, ``ValueType``;
 * + added ``Flag()``, ``Option()`` and ``AliasesFor()`` alias constructor functions;
 * + added ``Alias.SetHelp()``, ``Alias.SetAlias()``, ``Alias.SetAliases()``, ``Alias.SetExtra()`` methods;
 * + added ``Alias.Used()`` method;
 * + added ``Arguments.LookupFlag()`` method;
 * ~ renamed ``Arguments.CheckAllBitFlags()`` to ``Arguments.CheckUnusedBitFlags()``;
 * + added ``Arguments.CheckAllBitFlags()``, which does checks all bit flags and does not mark any as used;
 * ~ fixed defect whereby a value-option alias was attached to the argument, but now the underlying alias (without the option value) is attached; and
 * + added CHANGES.md

These are **destructive changes** and will not be compatible with any code depending on version *0.12.x* or earlier.


## 0.12.1 - 22nd March 2019

* + added ``Extras`` field to ``Alias`` struct, to support first released version of [**libCLImate.Go**](https://github.com/synesissoftware/libCLImate.Go).


## 0.11.2 - 8th March 2019

* ~ fixed to ``usage_test.go``.


## 0.11.1 - 8th March 2019

* + added ``String()`` methods for various CLASP.Go structures; and
* ~ examples improvement.


## 0.10.3 - 8th March 2019

* ~ handling combined flags involving options-with-values;
* ~ various improvements with examples; and
* README.md : + added related projects.


## 0.10.2 - 8th March 2019

FIRST PUBLIC RELEASE


<!-- ########################### end of file ########################### -->

