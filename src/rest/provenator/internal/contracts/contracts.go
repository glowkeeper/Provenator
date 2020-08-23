package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"go.uber.org/zap"
	"github.com/spf13/viper"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/types"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Jobs"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Spons"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Subbies"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"

    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

// Contracts - zeus ethereum contracts
var Contracts *types.Contracts
// Conn - ethereum connection
var Conn *ethclient.Client

// SponsContract - get  contract
func SponsContract () (*Spons.Spons) {

	contract, err := Spons.NewSpons(common.HexToAddress(viper.GetString("address.sponsContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorSponsContract, err)
		return nil
	}
	return contract
}

// JobsContract - get contract
func JobsContract () (*Jobs.Jobs) {

	contract, err := Jobs.NewJobs(common.HexToAddress(viper.GetString("address.jobsContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorJobsContract, err)
		return nil
	}
	return contract
}

// SubbiesContract - get contract
func SubbiesContract () (*Subbies.Subbies) {

	contract, err := Subbies.NewSubbies(common.HexToAddress(viper.GetString("address.subbiesContract")), Conn)
	if err != nil {
		log.Fatalf("%s: %v", text.ErrorJobsContract, err)
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
