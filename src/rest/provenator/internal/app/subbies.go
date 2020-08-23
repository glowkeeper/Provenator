package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"
    //"github.com/glowkeeper/Provenator/src/rest/provenator/utils"

    "go.uber.org/zap"
    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

func subbiesTotal() (int64) {

    num, err := contracts.Contracts.SubbiesContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorSubbiesNum, zap.Error(err))
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// SubbiesTotal - get total Subbies
func SubbiesTotal () ([]byte) {

	var result []byte
    num := subbiesTotal()
    thisJSON := &types.SubbiesTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorSubbiesNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Subbie - get a single subbie
func Subbie (ref [32]byte) ([]byte) {

	var result []byte

	address := common.BytesToAddress(ref[:])
	sAddress, aSubbie, err := contracts.Contracts.SubbiesContract.GetSubbie(&bind.CallOpts{}, address)
	//fmt.Sprintf("sRef %#x", sRef)
	if (err != nil || sAddress != address ) {
		pkgLog.SLogger.Error(text.ErrorSubbiesAll, zap.Error(err))
		return result
	}

	subbie := types.Subbie {
		Address: sAddress,
		MainType: aSubbie.MainType,
		Status: aSubbie.Status,
		Name: aSubbie.Name,
		Location: aSubbie.Location,
		Description: aSubbie.Description,
	}

	tResult, err := json.MarshalIndent(&subbie, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorSubbiesAll + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	return result
}

// SubbiesList - subbies list
func SubbiesList () ([]byte) {

	var result []byte
    subbiesAddresses := &types.SubbiesList{}
    num := subbiesTotal()

    var i int64
    for ; i < num; i++ {
        address, err := contracts.Contracts.SubbiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorSubbiesAddress, zap.Error(err))
	        return []byte(`{}`)
        }
        thisAddress := fmt.Sprintf("%#x", address[12:32])
        subbiesAddresses.Address = append(subbiesAddresses.Address, thisAddress)
    }

    tResult, err := json.MarshalIndent(subbiesAddresses, "", "")
    if err != nil {
		pkgLog.SLogger.Error(text.ErrorSubbiesList + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// SubbiesAll - get all subbies
func SubbiesAll () ([]byte) {

	var result []byte
	subbies := &types.SubbiesAll{}
	subbie := types.Subbie{}

	num := subbiesTotal()
	var i int64
	for i = 0; i < num; i++ {

	    ref, err := contracts.Contracts.SubbiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
	    //fmt.Printf("ref %#x", ref)
	    if err != nil {
	        pkgLog.SLogger.Error(text.ErrorSubbiesAddress, zap.Error(err))
	        return result
	    }

		address := common.BytesToAddress(ref[:])
	    sAddress, aSubbie, err := contracts.Contracts.SubbiesContract.GetSubbie(&bind.CallOpts{}, address)
	    //fmt.Sprintf("sRef %#x", sRef)
	    if (err != nil || sAddress != address ) {
	        pkgLog.SLogger.Error(text.ErrorSubbiesAll, zap.Error(err))
	        return result
	    }

		subbie = types.Subbie {
	        Address: sAddress,
			MainType: aSubbie.MainType,
	        Status: aSubbie.Status,
			Name: aSubbie.Name,
	        Location: aSubbie.Location,
	        Description: aSubbie.Description,
	    }
	    subbies.Subbies = append(subbies.Subbies, subbie)
	}

	tResult, err := json.MarshalIndent(&subbies, "", "")
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorSubbiesAll + " - " + text.ErrorUnMarshall, zap.Error(err))
	    return result
	}
	result = append(result, tResult...)
	return result
}
