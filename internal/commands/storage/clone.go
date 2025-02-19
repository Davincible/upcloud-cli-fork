package storage

import (
	"fmt"

	"github.com/UpCloudLtd/upcloud-cli/internal/commands"
	"github.com/UpCloudLtd/upcloud-cli/internal/completion"
	"github.com/UpCloudLtd/upcloud-cli/internal/output"
	"github.com/UpCloudLtd/upcloud-cli/internal/resolver"

	"github.com/UpCloudLtd/upcloud-go-api/v4/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/v4/upcloud/request"
	"github.com/spf13/pflag"
)

type cloneCommand struct {
	*commands.BaseCommand
	resolver.CachingStorage
	completion.Storage
	params cloneParams
}

type cloneParams struct {
	request.CloneStorageRequest
}

// CloneCommand creates the "storage clone" command
func CloneCommand() commands.Command {
	return &cloneCommand{
		BaseCommand: commands.New(
			"clone",
			"Clone a storage",
			"upctl storage clone 015899e0-0a68-4949-85bb-261a99de5fdd --title my_storage_clone --zone fi-hel1",
			"upctl storage clone 015899e0-0a68-4949-85bb-261a99de5fdd --title my_storage_clone2 --zone pl-waw1  --tier maxiops",
			`upctl storage clone "My Storage" --title my_storage_clone3 --zone pl-waw1  --tier maxiops`,
		),
	}
}

var defaultCloneParams = &cloneParams{
	CloneStorageRequest: request.CloneStorageRequest{
		Tier: upcloud.StorageTierHDD,
	},
}

// InitCommand implements Command.InitCommand
func (s *cloneCommand) InitCommand() {
	s.params = cloneParams{CloneStorageRequest: request.CloneStorageRequest{}}

	flagSet := &pflag.FlagSet{}
	flagSet.StringVar(&s.params.Tier, "tier", defaultCloneParams.Tier, "The storage tier to use.")
	flagSet.StringVar(&s.params.Title, "title", defaultCloneParams.Title, "A short, informational description.")
	flagSet.StringVar(&s.params.Zone, "zone", defaultCloneParams.Zone, "The zone in which the storage will be created, e.g. fi-hel1.")

	s.AddFlags(flagSet)
	s.Cobra().MarkFlagRequired("title") //nolint:errcheck
	s.Cobra().MarkFlagRequired("zone")  //nolint:errcheck
}

// Execute implements commands.MultipleArgumentCommand
func (s *cloneCommand) Execute(exec commands.Executor, uuid string) (output.Output, error) {
	svc := exec.Storage()
	req := s.params.CloneStorageRequest
	req.UUID = uuid

	msg := fmt.Sprintf("Cloning storage %v", uuid)
	logline := exec.NewLogEntry(msg)

	logline.StartedNow()

	res, err := svc.CloneStorage(&req)
	if err != nil {
		return commands.HandleError(logline, fmt.Sprintf("%s: failed", msg), err)
	}

	logline.SetMessage(fmt.Sprintf("%s: success", msg))
	logline.MarkDone()

	return output.OnlyMarshaled{Value: res}, nil
}
