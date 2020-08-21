#!/bin/bash

TARGETDIR="/home/steve/gitrepos/Provenator/src/rest/provenator/internal/contracts"

# Just change the names here, the rest will take care of itself
WORKS="Subbies"

# Need to stick each contract in its own directory, unfortunately, due to Strings library
WORKSDIR="${TARGETDIR}/${WORKS}/"

CONTRACTSDIR=(${WORKSDIR}})

# Clean up first
for DIR in "${CONTRACTSDIR[@]}"
do
    rm -rf $DIR 2>/dev/null
    mkdir $DIR
done

abigen --sol="./contracts/${WORKS}.sol" --pkg="${WORKS}" --out="${WORKSDIR}/${WORKS}.go"
