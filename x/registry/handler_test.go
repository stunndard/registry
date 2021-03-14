package registry_test

import (
	"strings"
	"testing"
	//"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"
	"github.com/stunndard/registry/x/registry"
	registrykeeper "github.com/stunndard/registry/x/registry/keeper"
	registrytypes "github.com/stunndard/registry/x/registry/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stunndard/registry/x/registry/keeper"
)

func TestInvalidMsg(t *testing.T) {
	k := registrykeeper.Keeper{}
	h := registry.NewHandler(k)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized registry message type"))
}

func TestMsgBuyName(t *testing.T) {
	// do need this?
	app := simapp.Setup(false)

	// TODO initialize keeper properly
	k := keeper.NewKeeper(
		app.BankKeeper,
		app.AppCodec(),
		app.GetKey(registrytypes.StoreKey),
		app.GetMemKey(registrytypes.StoreKey))


	h := registry.NewHandler(*k)

	// TODO initialize ctx properly
	// ctx := sdk.NewContext(nil, tmproto.Header{}, false, nil)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	pks := simapp.CreateTestPubKeys(1)
	addr := sdk.AccAddress(pks[0].Address()).String()

	res, err := h(ctx, &registrytypes.MsgBuyName{
		Creator: addr,
		Name:    "testname",
		Bid: sdk.Coins{sdk.NewCoin("token", sdk.NewInt(10)),
		},
		Onsale: false})

	require.NoError(t, err)
	require.NotNil(t, res)
	//require.True(t, strings.Contains(err.Error(), "unrecognized registry message type"))
}

