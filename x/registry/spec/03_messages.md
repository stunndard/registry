<!--
order: 3
-->

# Messages

In this section we describe the processing of the registration messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.

## MsgBuyName

A new name is created or updated using the `MsgBuyName` message.

```go
type MsgBuyName struct {
    Creator    string
    Name       string
    Bid        sdk.Coins
    OnSale     bool
}
```

This message is expected to fail if:

- Registrant address (`Creator`) is already an owner of this name
- this name is not for sale, namely `name.Onsale` is false
- the `Bid` amount is less than or equal to the current `name.Price`
- the Registrant address balance is less than `Bid` amount

This message either creates or updates the `Name` object and saves it to the store.


## MsgUpdateName

The `Owner`, `Price`, `Onsale` of a name can be updated using the
`MsgUpdateName`.

```go
type MsgUpdateName struct {
    Creator    string
    Name       string
    Owner      string
    Price      sdk.Coins
    OnSale     bool
}
```

This message is expected to fail if:

- there was no `Name` previously registered
- the `Creator` of this message is not the current owner of this name. `(Creator != name.Owner)`

This message stores the updated `Name` object.

## MsgDeleteName

This message deletes `Name` object from the store.

```go
type MsgDeleteName struct {
    Creator    string
    Name       string
}
```

This message is expected to fail if:

- there was no `Name` previously registered
- the `Creator` of this message is not the current owner of this name. `(Creator != name.Owner)`

