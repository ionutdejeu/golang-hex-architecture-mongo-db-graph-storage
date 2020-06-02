
import (
	"context"
	"errors"
	"strings"
	"time"

)

type UseCaseAdd struct {
	Saver saver
}

type Saver interface { 
	Store (*entity) (uuid string, err error)
}

