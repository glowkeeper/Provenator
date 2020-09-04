package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"
    works "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Artefacts"
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
    thisJSON := &types.WorksTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorArtefactsNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Artefact - get a single artefacts record
func Artefact (ref [32]byte) ([]byte) {

	var result []byte
	artefactsAddress, err := contracts.Contracts.ArtefactsContract.GetWorkContract(&bind.CallOpts{}, ref)
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	artefactsContract, err := works.NewArtefactNode(artefactsAddress, contracts.Conn)
	if err != nil {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	aArtefacts, err := artefactsContract.Get(&bind.CallOpts{})
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	artefact := types.Works {
		ID: fmt.Sprintf("%#x", ref),
		WorkType: aArtefacts.WorkType,
		License: aArtefacts.License,
	    DateCreated: utils.GetString(aArtefacts.DateCreated),
	    DateModified: utils.GetString(aArtefacts.DateModified),
	    URL: aArtefacts.Url,
	    Name: aArtefacts.Name,
	    Description: aArtefacts.Description,
	}

	authorIds, err := contracts.Contracts.ArtefactsContract.GetWorkAuthors(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

	fmt.Println("Author: ", authorIds)

    for i := 0; i < len(authorIds); i++ {

		author, err := contracts.Contracts.EntitiesContract.GetEntity(&bind.CallOpts{}, authorIds[i])
		if (err != nil) {
		    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
		    return result
		}

		thisAuthor := types.Entity {
		   ID: fmt.Sprintf("%#x", authorIds[i]),
		   Name: author.Name,
		   EMail: author.Email,
		   URL: author.Url,
		}

		artefact.Authors = append(artefact.Authors, thisAuthor)
	}

	/*author := types.Entity {
	   ID: fmt.Sprintf("%#x", aArtefacts.Author.Id),
	   Name: aArtefacts.Author.Name,
	   EMail: aArtefacts.Author.Email,
	   URL: aArtefacts.Author.Url,
	}

	copyrightHolder := types.Entity {
	   ID: fmt.Sprintf("%#x", aArtefacts.CopyrightHolder.Id),
 	   Name: aArtefacts.CopyrightHolder.Name,
 	   EMail: aArtefacts.CopyrightHolder.Email,
 	   URL: aArtefacts.CopyrightHolder.Url,
	}

	publisher := types.Entity {
	   ID: fmt.Sprintf("%#x", aArtefacts.Publisher.Id),
 	   Name: aArtefacts.Publisher.Name,
 	   EMail: aArtefacts.Publisher.Email,
 	   URL: aArtefacts.Publisher.Url,
	}*/

	tResult, err := json.MarshalIndent(&artefact, "", "")
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
    worksRefs := &types.WorksList{}
    num := artefactsTotal()

    var i int64
    for ; i < num; i++ {
        ref, err := contracts.Contracts.ArtefactsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsRef, zap.Error(err))
	        return []byte(`{}`)
        }
        thisRef := fmt.Sprintf("%#x", ref)
        worksRefs.ID = append(worksRefs.ID, thisRef)
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
	artefacts := &types.WorksAll{}

    num := artefactsTotal()
    var i int64
    for ; i < num; i++ {

		ref, err := contracts.Contracts.ArtefactsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsRef, zap.Error(err))
	        return result
        }

		thisArtefact := types.Works{}
		artefact := Artefact(ref)
		json.Unmarshal(artefact, &thisArtefact)
		artefacts.Works = append(artefacts.Works, thisArtefact)
	}

	tResult, err := json.MarshalIndent(&artefacts, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}
