package types

import (
	//"time"

	//"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Artefacts"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts/Entities"
)

// Contracts - keeps pointers to all ethereum contracts
type Contracts struct {
	ArtefactsContract    	*Artefacts.Artefacts
	EntitiesContract    	*Entities.Entities
}

// Entity - creative entities
type Entity struct {
	ID				string	`json:"id"`
    Name        	string  `json:"name"`
    EMail        	string  `json:"email"`
    URL			 	string  `json:"url"`
}

// EntitiesAll - get all entities
type EntitiesAll struct {
	Entities []Entity 		`json:"entities"`
}

// EntitiesTotal - get the total number of entities
type EntitiesTotal struct {
	Total	int64    `json:"total"`
}

// EntitiesList - a list if all entities
type EntitiesList struct {
	ID	[]string `json:"id"`
}

// Works - creative works data
type Works struct {
    WorkType			uint8	`json:"workType"`
    License		        uint8	`json:"license"`
    ID					string	`json:"id"`
    DateCreated			string	`json:"dateCreated"`
    DateModified		string	`json:"dateModified"`
    URL					string	`json:"url"`
    Name				string	`json:"name"`
    Description			string	`json:"description"`
    Author				Entity	`json:"author"`
    CopyrightHolder		Entity	`json:"copyrightHolder"`
    Publisher			Entity	`json:"publisher"`
}

// WorksAll - get all works
type WorksAll struct {
	Works []Works `json:"creativeWorks"`
}

// WorksTotal - get the total number of Cost Models
type WorksTotal struct {
	Total	int64    `json:"total"`
}

// WorksList - a list if all activities
type WorksList struct {
	ID	[]string `json:"id"`
}
