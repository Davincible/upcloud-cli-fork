package storage

import (
	"testing"

	"github.com/UpCloudLtd/upcloud-cli/internal/commands"
	"github.com/UpCloudLtd/upcloud-cli/internal/config"
	smock "github.com/UpCloudLtd/upcloud-cli/internal/mock"
	"github.com/UpCloudLtd/upcloud-cli/internal/mockexecute"
	internal "github.com/UpCloudLtd/upcloud-cli/internal/service"

	"github.com/UpCloudLtd/upcloud-go-api/v4/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/v4/upcloud/request"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommand(t *testing.T) {
	targetMethod := "CreateStorage"
	var Storage1 = upcloud.Storage{
		UUID:   UUID1,
		Title:  Title1,
		Access: "private",
		State:  "maintenance",
		Type:   "backup",
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
		expected request.CreateStorageRequest
	}{
		{
			name: "create with default values, no backup rule",
			args: []string{
				"--title", "create-storage-test", "" +
					"--zone", "abc"},
			expected: request.CreateStorageRequest{
				Size:       defaultCreateParams.Size,
				Tier:       defaultCreateParams.Tier,
				Title:      "create-storage-test",
				Zone:       "abc",
				BackupRule: nil,
			},
		},
		{
			name: "create with default values, with backup rule",
			args: []string{"--title", "create-storage-test", "--zone", "abc", "--backup-time", "09:00"},
			expected: request.CreateStorageRequest{
				Size:  defaultCreateParams.Size,
				Tier:  defaultCreateParams.Tier,
				Title: "create-storage-test",
				Zone:  "abc",
				BackupRule: &upcloud.BackupRule{
					Interval:  "daily",
					Time:      "0900",
					Retention: 7,
				},
			},
		},
		{
			name: "create with non default values",
			args: []string{
				"--title", "create-storage-test",
				"--zone", "abc",
				"--size", "30",
				"--tier", "xyz",
				"--backup-time", "09:00",
				"--backup-retention", "10",
				"--backup-interval", "mon",
			},
			expected: request.CreateStorageRequest{
				Size:  30,
				Tier:  "xyz",
				Title: "create-storage-test",
				Zone:  "abc",
				BackupRule: &upcloud.BackupRule{
					Interval:  "mon",
					Time:      "0900",
					Retention: 10,
				},
			},
		},
		{
			name: "title is missing",
			args: []string{
				"--size", "10",
				"--zone", "zone",
			},
			error: `required flag(s) "title" not set`,
		},
		{
			name: "zone is missing",
			args: []string{
				"--title", "title",
				"--size", "10",
			},
			error: `required flag(s) "zone" not set`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			conf := config.New()
			testCmd := CreateCommand()
			mService := new(smock.Service)

			conf.Service = internal.Wrapper{Service: mService}
			mService.On(targetMethod, &test.expected).Return(&details, nil)

			c := commands.BuildCommand(testCmd, nil, config.New())

			c.Cobra().SetArgs(test.args)
			_, err := mockexecute.MockExecute(c, mService, conf)

			if test.error != "" {
				assert.EqualError(t, err, test.error)
			} else {
				assert.NoError(t, err)
				mService.AssertNumberOfCalls(t, targetMethod, 1)
			}
		})
	}
}
