package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Artefacts"
    "github.com/glowkeeper/Provenator/src/rest/provenator/utils"

    "go.uber.org/zap"
    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

func artefactsTotal() (int64) {

    num, err := contracts.Contracts.ArtefactsContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsNum, zap.Error(err))
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// ArtefactsTotal - get total Artefacts
func ArtefactsTotal () ([]byte) {

	var result []byte
    num := artefactsTotal()
    thisJSON := &types.ArtefactsTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorArtefactsNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Artefacts - get a single artefacts record
func Artefacts (ref [32]byte) ([]byte) {

	var result []byte
	//entities := &types.EntitiesAll{}
	artefacts := types.Artefacts{}
	artefactsAddress, err := contracts.Contracts.ArtefactsContract.GetArtefactsContract(&bind.CallOpts{}, ref)
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	artefactsContract, err := Entities.NewArtefactsNode(artefactsAddress, contracts.Conn)
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	aArtefacts, err := artefactsContract.Get(&bind.CallOpts{})
	//sRef, aArtefacts, err := contracts.Contracts.EntitiesContract.GetArtefacts(&bind.CallOpts{}, ref)
	//fmt.Sprintf("sRef %#x", sRef)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	artefacts = types.Artefacts {
		WorkType: aArtefacts.WorkType,
		License: aArtefacts.License,
	    ID: fmt.Sprintf("%#x", ref),
	    DateCreated: aArtefacts.DateCreated,
	    DateModified: aArtefacts.DateModified,
	    URL: aArtefacts.URL,
	    Name: aArtefacts.Name,
	    Description: aArtefacts.Description,
	    Author: {
		   ID: fmt.Sprintf("%#x", ref),
		   Name: aEntity.Name,
		   EMail: aEntity.Location,
		   URL: aEntity.Description
	   },
	    CopyrightHolder:  {
		   ID: fmt.Sprintf("%#x", ref),
		   Name: aEntity.Name,
		   EMail: aEntity.Location,
		   URL: aEntity.Description
	   },
	    Publisher:  {
		   ID: fmt.Sprintf("%#x", ref),
		   Name: aEntity.Name,
		   EMail: aEntity.Location,
		   URL: aEntity.Description
	   }
	}

	tResult, err := json.MarshalIndent(&artefacts, "", "")
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll + " - " + text.ErrorUnMarshall, zap.Error(err))
	    return result
	}
	result = append(result, tResult...)
	return result
}

// ArtefactsList - artefactsref list
func ArtefactsList () ([]byte) {

	var result []byte
    worksRefs := &types.ArtefactsList{}
    num := artefactsTotal()

    var i int64
    for ; i < num; i++ {
        ref, err := contracts.Contracts.ArtefactsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsRef, zap.Error(err))
	        return []byte(`{}`)
        }
        thisRef := fmt.Sprintf("%#x", ref)
        worksRefs.Ref = append(worksRefs.Ref, thisRef)
    }

    tResult, err := json.MarshalIndent(worksRefs, "", "")
    if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsList + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// ArtefactsAll - get all artefacts
func ArtefactsAll () ([]byte) {

	var result []byte

	costModelArtefacts := &types.CostModelAll{}
	functionalUnitArtefacts := &types.FunctionalUnitAll{}
	perSMArtefacts := &types.PerSMAll{}
	measuredArtefactsArtefacts := &types.MeasuredArtefactsAll{}

    num := artefactsTotal()
    var i int64
    for ; i < num; i++ {

		ref, err := contracts.Contracts.ArtefactsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsRef, zap.Error(err))
	        return result
        }

		pType, err := contracts.Contracts.ArtefactsContract.GetType(&bind.CallOpts{}, ref);
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsType, zap.Error(err))
	        return result
        }

		switch pType {
		case types.COSTMODEL:

			tCostModel := CostModel(ref)
			costModelArtefacts.Artefacts = append(costModelArtefacts.Artefacts, tCostModel)

		case types.FUNCTIONALUNIT:

			tFunctionalUnit := FunctionalUnit(ref)
			functionalUnitArtefacts.Artefacts = append(functionalUnitArtefacts.Artefacts, tFunctionalUnit)

		case types.PERSM:

			tSM := PerSM(ref)
			perSMArtefacts.Artefacts = append(perSMArtefacts.Artefacts, tSM)

		case types.MEASUREDWORKS:

			tMeasuredArtefacts := MeasuredArtefacts(ref)
			measuredArtefactsArtefacts.Artefacts = append(measuredArtefactsArtefacts.Artefacts, tMeasuredArtefacts)

		default:

			return result
		}
    }

	if ( len(costModelArtefacts.Artefacts) != 0 ) {
		tResult, err := json.MarshalIndent(&costModelArtefacts, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(functionalUnitArtefacts.Artefacts) != 0 ) {
		tResult, err := json.MarshalIndent(&functionalUnitArtefacts, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorFunctionalUnit + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(perSMArtefacts.Artefacts) != 0 ) {
		tResult, err := json.MarshalIndent(&perSMArtefacts, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorSM + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(measuredArtefactsArtefacts.Artefacts) != 0 ) {
		tResult, err := json.MarshalIndent(&measuredArtefactsArtefacts, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorMeasuredArtefacts + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

    return result
}
