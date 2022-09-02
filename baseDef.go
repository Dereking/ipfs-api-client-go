package IPFSClient

type TChunkSize string

const (
	ChunkSize131072   TChunkSize = "size-131072"
	ChunkSize262144              = "size-262144"
	ChunkSize524288              = "size-524288"
	ChunkSize1048576             = "size-1048576"
	ChunkSize2097152             = "size-2097152"
	ChunkSize4194304             = "size-4194304"
	ChunkSize8388608             = "size-8388608"
	ChunkSize16777216            = "size-16777216"
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
	Quiet                 bool           `query:"quiet"`               //
	Quieter               bool           `query:"quieter"`             //
	Silent                bool           `query:"silent"`              //
	Progress              bool           `query:"progress"`            //
	TrickleDAGFormat      bool           `query:"trickle"`             //
	OnlyHash              bool           `query:"only-hash"`           //
	WapFilesWithDirectory bool           `query:"wrap-with-directory"` //
	ChunkSize             TChunkSize     `query:"chunker"`             //
	Pin                   bool           `query:"pin"`                 //
	RawLeaves             bool           `query:"raw-leaves"`          //
	NoCopy                bool           `query:"nocopy"`              //
	FsCache               bool           `query:"fscache"`             //
	CidVersion            TCidVersion    `query:"cid-version"`         //
	HashAlgorithm         THashAlgorithm `query:"inlhashine"`          //
	Inline                bool           `query:"inline"`              //
	InlineLimit           int            `query:"inline-limit"`        //
	TargetPath            string         //
	SrcFilePath           string         //
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

//////////////////////
type CatReq struct {

	// arg [string]: The path to the IPFS object(s) to be outputted. Required: yes.
	IpfsPath string `query:"arg" required:"true"`

	// offset [int64]: Byte offset to begin reading from. Required: no.
	ReadOffset int64 `query:"offset" required:"false"`

	// length [int64]: Maximum number of bytes to read. Required: no.
	MaxLength int64 `query:"length" required:"false"`

	// progress [bool]: Stream progress data. Default: true. Required: no.
	StreamProgress bool `query:"progress"  required:"false"`
}

func NewCatReq() *CatReq {
	return &CatReq{
		StreamProgress: true,
		MaxLength:      256144,
	}
}

///////////////////////////
type CommandsReq struct {
	ShowFlags bool `query:"flags"` //flags [bool]: Show command flags. Required: no.
}

type CommandsResp struct {
	Name        string
	Options     []CommandOption
	Subcommands []Subcommand
}

type CommandOption struct {
	Names []string
}

type Subcommand struct {
	Name        string
	Options     []CommandOption
	Subcommands []Subcommand
}

func NewCommandsReq() *CommandsReq {
	return &CommandsReq{
		ShowFlags: true,
	}
}

/*{
  "Extra": "<string>",
  "ID": "<peer-id>",
  "Responses": [
    {
      "Addrs": [
        "<multiaddr-string>"
      ],
      "ID": "peer-id"
    }
  ],
  "Type": "<int>"
}
*/
type DhtQueryResp struct {
	Extra     string
	ID        string
	Responses []DhtQueryRespResponse
	Type      int
}

type DhtQueryRespResponse struct {
	Addrs []string //
	ID    string   //
}

func NewDhtQueryResp() *DhtQueryResp {
	return &DhtQueryResp{
		Responses: make([]DhtQueryRespResponse, 0),
	}
}

type GetReq struct {
	IpfsPath         string `query:"arg" required:"true"`                // arg [string]: The path to the IPFS object(s) to be outputted. Required: yes.
	TargetPath       string `query:"output" required:"false"`            // output [string]: The path where the output should be stored. Required: no.
	IsTarFormat      bool   `query:"archive" required:"false"`           // archive [bool]: Output a TAR archive. Required: no.
	IsCompress       bool   `query:"compress" required:"false"`          // compress [bool]: Compress the output with GZIP compression. Required: no.
	IsCompressLevel  int    `query:"compression-level" required:"false"` // compression-level [int]: The level of compression (1-9). Required: no.
	IsStreamProgress bool   `query:"progress" required:"false"`          // progress [bool]: Stream progress data. Default: true. Required: no.
}

func NewGetReq() *GetReq {
	return &GetReq{
		IpfsPath:        "",
		IsCompress:      false,
		IsCompressLevel: 9,
	}
}
