package IPFSClient

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
)

/*
/api/v0/add
Add a file or directory to IPFS.

#Arguments
quiet [bool]: Write minimal output. Required: no.
quieter [bool]: Write only final hash. Required: no.
silent [bool]: Write no output. Required: no.
progress [bool]: Stream progress data. Required: no.
trickle [bool]: Use trickle-dag format for dag generation. Required: no.
only-hash [bool]: Only chunk and hash - do not write to disk. Required: no.
wrap-with-directory [bool]: Wrap files with a directory object. Required: no.
chunker [string]: Chunking algorithm, size-[bytes], rabin-[min]-[avg]-[max] or buzhash. Default: size-262144. Required: no.
pin [bool]: Pin this object when adding. Default: true. Required: no.
raw-leaves [bool]: Use raw blocks for leaf nodes. Required: no.
nocopy [bool]: Add the file using filestore. Implies raw-leaves. (experimental). Required: no.
fscache [bool]: Check the filestore for pre-existing blocks. (experimental). Required: no.
cid-version [int]: CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. Passing version 1 will cause the raw-leaves option to default to true. Required: no.
hash [string]: Hash function to use. Implies CIDv1 if not sha2-256. (experimental). Default: sha2-256. Required: no.
inline [bool]: Inline small blocks into CIDs. (experimental). Required: no.
inline-limit [int]: Maximum block size to inline. (experimental). Default: 32. Required: no.

#Request Body
Argument path is of file type. This endpoint expects one or several files
(depending on the command) in the body of the request as 'multipart/form-data'.

The add command not only allows adding files, but also uploading directories and complex hierarchies.

This happens as follows: Every part in the multipart request is a directory or a file to be added to IPFS.

Directory parts have a special content type application/x-directory.
These parts do not carry any data. The part headers look as follows:

Content-Disposition: form-data; name="file"; filename="folderName"
Content-Type: application/x-directory
File parts carry the file payload after the following headers:

Abspath: /absolute/path/to/file.txt
Content-Disposition: form-data; name="file"; filename="folderName%2Ffile.txt"
Content-Type: application/octet-stream

...contents...
The above file includes its path in the "folderName/file.txt" hierarchy
and IPFS will therefore be able to add it inside "folderName".
The parts declaring the directories are optional when they have files inside
and will be inferred from the filenames.
In any case, a depth-first traversal of the directory tree is recommended to
order the different parts making the request.

The Abspath header is included for filestore/urlstore features that are enabled
with the nocopy option and it can be set to the location of the file
in the filesystem (within the IPFS root), or to its full web URL.

#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Bytes": "<int64>",
  "Hash": "<string>",
  "Name": "<string>",
  "Size": "<string>"
}

#cURL Example
curl -X POST -F file=@myfile "http://127.0.0.1:5001/api/v0/add?quiet=<value>&quieter=<value>&silent=<value>&progress=<value>&trickle=<value>&only-hash=<value>&wrap-with-directory=<value>&chunker=size-262144&pin=true&raw-leaves=<value>&nocopy=<value>&fscache=<value>&cid-version=<value>&hash=sha2-256&inline=<value>&inline-limit=32"
*/
func (c *IPFSClient) Add(req *AddReq) (res *AddResp, err error) {

	if req.TargetPath == "" {
		return nil, errors.New("TargetPath can't be empty")
	}

	_, err = os.Stat(req.SrcFilePath)
	if os.IsNotExist(err) {
		return nil, errors.New("SrcFilePath :" + req.SrcFilePath + " not exists")
	}

	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return nil, err
	}

	//log.Println("StructToHttpDataMap : add = ", query, form)

	b, err := PostFormWithFile(c.Host+"/api/v0/add", query, form, req.TargetPath, req.SrcFilePath)
	if err != nil {
		return nil, err
	}

	var ret AddResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type TChunkSize string

const (
	ChunkSize131072   TChunkSize = "size-131072"
	ChunkSize262144   TChunkSize = "size-262144"
	ChunkSize524288   TChunkSize = "size-524288"
	ChunkSize1048576  TChunkSize = "size-1048576"
	ChunkSize2097152  TChunkSize = "size-2097152"
	ChunkSize4194304  TChunkSize = "size-4194304"
	ChunkSize8388608  TChunkSize = "size-8388608"
	ChunkSize16777216 TChunkSize = "size-16777216"
	ChunkRabinMin     TChunkSize = "rabin-min"
	ChunkRabinAvg     TChunkSize = "rabin-avg"
	ChunkRabinMax     TChunkSize = "rabin-max"
	ChunkBuzhash      TChunkSize = "buzhash"
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

type AddResp struct {
	ErrMessage
	Bytes int64  // 	  "Bytes": "<int64>",
	Hash  string //   "Hash": "<string>",
	Name  string //   "Name": "<string>",
	Size  string //   "Size": "<string>"
}

//api/v0/cat
/*
Show IPFS object data.

#Arguments
arg [string]: The path to the IPFS object(s) to be outputted. Required: yes.
offset [int64]: Byte offset to begin reading from. Required: no.
length [int64]: Maximum number of bytes to read. Required: no.
progress [bool]: Stream progress data. Default: true. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/cat?arg=<ipfs-path>&offset=<value>&length=<value>&progress=true"
*/

func (c *IPFSClient) Cat(req *CatReq) (string, error) {
	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return "", err
	}

	//log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/cat", query, form)
	if err != nil {
		return "", err
	}
	return string(b), nil
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

//api/v0/commands
/*/api/v0/commands
List all available commands.

#Arguments
flags [bool]: Show command flags. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Name": "<string>",
  "Options": [
    {
      "Names": [
        "<string>"
      ]
    }
  ],
  "Subcommands": [
    {
      "Name": "<string>",
      "Options": [
        {
          "Names": [
            "<string>"
          ]
        }
      ],
      "Subcommands": [
        "..."
      ]
    }
  ]
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/commands?flags=<value>"

#
*/
func (c *IPFSClient) Commands(req *CommandsReq) (*CommandsResp, error) {
	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return nil, err
	}

	//log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/commands", query, form)
	if err != nil {
		return nil, err
	}
	//log.Println(string(b))
	var ret CommandsResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

///////////////////////////
type CommandsReq struct {
	ShowFlags bool `query:"flags"` //flags [bool]: Show command flags. Required: no.
}

type CommandsResp struct {
	ErrMessage
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

//api/v0/commands/completion/bash
/*
/api/v0/commands/completion/bash
Generate bash shell completions.

#Arguments
This endpoint takes no arguments.

#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/commands/completion/bash"

#
*/

func (c *IPFSClient) CommandsCompletionBash() (string, error) {

	b, err := PostUrl(c.Host + "/api/v0/commands/completion/bash")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//api/v0/dht/query
/*
/api/v0/dht/query
Find the closest Peer IDs to a given Peer ID by querying the DHT.

#Arguments
arg [string]: The peerID to run the query against. Required: yes.
verbose [bool]: Print extra information. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
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

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/dht/query?arg=<peerID>&verbose=<value>"

#
*/
func (c *IPFSClient) DhtQuery(peerID string, bVerbose bool) (string /*DhtQueryResp*/, error) {

	b, err := PostUrl(c.Host + "/api/v0/dht/query?arg=" + peerID + "&verbose=" + strconv.FormatBool(bVerbose))
	if err != nil {
		return "", err
	}
	//log.Println(string(b))
	return string(b), nil
	// var ret DhtQueryResp
	// err = json.Unmarshal(b, &ret)
	// if err != nil {
	// 	return nil, err
	// }
	//return &ret, nil
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
	ErrMessage
	Addrs []string //
	ID    string   //
}

func NewDhtQueryResp() *DhtQueryResp {
	return &DhtQueryResp{
		Responses: make([]DhtQueryRespResponse, 0),
	}
}

//api/v0/get
/*Download IPFS objects.

#Arguments
arg [string]: The path to the IPFS object(s) to be outputted. Required: yes.
output [string]: The path where the output should be stored. Required: no.
archive [bool]: Output a TAR archive. Required: no.
compress [bool]: Compress the output with GZIP compression. Required: no.
compression-level [int]: The level of compression (1-9). Required: no.
progress [bool]: Stream progress data. Default: true. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/get?arg=<ipfs-path>&output=<value>&archive=<value>&compress=<value>&compression-level=<value>&progress=true"
*/
func (c *IPFSClient) Get(req *GetReq) (string, error) {

	req.TargetPath = "/tmp/q.tar"

	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return "", err
	}

	//log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/get", query, form)
	if err != nil {
		return "", err
	}
	//log.Println(b)
	return string(b), nil
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

//api/v0/id
/*Show IPFS node id info.

#Arguments
arg [string]: Peer.ID of node to look up. Required: no.
format [string]: Optional output format. Required: no.
peerid-base [string]: Encoding used for peer IDs:
   Can either be a multibase encoded CID or a base58btc encoded multihash.
    Takes {b58mh|base36|k|base32|b...}. Default: b58mh. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Addresses": [
    "<string>"
  ],
  "AgentVersion": "<string>",
  "ID": "<string>",
  "ProtocolVersion": "<string>",
  "Protocols": [
    "<string>"
  ],
  "PublicKey": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/id?arg=<peerid>&format=<value>&peerid-base=b58mh"
*/
func (c *IPFSClient) Id(req *IdReq) (*IdResp, error) {

	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return nil, err
	}

	//log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/id", query, form)
	if err != nil {
		return nil, err
	}

	//log.Println(string(b))

	var ret IdResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type PeerIdBase string

const (
	//b58mh|base36|k|base32 b....
	PeerIdBase_B58mh  PeerIdBase = "b58mh"
	PeerIdBase_Base36 PeerIdBase = "base36"
	PeerIdBase_Base32 PeerIdBase = "base32"
	PeerIdBase_Base64 PeerIdBase = "base64"
)

type IdReq struct {
	PeerId     string     `query:"arg" required:"false"`
	Fromat     string     `query:"format" required:"false"`
	PeerIdBase PeerIdBase `query:"peerid-base" required:"false"`
}

func NewIdReq() *IdReq {
	return &IdReq{
		PeerId:     "",
		PeerIdBase: PeerIdBase_B58mh,
	}
}

type IdResp struct {
	ErrMessage
	Addresses       []string // "Addresses": [	"<string>" ],
	AgentVersion    string   // "AgentVersion": "<string>",
	ID              string   // "ID": "<string>",
	ProtocolVersion string   // "ProtocolVersion": "<string>",
	Protocols       []string // "Protocols": [	"<string>"],
	PublicKey       string   // "PublicKey": "<string>"

}

//api/v0/log/level
/*Change the logging level.

#Arguments
arg [string]: The subsystem logging identifier. Use 'all' for all subsystems. Required: yes.
arg [string]: The log level, with 'debug' the most verbose and 'fatal' the least verbose.
One of: debug, info, warn, error, dpanic, panic, fatal.
Required: yes.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Message": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/log/level?arg=<subsystem>&arg=<level>"
*/
func (c *IPFSClient) LogLevel(req *LogLevelReq) (res *LogLevelResp, err error) {

	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return nil, err
	}

	log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/log/level", query, form)
	if err != nil {
		return nil, err
	}

	// log.Println(string(b))

	var ret LogLevelResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type LogLevelReq struct {
	LogSubsystemId string    `query:"arg" required:"true"`
	LogLevel       TLogLevel `query:"arg" required:"true"`
	// arg [string]: The subsystem logging identifier. Use 'all' for all subsystems. Required: yes.
	// arg [string]: The log level, with 'debug' the most verbose and 'fatal' the least verbose.
	// One of: debug, info, warn, error, dpanic, panic, fatal. Required: yes.
}

type TLogLevel string

const (
	LogLevel_Debug  TLogLevel = "debug"
	LogLevel_Info   TLogLevel = "info"
	LogLevel_Warn   TLogLevel = "warn"
	LogLevel_Error  TLogLevel = "error"
	LogLevel_Dpanic TLogLevel = "dpanic"
	LogLevel_Panic  TLogLevel = "panic"
	LogLevel_Fatal  TLogLevel = "fatal"
)

func NewLogLevelReq() *LogLevelReq {
	return &LogLevelReq{
		LogSubsystemId: "all",
		LogLevel:       LogLevel_Debug,
	}
}

type LogLevelResp struct {
	ErrMessage
	Message string
}

//api/v0/log/ls
/*List the logging subsystems.

#Arguments
This endpoint takes no arguments.

#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Strings": [
    "<string>"
  ]
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/log/ls"
*/
func (c *IPFSClient) LogLs() (res *LogLsResp, err error) {
	res = &LogLsResp{
		Strings: make([]string, 0),
	}

	b, err := PostUrl(c.Host + "/api/v0/log/ls")
	if err != nil {
		return nil, err
	}

	// log.Println(string(b))

	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	return
}

type LogLsResp struct {
	ErrMessage
	Strings []string
}

//api/v0/ls
/*List directory contents for Unix filesystem objects.

#Arguments
arg [string]: The path to the IPFS object(s) to list links from. Required: yes.
headers [bool]: Print table headers (Hash, Size, Name). Required: no.
resolve-type [bool]: Resolve linked objects to find out their types. Default: true. Required: no.
size [bool]: Resolve linked objects to find out their file size. Default: true. Required: no.
stream [bool]: Enable experimental streaming of directory entries as they are traversed. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Objects": [
    {
      "Hash": "<string>",
      "Links": [
        {
          "Hash": "<string>",
          "Name": "<string>",
          "Size": "<uint64>",
          "Target": "<string>",
          "Type": "<int32>"
        }
      ]
    }
  ]
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/ls?arg=<ipfs-path>&headers=<value>&resolve-type=true&size=true&stream=<value>"
*/
func (c *IPFSClient) Ls(req *LsReq) (res *LsResp, err error) {

	log.Println("ls : ", req)
	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return nil, err
	}

	log.Println("ls : ", query, form)

	b, err := PostForm(c.Host+"/api/v0/ls", query, form)
	if err != nil {
		return nil, err
	}

	log.Println("Ls res", string(b))

	var ret LsResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	log.Println("Ls res", ret)

	return &ret, nil
}

type LsReq struct {
	// arg [string]: The path to the IPFS object(s) to list links from. Required: yes.
	IpfsPath string `query:"arg" required:"true"`
	// headers [bool]: Print table headers (Hash, Size, Name). Required: no.
	ShowHeader bool `query:"headers" required:"false"`
	// resolve-type [bool]: Resolve linked objects to find out their types. Default: true. Required: no.
	ResolveLinkType bool `query:"resolve-type" required:"false"`
	// size [bool]: Resolve linked objects to find out their file size. Default: true. Required: no.
	ResolveLinkSize bool `query:"size" required:"false"`
	// stream [bool]: Enable experimental streaming of directory entries as they are traversed. Required: no.
	EnableStreamOfDir bool `query:"stream" required:"false"`
}

func NewLsReq() *LsReq {
	return &LsReq{
		IpfsPath:        "",
		ResolveLinkType: true,
		ResolveLinkSize: true,
	}
}

type LsResp struct {
	ErrMessage
	Objects []*LsRespObject
}

type LsRespObject struct {
	Hash  string
	Links []*LsRespLink
}

type LsRespLink struct {
	Hash   string //     "Hash": "<string>",
	Name   string //     "Name": "<string>",
	Size   uint64 //      "Size": "<uint64>",
	Target string //     "Target": "<string>",
	Type   int32  //"Type": "<int32>"
}

//api/v0/ping
/*Send echo request packets to IPFS hosts.

#Arguments
arg [string]: ID of peer to be pinged. Required: yes.
count [int]: Number of ping messages to send. Default: 10. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Success": "<bool>",
  "Text": "<string>",
  "Time": "<duration-ns>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/ping?arg=<peer ID>&count=10"
*/
func (c *IPFSClient) Ping(req *PingReq) (res *PingResp, err error) {
	query, form, err := StructToHttpDataMap(*req, false)
	if err != nil {
		return "", err
	}

	log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/ping", query, form)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type PingReq struct {
	PeerId  string `query:"arg" required:"true"`// arg [string]: ID of peer to be pinged. Required: yes.
	PingCount int `query:"count" required:"false"`// count [int]: Number of ping messages to send. Default: 10. Required: no.
}

func  NewPingReq(peerId string)*PingReq{
	return &PingReq{
		PeerId : peerId,
		PingCount : 10,
	}
}

type PingResp struct {
	ErrMessage
	Success bool // "Success": "<bool>",
	Text string //"Text": "<string>",
	Time  time.Duration //"Time": "<duration-ns>"
}

//api/v0/refs
/*List links (references) from an object.

#Arguments
arg [string]: Path to the object(s) to list refs from. Required: yes.
format [string]: Emit edges with given format. Available tokens: <src> <dst> <linkname>. Default: <dst>. Default: <dst>. Required: no.
edges [bool]: Emit edge format: &lt;from&gt; -&gt; &lt;to&gt;. Required: no.
unique [bool]: Omit duplicate refs from output. Required: no.
recursive [bool]: Recursively list links of child nodes. Required: no.
max-depth [int]: Only for recursive refs, limits fetch and listing to the given depth. Default: -1. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Err": "<string>",
  "Ref": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/refs?arg=<ipfs-path>&format=<dst>&edges=<value>&unique=<value>&recursive=<value>&max-depth=-1"
*/
func (c *IPFSClient) Refs() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/refs/local
/*List all local references.

#Arguments
This endpoint takes no arguments.

#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Err": "<string>",
  "Ref": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/refs/local"
*/
func (c *IPFSClient) RefsLocal() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/resolve
/*Resolve the value of names to IPFS.

#Arguments
arg [string]: The name to resolve. Required: yes.
recursive [bool]: Resolve until the result is an IPFS name. Default: true. Required: no.
dht-record-count [int]: Number of records to request for DHT resolution. Required: no.
dht-timeout [string]: Max time to collect values during DHT resolution eg "30s". Pass 0 for no timeout. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Path": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/resolve?arg=<name>&recursive=true&dht-record-count=<value>&dht-timeout=<value>"
*/
func (c *IPFSClient) Resolve() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/shutdown
/*Shut down the IPFS daemon.

#Arguments
This endpoint takes no arguments.

#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/shutdown"
*/
func (c *IPFSClient) Shutdown() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/update
/*Arguments
arg [string]: Arguments for subcommand. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/update?arg=<args>"
*/
func (c *IPFSClient) Update() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/version
/*Show IPFS version information.

#Arguments
number [bool]: Only show the version number. Required: no.
commit [bool]: Show the commit hash. Required: no.
repo [bool]: Show repo version. Required: no.
all [bool]: Show all version information. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Commit": "<string>",
  "Golang": "<string>",
  "Repo": "<string>",
  "System": "<string>",
  "Version": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/version?number=<value>&commit=<value>&repo=<value>&all=<value>"
*/
func (c *IPFSClient) Version() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//api/v0/version/deps
/*Shows information about dependencies used for build.

#Arguments
This endpoint takes no arguments.

#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Path": "<string>",
  "ReplacedBy": "<string>",
  "Sum": "<string>",
  "Version": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/version/deps"
*/
func (c *IPFSClient) VersionDeps() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

// - [ ]
//   - [ ] /api/v0/log/tail
func (c *IPFSClient) LogTail() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}

//  Experimental RPC commands
//  - [ ] /api/v0/mount

func (c *IPFSClient) Mount() (res string, err error) {
	err = errors.New("unimplement api")
	return "", err
}
