<!--
order: 2
-->

# State Transitions

This document describes the state transition operations:

## Name
State transitions for name are performed on every transaction that registers,
buys from different owner, updates or deletes the name registration.

### Register a new name

A new `Name` object is inserted in the store

- set `Name.Creator` and `Name.Owner` to the Registrant address
- set `Name.Name` to the registered name
- set `Name.Id` to the `counter + 1`
- set `Name.Price` to the amount that is passed by the Registrant
- set `Name.Onsale` to `true` or `false` that is passed by the Registrant
- set `counter` to `counter + 1`
- `10token` registration fee is substracted from the balance of Registrant

### Buy an already registered name

An existing `Name` object is modified in the following way :

- Bid amount is transfered from the Buyer address to the Seller address (current `Name.Owner`)
- set `Name.Owner` to the Buyer address
- set `Name.Price` to the Bid amount passed by the Buyer
- set `Name.Onsale` to `true` or `false` value passed by the Buyer

### Update a registered name

An existing `Name` object is modified in the following way :

- set `Name.Owner` to the new address that is passed by the current Owner
- set `Name.Price` to the new amount that is passed by the current Owner
- set `Name.Onsale` to `true` or `false` that is passed by the current Owner

### Delete a registered name

An existing `Name` object is deleted from the store

