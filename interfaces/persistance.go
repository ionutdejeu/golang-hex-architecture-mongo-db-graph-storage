package intefaces 

import (
	"github.com/ionutdejeu/golang-hex-architecture-mongo-db-graph-storage/internal/domain/models"
)

type Persistance interface { 
	Save()
	
}

