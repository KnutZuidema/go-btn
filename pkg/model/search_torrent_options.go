package model

type Category string

const (
	CategorySeason  Category = "Season"
	CategoryEpisode          = "Episode"
)

type Codec string

const (
	CodecXViD      Codec = "XViD"
	CodecX264            = "x264"
	CodecMPEG2           = "MPEG2"
	CodecDiVX            = "DiVX"
	CodecDVDR            = "DVDR"
	CodecVC1             = "VC-1"
	CodecH264            = "H.264"
	CodecWMV             = "WMV"
	CodecBR              = "BR"
	CodecX264Hi10P       = "x264-Hi10P"
)

type Container string

const (
	ContainerAVI  Container = "AVI"
	ContainerMKV            = "MKV"
	ContainerVOB            = "VOB"
	ContainerMPEG           = "MPEG"
	ContainerMP4            = "MP4"
	ContainerISO            = "ISO"
	ContainerWMV            = "WMV"
	ContainerTS             = "TS"
	ContainerM4V            = "M4V"
	ContainerM2TS           = "M2TS"
)

type Source string

const (
	SourceHDTV   Source = "HDTV"
	SourcePDTV          = "PDTV"
	SourceDSR           = "DSR"
	SourceDVDRip        = "DVDRip"
	SourceTVRip         = "TVRip"
	SourceVHSRip        = "VHSRip"
	SourceBluray        = "Bluray"
	SourceBDRip         = "BDRip"
	SourceBRRip         = "BRRip"
	SourceDVD5          = "DVD5"
	SourceDVD9          = "DVD9"
	SourceHDDVD         = "HDDVD"
	SourceWEB           = "WEB"
	SourceBD5           = "BD5"
	SourceBD9           = "BD9"
	SourceBD25          = "BD25"
	SourceBD50          = "BD50"
	SourceMixed         = "Mixed"
)

type Resolution string

const (
	ResolutionPortableDevice Resolution = "Portable Device"
	ResolutionSD                        = "SD"
	Resolution720p                      = "720p"
	Resolution1080i                     = "1080i"
	Resolution1080p                     = "1080p"
)

type Origin string

const (
	OriginScene Origin = "Scene"
	OriginP2P          = "P2P"
	OriginUser         = "User"
)

type SearchTorrentOptions struct {
	TorrentID  int         `json:"id,omitempty"`
	Series     string      `json:"series,omitempty"`
	Category   Category    `json:"category,omitempty"`
	GroupName  string      `json:"name,omitempty"`
	Search     string      `json:"search,omitempty"`
	Codecs     []Codec     `json:"codec,omitempty"`
	Containers []Container `json:"container,omitempty"`
	Hash       string      `json:"hash,omitempty"`
	TVDbID     string      `json:"tvdb,omitempty"`
	TVRageID   string      `json:"tvrage,omitempty"`
	Time       string      `json:"time,omitempty"`
	Age        string      `json:"age,omitempty"`
}
