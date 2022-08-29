package IPFSClient

import (
	"encoding/json"
	"strconv"
)

/*
/api/v0/files/chcid
/api/v0/files/cp
/api/v0/files/flush
/api/v0/files/ls
/api/v0/files/mkdir
/api/v0/files/mv
/api/v0/files/read
/api/v0/files/rm
/api/v0/files/stat
/api/v0/files/write*/

/*
/api/v0/files/chcid
Change the CID version or hash function of the root node of a given path.

#Arguments
arg [string]: Path to change. Default: '/'. Required: no.
cid-version [int]: Cid version to use. (experimental). Required: no.
hash [string]: Hash function to use. Will set Cid version to 1 if used. (experimental). Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/chcid?arg=<path>&cid-version=<value>&hash=<value>"

*/
func (c *IPFSClient) FilesChCid(arg string, cidVersion int, hash string) (string, error) {
	//This endpoint returns a `text/plain` response body.

	query := make(map[string]string)
	query["arg"] = arg
	query["cid-version"] = strconv.FormatInt(int64(cidVersion), 10)
	query["hash"] = hash
	form := make(map[string]string)

	b, err := PostForm("http://127.0.0.1:5001/api/v0/files/chcid", query, form)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

/*
/api/v0/files/cp
Add references to IPFS files and directories in MFS (or copy within MFS).

#Arguments
arg [string]: Source IPFS or MFS path to copy. Required: yes.
arg [string]: Destination within MFS. Required: yes.
parents [bool]: Make parent directories as needed. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/cp?arg=<source>&arg=<dest>&parents=<value>"
*/

func (c *IPFSClient) FilesCp(srcIPfsOrMFSPath string, targetMfsPath string, bMakeParents bool) (string, error) {
	//This endpoint returns a `text/plain` response body.

	query := make(map[string]string)
	query["arg"] = srcIPfsOrMFSPath
	query["arg"] = targetMfsPath
	query["parents"] = strconv.FormatBool(bMakeParents)
	form := make(map[string]string)

	b, err := PostForm("http://127.0.0.1:5001/api/v0/files/cp", query, form)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

/*

/api/v0/files/flush
Flush a given path's data to disk.

#Arguments
arg [string]: Path to flush. Default: '/'. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Cid": "<string>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/flush?arg=<path>"
*/
func (c *IPFSClient) FilesFlush(path string) (cid string, err error) {
	//This endpoint returns a `text/plain` response body.

	return "", nil
}

/*
#/api/v0/files/ls
List directories in the local mutable namespace.

#Arguments
arg [string]: Path to show listing for. Defaults to '/'. Required: no.
long [bool]: Use long listing format. Required: no.
U [bool]: Do not sort; list entries in directory order. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Entries": [
    {
      "Hash": "<string>",
      "Name": "<string>",
      "Size": "<int64>",
      "Type": "<int>"
    }
  ]
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/ls?arg=<path>&long=<value>&U=<value>"
*/
func (c *IPFSClient) FilesLs(path string, bLongFormat, bSortInDirOrder bool) (res *FilesLsRes, err error) {

	query := make(map[string]string)
	form := make(map[string]string)

	b, err := PostForm("http://127.0.0.1:5001/api/v0/files/ls", query, form)
	if err != nil {
		return nil, err
	}

	res = &FilesLsRes{}
	json.Unmarshal(b, res)

	return res, nil
}

/*
#/api/v0/files/mkdir
Make directories.

#Arguments
arg [string]: Path to dir to make. Required: yes.
parents [bool]: No error if existing, make parent directories as needed. Required: no.
cid-version [int]: Cid version to use. (experimental). Required: no.
hash [string]: Hash function to use. Will set Cid version to 1 if used. (experimental). Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/mkdir?arg=<path>&parents=<value>&cid-version=<value>&hash=<value>"
*/
func (c *IPFSClient) FilesMkdir(path string, bMakeParent bool, cidVersion int, hash string) (res string, err error) {
	//This endpoint returns a `text/plain` response body.

	return "", nil
}

/*
#/api/v0/files/mv
Move files.

#Arguments
arg [string]: Source file to move. Required: yes.
arg [string]: Destination path for file to be moved to. Required: yes.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/mv?arg=<source>&arg=<dest>"
*/
func (c *IPFSClient) FilesMv(srcFile, destPath string) (res string, err error) {
	//This endpoint returns a `text/plain` response body.

	return "", nil
}

/*
#/api/v0/files/read
Read a file in a given MFS.

#Arguments
arg [string]: Path to file to be read. Required: yes.
offset [int64]: Byte offset to begin reading from. Required: no.
count [int64]: Maximum number of bytes to read. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/read?arg=<path>&offset=<value>&count=<value>"
*/
func (c *IPFSClient) FilesRead(filePath string, offset int64, count int64) (res []byte, err error) {
	//This endpoint returns a `text/plain` response body.

	return []byte(""), nil
}

/*
#/api/v0/files/rm
Remove a file.

#Arguments
arg [string]: File to remove. Required: yes.
recursive [bool]: Recursively remove directories. Required: no.
force [bool]: Forcibly remove target at path; implies -r for directories. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/rm?arg=<path>&recursive=<value>&force=<value>"
*/
func (c *IPFSClient) FilesRm(filePath string, bRecursive bool, bForce bool) (res string, err error) {
	//This endpoint returns a `text/plain` response body.

	return "", nil
}

/*
#/api/v0/files/stat
Display file status.

#Arguments
arg [string]: Path to node to stat. Required: yes.
format [string]: Print statistics in given format.
	 Allowed tokens: <hash> <size> <cumulsize> <type> <childs>.
	 Conflicts with other format options. Default: <hash>
	Size: <size>
	CumulativeSize: <cumulsize>
	ChildBlocks: <childs>
	Type: <type>. Default: <hash> Size: <size> CumulativeSize: <cumulsize> ChildBlocks: <childs> Type: <type>. Required: no.
hash [bool]: Print only hash. Implies '--format=<hash>'. Conflicts with other format options. Required: no.
size [bool]: Print only size. Implies '--format=<cumulsize>'. Conflicts with other format options. Required: no.
with-local [bool]: Compute the amount of the dag that is local, and if possible the total size. Required: no.
#Response
On success, the call to this endpoint will return with 200 and the following body:

{
  "Blocks": "<int>",
  "CumulativeSize": "<uint64>",
  "Hash": "<string>",
  "Local": "<bool>",
  "Size": "<uint64>",
  "SizeLocal": "<uint64>",
  "Type": "<string>",
  "WithLocality": "<bool>"
}

#cURL Example
curl -X POST "http://127.0.0.1:5001/api/v0/files/stat?arg=<path>&format=<hash> Size: <size> CumulativeSize: <cumulsize> ChildBlocks: <childs> Type: <type>&hash=<value>&size=<value>&with-local=<value>"
*/
func (c *IPFSClient) FilesStat(nodePath string, format string,
	bPrintOnlyHash, bPrintOnlySize, bWithLocal bool) (res *FilesStatRes, err error) {
	//This endpoint returns a `text/plain` response body.
	res = &FilesStatRes{}
	return
}

/*
#/api/v0/files/write
Write to a mutable file in a given filesystem.

#Arguments
arg [string]: Path to write to. Required: yes.

offset [int64]: Byte offset to begin writing at. Required: no.

create [bool]: Create the file if it does not exist. Required: no.

parents [bool]: Make parent directories as needed. Required: no.

truncate [bool]: Truncate the file to size zero before writing. Required: no.

count [int64]: Maximum number of bytes to read. Required: no.

raw-leaves [bool]: Use raw blocks for newly created leaf nodes. (experimental). Required: no.

cid-version [int]: Cid version to use. (experimental). Required: no.

hash [string]: Hash function to use. Will set Cid version to 1 if used. (experimental). Required: no.

#Request Body
Argument data is of file type. This endpoint expects one or several files
 (depending on the command) in the body of the request as 'multipart/form-data'.

#Response
On success, the call to this endpoint will return with 200 and the following body:

This endpoint returns a `text/plain` response body.
#cURL Example
curl -X POST -F file=@myfile "http://127.0.0.1:5001/api/v0/files/write?arg=<path>&offset=<value>&create=<value>&parents=<value>&truncate=<value>&count=<value>&raw-leaves=<value>&cid-version=<value>&hash=<value>"
*/

func (c *IPFSClient) FilesWrit(filePath string, offset uint64,
	bAutoCreateFile, bAutoCreateParent, bTruncateBeforeWrite, bUseRawBlockForLeaves bool,
	maxReadBytes uint64, cidVersion int, hash string, srcFilePaths []string) (res string, err error) {
	//This endpoint returns a `text/plain` response body.

	return "", nil
}
