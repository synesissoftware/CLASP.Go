#! /bin/bash

ScriptPath=$0
Dir=$(cd $(dirname "$ScriptPath"); pwd)


ListOnly=0
Verbosity=${XTESTS_VERBOSITY:-${TEST_VERBOSITY:-3}}


# ##########################################################
# command-line handling

while [[ $# -gt 0 ]]; do

  case $1 in
    -l|--list-only)

      ListOnly=1
      ;;
    --verbosity)

      shift
      Verbosity=$1
      ;;
    --help)

      cat << EOF
ANGoLS provides Algorithms Not in the Go Language Standard library
Copyright (c) 2019-2025, Matthew Wilson and Synesis Information Systems
Runs all (matching) component and unit test programs

$ScriptPath [ ... flags/options ... ]

Flags/options:

    behaviour:

    -l
    --list-only
        lists the target programs but does not execute them

    --verbosity <verbosity>
        specifies an explicit verbosity for the unit-test(s)


    standard flags:

    --help
        displays this help and terminates

EOF

      exit 0
      ;;
    *)

      >&2 echo "$ScriptPath: unrecognised argument '$1'; use --help for usage"

      exit 1
      ;;
  esac

  shift
done


# ##########################################################
# main()


for f in $(find $Dir -type f -name '*_test.go')
do

    if [ $ListOnly -ne 0 ]; then

      echo "would execute $f:"

      continue
    fi

    if [ $Verbosity -ge 3 ]; then

      echo
    fi
    if [ $Verbosity -ge 2 ]; then

      echo "executing $f:"
    fi

    if [ $Verbosity -ge 2 ]; then

		go test -v "$f"
	else

		go test  "$f"
    fi

done


# ############################## end of file ############################# #

