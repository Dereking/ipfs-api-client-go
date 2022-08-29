package IPFSClient

type FilesLsRes struct {
	Entries []FilesLsResItem
}

type FilesLsResItem struct {
	Hash string
	Name string
	Size int64
	Type int
}

type FilesStatRes struct {
	Blocks         int
	CumulativeSize uint64
	Hash           string
	Local          bool
	Size           uint64
	SizeLocal      uint64
	Type           string
	WithLocality   bool
}
