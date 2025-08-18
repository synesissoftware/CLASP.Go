#! /bin/bash

ScriptPath=$0
Dir=$(cd $(dirname "$ScriptPath"); pwd)


mkdir -p "$Dir/scratch/build-artefacts/"


# library
echo "building library ..."
go build -v $Dir

# examples
if [ -d "$Dir/examples" ];then

	echo "building examples ..."
	cd "$Dir/scratch/build-artefacts/"
	find "$Dir/examples" -type f -name '*.go' -print -exec go build -v {} \;
	cd ->/dev/null
fi


# tests
if [ -d "$Dir/tests" ];then

	echo "building tests ..."
	cd "$Dir/scratch/build-artefacts/"
	find "$Dir/tests" -type f -name '*.go' -print -exec go build -v {} \;
	cd ->/dev/null
fi


# ############################## end of file ############################# #

