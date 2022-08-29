package IPFSClient

type TChunkSize string

const (
	ChunkSize131072   TChunkSize = "Size-131072"
	ChunkSize262144              = "Size-262144"
	ChunkSize524288              = "Size-524288"
	ChunkSize1048576             = "Size-1048576"
	ChunkSize2097152             = "Size-2097152"
	ChunkSize4194304             = "Size-4194304"
	ChunkSize8388608             = "Size-8388608"
	ChunkSize16777216            = "Size-16777216"
	ChunkRabinMin                = "rabin-min"
	ChunkRabinAvg                = "rabin-avg"
	ChunkRabinMax                = "rabin-max"
	ChunkBuzhash                 = "buzhash"
)

type TCidVersion int

const (
	CIDv0 TCidVersion = 0
	CIDv1             = 1
)

type THashAlgorithm string

const (
	Sha2_256 THashAlgorithm = "sha2-256"
)

// quiet [bool]: Write minimal output. Required: no.
// quieter [bool]: Write only final hash. Required: no.
// silent [bool]: Write no output. Required: no.
// progress [bool]: Stream progress data. Required: no.
// trickle [bool]: Use trickle-dag format for dag generation. Required: no.
// only-hash [bool]: Only chunk and hash - do not write to disk. Required: no.
// wrap-with-directory [bool]: Wrap files with a directory object. Required: no.
// chunker [string]: Chunking algorithm, size-[bytes], rabin-[min]-[avg]-[max] or buzhash. Default: size-262144. Required: no.
// pin [bool]: Pin this object when adding. Default: true. Required: no.
// raw-leaves [bool]: Use raw blocks for leaf nodes. Required: no.
// nocopy [bool]: Add the file using filestore. Implies raw-leaves. (experimental). Required: no.
// fscache [bool]: Check the filestore for pre-existing blocks. (experimental). Required: no.
// cid-version [int]: CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. Passing version 1 will cause the raw-leaves option to default to true. Required: no.
// hash [string]: Hash function to use. Implies CIDv1 if not sha2-256. (experimental). Default: sha2-256. Required: no.
// inline [bool]: Inline small blocks into CIDs. (experimental). Required: no.
// inline-limit [int]: Maximum block size to inline. (experimental). Default: 32. Required: no.
type AddReq struct {
	Quiet                 bool           //Write minimal output. Required: no.
	Quieter               bool           //Write only final hash. Required: no.
	Silent                bool           //Write no output. Required: no.
	Progress              bool           //Stream progress data. Required: no.
	TrickleDAGFormat      bool           //Use trickle-dag format for dag generation. Required: no.
	OnlyHash              bool           //Only chunk and hash - do not write to disk. Required: no.
	WapFilesWithDirectory bool           //Wrap files with a directory object. Required: no.
	ChunkSize             TChunkSize     //Chunking algorithm, size-[bytes], rabin-[min]-[avg]-[max] or buzhash. Default: size-262144. Required: no.
	Pin                   bool           //Pin this object when adding. Default: true. Required: no.
	RawLeaves             bool           //Use raw blocks for leaf nodes. Required: no.
	NoCopy                bool           //Add the file using filestore. Implies raw-leaves. (experimental). Required: no.
	FsCache               bool           //Check the filestore for pre-existing blocks. (experimental). Required: no.
	CidVersion            TCidVersion    // CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. Passing version 1 will cause the raw-leaves option to default to true. Required: no.
	HashAlgorithm         THashAlgorithm //Hash function to use. Implies CIDv1 if not sha2-256. (experimental). Default: sha2-256. Required: no.
	Inline                bool           //Inline small blocks into CIDs. (experimental). Required: no.
	InlineLimit           int            //: Maximum block size to inline. (experimental). Default: 32. Required: no.

	TargetPath  string // ipfs path
	SrcFilePath string //src file path to upload
}

func NewAddReq() *AddReq {
	return &AddReq{
		Quiet:                 false,           //bool           //Write minimal output. Required: no.
		Quieter:               false,           //bool           //Write only final hash. Required: no.
		Silent:                false,           //bool           //Write no output. Required: no.
		Progress:              false,           //bool           //Stream progress data. Required: no.
		TrickleDAGFormat:      false,           //bool           //Use trickle-dag format for dag generation. Required: no.
		OnlyHash:              false,           //bool           //Only chunk and hash - do not write to disk. Required: no.
		WapFilesWithDirectory: false,           //bool           //Wrap files with a directory object. Required: no.
		ChunkSize:             ChunkSize262144, //TChunkSize     //Chunking algorithm, size-[bytes], rabin-[min]-[avg]-[max] or buzhash. Default: size-262144. Required: no.
		Pin:                   true,            //bool           //Pin this object when adding. Default: true. Required: no.
		RawLeaves:             false,           //bool           //Use raw blocks for leaf nodes. Required: no.
		NoCopy:                false,           //bool           //Add the file using filestore. Implies raw-leaves. (experimental). Required: no.
		FsCache:               false,           //bool           //Check the filestore for pre-existing blocks. (experimental). Required: no.
		CidVersion:            CIDv0,           //TCidVersion    // CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. Passing version 1 will cause the raw-leaves option to default to true. Required: no.
		HashAlgorithm:         Sha2_256,        //THashAlgorithm //Hash function to use. Implies CIDv1 if not sha2-256. (experimental). Default: sha2-256. Required: no.
		Inline:                false,           //bool           //Inline small blocks into CIDs. (experimental). Required: no.
		InlineLimit:           32,              //int            //: Maximum block size to inline. (experimental). Default: 32. Required: no.

		TargetPath:  "",
		SrcFilePath: "",
	}
}
