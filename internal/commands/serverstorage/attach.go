package serverstorage

import (
	"fmt"
	"github.com/UpCloudLtd/cli/internal/commands"
	"github.com/UpCloudLtd/cli/internal/commands/server"
	"github.com/UpCloudLtd/cli/internal/commands/storage"
	"github.com/UpCloudLtd/cli/internal/ui"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/service"
	"github.com/spf13/pflag"
)

type attachCommand struct {
	*commands.BaseCommand
	serverSvc  service.Server
	storageSvc service.Storage
	params     attachParams
}

type attachParams struct {
	request.AttachStorageRequest
}

// AttachCommand creates the "server storage attach" command
func AttachCommand(serverSvc service.Server, storageSvc service.Storage) commands.Command {
	return &attachCommand{
		BaseCommand: commands.New("attach", "Attaches a storage as a device to a server"),
		serverSvc:   serverSvc,
		storageSvc:  storageSvc,
	}
}

var defaultAttachParams = &attachParams{
	AttachStorageRequest: request.AttachStorageRequest{
		Type:     "disk",
		BootDisk: 0,
		Address:  "virtio",
	},
}

// InitCommand implements Command.InitCommand
func (s *attachCommand) InitCommand() error {
	s.SetPositionalArgHelp(positionalArgHelp)
	s.ArgCompletion(server.GetServerArgumentCompletionFunction(s.serverSvc))
	s.params = attachParams{AttachStorageRequest: request.AttachStorageRequest{}}

	flagSet := &pflag.FlagSet{}
	flagSet.StringVar(&s.params.Type, "type", defaultAttachParams.Type, "The type of the attached storage.\nAvailable: disk, cdrom")
	flagSet.StringVar(&s.params.Address, "address", defaultAttachParams.Address, "The address where the storage device is attached on the server. \nSpecify only the bus name (ide/scsi/virtio) to auto-select next available address from that bus.")
	flagSet.StringVar(&s.params.StorageUUID, "storage", defaultAttachParams.StorageUUID, "The UUID of the storage to attach.\n[Required]")
	flagSet.IntVar(&s.params.BootDisk, "boot-disk", defaultAttachParams.BootDisk, "If the value is 1 the storage device will be used as a boot disk, unless overridden with the server boot_order attribute.")

	s.AddFlags(flagSet)

	return nil
}

// MakeExecuteCommand implements Command.MakeExecuteCommand
func (s *attachCommand) MakeExecuteCommand() func(args []string) (interface{}, error) {
	return func(args []string) (interface{}, error) {

		if s.params.StorageUUID == "" {
			return nil, fmt.Errorf("storage is required")
		}

		strg, err := storage.SearchSingleStorage(s.params.StorageUUID, s.storageSvc)
		if err != nil {
			return nil, err
		}
		s.params.StorageUUID = strg.UUID

		return server.Request{
			BuildRequest: func(uuid string) interface{} {
				req := s.params.AttachStorageRequest
				req.ServerUUID = uuid
				return &req
			},
			Service:    s.serverSvc,
			ExactlyOne: true,
			Handler: ui.HandleContext{
				InteractiveUI: s.Config().InteractiveUI(),
				MaxActions:    maxServerActions,
				MessageFn: func(in interface{}) string {
					req := in.(*request.AttachStorageRequest)
					return fmt.Sprintf("Attaching storage %q to server %q", req.StorageUUID, req.ServerUUID)
				},
				Action: func(req interface{}) (interface{}, error) {
					return s.storageSvc.AttachStorage(req.(*request.AttachStorageRequest))
				},
			},
		}.Send(args)
	}
}
