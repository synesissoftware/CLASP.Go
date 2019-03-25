# **CLASP.Go** Changes

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
