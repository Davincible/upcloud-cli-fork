package server

import (
	"fmt"

	"github.com/UpCloudLtd/upcloud-cli/internal/commands"
	"github.com/UpCloudLtd/upcloud-cli/internal/completion"
	"github.com/UpCloudLtd/upcloud-cli/internal/config"
	"github.com/UpCloudLtd/upcloud-cli/internal/output"
	"github.com/UpCloudLtd/upcloud-cli/internal/resolver"

	"github.com/UpCloudLtd/upcloud-go-api/v4/upcloud/request"
	"github.com/spf13/pflag"
)

// DeleteCommand creates the "server delete" command
func DeleteCommand() commands.Command {
	return &deleteCommand{
		BaseCommand: commands.New(
			"delete",
			"Delete a server",
			"upctl delete 00cbe2f3-4cf9-408b-afee-bd340e13cdd8",
			"upctl delete 00cbe2f3-4cf9-408b-afee-bd340e13cdd8 0053a6f5-e6d1-4b0b-b9dc-b90d0894e8d0",
			"upctl delete my_server",
		),
	}
}

type deleteCommand struct {
	*commands.BaseCommand
	resolver.CachingServer
	completion.Server
	deleteStorages config.OptionalBoolean
}

// InitCommand implements Command.InitCommand
func (s *deleteCommand) InitCommand() {
	flags := &pflag.FlagSet{}
	config.AddToggleFlag(flags, &s.deleteStorages, "delete-storages", false, "Delete storages that are attached to the server.")
	s.AddFlags(flags)
}

// Execute implements commands.MultipleArgumentCommand
func (s *deleteCommand) Execute(exec commands.Executor, uuid string) (output.Output, error) {
	svc := exec.Server()
	msg := fmt.Sprintf("Deleting server %v", uuid)
	logline := exec.NewLogEntry(msg)

	logline.StartedNow()

	var err error
	if s.deleteStorages.Value() {
		logline.SetMessage(fmt.Sprintf("%s: deleting server and related storages", msg))
		err = svc.DeleteServerAndStorages(&request.DeleteServerAndStoragesRequest{
			UUID: uuid,
		})
	} else {
		logline.SetMessage(fmt.Sprintf("%s: deleting server", msg))
		err = svc.DeleteServer(&request.DeleteServerRequest{
			UUID: uuid,
		})
	}
	if err != nil {
		return commands.HandleError(logline, fmt.Sprintf("%s: failed", msg), err)
	}

	logline.SetMessage(fmt.Sprintf("%s: done", msg))
	logline.MarkDone()

	return output.None{}, nil
}
