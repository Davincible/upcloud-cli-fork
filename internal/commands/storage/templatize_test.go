package storage

import (
	"github.com/UpCloudLtd/cli/internal/commands"
	"github.com/UpCloudLtd/cli/internal/config"
	"github.com/UpCloudLtd/cli/internal/mocks"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTemplatizeCommand(t *testing.T) {
	methodName := "TemplatizeStorage"
	var Storage1 = upcloud.Storage{
		UUID:   Uuid1,
		Title:  Title1,
		Access: "private",
		State:  "maintenance",
		Type:   "backup",
		Zone:   "fi-hel1",
		Size:   40,
		Tier:   "maxiops",
	}
	var Storage2 = upcloud.Storage{
		UUID:   Uuid2,
		Title:  Title2,
		Access: "private",
		State:  "online",
		Type:   "normal",
		Zone:   "fi-hel1",
		Size:   40,
		Tier:   "maxiops",
	}
	details := upcloud.StorageDetails{
		Storage: Storage1,
	}
	for _, test := range []struct {
		name     string
		args     []string
		error    string
		expected request.TemplatizeStorageRequest
	}{
		{
			name:  "Backend called with no args",
			args:  []string{},
			error: "title is required",
		},
		{
			name: "Backend called with title",
			args: []string{"--title", "test-title"},
			expected: request.TemplatizeStorageRequest{
				UUID:  Storage2.UUID,
				Title: "test-title",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			mss := mocks.MockStorageService{}
			mss.On("GetStorages", mock.Anything).Return(&upcloud.Storages{Storages: []upcloud.Storage{Storage1, Storage2}}, nil)
			mss.On(methodName, &test.expected).Return(&details, nil)

			tc := commands.BuildCommand(TemplatizeCommand(&mss), nil, config.New(viper.New()))
			mocks.SetFlags(tc, test.args)

			_, err := tc.MakeExecuteCommand()([]string{Storage2.UUID})
			if test.error != "" {
				assert.Errorf(t, err, "title is required")
			} else {
				mss.AssertNumberOfCalls(t, methodName, 1)
			}
		})
	}
}
