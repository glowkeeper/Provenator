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

// ActivitiesContract - get  contract
func Artefacts () (*Spons.Spons) {

	contract, err := Spons.NewSpons(common.HexToAddress(viper.GetString("address.activitiesContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorEntitiesContract, err)
		return nil
	}
	return contract
}

// EntitiesContract - get contract
func EntitiesContract () (*Jobs.Jobs) {

	contract, err := Jobs.NewJobs(common.HexToAddress(viper.GetString("address.entitiesContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorActivitiesContract, err)
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
    Contracts.JobsContract = JobsContract()
	Contracts.SponsContract = SponsContract()
	Contracts.SubbiesContract = SubbiesContract()
}
