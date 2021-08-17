module github.com/filecoin-project/specs-actors/v5

go 1.16

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-amt-ipld/v3 v3.1.0
	github.com/filecoin-project/go-bitfield v0.2.3
	github.com/filecoin-project/go-hamt-ipld/v3 v3.1.0
	github.com/filecoin-project/go-state-types v0.1.1-0.20210810190654-139e0e79e69e
	github.com/filecoin-project/specs-actors v0.9.13
	github.com/filecoin-project/specs-actors/v2 v2.3.5-0.20210114162132-5b58b773f4fb
	github.com/filecoin-project/specs-actors/v3 v3.1.0
	github.com/filecoin-project/specs-actors/v4 v4.0.0
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-ipld-cbor v0.0.5
	github.com/ipfs/go-ipld-format v0.0.2
	github.com/ipld/go-car v0.1.0
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/minio/sha256-simd v0.1.1
	github.com/multiformats/go-multihash v0.0.14
	github.com/nats-io/nats-server/v2 v2.3.2 // indirect
	github.com/nats-io/nats.go v1.11.1-0.20210623165838-4b75fc59ae30
	github.com/pkg/errors v0.9.1
	github.com/rickiey/loggo v0.4.0
	github.com/stretchr/testify v1.7.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20210118024343-169e9d70c0c2
	github.com/xorcare/golden v0.6.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/text v0.3.3
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

retract v5.0.0
