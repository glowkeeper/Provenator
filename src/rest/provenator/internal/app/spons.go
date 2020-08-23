package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"

    "go.uber.org/zap"
    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

// CostModel - get Cost Model works
func CostModel (ref [32]byte) (types.CostModel) {

	tCostModel := types.CostModel{}
	sRef, aCostModel, err := contracts.Contracts.SponsContract.GetCostModel(&bind.CallOpts{}, ref)
	if (err != nil || sRef != ref ) {
		pkgLog.SLogger.Error(text.ErrorCostModel, zap.Error(err))
		return tCostModel
	}

	tCostModel = types.CostModel {
		Ref: fmt.Sprintf("%#x", ref),
		Unit: aCostModel.Unit,
		Rate: aCostModel.Rate.Int64(),
		Description: aCostModel.Description,
	}

	return tCostModel
}

// FunctionalUnit - get functional unit works
func FunctionalUnit (ref [32]byte) (types.FunctionalUnit) {

	tFunctionalUnit := types.FunctionalUnit{}
	sRef, aFunctionalUnit, err := contracts.Contracts.SponsContract.GetFunctionalUnit(&bind.CallOpts{}, ref)
	if (err != nil || sRef != ref ) {
		pkgLog.SLogger.Error(text.ErrorFunctionalUnit, zap.Error(err))
		return tFunctionalUnit
	}

	tFunctionalUnit = types.FunctionalUnit {
		Ref: fmt.Sprintf("%#x", ref),
		FunctionalUnit: aFunctionalUnit.Unit,
		LowerBound: aFunctionalUnit.LowerBound.Int64(),
		UpperBound: aFunctionalUnit.UpperBound.Int64(),
		Description: aFunctionalUnit.Description,
	}

	return tFunctionalUnit
}

// PerSM - get per SM works
func PerSM (ref [32]byte) (types.PerSM) {

	tSM := types.PerSM{}
	sRef, aSM, err := contracts.Contracts.SponsContract.GetSM(&bind.CallOpts{}, ref)
	if (err != nil || sRef != ref ) {
		pkgLog.SLogger.Error(text.ErrorSM, zap.Error(err))
		return tSM
	}

	tSM = types.PerSM {
		Ref: fmt.Sprintf("%#x", ref),
		LowerBound: aSM.LowerBound.Int64(),
		UpperBound: aSM.UpperBound.Int64(),
		Description: aSM.Description,
	}

	return tSM
}

// MeasuredWorks - get measured works
func MeasuredWorks (ref [32]byte) (types.MeasuredWorks) {

	tMeasuredWorks := types.MeasuredWorks{}
	sRef, aMeasuredWorks, err := contracts.Contracts.SponsContract.GetMeasuredWorks(&bind.CallOpts{}, ref)
	if (err != nil || sRef != ref ) {
		pkgLog.SLogger.Error(text.ErrorMeasuredWorks, zap.Error(err))
		return tMeasuredWorks
	}

	tMeasuredWorks = types.MeasuredWorks {
		Ref: fmt.Sprintf("%#x", ref),
		Unit: aMeasuredWorks.Unit,
		PrimeCost: aMeasuredWorks.PrimeCost.Int64(),
	    LabourMinutes: aMeasuredWorks.LabourMinutes.Int64(),
	    LaboutCost: aMeasuredWorks.LaboutCost.Int64(),
	    Plant: aMeasuredWorks.Plant.Int64(),
	    Material: aMeasuredWorks.Material.Int64(),
		Description: aMeasuredWorks.Description,
	}

	return tMeasuredWorks
}

func sponsTotal() (int64) {

    num, err := contracts.Contracts.SponsContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorWorksNum, zap.Error(err))
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// SponsTotal - get total Works
func SponsTotal () ([]byte) {

	var result []byte
    num := sponsTotal()
    thisJSON := &types.WorksTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorWorksNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Spons - get a single spons record
func Spons (ref [32]byte) ([]byte) {

	var result []byte

	pType, err := contracts.Contracts.SponsContract.GetType(&bind.CallOpts{}, ref);
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorWorksType, zap.Error(err))
		return result
	}

	switch pType {
	case types.COSTMODEL:

		tCostModel := CostModel(ref)
		tResult, err := json.MarshalIndent(&tCostModel, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)

	case types.FUNCTIONALUNIT:

		tFunctionalUnit := FunctionalUnit(ref)
		tResult, err := json.MarshalIndent(&tFunctionalUnit, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)

	case types.PERSM:

		tSM := PerSM(ref)
		tResult, err := json.MarshalIndent(&tSM, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)

	case types.MEASUREDWORKS:

		tMeasuredWorks := MeasuredWorks(ref)
		tResult, err := json.MarshalIndent(&tMeasuredWorks, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)

	default:

		return result
	}

    return result
}

// SponsList - sponsref list
func SponsList () ([]byte) {

	var result []byte
    worksRefs := &types.WorksList{}
    num := sponsTotal()

    var i int64
    for ; i < num; i++ {
        ref, err := contracts.Contracts.SponsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorWorksRef, zap.Error(err))
	        return []byte(`{}`)
        }
        thisRef := fmt.Sprintf("%#x", ref)
        worksRefs.Ref = append(worksRefs.Ref, thisRef)
    }

    tResult, err := json.MarshalIndent(worksRefs, "", "")
    if err != nil {
		pkgLog.SLogger.Error(text.ErrorWorksList + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// SponsAll - get all spons
func SponsAll () ([]byte) {

	var result []byte

	costModelWorks := &types.CostModelAll{}
	functionalUnitWorks := &types.FunctionalUnitAll{}
	perSMWorks := &types.PerSMAll{}
	measuredWorksWorks := &types.MeasuredWorksAll{}

    num := sponsTotal()
    var i int64
    for ; i < num; i++ {

		ref, err := contracts.Contracts.SponsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorWorksRef, zap.Error(err))
	        return result
        }

		pType, err := contracts.Contracts.SponsContract.GetType(&bind.CallOpts{}, ref);
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorWorksType, zap.Error(err))
	        return result
        }

		switch pType {
		case types.COSTMODEL:

			tCostModel := CostModel(ref)
			costModelWorks.Works = append(costModelWorks.Works, tCostModel)

		case types.FUNCTIONALUNIT:

			tFunctionalUnit := FunctionalUnit(ref)
			functionalUnitWorks.Works = append(functionalUnitWorks.Works, tFunctionalUnit)

		case types.PERSM:

			tSM := PerSM(ref)
			perSMWorks.Works = append(perSMWorks.Works, tSM)

		case types.MEASUREDWORKS:

			tMeasuredWorks := MeasuredWorks(ref)
			measuredWorksWorks.Works = append(measuredWorksWorks.Works, tMeasuredWorks)

		default:

			return result
		}
    }

	if ( len(costModelWorks.Works) != 0 ) {
		tResult, err := json.MarshalIndent(&costModelWorks, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorCostModel + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(functionalUnitWorks.Works) != 0 ) {
		tResult, err := json.MarshalIndent(&functionalUnitWorks, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorFunctionalUnit + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(perSMWorks.Works) != 0 ) {
		tResult, err := json.MarshalIndent(&perSMWorks, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorSM + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

	if ( len(measuredWorksWorks.Works) != 0 ) {
		tResult, err := json.MarshalIndent(&measuredWorksWorks, "", "")
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorMeasuredWorks + " - " + text.ErrorUnMarshall, zap.Error(err))
	        return result
	    }
		result = append(result, tResult...)
	}

    return result
}
