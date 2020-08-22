package app

import (
	"fmt"
	"encoding/json"
  	"math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

    "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/types"
	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/text"
    "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/contracts"
	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/contracts/Jobs"
    "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/utils"

    "go.uber.org/zap"
    pkgLog "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/pkg/log"
)

func getWorkStatus(j *types.Job, a common.Address, s uint8) ([]byte) {

	var result []byte
	job := types.Job {}

	for i := 0; i < len(j.Works); i++ {
		work := j.Works[i]
		if ( ( work.Subbie.Address == a ) && ( work.Status == s) ) {
			job.Ref = j.Ref
			job.Status = j.Status
			job.Location = j.Location
			job.Description = j.Description
			job.Works = append(job.Works, work)
		}
	}

	tResult, err := json.MarshalIndent(&job, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsAll + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	//fmt.Println(result)

	return result
}

func hasStatus(j *types.Job, s uint8) (bool) {

	contains := false
	if ( j.Status == s ) {
		contains = true
	}
	return contains
}

func hasAddress(j *types.Job, a common.Address) (bool) {

	contains := false
	for i := 0; i < len(j.Works); i++ {
		work := j.Works[i]
		if ( work.Subbie.Address == a ) {
			contains = true
			break
		}
	}

	return contains
}

func jobsTotal() (int64) {

    num, err := contracts.Contracts.JobsContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsNum, zap.Error(err))
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// JobsTotal - get total Jobs
func JobsTotal () ([]byte) {

	var result []byte
    num := jobsTotal()
    thisJSON := &types.JobsTotal{Total: num}

    tResult, err := json.MarshalIndent(thisJSON, "", "\t")
    if err != nil {
        pkgLog.SLogger.Error(text.ErrorJobsNum + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// Job - get a single job
func Job (ref [32]byte) ([]byte) {

	var result []byte
	//jobs := &types.JobsAll{}
	job := types.Job{}
	jobAddress, err := contracts.Contracts.JobsContract.GetJobContract(&bind.CallOpts{}, ref)
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
		return result
	}

	jobContract, err := Jobs.NewJobNode(jobAddress, contracts.Conn)
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
		return result
	}

	sRef, aJob, err := jobContract.Get(&bind.CallOpts{})
	//sRef, aJob, err := contracts.Contracts.JobsContract.GetJob(&bind.CallOpts{}, ref)
	//fmt.Sprintf("sRef %#x", sRef)
	if (err != nil || sRef != ref ) {
		pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
		return result
	}

	job = types.Job {
		Ref: fmt.Sprintf("%#x", ref),
		Status: aJob.Status,
		Location: aJob.Location,
		Description: aJob.Description,
	}

	workNum, err := jobContract.GetNum(&bind.CallOpts{})
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
		return result
	}
	smallNum := workNum.Int64()
	//fmt.Println("num work: ", smallNum)

	var i, j int64
	//var i int64
	for i = 0; i < smallNum; i++ {

		workRef, err := jobContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
		//fmt.Printf("workref %#x\n", workRef)
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
			return result
		}

		workAddress, err := jobContract.GetWorkContract(&bind.CallOpts{}, ref, workRef)
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
			return result
		}

		workContract, err := Jobs.NewWorkNode(workAddress, contracts.Conn)
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
			return result
		}

		wRef, aWork, err := workContract.Get(&bind.CallOpts{})
		if (err != nil || wRef != workRef) {
			pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
			return result
		}

		startDate := utils.GetString(aWork.StartDate)
		endDate := utils.GetString(aWork.EndDate)
		startTime := utils.GetString(aWork.StartTime)
		endTime := utils.GetString(aWork.EndTime)

		subbie := types.Subbie{}
		if ( aWork.Subbie[0] == 0) {
			address := new(common.Address)
			subbie.Address = *address
		} else {

			sAddress, aSubbie, err := contracts.Contracts.SubbiesContract.GetSubbie(&bind.CallOpts{}, aWork.Subbie)
			//fmt.Println("sAddress %#x", sAddress)
			if (err != nil || sAddress != aWork.Subbie ) {
				pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
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
		}

		spons := CostModel(aWork.SponsRef)

		work := types.Work {
			Ref: fmt.Sprintf("%#x", workRef),
			Status: aWork.Status,
			Subbie: subbie,
			StartDate: string(startDate),
			EndDate: string(endDate),
			StartTime: string(startTime),
			EndTime: string(endTime),
			Spons: spons,
			Quantity: aWork.Quantity.Int64(),
			Cost: aWork.Cost.Int64(),
			Description: aWork.Description,
		}

		infoNum, err := workContract.GetNum(&bind.CallOpts{})
		if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
			return result
		}
		smallInfoNum := infoNum.Int64()
		//fmt.Println("num info: ", smallInfoNum)

		for j = 0; j < smallInfoNum; j++ {
			infoRef, err := workContract.GetReference(&bind.CallOpts{}, big.NewInt(j))
			//fmt.Printf("workref %#x\n", workRef)
			if err != nil {
				pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
				return result
			}

			infoAddress, err := workContract.GetInfoContract(&bind.CallOpts{}, workRef, infoRef)
			if err != nil {
				pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
				return result
			}

			infoContract, err := Jobs.NewInfoNode(infoAddress, contracts.Conn)
			if err != nil {
				pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
				return result
			}

			iRef, aInfo, err := infoContract.Get(&bind.CallOpts{})
			if (err != nil || iRef != infoRef) {
				pkgLog.SLogger.Error(text.ErrorJobsAll, zap.Error(err))
				return result
			}

			thisInfo := types.ExtraInfo {
				Ref: fmt.Sprintf("%#x", infoRef),
				FileHash: aInfo.FileHash,
			    FileName: aInfo.FileName,
			    FileUrl: aInfo.FileUrl,
			    Description: aInfo.Description,
			}
			work.Info = append(work.Info, thisInfo)
		}

		job.Works = append(job.Works, work)
	}

	tResult, err := json.MarshalIndent(&job, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsAll + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	return result
}

// JobsList - jobs list
func JobsList () ([]byte) {

	var result []byte
    jobsRefs := &types.JobsList{}
    num := jobsTotal()

    var i int64
    for ; i < num; i++ {
		ref, err := contracts.Contracts.JobsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
	        return result
        }
        thisRef := fmt.Sprintf("%#x", ref)
        jobsRefs.Ref = append(jobsRefs.Ref, thisRef)
    }

    tResult, err := json.MarshalIndent(jobsRefs, "", "")
    if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsList + " - " + text.ErrorUnMarshall, zap.Error(err))
        return result
    }
	result = append(result, tResult...)

    return result
}

// JobsForAddress - Get all jobs assigned to a given address
func JobsForAddress (a common.Address) ([]byte) {

	var result []byte
	jobs := &types.JobsAll{}

	num := jobsTotal()
    var i int64
    for i = 0; i < num; i++ {

		ref, err := contracts.Contracts.JobsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
	        return result
        }

		thisJob := types.Job{}
		job := Job(ref)
		json.Unmarshal(job, &thisJob)
		if ( hasAddress(&thisJob, a) ) {

			subbieJob := types.Job{}

			for i := 0; i < len(thisJob.Works); i++ {
				work := thisJob.Works[i]
				if ( work.Subbie.Address == a ) {
					//fmt.Println(work.Subbie, a)
					subbieJob.Ref = thisJob.Ref
					subbieJob.Status = thisJob.Status
					subbieJob.Location = thisJob.Location
					subbieJob.Description = thisJob.Description
					subbieJob.Works = append(subbieJob.Works, work)
				}
			}
			jobs.Jobs = append(jobs.Jobs, subbieJob)
		}
	}

	tResult, err := json.MarshalIndent(&jobs, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	return result
}

// JobsForAddressStatus - get jobs for given address with a given status
func JobsForAddressStatus (a common.Address, s uint8) ([]byte) {

	var result []byte
	jobs := &types.JobsAll{}

	num := jobsTotal()
    var i int64
    for i = 0; i < num; i++ {

		ref, err := contracts.Contracts.JobsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
	        return result
        }

		thisJob := types.Job{}
		job := Job(ref)
		json.Unmarshal(job, &thisJob)
		if ( hasAddress(&thisJob, a) ) {
			//fmt.Println(i, a, thisJob)
			newJob := types.Job{}
			work := getWorkStatus(&thisJob, a, s)
			json.Unmarshal(work, &newJob)
			jobs.Jobs = append(jobs.Jobs, newJob)
		}
	}

	tResult, err := json.MarshalIndent(&jobs, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
	return result
}

// JobsForStatus - get jobs with a given status
func JobsForStatus (s uint8) ([]byte) {

	var result []byte
	jobs := &types.JobsAll{}

	num := jobsTotal()
    var i int64
    for i = 0; i < num; i++ {

		ref, err := contracts.Contracts.JobsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
	        return result
        }

		thisJob := types.Job{}
		job := Job(ref)
		json.Unmarshal(job, &thisJob)
		if ( hasStatus(&thisJob, s) ) {
			jobs.Jobs = append(jobs.Jobs, thisJob)
		}
	}

	tResult, err := json.MarshalIndent(&jobs, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}

// JobsAll - get all jobs
func JobsAll () ([]byte) {

	var result []byte
	jobs := &types.JobsAll{}

	num := jobsTotal()
	//fmt.Println("Jobs total: ", num)
    var i int64
    for i = 0; i < num; i++ {

		ref, err := contracts.Contracts.JobsContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
			pkgLog.SLogger.Error(text.ErrorJobsRef, zap.Error(err))
	        return result
        }

		thisJob := types.Job{}
		job := Job(ref)
		json.Unmarshal(job, &thisJob)
		jobs.Jobs = append(jobs.Jobs, thisJob)
	}

	tResult, err := json.MarshalIndent(&jobs, "", "")
	if err != nil {
		pkgLog.SLogger.Error(text.ErrorJobsRef + " - " + text.ErrorUnMarshall, zap.Error(err))
		return result
	}
	result = append(result, tResult...)
    return result
}
