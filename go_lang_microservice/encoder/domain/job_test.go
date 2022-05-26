package domain_test

import (
	"testing"
	"encoder/domain"	
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)  

func TestNewJob(t *testing.T){
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path/to/file"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}