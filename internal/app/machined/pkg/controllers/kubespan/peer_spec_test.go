// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package kubespan_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/stretchr/testify/suite"
	"github.com/talos-systems/go-retry/retry"
	"inet.af/netaddr"

	kubespanctrl "github.com/talos-systems/talos/internal/app/machined/pkg/controllers/kubespan"
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1/machine"
	"github.com/talos-systems/talos/pkg/machinery/constants"
	"github.com/talos-systems/talos/pkg/resources/cluster"
	"github.com/talos-systems/talos/pkg/resources/config"
	"github.com/talos-systems/talos/pkg/resources/kubespan"
	runtimeres "github.com/talos-systems/talos/pkg/resources/runtime"
	"github.com/talos-systems/talos/pkg/resources/v1alpha1"
)

type PeerSpecSuite struct {
	KubeSpanSuite

	statePath string
}

func (suite *PeerSpecSuite) TestReconcile() {
	suite.statePath = suite.T().TempDir()

	suite.Require().NoError(suite.runtime.RegisterController(&kubespanctrl.PeerSpecController{}))

	suite.startRuntime()

	stateMount := runtimeres.NewMountStatus(v1alpha1.NamespaceName, constants.StatePartitionLabel)

	suite.Assert().NoError(suite.state.Create(suite.ctx, stateMount))

	cfg := kubespan.NewConfig(config.NamespaceName, kubespan.ConfigID)
	cfg.TypedSpec().Enabled = true

	suite.Require().NoError(suite.state.Create(suite.ctx, cfg))

	nodeIdentity := cluster.NewIdentity(cluster.NamespaceName, cluster.LocalIdentity)
	suite.Require().NoError(nodeIdentity.TypedSpec().Generate())
	suite.Require().NoError(suite.state.Create(suite.ctx, nodeIdentity))

	affiliate1 := cluster.NewAffiliate(cluster.NamespaceName, "7x1SuC8Ege5BGXdAfTEff5iQnlWZLfv9h1LGMxA2pYkC")
	*affiliate1.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      "7x1SuC8Ege5BGXdAfTEff5iQnlWZLfv9h1LGMxA2pYkC",
		Hostname:    "foo.com",
		Nodename:    "bar",
		MachineType: machine.TypeControlPlane,
		Addresses:   []netaddr.IP{netaddr.MustParseIP("192.168.3.4")},
		KubeSpan: cluster.KubeSpanAffiliateSpec{
			PublicKey:           "PLPNBddmTgHJhtw0vxltq1ZBdPP9RNOEUd5JjJZzBRY=",
			Address:             netaddr.MustParseIP("fd50:8d60:4238:6302:f857:23ff:fe21:d1e0"),
			AdditionalAddresses: []netaddr.IPPrefix{netaddr.MustParseIPPrefix("10.244.3.1/24"), netaddr.MustParseIPPrefix("10.244.3.0/32")},
			Endpoints:           []netaddr.IPPort{netaddr.MustParseIPPort("10.0.0.2:51820"), netaddr.MustParseIPPort("192.168.3.4:51820")},
		},
	}

	affiliate2 := cluster.NewAffiliate(cluster.NamespaceName, "9dwHNUViZlPlIervqX9Qo256RUhrfhgO0xBBnKcKl4F")
	*affiliate2.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      "9dwHNUViZlPlIervqX9Qo256RUhrfhgO0xBBnKcKl4F",
		Hostname:    "worker-1",
		Nodename:    "worker-1",
		MachineType: machine.TypeWorker,
		Addresses:   []netaddr.IP{netaddr.MustParseIP("192.168.3.5")},
	}

	affiliate3 := cluster.NewAffiliate(cluster.NamespaceName, "xCnFFfxylOf9i5ynhAkt6ZbfcqaLDGKfIa3gwpuaxe7F")
	*affiliate3.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      "xCnFFfxylOf9i5ynhAkt6ZbfcqaLDGKfIa3gwpuaxe7F",
		MachineType: machine.TypeWorker,
		Nodename:    "worker-2",
		Addresses:   []netaddr.IP{netaddr.MustParseIP("192.168.3.6")},
		KubeSpan: cluster.KubeSpanAffiliateSpec{
			PublicKey:           "mB6WlFOR66Jx5rtPMIpxJ3s4XHyer9NCzqWPP7idGRo",
			Address:             netaddr.MustParseIP("fdc8:8aee:4e2d:1202:f073:9cff:fe6c:4d67"),
			AdditionalAddresses: []netaddr.IPPrefix{netaddr.MustParseIPPrefix("10.244.4.1/24")},
			Endpoints:           []netaddr.IPPort{netaddr.MustParseIPPort("192.168.3.6:51820")},
		},
	}

	// local node affiliate, should be skipped as a peer
	affiliate4 := cluster.NewAffiliate(cluster.NamespaceName, nodeIdentity.TypedSpec().NodeID)
	*affiliate4.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      nodeIdentity.TypedSpec().NodeID,
		MachineType: machine.TypeWorker,
		Addresses:   []netaddr.IP{netaddr.MustParseIP("192.168.3.7")},
		KubeSpan: cluster.KubeSpanAffiliateSpec{
			PublicKey:           "27E8I+ekrqT21cq2iW6+fDe+H7WBw6q9J7vqLCeswiM=",
			Address:             netaddr.MustParseIP("fdc8:8aee:4e2d:1202:f073:9cff:fe6c:4d67"),
			AdditionalAddresses: []netaddr.IPPrefix{netaddr.MustParseIPPrefix("10.244.5.1/24")},
			Endpoints:           []netaddr.IPPort{netaddr.MustParseIPPort("192.168.3.7:51820")},
		},
	}

	for _, r := range []resource.Resource{affiliate1, affiliate2, affiliate3, affiliate4} {
		suite.Require().NoError(suite.state.Create(suite.ctx, r))
	}

	// affiliate2 shouldn't be rendered as a peer, as it doesn't have kubespan data
	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertResourceIDs(resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, "", resource.VersionUndefined),
			[]resource.ID{
				affiliate1.TypedSpec().KubeSpan.PublicKey,
				affiliate3.TypedSpec().KubeSpan.PublicKey,
			},
		),
	))

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertResource(
			resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, affiliate1.TypedSpec().KubeSpan.PublicKey, resource.VersionUndefined),
			func(res resource.Resource) error {
				spec := res.(*kubespan.PeerSpec).TypedSpec()

				suite.Assert().Equal("fd50:8d60:4238:6302:f857:23ff:fe21:d1e0", spec.Address.String())
				suite.Assert().Equal("[10.244.3.0/24 fd50:8d60:4238:6302:f857:23ff:fe21:d1e0/128]", fmt.Sprintf("%v", spec.AllowedIPs))
				suite.Assert().Equal([]netaddr.IPPort{netaddr.MustParseIPPort("10.0.0.2:51820"), netaddr.MustParseIPPort("192.168.3.4:51820")}, spec.Endpoints)
				suite.Assert().Equal("bar", spec.Label)

				return nil
			},
		),
	))

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertResource(
			resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, affiliate3.TypedSpec().KubeSpan.PublicKey, resource.VersionUndefined),
			func(res resource.Resource) error {
				spec := res.(*kubespan.PeerSpec).TypedSpec()

				suite.Assert().Equal("fdc8:8aee:4e2d:1202:f073:9cff:fe6c:4d67", spec.Address.String())
				suite.Assert().Equal("[10.244.4.0/24 fdc8:8aee:4e2d:1202:f073:9cff:fe6c:4d67/128]", fmt.Sprintf("%v", spec.AllowedIPs))
				suite.Assert().Equal([]netaddr.IPPort{netaddr.MustParseIPPort("192.168.3.6:51820")}, spec.Endpoints)
				suite.Assert().Equal("worker-2", spec.Label)

				return nil
			},
		),
	))

	// disabling kubespan should remove all peers
	oldVersion := cfg.Metadata().Version()
	cfg.TypedSpec().Enabled = false
	cfg.Metadata().BumpVersion()

	suite.Require().NoError(suite.state.Update(suite.ctx, oldVersion, cfg))

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertNoResource(
			resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, affiliate1.TypedSpec().KubeSpan.PublicKey, resource.VersionUndefined),
		),
	))
	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertNoResource(
			resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, affiliate3.TypedSpec().KubeSpan.PublicKey, resource.VersionUndefined),
		),
	))
}

func (suite *PeerSpecSuite) TestIPOverlap() {
	suite.statePath = suite.T().TempDir()

	suite.Require().NoError(suite.runtime.RegisterController(&kubespanctrl.PeerSpecController{}))

	suite.startRuntime()

	stateMount := runtimeres.NewMountStatus(v1alpha1.NamespaceName, constants.StatePartitionLabel)

	suite.Assert().NoError(suite.state.Create(suite.ctx, stateMount))

	cfg := kubespan.NewConfig(config.NamespaceName, kubespan.ConfigID)
	cfg.TypedSpec().Enabled = true

	suite.Require().NoError(suite.state.Create(suite.ctx, cfg))

	nodeIdentity := cluster.NewIdentity(cluster.NamespaceName, cluster.LocalIdentity)
	suite.Require().NoError(nodeIdentity.TypedSpec().Generate())
	suite.Require().NoError(suite.state.Create(suite.ctx, nodeIdentity))

	affiliate1 := cluster.NewAffiliate(cluster.NamespaceName, "7x1SuC8Ege5BGXdAfTEff5iQnlWZLfv9h1LGMxA2pYkC")
	*affiliate1.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      "7x1SuC8Ege5BGXdAfTEff5iQnlWZLfv9h1LGMxA2pYkC",
		Nodename:    "bar",
		MachineType: machine.TypeControlPlane,
		KubeSpan: cluster.KubeSpanAffiliateSpec{
			PublicKey:           "PLPNBddmTgHJhtw0vxltq1ZBdPP9RNOEUd5JjJZzBRY=",
			Address:             netaddr.MustParseIP("fd50:8d60:4238:6302:f857:23ff:fe21:d1e0"),
			AdditionalAddresses: []netaddr.IPPrefix{netaddr.MustParseIPPrefix("10.244.3.1/24"), netaddr.MustParseIPPrefix("10.244.3.0/32")},
			Endpoints:           []netaddr.IPPort{netaddr.MustParseIPPort("10.0.0.2:51820"), netaddr.MustParseIPPort("192.168.3.4:51820")},
		},
	}

	affiliate2 := cluster.NewAffiliate(cluster.NamespaceName, "9dwHNUViZlPlIervqX9Qo256RUhrfhgO0xBBnKcKl4F")
	*affiliate2.TypedSpec() = cluster.AffiliateSpec{
		NodeID:      "9dwHNUViZlPlIervqX9Qo256RUhrfhgO0xBBnKcKl4F",
		Hostname:    "worker-1",
		Nodename:    "worker-1",
		MachineType: machine.TypeWorker,
		KubeSpan: cluster.KubeSpanAffiliateSpec{
			PublicKey:           "Zr5ewpUm2Ywo1c+/59WFKIBjZ3c/nVbIWsT5elbjwCU=",
			Address:             netaddr.MustParseIP("fd50:8d60:4238:6302:f857:23ff:fe21:d1e0"),
			AdditionalAddresses: []netaddr.IPPrefix{netaddr.MustParseIPPrefix("10.244.3.128/28"), netaddr.MustParseIPPrefix("192.168.3.0/24")},
			Endpoints:           []netaddr.IPPort{netaddr.MustParseIPPort("10.0.0.2:51820"), netaddr.MustParseIPPort("192.168.3.4:51820")},
		},
	}

	for _, r := range []resource.Resource{affiliate1, affiliate2} {
		suite.Require().NoError(suite.state.Create(suite.ctx, r))
	}

	// affiliate2 shouldn't be rendered as a peer, as its AdditionalAddresses overlap with affiliate1 addresses
	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertResourceIDs(resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, "", resource.VersionUndefined),
			[]resource.ID{
				affiliate1.TypedSpec().KubeSpan.PublicKey,
			},
		),
	))

	suite.Assert().NoError(retry.Constant(3*time.Second, retry.WithUnits(100*time.Millisecond)).Retry(
		suite.assertNoResource(
			resource.NewMetadata(kubespan.NamespaceName, kubespan.PeerSpecType, affiliate2.TypedSpec().KubeSpan.PublicKey, resource.VersionUndefined),
		),
	))
}

func TestPeerSpecSuite(t *testing.T) {
	suite.Run(t, new(PeerSpecSuite))
}
