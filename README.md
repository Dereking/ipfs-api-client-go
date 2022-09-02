# ipfs-api-client-go

#### 介绍
本项目是 golang 实现的 IPFS （Kubo） HTTP RPC API 客户端library。功能正在逐步完善中。已实现代码见下方列表，欢迎大家补充并提交代码

#### 使用教程

1.  引入 library
```go
import （
    "github.com/Dereking/ipfs-api-client-go"
)
```
2.  使用示例

  列出 Unix filesystem 目录内容：
```go
client := NewIPFSClientLocal()
res, err := client.FilesLs("/", false, false)
if err != nil {
  t.Fatalf(err.Error())
}
fmt.Println("Test_FilesLs", res) 
```

  添加文件到IPFS: 
```go
req := NewAddReq()
req.SrcFilePath = "./test.txt"
req.TargetPath = "/test.txt"

client := NewIPFSClientLocal()
res, err := client.Add(req)
if err != nil {
  fmt.Fatal(err.Error())
}
fmt.Log("add ", res)
```

  各模块示例代码可参考*_test.go



#### 代码风格
- 如果 api 参数超过 3 个，使用结构体定义，并用 NewXXX() 的方式指定默认值。
- 

#### 开发进度

- [X] IPFSClient: Kubo RPC API client 
  - [x] /api/v0/add
  - [ ] /api/v0/cat 
  - [ ] /api/v0/commands
  - [ ] /api/v0/commands/completion/bash
  - [ ] /api/v0/dht/query  
  - [ ] /api/v0/get
  - [ ] /api/v0/id
  - [ ] /api/v0/log/level
  - [ ] /api/v0/log/ls
  - [ ] /api/v0/ls
  - [ ] /api/v0/ping 
  - [ ] /api/v0/refs
  - [ ] /api/v0/refs/local
  - [ ] /api/v0/resolve 
  - [ ] /api/v0/shutdown
  - [ ] /api/v0/update
  - [ ] /api/v0/version
  - [ ] /api/v0/version/deps 
- [ ] bitswap
  - [ ] /api/v0/bitswap/ledger
  - [ ] /api/v0/bitswap/reprovide
  - [ ] /api/v0/bitswap/stat
  - [ ] /api/v0/bitswap/wantlist
- [ ] block
  - [ ] /api/v0/block/get
  - [ ] /api/v0/block/put
  - [ ] /api/v0/block/rm
  - [ ] /api/v0/block/stat
- [ ] bootstrap
  - [ ] /api/v0/bootstrap
  - [ ] /api/v0/bootstrap/add
  - [ ] /api/v0/bootstrap/add/default
  - [ ] /api/v0/bootstrap/list
  - [ ] /api/v0/bootstrap/rm
  - [ ] /api/v0/bootstrap/rm/all
- [ ] cid
  - [ ] /api/v0/cid/base32
  - [ ] /api/v0/cid/bases
  - [ ] /api/v0/cid/codecs
  - [ ] /api/v0/cid/format
  - [ ] /api/v0/cid/hashes
- [ ] config
  - [ ] /api/v0/config
  - [ ] /api/v0/config/edit
  - [ ] /api/v0/config/profile/apply
  - [ ] /api/v0/config/replace
  - [ ] /api/v0/config/show
- [ ] dag
  - [ ] /api/v0/dag/export
  - [ ] /api/v0/dag/get
  - [ ] /api/v0/dag/import
  - [ ] /api/v0/dag/put
  - [ ] /api/v0/dag/resolve
  - [ ] /api/v0/dag/stat
- [ ] diag
  - [ ] /api/v0/diag/cmds
  - [ ] /api/v0/diag/cmds/clear
  - [ ] /api/v0/diag/cmds/set-time
  - [ ] /api/v0/diag/profile
  - [ ] /api/v0/diag/sys
- [ ] files
  - [ ] /api/v0/files/chcid
  - [ ] /api/v0/files/cp
  - [ ] /api/v0/files/flush
  - [ ] /api/v0/files/ls
  - [ ] /api/v0/files/mkdir
  - [ ] /api/v0/files/mv
  - [ ] /api/v0/files/read
  - [ ] /api/v0/files/rm
  - [ ] /api/v0/files/stat
  - [ ] /api/v0/files/write
- [ ] filestore
  - [ ] /api/v0/filestore/dups
  - [ ] /api/v0/filestore/ls
  - [ ] /api/v0/filestore/verify
- [ ] key
  - [ ] /api/v0/key/export
  - [ ] /api/v0/key/gen
  - [ ] /api/v0/key/import
  - [ ] /api/v0/key/list
  - [ ] /api/v0/key/rename
  - [ ] /api/v0/key/rm
  - [ ] /api/v0/key/rotate
- [ ] multibase
  - [ ] /api/v0/multibase/decode
  - [ ] /api/v0/multibase/encode
  - [ ] /api/v0/multibase/list
  - [ ] /api/v0/multibase/transcode
- [ ] name
  - [ ] /api/v0/name/publish
  - [ ] /api/v0/name/resolve
- [ ] pin
  - [ ] /api/v0/pin/add
  - [ ] /api/v0/pin/ls
  - [ ] /api/v0/pin/remote/add
  - [ ] /api/v0/pin/remote/ls
  - [ ] /api/v0/pin/remote/rm
  - [ ] /api/v0/pin/remote/service/add
  - [ ] /api/v0/pin/remote/service/ls
  - [ ] /api/v0/pin/remote/service/rm
  - [ ] /api/v0/pin/rm
  - [ ] /api/v0/pin/update
  - [ ] /api/v0/pin/verify
- [ ] repo
  - [ ] /api/v0/repo/gc
  - [ ] /api/v0/repo/migrate
  - [ ] /api/v0/repo/stat
  - [ ] /api/v0/repo/verify
  - [ ] /api/v0/repo/version
- [ ] routing
  - [ ] /api/v0/routing/findpeer
  - [ ] /api/v0/routing/findprovs
  - [ ] /api/v0/routing/get
  - [ ] /api/v0/routing/provide
  - [ ] /api/v0/routing/put
- [ ] stats
  - [ ] /api/v0/stats/bitswap
  - [ ] /api/v0/stats/bw
  - [ ] /api/v0/stats/dht
  - [ ] /api/v0/stats/provide
  - [ ] /api/v0/stats/repo
- [ ] swarm
  - [ ] /api/v0/swarm/addrs
  - [ ] /api/v0/swarm/addrs/listen
  - [ ] /api/v0/swarm/addrs/local
  - [ ] /api/v0/swarm/connect
  - [ ] /api/v0/swarm/disconnect
  - [ ] /api/v0/swarm/filters
  - [ ] /api/v0/swarm/filters/add
  - [ ] /api/v0/swarm/filters/rm
  - [ ] /api/v0/swarm/peering/add
  - [ ] /api/v0/swarm/peering/ls
  - [ ] /api/v0/swarm/peering/rm
  - [ ] /api/v0/swarm/peers
- [ ] Experimental RPC commands
  - [ ] /api/v0/log/tail
  - [ ] /api/v0/mount
  - [ ] /api/v0/name/pubsub/cancel
  - [ ] /api/v0/name/pubsub/state
  - [ ] /api/v0/name/pubsub/subs
  - [ ] /api/v0/p2p/close
  - [ ] /api/v0/p2p/forward
  - [ ] /api/v0/p2p/listen
  - [ ] /api/v0/p2p/ls
  - [ ] /api/v0/p2p/stream/close
  - [ ] /api/v0/p2p/stream/ls
  - [ ] /api/v0/pubsub/ls
  - [ ] /api/v0/pubsub/peers
  - [ ] /api/v0/pubsub/pub
  - [ ] /api/v0/pubsub/sub
  - [ ] /api/v0/swarm/limit
  - [ ] /api/v0/swarm/stats
- [ ] 以下废弃的 rpc 命令，不再实现。 
  - [ ] ~~/api/v0/dht/findpeer~~
  - [ ] ~~/api/v0/dht/findprovs~~
  - [ ] ~~/api/v0/dht/get~~
  - [ ] ~~/api/v0/dht/provide~~
  - [ ] ~~/api/v0/dht/put~~
  - [ ] ~~/api/v0/dnst~~
  - [ ] ~~/api/v0/file/ls~~
  - [ ] ~~/api/v0/object/data~~
  - [ ] ~~/api/v0/object/diff~~
  - [ ] ~~/api/v0/object/get~~
  - [ ] ~~/api/v0/object/links~~
  - [ ] ~~/api/v0/object/new~~
  - [ ] ~~/api/v0/object/patch/add-link~~
  - [ ] ~~/api/v0/object/patch/append-data~~
  - [ ] ~~/api/v0/object/patch/rm-link~~
  - [ ] ~~/api/v0/object/patch/set-data~~
  - [ ] ~~/api/v0/object/put~~
  - [ ] ~~/api/v0/object/stat~~
  - [ ] ~~/api/v0/repo/fsck~~
  - [ ] ~~/api/v0/tar/add~~
  - [ ] ~~/api/v0/tar/cat~~
  - [ ] ~~/api/v0/urlstore/add~~
 

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

