/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package etcd

import (
	"errors"

	"github.com/HeavyHorst/easyKV"
	"github.com/HeavyHorst/easyKV/etcd/etcdv2"
	"github.com/HeavyHorst/easyKV/etcd/etcdv3"
)

// ErrUnknownApiLevel is returned if no valid api level is given
var ErrUnknownApiLevel = errors.New("unknown etcd api level - must be 2 or 3")

// NewEtcdClient returns an *etcd{2,3}.Client with a connection to named machines.
func NewEtcdClient(opts ...Option) (easyKV.StoreClient, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	if options.Version == 3 {
		return etcdv3.NewEtcdClient(options.Nodes, options.TLS.ClientCert, options.TLS.ClientKey, options.TLS.ClientCaKeys, options.Auth.BasicAuth, options.Auth.Username, options.Auth.Password)
	}

	if options.Version == 2 {
		return etcdv2.NewEtcdClient(options.Nodes, options.TLS.ClientCert, options.TLS.ClientKey, options.TLS.ClientCaKeys, options.Auth.BasicAuth, options.Auth.Username, options.Auth.Password)
	}

	return nil, ErrUnknownApiLevel
}
