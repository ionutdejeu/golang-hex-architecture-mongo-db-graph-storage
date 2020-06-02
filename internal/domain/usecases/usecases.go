
import (
	"context"
	"errors"
	"strings"
	"time"

)

type UseCaseAdd struct {
	Saver saver
}

type UseCaseRead struct {
	Reader reader
}

type Saver interface { 
	Store (*entity) (uuid string, err error)
}
type Reader interface {
	ReadOne()
	ReadList()
}



