<!--
order: 1
-->

# State

## Name

Each registered name state is stored in a `Name` struct:

```go
type Name struct {
  Creator string        // address that created (registered this name for the first time)
  Id      string        // numerical id
  Name    string        // registered name
  Owner   string        // address of the current owner of this name
  Price   sdk.Coins     // price this name was last bought for, or the sale price if this name is for sale
  OnSale  bool          // if this name is for sale by the current owner
}
```

## Name registration

Name registrations are identified by the presence in the store of Name object.
The name is indexed in the store as follows:

- Name: `"Name-value-" + Name -> amino(name)`

Id is a monotonously increasing counter, incremented by 1 every time a new `Name` is insereted in the store.


## Id counter
Counter for name Id's is indexed in the store as follows:

- Counter: `"Name-count-" -> counter`