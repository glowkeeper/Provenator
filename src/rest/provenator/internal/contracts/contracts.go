package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"go.uber.org/zap"
	"github.com/spf13/viper"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Artefacts"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Entities"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"

    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

// Contracts - provenator ethereum contracts
var Contracts *types.Contracts
// Conn - ethereum connection
var Conn *ethclient.Client

// ArtefactsContract - get  contract
func ArtefactsContract () (*Artefacts.Artefacts) {

	contract, err := Artefacts.NewArtefacts(common.HexToAddress(viper.GetString("address.artefactsContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorArtefactsContract, err)
		return nil
	}
	return contract
}

// EntitiesContract - get contract
func EntitiesContract () (*Entities.Entities) {

	contract, err := Entities.NewEntities(common.HexToAddress(viper.GetString("address.entitiesContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorEntitiesContract, err)
		return nil
	}
	return contract
}

// InitContracts - setup the Ether' contracts
func InitContracts () {

	aEther := viper.GetString("address.ethereum")
	conn, err := ethclient.Dial(aEther)
	if err != nil {
        pkgLog.SLogger.Fatal(text.ErrorEthereum, zap.Error(err))
	}

	Conn = conn
	Contracts = new(types.Contracts)
    Contracts.ArtefactsContract = ArtefactsContract()
	Contracts.EntitiesContract = EntitiesContract()
}
