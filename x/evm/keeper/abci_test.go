package keeper_test

import (
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) TestEndBlock() {
	em := suite.ctx.EventManager()
	// should emit an event for property NFT mint
	suite.Require().Equal(1, len(em.Events()))

	res := suite.app.EvmKeeper.EndBlock(suite.ctx, types.RequestEndBlock{})
	suite.Require().Equal([]types.ValidatorUpdate{}, res)

	// should emit 1 event for property NFT mint
	// and 1 EventTypeBlockBloom event on EndBlock
	suite.Require().Equal(2, len(em.Events()))
	suite.Require().Equal(evmtypes.EventTypeBlockBloom, em.Events()[1].Type)
}
