package attachment

// attachment is simple interface
// that can represent any type of data
// that could be uploaded with telegram bot api.
type Attachment interface {
	GetAttachmentPath() string
}
