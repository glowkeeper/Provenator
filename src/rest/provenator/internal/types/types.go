package types

import (
	//"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/contracts/Jobs"
	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/contracts/Spons"
	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/contracts/Subbies"
)

// Work Models
const (
	COSTMODEL uint8 = iota + 1
	FUNCTIONALUNIT
	PERSM
	MEASUREDWORKS
)

// Status
const (
	AVAILABLE uint8 = iota + 1
    ISBOOKED
    INPROGRESS
    COMPLETE
    SIGNEDOFF
    SETTLED
)

// Contracts - keeps pointers to all ethereum contracts
type Contracts struct {
	JobsContract    	*Jobs.Jobs
	SponsContract    	*Spons.Spons
	SubbiesContract    	*Subbies.Subbies
}

//CostModel - works type
type CostModel struct {
	Ref				string	`json:"ref"`
    Unit        	uint8  	`json:"unit"`
    Rate        	int64  	`json:"rate"`
    Description 	string  `json:"description"`
}

//FunctionalUnit - works type
type FunctionalUnit struct {
	Ref					string	`json:"ref"`
    FunctionalUnit		uint8  	`json:"unit"`
    LowerBound			int64   `json:"lowerBound"`
    UpperBound			int64   `json:"upperBound"`
    Description 		string  `json:"description"`
}

//PerSM  - works type
type PerSM struct {
	Ref					string	`json:"ref"`
	LowerBound			int64   `json:"lowerBound"`
	UpperBound			int64   `json:"upperBound"`
    Description 		string  `json:"description"`
}

//MeasuredWorks  - works type
 type MeasuredWorks struct {
 	Ref					string	`json:"ref"`
    Unit            	uint8  	`json:"unit"`
    PrimeCost			int64   `json:"primeCost"`
    LabourMinutes		int64   `json:"labourMinutes"`
    LaboutCost			int64   `json:"laboutCost"`
    Plant				int64   `json:"plant"`
    Material			int64   `json:"material"`
    Description 		string  `json:"description"`
}

// CostModelAll - get all works
type CostModelAll struct {
	Works []CostModel `json:"costModelWorks"`
}

// FunctionalUnitAll - get all works
type FunctionalUnitAll struct {
	Works []FunctionalUnit `json:"functionalUnitWorks"`
}

// PerSMAll - get all works
type PerSMAll struct {
	Works []PerSM `json:"works"`
}

// MeasuredWorksAll - get all works
type MeasuredWorksAll struct {
	Works []MeasuredWorks `json:"works"`
}

//ExtraInfo - extra stuff confiirming works
type ExtraInfo struct {
	Ref				string		`json:"infoRef"`
    FileHash 		string      `json:"fileHash"`
    FileName		string      `json:"fileName"`
    FileUrl			string      `json:"fileUrl"`
    Description		string		`json:"description"`
}

// WorksTotal - get the total number of Cost Models
type WorksTotal struct {
	Total	int64    `json:"total"`
}

// WorksList - a list if all activities
type WorksList struct {
	Ref	[]string `json:"ref"`
}

// Work - info output by API
type Work struct {
	Ref				string			`json:"workRef"`
	Status 			uint8  			`json:"status"`
    Subbie          Subbie			`json:"subbie"`
	StartDate		string			`json:"startDate"`
	EndDate			string			`json:"endDate"`
	StartTime		string			`json:"startTime"`
	EndTime			string			`json:"endTime"`
	Spons			CostModel		`json:"spons"`
    Quantity		int64   		`json:"quantity"`
    Cost			int64   		`json:"cost"`
    Description		string			`json:"description"`
	Info			[]ExtraInfo 	`json:"extraInfo"`
}

// Job - info' output about a job
type Job struct {
	Ref			string		`json:"ref"`
	Status 		uint8  		`json:"status"`
	Location	string		`json:"location"`
	Description	string		`json:"description"`
	Works       []Work		`json:"works"`
}

// JobsAll - list of all Job info
type JobsAll struct {
	Jobs []Job `json:"jobs"`
}

// JobsTotal - get the total number of jobs
type JobsTotal struct {
	Total 		int64    `json:"total"`
}

// JobsList - a list if all jobs
type JobsList struct {
	Ref			[]string `json:"ref"`
}

// Subbie - info' about a job
type Subbie struct {
	Address		common.Address	`json:"address"`
	MainType	uint8  			`json:"mainType"`
	Status 		uint8  			`json:"status"`
	Name		string			`json:"name"`
	Location	string			`json:"location"`
	Description	string			`json:"description"`
}

// SubbiesAll - list of all Subbie info
type SubbiesAll struct {
	Subbies []Subbie `json:"subbies"`
}

// SubbiesTotal - get the total number of jobs
type SubbiesTotal struct {
	Total 		int64    `json:"total"`
}

// SubbiesList - a list if all jobs
type SubbiesList struct {
	Address			[]string `json:"address"`
}
