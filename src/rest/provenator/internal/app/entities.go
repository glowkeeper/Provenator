package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Entities"
    //"github.com/glowkeeper/Provenator/src/rest/provenator/utils"

    "go.uber.org/zap"
    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

func entitiesTotal() (int64) {

    num, err := contracts.Contracts.EntitiesContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesNum, zap.Error(err))
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// EntitiesTotal - get total Entities
func EntitiesTotal () ([]byte) {

	var result []byte
    num := entitiesTotal()
    thisJSON := &types.EntitiesTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorEntitiesNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// EntityTypes - get the type of entity
func EntityTypes (ref [32]byte) ([]byte) {

	var result []byte
    entityTypes := &types.EntityTypes{}

	eTypes, err := contracts.Contracts.EntitiesContract.GetEntityTypes(&bind.CallOpts{}, ref)
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesNum, zap.Error(err))
		return result
	}

	var i uint8
	for  ; i < uint8(len(eTypes)); i++ {
		if (eTypes[i]) {
			thisJSON := types.EntityType{Type: i}
			entityTypes.Types = append(entityTypes.Types, thisJSON)
		}
	}

    tResult, err := json.MarshalIndent(entityTypes, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorEntitiesType + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Entity - get a single entity
func Entity (ref [32]byte) ([]byte) {

	var result []byte

	//entities := &types.EntitiesAll{}
	entityAddress, err := contracts.Contracts.EntitiesContract.GetEntityContract(&bind.CallOpts{}, ref)
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesAll, zap.Error(err))
		return result
	}

	entityContract, err := Entities.NewEntityNode(entityAddress, contracts.Conn)
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesAll, zap.Error(err))
		return result
	}

	aEntity, err := entityContract.Get(&bind.CallOpts{})
	if (err != nil) {
		pkgLog.SLogger.Error(text.ErrorEntitiesAll, zap.Error(err))
		return result
	}

	entity := types.Entity {
		ID: fmt.Sprintf("%#x", ref),
		Name: aEntity.Name,
		EMail: aEntity.Email,
		URL: aEntity.Url,
	}

	tResult, err := json.MarshalIndent(&entity, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesAll + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	return result
}

// EntitiesList - entities list
func EntitiesList () ([]byte) {

	var result []byte
    entitiesRefs := &types.EntitiesList{}
    num := entitiesTotal()

    var i int64
    for ; i < num; i++ {
		ref, err := contracts.Contracts.EntitiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorEntitiesRef, zap.Error(err))
	        return result
        }
        thisRef := fmt.Sprintf("%#x", ref)
        entitiesRefs.ID = append(entitiesRefs.ID, thisRef)
    }

    tResult, err := json.MarshalIndent(entitiesRefs, "", "")
    if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesList + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// EntitiesAll - get all entities
func EntitiesAll () ([]byte) {

	var result []byte
	entities := &types.EntitiesAll{}

	num := entitiesTotal()
	//fmt.Println("Entities total: ", num)
    var i int64
    for i = 0; i < num; i++ {

		ref, err := contracts.Contracts.EntitiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorEntitiesRef, zap.Error(err))
	        return result
        }

		thisEntity := types.Entity{}
		entity := Entity(ref)
		json.Unmarshal(entity, &thisEntity)
		entities.Entities = append(entities.Entities, thisEntity)
	}

	tResult, err := json.MarshalIndent(&entities, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorEntitiesRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}
