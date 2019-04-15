package local_repository

import (
	"fmt"
	"os"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

type (
	LocalRepository struct{}
)

func (*LocalRepository) Upload(val string, path models.Path) error {
	file, err := os.Create(fmt.Sprintf(path.GetFullFilePath()))
	if err != nil {
		return errors.Wrap(err, "os.Create error")
	}
	defer file.Close()

	_, err = file.Write([]byte(val))
	if err != nil {
		return errors.Wrap(err, "file.Write err")
	}

	return nil
}
