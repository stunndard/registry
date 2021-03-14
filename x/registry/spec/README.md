<!--
order: 0
title: Registry Overview
parent:
  title: "registry"
-->

# `registry`

## Abstract

This documentation specifies the Registry module.

The module enables Cosmos-SDK based blockchain to support a registry of names.
In this system, anyone can register a name, buy it from another owner or sell it.

## Contents

1. **[State](01_state.md)**
    - [Name](01_state.md#name)
2. **[State Transitions](02_state_transitions.md)**
    - [Register new name](02_state_transitions.md#register-a-new-name)
    - [Buy a registered name](02_state_transitions.md#buy-an-already-registered-name) 
    - [Update a registered name](02_state_transitions.md#update-a-registered-name) 
    - [Delete a registered name](02_state_transitions.md#delete-a-registered-name) 
3. **[Messages](03_messages.md)**
    - [MsgBuyname](03_messages.md#msgbuyname)
    - [MsgUpdateName](03_messages.md#msgupdatename)
    - [MsgDeleteName](03_messages.md#msgdeletename)
4. **[Begin-Block](04_begin_block.md)**
4. **[End-Block ](05_end_block.md)**
5. **[Hooks](06_hooks.md)**
6. **[Events](07_events.md)**
7. **[Parameters](08_params.md)**
