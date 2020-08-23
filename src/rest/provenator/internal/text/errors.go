package text

const (
    // ErrorConfig - error
    ErrorConfig = "No configuration yaml (.provenator.yaml)"

    // ErrorServer - error
    ErrorServer = "Server failed"

    // ErrorEthereum - error
    ErrorEthereum = "Failed to connect to Etehreum network"

    // ErrorUnMarshall - unmarshalling problems
    ErrorUnMarshall = "Failed to unmashall struct"

    // ErrorNotFound - ErrorRinkeby
    ErrorNotFound = "Error 404 - Page not found"
    // ErrorPostDisallowed - error
    ErrorPostDisallowed = "This server only allows GET requests"

    // ErrorZap - error init'ing zap
    ErrorZap = "Can't initialize zap logger"

    // ErrorArtefactsContract - error
    ErrorArtefactsContract = "Couldn't create artefacts contract instance"

    // ErrorArtefactsRef - error
    ErrorArtefactsRef = "Error getting the activity reference"

    // ErrorArtefactsList - error
    ErrorArtefactsList = "Error getting activities List"

    // ErrorArtefactsAll - error
    ErrorArtefactsAll = "Error getting all activities"

    // ErrorArtefactsNum - Works error
    ErrorArtefactsNum = "Can't get activities total"
    
    // ErrorEntitiesContract - error
    ErrorEntitiesContract = "Couldn't create entity contract instance"

    // ErrorEntitiesRef - error
    ErrorEntitiesRef = "Error getting the entity reference"

    // ErrorEntitiesType - error
    ErrorEntitiesType = "Error getting the entity type"

    // ErrorEntitiesList - error
    ErrorEntitiesList = "Error getting entities list"

    // ErrorEntitiesAll - error
    ErrorEntitiesAll = "Error getting all entities"

    // ErrorEntitiesNum - Jobs error
    ErrorEntitiesNum = "Can't get entities total"
)
