package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/types/module/server"
	datamodule "github.com/regen-network/regen-ledger/x/data/module"
)

func TestGenesis(t *testing.T) {
	ff := server.NewFixtureFactory(t, 2)
	ff.SetModules([]module.Module{datamodule.Module{}})
	s := NewGenesisTestSuite(ff)
	suite.Run(t, s)
}
