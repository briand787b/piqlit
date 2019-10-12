package obj

// UploadStatus represents the status of the process
// that uploads user-generated content into piqlit
type UploadStatus string

const (
	// UploadDone is the status of a successfully completed upload
	UploadDone UploadStatus = "done"

	// UploadFailed is the status of an unsuccessfully completed upload
	UploadFailed UploadStatus = "failed"

	// UploadInProgress is the status of a successfully completed upload
	UploadInProgress UploadStatus = "in_progress"

	// UploadNotStarted is the status of an upload that has not yet begun
	UploadNotStarted UploadStatus = "not_started"
)
