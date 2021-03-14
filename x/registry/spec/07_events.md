<!--
order: 6
-->

# Events

The registry module emits the following events:

## Handlers

### MsgBuyName

| Type             | Attribute Key | Attribute Value    |
| ---------------- | ------------- | ------------------ |
| register         | registrant    | registrant address  |
| register         | name          | registered name     |
| register         | amount        | registration fee paid      |
| register         | onsale        | if registered name is on sale     |
| buy              | buyer         | buyer address       |
| buy              | seller        | seller address      |
| buy              | name          | registered name |          |
| buy              | amount        | amount for which name is sold |          |
| buy              | onsale        | if sold name is on sale again         |


### MsgUpdateName

| Type           | Attribute Key       | Attribute Value     |
| -------------- | ------------------- | ------------------- |
| update         | owner               | owner address        |
| update         | new_owner           | new owner address     |
| update         | name                | registered name      |       |
| update         | amount              | name new price            |
| update         | onsale              | if name is on sale             |

### MsgDeleteName

| Type           | Attribute Key       | Attribute Value     |
| -------------- | ------------------- | ------------------- |
| delete         | owner               | owner address        |
| delete         | name                | name that is deleted |
|
