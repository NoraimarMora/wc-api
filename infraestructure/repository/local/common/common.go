package common

import (
	"context"
	"fmt"
	"os"
)

type Common struct {
	path string
}

func New() Common {
	return Common{
		path: "./assets/",
	}
}

func (c *Common) GetFile(ctx context.Context, fileName string) ([]byte, error) {
	readFile, err := os.ReadFile(c.path + fileName)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(): %w", err)
	}

	return readFile, nil

}
