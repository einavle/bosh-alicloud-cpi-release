package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"alibaba/bosh-alicloud-cpi/src/bosh-alicloud-cpi/alicloud"
)

type Factory struct {
	config alicloud.AlicloudConfig
}


type CPI struct {
	InfoMethod

	CreateStemcellMethod
	DeleteStemcellMethod

	CreateVMMethod
	DeleteVMMethod
	HasVMMethod
	RebootVMMethod
	SetVMMetadataMethod
	GetDisksMethod

	CreateDiskMethod
	DeleteDiskMethod
	AttachDiskMethod
	DetachDiskMethod
	HasDiskMethod
}

func NewFactory(config alicloud.AlicloudConfig) Factory {

	//var decompressor bwcutil.Decompressor
	//if opts.ExpandStemcellTarball {
	//	compressor := boshcmd.NewTarballCompressor(cmdRunner, fs)
	//	decompressor = bwcutil.NewTarDecompressor(fs, compressor)
	//} else {
	//	decompressor = bwcutil.NewGzipDecompressor(fs, cmdRunner)
	//}
	//
	//stemcellImporter := bwcstem.NewFSImporter(
	//	opts.StemcellsDir, fs, uuidGen, decompressor, logger)
	//
	//stemcellFinder := bwcstem.NewFSFinder(opts.StemcellsDir, fs, logger)
	//
	//ports := bwcvm.NewIPTablesPorts(cmdRunner)
	//sleeper := bwcutil.RealSleeper{}
	//
	//hostBindMounts := bwcvm.NewFSHostBindMounts(
	//	opts.HostEphemeralBindMountsDir, opts.HostPersistentBindMountsDir,
	//	sleeper, fs, cmdRunner, logger)
	//
	//guestBindMounts := bwcvm.NewFSGuestBindMounts(
	//	opts.GuestEphemeralBindMountPath, opts.GuestPersistentBindMountsDir, logger)
	//
	//systemResolvConfProvider := func() (bwcvm.ResolvConf, error) {
	//	return bwcvm.NewSystemResolvConfFromPath(fs)
	//}
	//
	//metadataService := bwcvm.NewMetadataService(
	//	opts.AgentEnvService, opts.Registry, logger)
	//
	//agentEnvServiceFactory := bwcvm.NewWardenAgentEnvServiceFactory(
	//	opts.AgentEnvService, opts.Registry, logger)
	//
	//vmCreator := bwcvm.NewWardenCreator(
	//	uuidGen, wardenClient, metadataService, agentEnvServiceFactory, ports,
	//	hostBindMounts, guestBindMounts, systemResolvConfProvider, opts.Agent, logger)
	//
	//vmFinder := bwcvm.NewWardenFinder(
	//	wardenClient, agentEnvServiceFactory, ports, hostBindMounts, guestBindMounts, logger)
	//
	//diskFactory := bwcdisk.NewFSFactory(opts.DisksDir, fs, uuidGen, cmdRunner, logger)

	return Factory{config:config}
}

//type TestCPI interface {
//	Info() (apiv1.Info, error)
//	apiv1.Stemcells
//}

func (this Factory) New(_ apiv1.CallContext) (apiv1.CPI, error) {
	cpi := CPI {
		NewInfoMethod(),

		NewCreateStemcellMethod(this.config),
		NewDeleteStemcellMethod(this.config),

		NewCreateVMMethod(this.config),
		NewDeleteVMMethod(this.config),
		NewHasVMMethod(this.config),
		NewRebootVMMethod(),
		NewSetVMMetadataMethod(),
		NewGetDisksMethod(this.config),

		NewCreateDiskMethod(this.config),
		NewDeleteDiskMethod(this.config),
		NewAttachDiskMethod(this.config),
		NewDetachDiskMethod(this.config),
		NewHasDiskMethod(this.config),
	}

	return cpi, nil
}



/*
type Stemcells interface {
	CreateStemcell(string, StemcellCloudProps) (StemcellCID, error)
	DeleteStemcell(StemcellCID) error
}

type VMs interface {
	CreateVM(AgentID, StemcellCID, VMCloudProps, Networks, []DiskCID, VMEnv) (VMCID, error)
	DeleteVM(VMCID) error

	SetVMMetadata(VMCID, VMMeta) error
	HasVM(VMCID) (bool, error)

	RebootVM(VMCID) error
	GetDisks(VMCID) ([]DiskCID, error)
}

type Disks interface {
	CreateDisk(int, DiskCloudProps, *VMCID) (DiskCID, error)
	DeleteDisk(DiskCID) error

	AttachDisk(VMCID, DiskCID) error
	DetachDisk(VMCID, DiskCID) error

	HasDisk(DiskCID) (bool, error)
}





 */
//func NewStemCells(_ apiv1.CallContext) (apiv1.Stemcells) {
//	return apiv1.Stemcells{
//		NewCreateStemcellMethod(),
//		nil,
//	};
//}
