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

	copyrightHolderIds, err := contracts.Contracts.ArtefactsContract.GetWorkAuthors(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

    for i := 0; i < len(copyrightHolderIds); i++ {

		copyrightHolder, err := contracts.Contracts.EntitiesContract.GetEntity(&bind.CallOpts{}, copyrightHolderIds[i])
		if (err != nil) {
		    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
		    return result
		}

		thisCopyrightHolder := types.Entity {
		   ID: fmt.Sprintf("%#x", copyrightHolderIds[i]),
		   Name: copyrightHolder.Name,
		   EMail: copyrightHolder.Email,
		   URL: copyrightHolder.Url,
		}

		artefact.CopyrightHolders = append(artefact.CopyrightHolders, thisCopyrightHolder)
	}

	publisherIds, err := contracts.Contracts.ArtefactsContract.GetWorkPublishers(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

    for i := 0; i < len(publisherIds); i++ {

		publisher, err := contracts.Contracts.EntitiesContract.GetEntity(&bind.CallOpts{}, publisherIds[i])
		if (err != nil) {
		    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
		    return result
		}

		thisPublisher := types.Entity {
		   ID: fmt.Sprintf("%#x", publisherIds[i]),
		   Name: publisher.Name,
		   EMail: publisher.Email,
		   URL: publisher.Url,
		}

		artefact.Publishers = append(artefact.Publishers, thisPublisher)
	}

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

// ArtefactEntity - Artefacts for a given entity
func ArtefactEntity (ref [32]byte) ([]byte) {

	var result []byte
	artefacts := &types.WorksAll{}
	sRef := fmt.Sprintf("%#x", ref)
	//fmt.Println(sRef)

    num := artefactsTotal()
    var i int64
    for ; i < num; i++ {

		thisRef, err := contracts.Contracts.ArtefactsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorArtefactsRef, zap.Error(err))
	        return result
        }

		thisArtefact := types.Works{}
		artefact := Artefact(thisRef)
		json.Unmarshal(artefact, &thisArtefact)
		for i := 0; i < len(thisArtefact.Authors); i++ {

			if ( thisArtefact.Authors[i].ID == sRef ) {
				artefacts.Works = append(artefacts.Works, thisArtefact)
			}
		}
	}

	tResult, err := json.MarshalIndent(&artefacts, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}

// ArtefactAuthors - authors for a given artefact
func ArtefactAuthors (ref [32]byte) ([]byte) {

	var result []byte
	authors := &types.WorksAuthors{}
	//sRef := fmt.Sprintf("%#x", ref)
	//fmt.Println(sRef)

	authorIds, err := contracts.Contracts.ArtefactsContract.GetWorkAuthors(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

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

		authors.Authors = append(authors.Authors, thisAuthor)
	}

	tResult, err := json.MarshalIndent(&authors, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}

// ArtefactCopyrightHolders - copyright holders for a given artefact
func ArtefactCopyrightHolders (ref [32]byte) ([]byte) {

	var result []byte
	copyrightHolders := &types.WorksCopyrightHolders{}

	copyrightHolderIds, err := contracts.Contracts.ArtefactsContract.GetWorkAuthors(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

    for i := 0; i < len(copyrightHolderIds); i++ {

		copyrightHolder, err := contracts.Contracts.EntitiesContract.GetEntity(&bind.CallOpts{}, copyrightHolderIds[i])
		if (err != nil) {
		    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
		    return result
		}

		thisCopyrightHolder := types.Entity {
		   ID: fmt.Sprintf("%#x", copyrightHolderIds[i]),
		   Name: copyrightHolder.Name,
		   EMail: copyrightHolder.Email,
		   URL: copyrightHolder.Url,
		}

		copyrightHolders.CopyrightHolders = append(copyrightHolders.CopyrightHolders, thisCopyrightHolder)
	}

	tResult, err := json.MarshalIndent(&copyrightHolders, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}

// ArtefactPublishers - publishers for a given artefact
func ArtefactPublishers (ref [32]byte) ([]byte) {

	var result []byte
	publishers := &types.WorksPublishers{}

	publisherIds, err := contracts.Contracts.ArtefactsContract.GetWorkPublishers(&bind.CallOpts{}, ref)
	if (err != nil) {
	    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
	    return result
	}

    for i := 0; i < len(publisherIds); i++ {

		publisher, err := contracts.Contracts.EntitiesContract.GetEntity(&bind.CallOpts{}, publisherIds[i])
		if (err != nil) {
		    pkgLog.SLogger.Error(text.ErrorArtefactsAll, zap.Error(err))
		    return result
		}

		thisPublisher := types.Entity {
		   ID: fmt.Sprintf("%#x", publisherIds[i]),
		   Name: publisher.Name,
		   EMail: publisher.Email,
		   URL: publisher.Url,
		}

		publishers.Publishers = append(publishers.Publishers, thisPublisher)
	}

	tResult, err := json.MarshalIndent(&publishers, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorArtefactsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}
