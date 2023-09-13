# Strike's Lightning Network Daemon

![Build Status](https://github.com/LN-Zap/lnd-strike/actions/workflows/main.yml/badge.svg)

![Logo](logo.png)

## Customisations

- `/lnrpc.Lightning/IsOurAddress` RPC to check whether an address is controlled by the node's wallet or not
  - NOTE: Merged upstream in https://github.com/lightningnetwork/lnd/pull/5476
- `/lnrpc.Lightning/GetTransaction` RPC to return data for any transaction given its id
  - Note: `btcwallet` dependency replaced with our own `btcwallet-strike` fork
  - See https://github.com/btcsuite/btcwallet/pull/779 for details
- `/walletrpc.WalletKit/SignMessageFromAddress` RPC that signs a given message with the private key for the given address
- `Dockerfile` customisations to allow fetching private fork of btcwallet dependency
- `.github/workflows/codeql.yml` githib actions workflow added to check code ssecurity
- `.github/workflows/automated-release.yml` github action added to create new releases for Strike
- `.github/workflows/docker.yml` github actions removed to prevent nightly builds
- `.github/workflows/main.yml` github actions workflow removed
- `.github/workflows/release.yaml` github actions workflow removed to prevent creating release on tag pushes
- `.github/workflows/version-check-on-pr.yaml` github actions workflow added to check version on PR to stable branch
- `--strike-version` lnd cli flag to print the LND-Strike version
- `Makefile` updated to support builds that include private forks of dependencies
- `README` updated with Strike specific docs

## Versioning

LND-Strike uses [semantic versioning](https://semver.org/) and incorporates the LND version it is based on via the _pre-release_ label. However, each time the LND version is increased it also results in a major, minor or patch level increase of LND-Strike (and the pre-release tag is therefore considered as informational only).

This allows updating LND-Strike as a dependency and knowing exactly what kind of changes are introduced (breaking changes, features or fixes).

LND-Strike is currently being supported in the following distinct versions:

 * `2.0.5-lnd.0.14.2.beta` following LND 0.14.2
 * `3.0.4-lnd.0.15.5.beta` following LND 0.15.5

The changes in each can be obtained from their releases.

## Branch Structure

### The Main Branches

#### `master`

We consider `master` to be the main branch where the source code of HEAD always reflects a state with the latest delivered development changes for the next release.

#### `stable`

We consider `stable` to be the main branch where the source code of HEAD always reflects a production-ready state. Each time when changes are merged back into `stable`, this is a new production release _by definition_.

#### `upstream/master`

We consider `upstream/master` to be the main branch where the source code of HEAD reflects the `master` branch of the upstream LND repository. This is tracked to merge new versions of LND into a production release.

### Suporting Branches

#### Feature Branches

Feature branches are used to develop new features and fixes for the upcoming or a distant future release.

* Naming convention: `feature-*` or `fix-*` to track a specific issue
* May branch off from: `master`
* Must merge back into: `master`

#### Release Branches

Release branches support preparation of a new production release. They allow for last-minute dotting of i’s and crossing t’s. Furthermore, they allow for minor bug fixes and preparing meta-data for a release (version number, build dates, etc.). By doing all of this work on a release branch, the `master` branch is cleared to receive features for the next big release.

The key moment to branch off a new release branch from `master` is when `master` (almost) reflects the desired state of the new release.

* Naming convention: `release-*`
* May branch off from: `master`
* Must merge back into: `master` and `stable`
* It is exactly at the start of a release branch that the upcoming release gets assigned a version number &ndash; not any earlier.

#### Rebased Branch

To keep better track of our custom changes made to LND we also maintain a branch `rebase-history`. With each merge of `upstream-master` into `master` all our custom changes are also rebased onto the `upstream-master` branch with `rebase-history`.

> **_NOTE:_** This branch should not be used other than as help to get an overview of all changes we made, as it will continuously rewrite history.
