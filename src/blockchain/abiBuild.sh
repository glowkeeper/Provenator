#!/bin/bash

TARGETDIR="/home/steve/gitrepos/Provenator/src/rest/provenator/internal/contracts"

# Just change the names here, the rest will take care of itself
ARTEFACTS="Works"
ENTITIES="Entities"

# Need to stick each contract in its own directory, unfortunately, due to Strings library
ARTEFACTSDIR="${TARGETDIR}/${ARTEFACTS}/"
ENTITIESDIR="${TARGETDIR}/${ENTITIES}/"

CONTRACTSDIR=(${ARTEFACTSDIR} ${ENTITIESDIR})

# Clean up first
for DIR in "${CONTRACTSDIR[@]}"
do
    rm -rf $DIR 2>/dev/null
    mkdir $DIR
done

abigen --sol="./contracts/${ARTEFACTS}.sol" --pkg="${ARTEFACTS}" --out="${ARTEFACTSDIR}/${ARTEFACTS}.go"
abigen --sol="./contracts/${ENTITIES}.sol" --pkg="${ENTITIES}" --out="${ENTITIESDIR}/${ENTITIES}.go"
