package IPFSClient

import (
	"encoding/json"
	"errors"
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
func (c *IPFSClient) Add(req *AddReq) (string, error) {
	//This endpoint returns a `text/plain` response body.

	// if req.TargetPath == "" {
	// 	return "", errors.New("TargetPath can't be empty")
	// }

	// _, err := os.Stat(req.SrcFilePath)

	// if os.IsNotExist(err) {
	// 	return "", errors.New("SrcFilePath :" + req.SrcFilePath + " not exists")
	// }

	// query := make(map[string]string)
	// query["quiet"] = strconv.FormatBool(req.Quiet)
	// query["quieter"] = strconv.FormatBool(req.Quieter)
	// query["silent"] = strconv.FormatBool(req.Silent)
	// query["progress"] = strconv.FormatBool(req.Progress)
	// query["trickle"] = strconv.FormatBool(req.TrickleDAGFormat)
	// query["only-hash "] = strconv.FormatBool(req.OnlyHash)
	// query["wrap-with-directory"] = strconv.FormatBool(req.WapFilesWithDirectory)
	// query["chunker"] = string(req.ChunkSize)
	// query["pin"] = strconv.FormatBool(req.Pin)
	// query["raw-leaves"] = strconv.FormatBool(req.RawLeaves)
	// query["nocopy"] = strconv.FormatBool(req.NoCopy)
	// query["fscache"] = strconv.FormatBool(req.FsCache)
	// query["cid-version"] = strconv.FormatInt(int64(req.CidVersion), 10) ///int
	// query["hash"] = string(req.HashAlgorithm)
	// query["inline"] = strconv.FormatBool(req.Inline)
	// query["inline-limit"] = strconv.FormatInt(int64(req.InlineLimit), 10) //int

	// form := make(map[string]string)

	// //b, err := PostForm("http://127.0.0.1:5001/api/v0/files/chcid", query, form)
	// b, err := PostFormWithFile(c.Host+"/api/v0/add", query, form,
	// 	req.TargetPath, req.SrcFilePath)
	// if err != nil {
	// 	return "", err
	// }

	// return string(b), nil

	if req.TargetPath == "" {
		return "", errors.New("TargetPath can't be empty")
	}

	_, err := os.Stat(req.SrcFilePath)

	if os.IsNotExist(err) {
		return "", errors.New("SrcFilePath :" + req.SrcFilePath + " not exists")
	}

	query, form := StructToHttpDataMap(*req)

	//log.Println("StructToHttpDataMap : add = ", query, form)

	b, err := PostFormWithFile(c.Host+"/api/v0/add", query, form, req.TargetPath, req.SrcFilePath)
	if err != nil {
		return "", err
	}

	return string(b), nil
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
	query, form := StructToHttpDataMap(*req)

	//log.Println(query, form)

	b, err := PostForm(c.Host+"/api/v0/cat", query, form)
	if err != nil {
		return "", err
	}
	return string(b), nil
	//This endpoint returns a `text/plain` response body.
	return "", nil
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
	query, form := StructToHttpDataMap(*req)

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
func (c *IPFSClient) DhtQuery(peerID string, bVerbose bool) (*DhtQueryResp, error) {

	b, err := PostUrl(c.Host + "/api/v0/dht/query?arg=" + peerID + "&verbose=" + strconv.FormatBool(bVerbose))
	if err != nil {
		return nil, err
	}
	//log.Println(string(b))
	var ret DhtQueryResp
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

//api/v0/get
//api/v0/id
//api/v0/log/level
//api/v0/log/ls
//api/v0/ls
//api/v0/ping
//api/v0/refs
//api/v0/refs/local
//api/v0/resolve
//api/v0/shutdown
//api/v0/update
//api/v0/version
//api/v0/version/deps
