package coreapi

import (
	"context"

	core "github.com/ipfs/go-ipfs/core"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	path "github.com/ipfs/go-ipfs/path"

	ipld "gx/ipfs/QmUsVJ7AEnGyjX8YWnrwq9vmECVGwBQNAKPpgz5KSg8dcq/go-ipld-node"
	cid "gx/ipfs/QmcEcrBAMrwMyhSjXt4yfyPpzgSuV8HLHavnfmiKCSRqZU/go-cid"
)

type CoreAPI struct {
	node *core.IpfsNode
}

func NewCoreAPI(n *core.IpfsNode) coreiface.CoreAPI {
	api := &CoreAPI{n}
	return api
}

func (api *CoreAPI) Unixfs() coreiface.UnixfsAPI {
	return (*UnixfsAPI)(api)
}

func resolve(ctx context.Context, n *core.IpfsNode, ref coreiface.Ref) (ipld.Node, error) {
	c, err := cid.Parse(ref)
	if err != nil {
		return nil, err
	}

	dagnode, err := core.Resolve(ctx, n.Namesys, n.Resolver, path.FromCid(c))
	if err == core.ErrNoNamesys {
		return nil, coreiface.ErrOffline
	} else if err != nil {
		return nil, err
	}
	return dagnode, nil
}
