package config

var (
	Version bool // is print version command
	// Con           string // config file--env.yml
	Authorization string // authorization string
)

var Conf = new(Config)

const (
	VERSION               = "v1.0.0"
	ImageThumbnailProcess = "image/resize,w_50"
	VideoThumbnailProcess = "video/snapshot,t_0,f_jpg,w_50"
	ImageUrlProcess       = "image/resize,w_1920/format,jpeg"
	ASC                   = "ASC"
	DESC                  = "DESC"
	OrderUpdatedAt        = "updated_at"
	OrderCreatedAt        = "created_at"
	OrderSize             = "size"
	OrderName             = "name"
	OrderSearch           = "type ASC,updated_at DESC"
	AccessTokenInvalid    = "AccessTokenInvalid"
	Bearer                = "Bearer\t"
)
