# desmondtzh
**desmondtzh** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Commands

`./desmondtzhd q desmondtzh create-post "XXXX" "YYYY" --from alice` create post with title `XXXX` and body of `YYYY`.

`./desmondtzhd q desmondtzh show-post 0` show the 0th post.

`./desmondtzhd q desmondtzh list-by-title XXXX` get post with title with 'XXXX' text.

`./desmondtzhd q desmondtzh list-post` show all posts.

`./desmondtzhd q desmondtzh update-post "AAAA" "BBBBB" 0 --from alice` update the 0th post with title `AAAA` and body of 'BBBB'.

`./desmondtzhd q desmondtzh delete-post 0 --from alice` delete the 0th post.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/desmondtzh@latest! | sudo bash
```
`username/desmondtzh` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
