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

	keys := sdk.NewKVStoreKeys(registrytypes.StoreKey, registrytypes.MemStoreKey)
	//tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	//memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// initialize stores
  //app.MountKVStores(keys)
	//app.MountTransientStores(tkeys)
	//app.MountMemoryStores(memKeys)

	// TODO initialize keeper properly
	rk := keeper.NewKeeper(
		app.BankKeeper,
		app.AppCodec(),
		keys[registrytypes.StoreKey],
		keys[registrytypes.MemStoreKey])

	handler := registry.NewHandler(*rk)

	// TODO initialize ctx properly
	// ctx := sdk.NewContext(nil, tmproto.Header{}, false, nil)
	ctx := app.NewContext(false, tmproto.Header{})


	pks := simapp.CreateTestPubKeys(1)
	addr := sdk.AccAddress(pks[0].Address()).String()

	/*
		below currently panics when getting KVStore for the registry keeper in store/cachemulti/store.go:

	// GetKVStore returns an underlying KVStore by key.
	func (cms Store) GetKVStore(key types.StoreKey) types.KVStore {
--->		store := cms.stores[key]
cms.stores["registry"] is nil.

		if key == nil {
			panic(fmt.Sprintf("kv store with key %v has not been registered in stores", key))
		}
		return store.(types.KVStore)
	}
	There should be some way to initialize a store for a non standard keeper, such as "registry"
	*/
	res, err := handler(ctx, &registrytypes.MsgBuyName{
		Creator: addr,
		Name:    "testname",
		Bid: sdk.Coins{sdk.NewCoin("token", sdk.NewInt(10)),
		},
		Onsale: false})

	require.NoError(t, err)
	require.NotNil(t, res)

	/*
		so after MsgBuyName is successful, we can check
		if name was registered:
		require.True(t, rk.HasName(ctx, "testname"))

		then can check more by trying to register again the same name:
		...
		...
		and check for error:
		require.Error(t, err)

		then do other testcases, like trying to call MsgBuy name again from a different
		address and check for error:
		...
		...
		require.Error(t, err)

		Same with all testcases:
		- Try to buy a name that is not for sale: should fail
		- Try to buy a name that is for sale, but bid amount is less than the name price: should fail
		- Try to buy a name that is for sale but account holder address doesn't have enough coins for bid: should fail
		- Try to buy a name with bid that is ok and token balance is enough: should succees. after success: check rk.HasName for correctness.
			check buyer coin balance is correctly decreased. check seller coin balance correctly increased.

		- Try to update a name that doesn't belong to the account holder address: should fail
		- Try to update a name from correct owner: should success. check that the updated name reflected changes
			...
			...

		- Try to delete a name that doesn't belong to the account holder address: should fail
		- Try to delete a name from correct owner: should success. after success, check that there's
			no more such name returned by keeper:
			require.False(t, rk.HasName("xxx"))

			.. and so on
	*/

}

