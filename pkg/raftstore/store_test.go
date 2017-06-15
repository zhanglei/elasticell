// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package raftstore

import (
	"testing"

	. "github.com/pingcap/check"
)

var _ = Suite(&testUtilSuite{})
var _ = Suite(&testStoreSuite{})
var _ = Suite(&testSplitSuite{})
var _ = Suite(&testRemoveSuite{})
var _ = Suite(&testRedisKVSuite{})
var _ = Suite(&testTransportSuite{})

func TestRaftStore(t *testing.T) {
	t.Parallel()
	TestingT(t)
}

type testStoreSuite struct {
	baseSuite
}

func (s *testStoreSuite) SetUpSuite(c *C) {
}

func (s *testStoreSuite) TearDownSuite(c *C) {
	s.stopMultiPDServers(c)
}

func (s *testStoreSuite) TestStart(c *C) {
	clusterID := s.restartMultiPDServer(c, 3)
	defer s.stopMultiPDServers(c)

	storeID := s.AllocID(c)
	store := s.bootstrap(c, clusterID, storeID)
	store.Stop()
}
