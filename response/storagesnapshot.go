package response

type StorageSnapshot struct {

	// Account to use for snapshots
	Account *string `json:"account,omitempty"`

	// Description is the description of the storage
	// snapshot
	Description string `json:"description,omitempty"`

	// Machineimage_name is the name of the machine
	// image that's used in the boot volume
	// from which this snapshot is taken
	Machineimage_name *string `json:"machineimage_name,omitempty"`

	// Name is the name of the snapshot
	Name string `json:"name"`

	// Property that describes:
	// If you don't specify a value, a remote snapshot is created.
	// Remote snapshots aren't stored in the
	// same location as the original storage volume. Instead,
	// they are reduced and stored in the associated Oracle
	// Storage Cloud Service instance. Remote snapshots are useful
	// if your domain spans multiple sites.
	// With remote snapshots, you can create a snapshot in one
	// site, then switch to another site and create a copy of the
	// storage volume on that site. However, creating a
	// remote snapshot and restoring a storage volume from a remote
	// snapshot can take quite a long time depending on the size
	// of the storage volume, as data is written to and from
	// the Oracle Storage Cloud Service instance.
	// Specify /oracle/private/storage/snapshot/collocated to
	// create a collocated snapshot. Colocated snapshots
	// are stored in the same physical location as the
	// original storage volume and each snapshot uses the
	// same amount of storage as the original volume.
	// Colocated snapshots and volumes from colocated snapshots can
	// be created very quickly. Colocated snapshots are useful
	// for quickly cloning storage volumes within a site.
	// However, you can't restore volumes across sites using colocated snapshots
	Property string `json:"property"`

	// Size is the size of the storage snapshot in bytes
	Size string `json:"size"`

	// Snapshot_id is a unique id for this snapshot
	// auto-generated by the server
	Snapshot_id string `json:"snapshot_id"`

	// Snapshot_timestamp is the timestamp of the storage snapshot,
	// generated by storage server. The snapshot will contain
	// data written to the original volume before this time
	Snapshot_timestamp *string `json:"snapshot_timestamp,omitempty"`

	// Start_timestamp is the timestamp when the operation is started
	Start_timestamp string `json:"start_timestamp"`

	// Status is the current status of the storage snapshot
	Status string `json:"status"`

	// Status_details indicates details about the latest
	// state of the storage volume snapshot.
	Status_detail string `json:"status_detail"`

	// Status_timestamp indicates the time that the current view of
	// the storage volume snapshot was generated.
	Status_timestamp string `json:"status_timestamp"`

	// Tags are strings that describe the
	// storage snapshot and help you identify it
	Tags []string `json:"tags,omitempty"`

	// Volume is the volume name you wish to create
	// the snapshot
	Volume string `json:"volume"`

	// Uri is the Uniform Resource Identifier
	Uri string `json:"uri"`
}

type AllStorageSnapshots struct {
	Result []StorageSnapshot `json:"result,omitempty"`
}