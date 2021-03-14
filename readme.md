# registry

**registry** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Let's play

Alice wants to buy a new name "myname".
```
$ registryd tx registry buy-name myname 6token false --from alice --chain-id registry
```
"myname" is not yet registered, so the standard name registration fee is `10token`.

Alice balance is substracted by `10tokens` after buying this name.

Alice also passed the new price for "myname" name which is set to `6tokens`.
Which is not wise, because it means selling it for less than it was bought.

Alice also passed `false` for `onsale` which means Alice doesn't want to make this name available for sell to anyone else.

Let's check the current name:

```
$ registryd query registry list-name

Name:
- creator: <Alice address>
  id: "0"
  name: myname
  onsale: false
  owner: <Alice address>
  price:
  - amount: "6"
    denom: token

```

Now Alice wants to make money and sell this name to someone who wants to buy.
But before, Alice needs to update the name price and make this name available for sale:

```
$ registryd tx registry update-name myname <Alice address> 100token true --from alice --chain-id registry
```
Alice passed her own address for the new owner, that means the ownership didn't change and she still owns it.

Alice also passed `true` for `onsale`, that means that this name will be available for anyone to buy.

Alice also passed `100token` as the new starting price for this name. It means that the new buyer needs to supply a higher bid that `100token` to by this name.

Let's check for changes:

```
$ registryd query registry list-name

Name:
- creator: <Alice address>
  id: "0"
  name: myname
  onsale: true
  owner: <Alice address>
  price:
  - amount: "100"
    denom: token

```

"myname" is available to sell for more than `100tokens` now

Bob now wants to buy "myname":

```
registryd tx registry buy-name myname 101token false --from bob --chain-id registry
```

Bob decided to pay `101token` and set "myname" to not available for further sell.

After sell, `101token` will be sent from Bob to Alice.


Let's check:

```
$ registryd query registry list-name

Name:
- creator: <Alice address>
  id: "0"
  name: myname
  onsale: false
  owner: <Bob address>
  price:
  - amount: "101"
    denom: token
```

We can see that the owner's address now changed to Bob's.
We can see that the current price is set to `101token` which is the last price Bob paid for this name.
Bob also doesn't want to sell this name, so `onsale` is `false`. If anyone wants to buy it now, it's impossible unless Bob makes it available for sale.

After 5 years of owning this name, Bob decided that he doesn't need this name.
So he proceeds to deleting it:

```
registryd tx registry delete-name myname --from bob --chain-id registry
```

Let's check:

```
$ registryd query registry list-name

Name: []
$
```

There's no more "myname" registered.
It can bre registered again by anyone else.


## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
