# Changelog

All notable changes to this project will be documented in this file.

## [1.36.0](https://github.com/equinor/radix-cli/compare/v1.35.2..v1.36.0) - 2026-01-09

### 🚀 Features

- Show radixconfig validation warnings (#153) - ([b97ed49](https://github.com/equinor/radix-cli/commit/b97ed49f049eb9b1082a571777642ead35653c66)) by @nilsgstrabo in [#153](https://github.com/equinor/radix-cli/pull/153)


## [1.35.2](https://github.com/equinor/radix-cli/compare/v1.35.1..v1.35.2) - 2025-11-26

### 🐛 Bug Fixes

- Support new fields in radixconfig.yaml with validate radix-config command - ([ddd71e9](https://github.com/equinor/radix-cli/commit/ddd71e9b2c454c793f3f8bf45d1064bde7daaf3f)) by @nilsgstrabo in [#149](https://github.com/equinor/radix-cli/pull/149)


## [1.35.1](https://github.com/equinor/radix-cli/compare/v1.35.0..v1.35.1) - 2025-11-17

### 🐛 Bug Fixes

- Always return an error message when context is invalid. (#145) - ([2b388d8](https://github.com/equinor/radix-cli/commit/2b388d85d56a11fbf6cd6a5dea377b5522f163a8)) by @jacobsolbergholm in [#145](https://github.com/equinor/radix-cli/pull/145)


### ⚙️ Miscellaneous Tasks

- Refactor release process (#144) - ([2baf5d3](https://github.com/equinor/radix-cli/commit/2baf5d30a20c0a45483df0be878ab9cea31d88aa)) by @nilsgstrabo in [#144](https://github.com/equinor/radix-cli/pull/144)


## New Contributors ❤️

* @github-actions[bot] made their first contribution in [#146](https://github.com/equinor/radix-cli/pull/146)
* @jacobsolbergholm made their first contribution in [#145](https://github.com/equinor/radix-cli/pull/145)
## [1.35.0](https://github.com/equinor/radix-cli/compare/v1.34.0..v1.35.0) - 2025-10-07

### 🚀 Features

- Get cluster config (#143) - ([ff462c8](https://github.com/equinor/radix-cli/commit/ff462c8fcefd033b973db841ef8020de6319235d)) by @Richard87 in [#143](https://github.com/equinor/radix-cli/pull/143)


## [1.34.0](https://github.com/equinor/radix-cli/compare/v1.33.4..v1.34.0) - 2025-09-30

### 🚀 Features

- Stream logs (#139) - ([9b2d03b](https://github.com/equinor/radix-cli/commit/9b2d03b81eaf72a6371a5c4acb0f640ceb982bf9)) by @Richard87 in [#139](https://github.com/equinor/radix-cli/pull/139)


## [1.33.4](https://github.com/equinor/radix-cli/compare/v1.33.3..v1.33.4) - 2025-08-11

### 🐛 Bug Fixes

- Panic when using --token-environment and APP_SERVICE_ACCOUNT_TOKEN not set - ([27e6f1c](https://github.com/equinor/radix-cli/commit/27e6f1c2cd020baa76dab321080b9a020f4c38c3)) by @nilsgstrabo in [#137](https://github.com/equinor/radix-cli/pull/137)


## [1.33.3](https://github.com/equinor/radix-cli/compare/v1.33.2..v1.33.3) - 2025-08-08

### 🐛 Bug Fixes

- Remove Prometheus version replacement from go.mod (#136) - ([1b5c7a1](https://github.com/equinor/radix-cli/commit/1b5c7a1d4564e049b6aab994b5152d29ded635df)) by @Richard87 in [#136](https://github.com/equinor/radix-cli/pull/136)


## [1.33.2](https://github.com/equinor/radix-cli/compare/v1.33.1..v1.33.2) - 2025-07-29

### ⚙️ Miscellaneous Tasks

- Updated refs to support command and args in job-descriptions (#135) - ([e60f92c](https://github.com/equinor/radix-cli/commit/e60f92c783c42ddf513d067cde4340c00b6b41f1)) by @satr in [#135](https://github.com/equinor/radix-cli/pull/135)


## [1.33.1](https://github.com/equinor/radix-cli/compare/v1.33.0..v1.33.1) - 2025-07-09

### 💼 Other

- Updated refs (#133) - ([444ed43](https://github.com/equinor/radix-cli/commit/444ed4381cf3e503cf1a24e1b0ecf9fecc5153bc)) by @satr in [#133](https://github.com/equinor/radix-cli/pull/133)


## [1.33.0](https://github.com/equinor/radix-cli/compare/v1.32.0..v1.33.0) - 2025-06-24

### 💼 Other

- Updated refs, Regenerated radix-app client (#131) - ([45e032c](https://github.com/equinor/radix-cli/commit/45e032c3250330c58fe79285351d6e084c196b8a)) by @satr in [#131](https://github.com/equinor/radix-cli/pull/131)


## [1.32.0](https://github.com/equinor/radix-cli/compare/v1.31.0..v1.32.0) - 2025-05-27

### 💼 Other

- Added options tag and from-type to build and build-deploy pipeline jobs (#130)

* Added options tag and from-type to build and build-deploy pipeline jobs

* Cleanup

* Cleanup

* Removed from-type option - ([b904bfe](https://github.com/equinor/radix-cli/commit/b904bfe9edd7486543794873cd996050628fb64b)) by @satr in [#130](https://github.com/equinor/radix-cli/pull/130)


## [1.31.0](https://github.com/equinor/radix-cli/compare/v1.30.2..v1.31.0) - 2025-05-14

### 💼 Other

- 1205-refresh-build-kit-cache (#128)

* Regenerated, added triggered from webhook and build-kit and cache related fields

* Added option refresh-build-cache

* Updated version

* Regenerated - ([4c96fd2](https://github.com/equinor/radix-cli/commit/4c96fd2c324c5c57935e2718dd0e2d75b2e5c976)) by @satr in [#128](https://github.com/equinor/radix-cli/pull/128)


## [1.30.2](https://github.com/equinor/radix-cli/compare/v1.30.1..v1.30.2) - 2025-04-23

### 🐛 Bug Fixes

- Pass AzureClientId to Github Provider, fix expiry (#126) - ([a218b3e](https://github.com/equinor/radix-cli/commit/a218b3eb48e61e670df10495686987c540a47f65)) by @Richard87 in [#126](https://github.com/equinor/radix-cli/pull/126)


## [1.30.1](https://github.com/equinor/radix-cli/compare/v1.30.0..v1.30.1) - 2025-04-23

### 🐛 Bug Fixes

- Authenticate with github federated credentials (#125) - ([578e62d](https://github.com/equinor/radix-cli/commit/578e62d3654e33ce8f220c364e195575a6c524e1)) by @Richard87 in [#125](https://github.com/equinor/radix-cli/pull/125)


## [1.30.0](https://github.com/equinor/radix-cli/compare/v1.29.0..v1.30.0) - 2025-04-22

### 🚀 Features

- Login with client secret or federated credentials (#123) - ([27c05d2](https://github.com/equinor/radix-cli/commit/27c05d243db0fd460dc7575a8df995d8748b5325)) by @Richard87 in [#123](https://github.com/equinor/radix-cli/pull/123)


## [1.29.0](https://github.com/equinor/radix-cli/compare/v1.28.0..v1.29.0) - 2025-04-09

### 🚀 Features

- Add commands set private-image-hub-secret and set build-secret (#121) - ([1690bfd](https://github.com/equinor/radix-cli/commit/1690bfd45b492fdc020b855376151915dd01f801)) by @Richard87 in [#121](https://github.com/equinor/radix-cli/pull/121)


## [1.28.0](https://github.com/equinor/radix-cli/compare/v1.27.0..v1.28.0) - 2025-04-08

### 💼 Other

- Added an option commit-id to the promote job (#122)

* Added an option commit-id to the promote job

* Added an option commit-id to the promote job - ([d155ed6](https://github.com/equinor/radix-cli/commit/d155ed6296f334956fdebb9f3f36db724f2e9106)) by @satr in [#122](https://github.com/equinor/radix-cli/pull/122)


## [1.27.0](https://github.com/equinor/radix-cli/compare/v1.26.1..v1.27.0) - 2025-02-28

### 💼 Other

- Fixed stop env batches and jobs (#118) - ([ab6e095](https://github.com/equinor/radix-cli/commit/ab6e095a86e0250de2e38fac14c7c77bf7668cbc)) by @satr in [#118](https://github.com/equinor/radix-cli/pull/118)


## [1.26.1](https://github.com/equinor/radix-cli/compare/v1.26.0..v1.26.1) - 2025-02-28

### 💼 Other

- Added stop scheduled-job commands (#116)

* Added stop scheduled-job commands

* Added examples for the scheduled-job commands

* Updated versions

* Added autocomplete for jobs and batches

* Update cmd/stopScheduledJobs.go

Co-authored-by: Richard Hagen <richard.hagen@gmail.com>

* Added endpoints to stop jobs and batches for entire environment

* Added command to get batches and jobs

* Extended command to get batches and jobs

* Added silence error option

* Added batch summary test

* Added job summary text

---------

Co-authored-by: Richard Hagen <richard.hagen@gmail.com> - ([bc940f1](https://github.com/equinor/radix-cli/commit/bc940f1fc3a890b7ec72e140c2339add009ceccc)) by @satr in [#116](https://github.com/equinor/radix-cli/pull/116)


## [1.26.0](https://github.com/equinor/radix-cli/compare/v1.25.0..v1.26.0) - 2025-02-21

### 💼 Other

- 154 vulnerability scans accessible through api or grafana (#115)

* move the generated api client to new folder

* Add vulnscanapi

* Rename GetForCommand to GetRadixApiForCommand for clarity

* refactor client for different endpoints

* refactor client for different endpoints

* Add vulnerability command to retrieve scan results for applications

* Update cmd/getVulnerability.go

Co-authored-by: Sergey Smolnikov <ssmol@equinor.com>

* Fix error handling in getVulnerability.go for improved clarity

* Update cmd/getVulnerability.go

Co-authored-by: Sergey Smolnikov <ssmol@equinor.com>

* Enhance getVulnerability.go to support job and component detection for vulnerability scans

* Refactor getVulnerability.go to enhance job and component detection for vulnerability scans

* Refactor getVulnerability.go to update component model references for improved clarity

* remove quoted names

---------

Co-authored-by: Sergey Smolnikov <ssmol@equinor.com> - ([13a6146](https://github.com/equinor/radix-cli/commit/13a6146a67e24eaa353fbac2e3a5d6553b2642d2)) by @Richard87 in [#115](https://github.com/equinor/radix-cli/pull/115)


## [1.25.0](https://github.com/equinor/radix-cli/compare/v1.24.1..v1.25.0) - 2025-02-18

### 💼 Other

- Updated version and api (#114)

* Updated version and api

* Updated api and optional descriptions

* Updated versions

* Cleanup

* Cleanup

* Cleanup - ([24bb9ee](https://github.com/equinor/radix-cli/commit/24bb9ee18b99c46462c0211ae4fefb97504e11b5)) by @satr in [#114](https://github.com/equinor/radix-cli/pull/114)


## [1.24.1](https://github.com/equinor/radix-cli/compare/v1.24.0..v1.24.1) - 2024-12-17

### 💼 Other

- Bump golang.org/x/crypto from 0.26.0 to 0.31.0

Bumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.26.0 to 0.31.0.
- [Commits](https://github.com/golang/crypto/compare/v0.26.0...v0.31.0)

---
updated-dependencies:
- dependency-name: golang.org/x/crypto
  dependency-type: indirect
...

Signed-off-by: dependabot[bot] <support@github.com> - ([fa24e7d](https://github.com/equinor/radix-cli/commit/fa24e7d982c6dfecdee6253c4a05e8221559fb84)) by @dependabot[bot]

- Update github.com/equinor/radix-operator to v1.68.1 - ([854e605](https://github.com/equinor/radix-cli/commit/854e6054b45d52378d40290d255780b2201d22f9)) by @nilsgstrabo

- Merge pull request #111 from equinor/dependabot/go_modules/golang.org/x/crypto-0.31.0

Bump golang.org/x/crypto from 0.26.0 to 0.31.0 - ([a44a3be](https://github.com/equinor/radix-cli/commit/a44a3be1c86987b42575b3ede02d6a5bb971f33a)) by @nilsgstrabo in [#111](https://github.com/equinor/radix-cli/pull/111)


## [1.24.0](https://github.com/equinor/radix-cli/compare/v1.23.0..v1.24.0) - 2024-11-12

### 💼 Other

- Added build deploy option environment (#109)

* Added build deploy option environment

* Added external DNS aliases option to the apply-config

* Update cmd/createApplyConfigPipelineJob.go

Co-authored-by: Richard Hagen <richard.hagen@gmail.com>

* Added error handling

---------

Co-authored-by: Richard Hagen <richard.hagen@gmail.com> - ([9fd4fce](https://github.com/equinor/radix-cli/commit/9fd4fcebd80b50edf70ecf229fb4d66ef123978a)) by @satr in [#109](https://github.com/equinor/radix-cli/pull/109)


## [1.23.0](https://github.com/equinor/radix-cli/compare/v1.22.0..v1.23.0) - 2024-10-22

### 💼 Other

- Added build pipeline jobs (#108)

* Added build pipeline jobs

* Update cmd/createBuildPipelineJob.go

Co-authored-by: Richard Hagen <richard.hagen@gmail.com>

* Update cmd/createBuildPipelineJob.go

Co-authored-by: Richard Hagen <richard.hagen@gmail.com>

---------

Co-authored-by: Richard Hagen <richard.hagen@gmail.com> - ([87e887b](https://github.com/equinor/radix-cli/commit/87e887baa7b24909b5dd1b6a12e8e35460cf218d)) by @satr in [#108](https://github.com/equinor/radix-cli/pull/108)


## [1.22.0](https://github.com/equinor/radix-cli/compare/v1.21.0..v1.22.0) - 2024-10-02

### 💼 Other

- Get-resources (#106)

* Added resources endpoint

* Added resources endpoint

* Corrected resources command arguments

* Removed outdated const

* Added flag completion

* Added flag completion

* Added flag completion - ([2c7d3d3](https://github.com/equinor/radix-cli/commit/2c7d3d39d9555f3d65e8ccada5fa3bc45df2a4a4)) by @satr in [#106](https://github.com/equinor/radix-cli/pull/106)


## [1.21.0](https://github.com/equinor/radix-cli/compare/v1.20.1..v1.21.0) - 2024-09-26

### 💼 Other

- Autocomplete flags (#107)

* Add Autocomplete to flags

* Generic cache

* Add Application autocompletion to all commands

* Fix Application autocompletion to all commands, Add Component autocomplete

* Add Component autocomplete

* remove appName pointer to cleanup code

* Add autocomplete of Deployments

* Add autocomplete of Variables and Secrets

* better failsafe

* Set n to shorthand for component everywhere

* Add Jobs autocompletion

* revert "n" as shortcut for component

* remove RADIX_ from available variables

* use slice.FindAll() instead of slices.Filter() - ([3afbaad](https://github.com/equinor/radix-cli/commit/3afbaada9ebec8775abf5719e48ea4f1b0b55cb6)) by @Richard87 in [#107](https://github.com/equinor/radix-cli/pull/107)


## [1.20.1](https://github.com/equinor/radix-cli/compare/v1.20.0..v1.20.1) - 2024-09-12

### 💼 Other

- Scale components (#105)

* Manually Scale components, add reset command

* deprecate start component command

* cleanup, use deprecated and example fields

* better description of reset command - ([c0ab9f4](https://github.com/equinor/radix-cli/commit/c0ab9f43e825bbd7200d9bd4f385b97800b365f5)) by @Richard87 in [#105](https://github.com/equinor/radix-cli/pull/105)


## [1.20.0](https://github.com/equinor/radix-cli/compare/v1.19.2..v1.20.0) - 2024-08-26

### 💼 Other

- Add-pipeline-job-option-to-control-using-build-cache (#104)

* Added option

* Set ref

* Re-generated client

* Added override use cache option - ([2160624](https://github.com/equinor/radix-cli/commit/2160624cbfafcca2efd6c84b2d177275a47b34e8)) by @satr in [#104](https://github.com/equinor/radix-cli/pull/104)


## [1.19.2](https://github.com/equinor/radix-cli/compare/v1.19.1..v1.19.2) - 2024-07-26

### 💼 Other

- Update-refs (#101)

* Updated radix-operator version

* Cleanup - ([a8e1249](https://github.com/equinor/radix-cli/commit/a8e124925d159750712342ee827acab2d677c7b5)) by @satr in [#101](https://github.com/equinor/radix-cli/pull/101)


## [1.19.1](https://github.com/equinor/radix-cli/compare/v1.19.0..v1.19.1) - 2024-06-19

### 🐛 Bug Fixes

- Get appllication list now prints to stdout - ([c3a8a1b](https://github.com/equinor/radix-cli/commit/c3a8a1b7251e863bb0a84cb7d34d2aafe9f35809)) by @FredrikMWold


### 💼 Other

- Merge pull request #99 from FredrikMWold/fix/get-application-list-bug

fix: get appllication list now prints to stdout - ([f0cc40c](https://github.com/equinor/radix-cli/commit/f0cc40cf11014b791563d298e6cbeb4e0522df42)) by @nilsgstrabo in [#99](https://github.com/equinor/radix-cli/pull/99)


## New Contributors ❤️

* @FredrikMWold made their first contribution
## [1.19.0](https://github.com/equinor/radix-cli/compare/v1.18.1..v1.19.0) - 2024-06-18

### 💼 Other

- Update module references (#98)

validation support for `runtime` section in radixconfig
add flag `--strict-validation` in command validate `radix-config` - ([723000d](https://github.com/equinor/radix-cli/commit/723000d1f9063073dc3feec7e3891683f51dfc2d)) by @nilsgstrabo in [#98](https://github.com/equinor/radix-cli/pull/98)


## [1.18.1](https://github.com/equinor/radix-cli/compare/v1.18.0..v1.18.1) - 2024-06-04

### 💼 Other

- Update golangci-lint, goreleaser, remove shorthand flag, add comment about local file and file:// (#96) - ([e73944f](https://github.com/equinor/radix-cli/commit/e73944fa1371f43ec04f909b2a980dc0d0d32750)) by @Richard87 in [#96](https://github.com/equinor/radix-cli/pull/96)


## [1.18.0](https://github.com/equinor/radix-cli/compare/v1.17.0..v1.18.0) - 2024-06-03

### 💼 Other

- Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 (#88)

Bumps google.golang.org/protobuf from 1.31.0 to 1.33.0.

---
updated-dependencies:
- dependency-name: google.golang.org/protobuf
  dependency-type: indirect
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
Co-authored-by: Sergey Smolnikov <ssmol@equinor.com> - ([466e58f](https://github.com/equinor/radix-cli/commit/466e58f77de29a92951c29ebb4ee9641ddfae0f1)) by @dependabot[bot] in [#88](https://github.com/equinor/radix-cli/pull/88)

- Schema and strict validation radixconfig (#95)

* validate against crd and strictly

* cleanup, use jsonschema directly

* update go.mod

* updated operator

* update golangci-lint, goreleaser, remove shorthand flag, add comment about local file and file:// - ([103a074](https://github.com/equinor/radix-cli/commit/103a074a9747b8f6a3204dfddfdb08a596a86878)) by @Richard87 in [#95](https://github.com/equinor/radix-cli/pull/95)


## [1.17.0](https://github.com/equinor/radix-cli/compare/v1.16.1..v1.17.0) - 2024-04-19

### 💼 Other

- Add-pipeline-job-apply-config (#93)

* Added apply-config pipeline-job type

* Updated radix-operator version - ([6ba1912](https://github.com/equinor/radix-cli/commit/6ba1912b80ea2de724f8d75f5800a43aa7b42447)) by @satr in [#93](https://github.com/equinor/radix-cli/pull/93)


## [1.16.1](https://github.com/equinor/radix-cli/compare/v1.16.0..v1.16.1) - 2024-04-15

### 💼 Other

- Update radix-operator version - ([aa2df2f](https://github.com/equinor/radix-cli/commit/aa2df2fdd175e70fad300a58955c63e09b2c0ac5)) by @nilsgstrabo

- Merge pull request #91 from equinor/update-radix-operator-version

update radix-operator version - ([dd191e7](https://github.com/equinor/radix-cli/commit/dd191e74c1324060178cf4a3b6881d41cac1bc98)) by @nilsgstrabo in [#91](https://github.com/equinor/radix-cli/pull/91)


## [1.16.0](https://github.com/equinor/radix-cli/compare/v1.15.0..v1.16.0) - 2024-04-12

### 🐛 Bug Fixes

- Dockerfile.goreleaser to reduce vulnerabilities (#86) - ([8f0dd69](https://github.com/equinor/radix-cli/commit/8f0dd69cb32faf2e793b12dc5f0003bb509156c7)) by @satr in [#86](https://github.com/equinor/radix-cli/pull/86)


### 💼 Other

- Easy RX CLI Install (#87)

* move main and read build info

* set version

* tweaks

* Tweak readme - ([01969bf](https://github.com/equinor/radix-cli/commit/01969bff441d97146d32b5c5ffb4ad12b0602afe)) by @Richard87 in [#87](https://github.com/equinor/radix-cli/pull/87)

- Replace device code login with interactive login (#90)

* replace device code login with interactive login

* add flag --use-device-code for login command

* add SilenceUsage and fix usage for flag

* extend help for use-device-code

* update help - ([f26d0a8](https://github.com/equinor/radix-cli/commit/f26d0a88d66fc615a88a9b600b0658d748914d27)) by @nilsgstrabo in [#90](https://github.com/equinor/radix-cli/pull/90)


## [1.15.0](https://github.com/equinor/radix-cli/compare/v1.14.1..v1.15.0) - 2024-01-22

### 💼 Other

- Added an option component to the create deploy pipeline job (#84)

* Added an option component to the create deploy pipeline job

* Fixed error handling

* Merged

* Changed api option componentToDeploy to array

* Using example field - ([bc56b79](https://github.com/equinor/radix-cli/commit/bc56b795a0aebf6d9962c748b383dcf040e04496)) by @satr in [#84](https://github.com/equinor/radix-cli/pull/84)


## [1.14.1](https://github.com/equinor/radix-cli/compare/v1.14.0..v1.14.1) - 2024-01-18

### 💼 Other

- Add environment flag to deploy pipeline job - ([4d4b89d](https://github.com/equinor/radix-cli/commit/4d4b89dc3521ff08c2bf07ef2c9fc2b9dfaabedd)) by @nilsgstrabo

- Merge pull request #85 from equinor/add-missing-flag

add environment flag to deploy pipeline job - ([f1e55fb](https://github.com/equinor/radix-cli/commit/f1e55fb9e12bcb9c7d149541401670376d5c6944)) by @nilsgstrabo in [#85](https://github.com/equinor/radix-cli/pull/85)


## [1.14.0](https://github.com/equinor/radix-cli/compare/v1.13.3..v1.14.0) - 2024-01-17

### 💼 Other

- Add command to set tls certificate and private key (#83)

* add command to set tls certificate and private key

* add command for setting external dns tls secrets

* update module versions

* fix example

* add flag to get certificate and private key from file

* remove unneccessary else-if block - ([b2042d3](https://github.com/equinor/radix-cli/commit/b2042d34711b025a91db34022446b67c379c2197)) by @nilsgstrabo in [#83](https://github.com/equinor/radix-cli/pull/83)


## [1.13.3](https://github.com/equinor/radix-cli/compare/v1.13.2..v1.13.3) - 2024-01-02

### 💼 Other

- Upgrade dependencies (#81)

* upgrade dependencies and fix lint bugs
* cleanup dockerfile

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([81396e9](https://github.com/equinor/radix-cli/commit/81396e9ac3bf71d5efcb35d6dca75a59a077e7eb)) by @Richard87 in [#81](https://github.com/equinor/radix-cli/pull/81)


## [1.13.2](https://github.com/equinor/radix-cli/compare/v1.13.1..v1.13.2) - 2024-01-02

### 💼 Other

- Add set env var (#80)

* Added a command set env var

* Added a command set env var, updated radix-api client

* Updated radix-api client - ([7357b04](https://github.com/equinor/radix-cli/commit/7357b04f813b6f4079eaab96e18eaea8d33fca15)) by @satr in [#80](https://github.com/equinor/radix-cli/pull/80)


## [1.13.1](https://github.com/equinor/radix-cli/compare/v1.13.0..v1.13.1) - 2023-11-29

### 💼 Other

- Return error if pipeline job fails when using --follow (#78)

* return error when  pipeline job fails with --follow

* rename func

* change info message logging - ([04b58c4](https://github.com/equinor/radix-cli/commit/04b58c446a5bbc12af2cf1694632370b84d21a4e)) by @nilsgstrabo in [#78](https://github.com/equinor/radix-cli/pull/78)


## [1.13.0](https://github.com/equinor/radix-cli/compare/v1.12.1..v1.13.0) - 2023-11-20

### 💼 Other

- Set get logs since duration (#77)

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([b18133a](https://github.com/equinor/radix-cli/commit/b18133a1d7b127b3d1d7a5270db6c8adff69dadc)) by @Richard87 in [#77](https://github.com/equinor/radix-cli/pull/77)


## [1.12.1](https://github.com/equinor/radix-cli/compare/v1.12.0..v1.12.1) - 2023-11-10

### 💼 Other

- Allow, but hide token-env flag in validate cmd (#76)

* hide token-env flag
---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([f691481](https://github.com/equinor/radix-cli/commit/f69148194f81d7508d9c0f3c0dd679e5fa7c975f)) by @Richard87 in [#76](https://github.com/equinor/radix-cli/pull/76)


## [1.12.0](https://github.com/equinor/radix-cli/compare/v1.11.0..v1.12.0) - 2023-11-10

### 💼 Other

- Merge branch 'master' into update-readme - ([0292f9b](https://github.com/equinor/radix-cli/commit/0292f9bc500823b0ed378c706de134ad76237687)) by @satr

- Merge pull request #72 from equinor/update-readme

Update readme - ([b3e1931](https://github.com/equinor/radix-cli/commit/b3e1931d4663dba72bef66198f95c2ab86c40e6b)) by @emirgens in [#72](https://github.com/equinor/radix-cli/pull/72)

- Removed sample repo - ([d64867b](https://github.com/equinor/radix-cli/commit/d64867b2ffda7747b75f8c3c5b453f249763bda0)) by @emirgens

- Merge pull request #73 from equinor/cleanup-launch-vscode

Cleanup launch vscode - ([3837384](https://github.com/equinor/radix-cli/commit/38373848311d69cbd7c02ed1fcf14e8401e137b0)) by @emirgens in [#73](https://github.com/equinor/radix-cli/pull/73)

- Bump golang.org/x/net from 0.10.0 to 0.17.0

Bumps [golang.org/x/net](https://github.com/golang/net) from 0.10.0 to 0.17.0.
- [Commits](https://github.com/golang/net/compare/v0.10.0...v0.17.0)

---
updated-dependencies:
- dependency-name: golang.org/x/net
  dependency-type: indirect
...

Signed-off-by: dependabot[bot] <support@github.com> - ([3860d68](https://github.com/equinor/radix-cli/commit/3860d685c97e9b36d451d34c3f62bd9aabf861fd)) by @dependabot[bot]

- Merge pull request #64 from equinor/dependabot/go_modules/golang.org/x/net-0.17.0

Bump golang.org/x/net from 0.10.0 to 0.17.0 - ([55ddb01](https://github.com/equinor/radix-cli/commit/55ddb01be3783308235b767ddca604d0ef3247dc)) by @emirgens in [#64](https://github.com/equinor/radix-cli/pull/64)

- Create validate command (#75)

* Create validate command
* make sure all output goes to stdout, except printing yaml
* create a validate subcommand

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no>
Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com> - ([3ec779b](https://github.com/equinor/radix-cli/commit/3ec779b40bc94fe9437f6ca2da21f25b7d6e4968)) by @Richard87 in [#75](https://github.com/equinor/radix-cli/pull/75)


## New Contributors ❤️

* @dependabot[bot] made their first contribution
## [1.11.0](https://github.com/equinor/radix-cli/compare/v1.10.0..v1.11.0) - 2023-10-20

### 💼 Other

- Add rerun pipeline job command (#71)

* Added rerun command

* Added rerun pipeline-job command

* Cleanup - ([2f43b17](https://github.com/equinor/radix-cli/commit/2f43b17b737745eba3568588ab3ea3c0414119bc)) by @satr in [#71](https://github.com/equinor/radix-cli/pull/71)


## [1.10.0](https://github.com/equinor/radix-cli/compare/v1.9.3..v1.10.0) - 2023-10-18

### 💼 Other

- Remove unused embed package - ([179f372](https://github.com/equinor/radix-cli/commit/179f372b71348b5b3e5d7fafb47e73529f38f1b2)) by @nilsgstrabo

- Add ldflags to reduce size of compiled binary - ([6115b32](https://github.com/equinor/radix-cli/commit/6115b32c5521ab5e7f7bbad388008a2e6d2d4412)) by @nilsgstrabo

- Remove unused var - ([e187e63](https://github.com/equinor/radix-cli/commit/e187e630d77d228e6453e87b01f728d510c319e3)) by @nilsgstrabo

- Merge pull request #69 from equinor/goreleaser-ldflags

Misc changes - ([2e74c35](https://github.com/equinor/radix-cli/commit/2e74c356dc2357b3b3c56936e1b29d673a0d0aff)) by @nilsgstrabo in [#69](https://github.com/equinor/radix-cli/pull/69)

- Added option commitID for deploy pipeline job (#70)

* Added option commitID for deploy pipeline job

* Corrected commit-id option description - ([1cfb0a4](https://github.com/equinor/radix-cli/commit/1cfb0a4e197c75edff33f913cbc3cb2c10e809a0)) by @satr in [#70](https://github.com/equinor/radix-cli/pull/70)


## [1.9.3](https://github.com/equinor/radix-cli/compare/v1.9.2..v1.9.3) - 2023-10-12

### 💼 Other

- Fixed the error messages (#68) - ([50b83a2](https://github.com/equinor/radix-cli/commit/50b83a219d79b8ea3d91a21a3790c137aac15a07)) by @satr in [#68](https://github.com/equinor/radix-cli/pull/68)


## [1.9.2](https://github.com/equinor/radix-cli/compare/v1.9.1..v1.9.2) - 2023-10-12

### 💼 Other

- Upate release readme - ([d330a83](https://github.com/equinor/radix-cli/commit/d330a831cfddef1d8c71e14d8c86458352aba5a9)) by @Richard87

- Added contributing.md
Added security.md
Updated readme.md - ([73e1423](https://github.com/equinor/radix-cli/commit/73e142355bcb4bc6e9428a2bdd6a8946dfa37859)) by @emirgens

- Version from env (#67)

* read version from git tag
* upate release readme

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([4566420](https://github.com/equinor/radix-cli/commit/4566420ea3ecaba5096eaf001cf60638c037e873)) by @Richard87 in [#67](https://github.com/equinor/radix-cli/pull/67)


## New Contributors ❤️

* @emirgens made their first contribution
## [1.9.1](https://github.com/equinor/radix-cli/compare/v1.9.0..v1.9.1) - 2023-10-12

### 💼 Other

- Include dependencies (#63)

* add api client
* remove swagger-gen from release
* test goreleaser build and golang ci lint
* add CD, remove lint test

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([0d8bfe4](https://github.com/equinor/radix-cli/commit/0d8bfe41563c840ad6f55f7c46f26ab7e257dd90)) by @Richard87 in [#63](https://github.com/equinor/radix-cli/pull/63)

- Login to docker registry (#65)

* login to docker registry
* run fast golangci-lint

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([bfe9f22](https://github.com/equinor/radix-cli/commit/bfe9f22bb516b0ce3c72e90ef9db910c969079ce)) by @Richard87 in [#65](https://github.com/equinor/radix-cli/pull/65)

- Add scope (#66)

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([9df533a](https://github.com/equinor/radix-cli/commit/9df533a9fc32e11eb20aef620db0b796a86e1c67)) by @Richard87 in [#66](https://github.com/equinor/radix-cli/pull/66)


## [1.9.0](https://github.com/equinor/radix-cli/compare/v1.8.1..v1.9.0) - 2023-10-11

### 💼 Other

- Promote active deployment (#62)

* Promote active deployment
* print branch name without newline
* add new version 1.9.0

---------

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([7991aae](https://github.com/equinor/radix-cli/commit/7991aaef21ae9b12ddf4ffd5a1db0b91e9fcdd1f)) by @Richard87 in [#62](https://github.com/equinor/radix-cli/pull/62)


## [1.8.1](https://github.com/equinor/radix-cli/compare/v1.8.0..v1.8.1) - 2023-09-28

### 🐛 Bug Fixes

- Dockerfile.goreleaser to reduce vulnerabilities - ([6fbf2a0](https://github.com/equinor/radix-cli/commit/6fbf2a0a799acb83f194bda1628cf00811828b4c)) by @snyk-bot


### 💼 Other

- Update Dockerfile.goreleaser

Co-authored-by: Sergey Smolnikov <ssmol@equinor.com> - ([07ec4c0](https://github.com/equinor/radix-cli/commit/07ec4c0527ae04477b7e3c184d0cf840f3a26e2e)) by @nilsgstrabo

- Merge pull request #53 from equinor/snyk-fix-05aa567a3e60e868aa7f3d5d2ef69921

[Snyk] Security upgrade alpine from 3.17 to 3 - ([fbc9e57](https://github.com/equinor/radix-cli/commit/fbc9e57cc4dc67984d181e49afdcfcce31b346a8)) by @nilsgstrabo in [#53](https://github.com/equinor/radix-cli/pull/53)

- Upgrade Go (1.21) an ApiMachinery(0.27.6) (#57)

* Upgrade Go (1.21) an ApiMachinery(0.27.6) - ([3808a5e](https://github.com/equinor/radix-cli/commit/3808a5e842305cb6c799fbb3d0c9021e4ff3a5ce)) by @Richard87 in [#57](https://github.com/equinor/radix-cli/pull/57)

- Upgrade name_template (#58)

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([54c8d45](https://github.com/equinor/radix-cli/commit/54c8d4534a4779f8339b338ae708138f4ecff38b)) by @Richard87 in [#58](https://github.com/equinor/radix-cli/pull/58)

- Tweaking-name-templae (#59)

Co-authored-by: Richard Hagen <richard.hagen@bouvet.no> - ([da57ba0](https://github.com/equinor/radix-cli/commit/da57ba0b4e9d721fcb4697720826febcb9529291)) by @Richard87 in [#59](https://github.com/equinor/radix-cli/pull/59)


## New Contributors ❤️

* @Richard87 made their first contribution in [#59](https://github.com/equinor/radix-cli/pull/59)
## [1.8.0](https://github.com/equinor/radix-cli/compare/v1.7.1..v1.8.0) - 2023-08-21

### 💼 Other

- Suppress warning (#49)

* Suppressed client-go warnings

* Fixed creating of an app

* Removed extra logs - ([b8b17f1](https://github.com/equinor/radix-cli/commit/b8b17f107f98b7b57ca7e4505295f686b31950cd)) by @satr in [#49](https://github.com/equinor/radix-cli/pull/49)

- Add delay loop on app registration (#50)

* Added delay between registered app and returned deploy key

* Set version

* Fixed issue - ([258648c](https://github.com/equinor/radix-cli/commit/258648c9fd929cba6071784ef367a4ad2986dd1a)) by @satr in [#50](https://github.com/equinor/radix-cli/pull/50)

- Replace authentication with msal in Radix CLI (#55)

* Added get deployment command

* Added get deployments command

* Updated refs

* Return error always

* Added msal provider.

* Extended msal provider.

* Saved old radix config from msal struct

* Saved radix config without token, saved msal struct

* Rearranged code

* Saved contract to contract file

* Added logout

* Updated copyright

* Removed outdated file

* Removed outdated code, cleaned up

* Cleaned up

* Set version

* Added reader ad group option

* Added create app example

* Removed outdated code. Added troubleshooting for failure scenarios.

* Moved msal contract data structure to radix config file as a property

* cleanup code

* store msal cache as base64

* cleanup code

* misc fixes

* Extracted context related flags to an extra method

* Set context related flags only to context specific commands

* Removed not used constants

* Suppressed help output on error. Corrected get job log.

* Fixed preserved context

---------

Co-authored-by: Nils Gustav Stråbø <nst@equinor.com> - ([6a9ab49](https://github.com/equinor/radix-cli/commit/6a9ab49c4605253565f749449d9e6ecf4bba7b56)) by @satr in [#55](https://github.com/equinor/radix-cli/pull/55)

- Joined get deploy commands to two. Corrected command descriptions (#56)

* Joined get deploy commands to two. Corrected command descriptions

* Update cmd/logsJob.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/logsJob.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/logsEnvironmentComponent.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/logsEnvironmentComponent.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/logsEnvironment.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/logsEnvironment.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

---------

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com> - ([5bc27ad](https://github.com/equinor/radix-cli/commit/5bc27adb4003415049f469dbf661dc8db6020071)) by @satr in [#56](https://github.com/equinor/radix-cli/pull/56)


## [1.7.1](https://github.com/equinor/radix-cli/compare/v1.7.0..v1.7.1) - 2023-04-03

### 💼 Other

- Bump dependencies - ([429e1c7](https://github.com/equinor/radix-cli/commit/429e1c78cb1850c5cb1d15e139e09167b42e8c5c)) by @nilsgstrabo

- Merge pull request #46 from equinor/security/update-packages

bump dependencies - ([8b706aa](https://github.com/equinor/radix-cli/commit/8b706aa34eae5e1ff3d01d6e47a5ce35c375445c)) by @nilsgstrabo in [#46](https://github.com/equinor/radix-cli/pull/46)


## [1.7.0](https://github.com/equinor/radix-cli/compare/v1.4.0..v1.7.0) - 2023-03-27

### 💼 Other

- Add prev log (#44)

* Cleanup

* Added prev log

* Removed generation of machine user token - ([8da66eb](https://github.com/equinor/radix-cli/commit/8da66eb66a7021e6d8679eb9a43720f3f0bc5c9f)) by @satr in [#44](https://github.com/equinor/radix-cli/pull/44)

- Scale replica (#45)

* Added scale command and verbose output option

* Set version - ([082d463](https://github.com/equinor/radix-cli/commit/082d463bcf99b62b688aa432eeb3fcdfbc1ec7d2)) by @satr in [#45](https://github.com/equinor/radix-cli/pull/45)

- Add option image-tags (#43)

* Added imagetags option to deploy job

* Cleanup

* Renamed image-tag-names

* Updated ref

* Set version - ([0a754a9](https://github.com/equinor/radix-cli/commit/0a754a9d6fe4196d9cd45879d0351d88602b9998)) by @satr in [#43](https://github.com/equinor/radix-cli/pull/43)


## [1.4.0](https://github.com/equinor/radix-cli/compare/v1.3.0..v1.4.0) - 2023-01-12

### 💼 Other

- Updated readme - ([a786e25](https://github.com/equinor/radix-cli/commit/a786e25061726fbf5e61f5fa2e28fbfda3e83c71)) by @nilsgstrabo

- Merge pull request #41 from equinor/update-readme

updated readme - ([4131d49](https://github.com/equinor/radix-cli/commit/4131d497aeb793ef1c7c8e295216d85d27a6bdf3)) by @nilsgstrabo in [#41](https://github.com/equinor/radix-cli/pull/41)

- Added command to create a promote pipeline job. Changed the version c… (#42)

* Added command to create a promote pipeline job. Changed the version command

* Update cmd/promoteApplication.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/promoteApplication.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/promoteApplication.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update cmd/promoteApplication.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com>

* Update buildDeployApplication.go

* Create deployApplication.go

Co-authored-by: Nils Gustav Stråbø <65334626+nilsgstrabo@users.noreply.github.com> - ([3b532f3](https://github.com/equinor/radix-cli/commit/3b532f33631972003f17a53256c832fd3e1d0a3b)) by @satr in [#42](https://github.com/equinor/radix-cli/pull/42)


## [1.3.0](https://github.com/equinor/radix-cli/compare/v1.2.1..v1.3.0) - 2022-11-11

### 💼 Other

- Add configuration-item argument for create app - ([5004d8c](https://github.com/equinor/radix-cli/commit/5004d8cc2b742a90dbe663318d05edaac333162d)) by @nilsgstrabo

- Merge pull request #40 from equinor/46539-require-configurationItem-for-createApp

add configuration-item argument for create app - ([7e3734e](https://github.com/equinor/radix-cli/commit/7e3734e82a8f38bd59bf790bc4def1fa311bbc13)) by @nilsgstrabo in [#40](https://github.com/equinor/radix-cli/pull/40)


## [1.2.1](https://github.com/equinor/radix-cli/compare/v1.1.1..v1.2.1) - 2022-10-12

### 💼 Other

- Added docker file (#35)

* Fixed get job logs

* Added new dockerfile code

* Corrected dockerfile

* Corrected readme

* Added make command - ([b2d5c37](https://github.com/equinor/radix-cli/commit/b2d5c37b6a6e41f8fecdcde5f23d2d757828b0b1)) by @satr in [#35](https://github.com/equinor/radix-cli/pull/35)

- Added commands start, stop component, env, app (#36)

* Added commands start, stop component, env, app

* Corrected doc message - ([02af20d](https://github.com/equinor/radix-cli/commit/02af20d8f4e97b8a59a4d28ee49dcacf78e60ef0)) by @satr in [#36](https://github.com/equinor/radix-cli/pull/36)

- Fixed dockerfile (#37) - ([000a24d](https://github.com/equinor/radix-cli/commit/000a24d7b27b8a5b5d0ded970b11e983cf75b246)) by @satr in [#37](https://github.com/equinor/radix-cli/pull/37)

- Added acknowledge warnings flag and version command (#39) - ([23e4869](https://github.com/equinor/radix-cli/commit/23e48694940b219d375a4adb9222cd9758744922)) by @satr in [#39](https://github.com/equinor/radix-cli/pull/39)

- Fix dockerfile2 (#38)

* Fixed dockerfile

* Restored dockerfile

* Created goreleaser dockerfile - ([679973d](https://github.com/equinor/radix-cli/commit/679973d07793220f4458eb088e47be008cb3335d)) by @satr in [#38](https://github.com/equinor/radix-cli/pull/38)


## [1.1.1](https://github.com/equinor/radix-cli/compare/v1.1.0..v1.1.1) - 2022-09-14

### 💼 Other

- Fixed get job logs (#34) - ([be31f57](https://github.com/equinor/radix-cli/commit/be31f57ac8465aa0a541550d8b5150deea30cb3d)) by @satr in [#34](https://github.com/equinor/radix-cli/pull/34)


## [1.1.0](https://github.com/equinor/radix-cli/compare/v1.0.8..v1.1.0) - 2022-04-20

### 🐛 Bug Fixes

- Dockerfile to reduce vulnerabilities (#29) - ([a084368](https://github.com/equinor/radix-cli/commit/a08436836d424452524efd389b7ef2afe711ec1b)) by @snyk-bot in [#29](https://github.com/equinor/radix-cli/pull/29)

- Dockerfile to reduce vulnerabilities (#30) - ([b4acbb6](https://github.com/equinor/radix-cli/commit/b4acbb6fa19783e001a013e94faab38078cee376)) by @snyk-bot in [#30](https://github.com/equinor/radix-cli/pull/30)


### 💼 Other

- Added cluster C2 (#31) - ([7910497](https://github.com/equinor/radix-cli/commit/7910497f75cf44e02419a1300b289d3b53e70cd6)) by @satr in [#31](https://github.com/equinor/radix-cli/pull/31)


## New Contributors ❤️

* @snyk-bot made their first contribution in [#30](https://github.com/equinor/radix-cli/pull/30)
## [1.0.8](https://github.com/equinor/radix-cli/compare/v1.0.7..v1.0.8) - 2022-02-23

### 💼 Other

- Update go modules - ([f713f03](https://github.com/equinor/radix-cli/commit/f713f0358ac84189cb9982b2b5a145abcb3a517e)) by @nilsgstrabo

- Merge pull request #28 from equinor/185798-update-go-modules

update go modules - ([8bf9b32](https://github.com/equinor/radix-cli/commit/8bf9b32ea292f28209c0897070406b3293674b8a)) by @nilsgstrabo in [#28](https://github.com/equinor/radix-cli/pull/28)


## [1.0.7](https://github.com/equinor/radix-cli/compare/v1.0.6..v1.0.7) - 2021-10-14

### 💼 Other

- Added login/logout commands (#26)

* Added login/logout commands

* Added login/logout commands - ([5c2f016](https://github.com/equinor/radix-cli/commit/5c2f016e8ec4d3dbfd75acd3197333553a884d24)) by @satr in [#26](https://github.com/equinor/radix-cli/pull/26)

- Fix releaser (#27) - ([2a1382a](https://github.com/equinor/radix-cli/commit/2a1382abd750555f5799f62f7bdd425b780f00e7)) by @satr in [#27](https://github.com/equinor/radix-cli/pull/27)


## [1.0.6](https://github.com/equinor/radix-cli/compare/v1.0.5..v1.0.6) - 2021-10-04

### 🐛 Bug Fixes

- Fix login error - ([bdb1acd](https://github.com/equinor/radix-cli/commit/bdb1acd5f290d2c4aa83ac72047a04938e936728)) by @nilsgstrabo


### 💼 Other

- Fix command example in readme (#23) - ([0c47dc8](https://github.com/equinor/radix-cli/commit/0c47dc895d0b354380790afc429f42f6c89114e8)) by @satr in [#23](https://github.com/equinor/radix-cli/pull/23)

- Merge pull request #24 from equinor/set-configmode-for-audience

fix login error - ([1abee3f](https://github.com/equinor/radix-cli/commit/1abee3fd17e8c840a32e0e60ec745c9a93ea6d56)) by @nilsgstrabo in [#24](https://github.com/equinor/radix-cli/pull/24)


## [1.0.5](https://github.com/equinor/radix-cli/compare/v1.0.4..v1.0.5) - 2021-07-14

### 💼 Other

- Add error message when error occurs (#13)

Co-authored-by: INGE KNUDSEN <ingeknudsen@INGEs-MacBook-Pro.local> - ([8d01cbc](https://github.com/equinor/radix-cli/commit/8d01cbced78bdb943d3c5f8e87e18b4996a475b3)) by @ingeknudsen in [#13](https://github.com/equinor/radix-cli/pull/13)

- Updated dependencies (#22) - ([4840236](https://github.com/equinor/radix-cli/commit/48402368682d0463ed121fcecc3e45d170fbcb90)) by @satr in [#22](https://github.com/equinor/radix-cli/pull/22)


## [1.0.4](https://github.com/equinor/radix-cli/compare/v1.0.3..v1.0.4) - 2020-11-03

### 💼 Other

- Require configBranch on create application - ([ae9aa45](https://github.com/equinor/radix-cli/commit/ae9aa4507949ec0800e611518011553780414169)) by @nilsgstrabo

- Merge pull request #17 from equinor/48324_update_radixapi_reference

Require configBranch on create application - ([a74048f](https://github.com/equinor/radix-cli/commit/a74048fc03fb26aebfe4f31796474541274ecc7f)) by @nilsgstrabo in [#17](https://github.com/equinor/radix-cli/pull/17)


## New Contributors ❤️

* @nilsgstrabo made their first contribution in [#17](https://github.com/equinor/radix-cli/pull/17)
## [1.0.3](https://github.com/equinor/radix-cli/compare/v1.0.2..v1.0.3) - 2020-10-27

### 💼 Other

- Ra 1909 machine user token reset function (#14)

* RA-1909 Added command regenerate machine user token. Work in progress

* RA-1909 Changed command ntype

* RA-1909 Added launch entry

* RA-1909 Implemented set machine user token

* RA-1909 Printed machine user token

* Removed not needed reference from go.mod - ([b4b2e57](https://github.com/equinor/radix-cli/commit/b4b2e5775f2d1ea7dc91b11a63902b3a44aa973c)) by @satr in [#14](https://github.com/equinor/radix-cli/pull/14)

- Changed ignorefile - ([6f37a6e](https://github.com/equinor/radix-cli/commit/6f37a6e0f13a306d7c2b9d9211f8d5079e6c3b16)) by @satr

- Extended push release documentation (#15) - ([6aa40fa](https://github.com/equinor/radix-cli/commit/6aa40fabd8bb9b8bc1dfe0caafad1997d64c4a6d)) by @satr in [#15](https://github.com/equinor/radix-cli/pull/15)

- Removed extra text from output of machine-user token (#16) - ([2fd65a7](https://github.com/equinor/radix-cli/commit/2fd65a7b975b212d1481bce626e065c25839019c)) by @satr in [#16](https://github.com/equinor/radix-cli/pull/16)


## New Contributors ❤️

* @satr made their first contribution in [#16](https://github.com/equinor/radix-cli/pull/16)
## [1.0.2](https://github.com/equinor/radix-cli/compare/v1.0.1..v1.0.2) - 2020-06-18

### 💼 Other

- Added triggeredBy argument to deploy command - ([6cc05f7](https://github.com/equinor/radix-cli/commit/6cc05f7ad221c21ba29c066408efaaca293d4292))

- Merge pull request #12 from equinor/RA-1478

Added triggeredBy argument to deploy command - ([2a358ea](https://github.com/equinor/radix-cli/commit/2a358ea323cefa79da7bfa2b8773e83a1712d808)) in [#12](https://github.com/equinor/radix-cli/pull/12)


## [1.0.1](https://github.com/equinor/radix-cli/compare/v1.0.0..v1.0.1) - 2020-05-12

### 💼 Other

- Update README.md - ([84c15d9](https://github.com/equinor/radix-cli/commit/84c15d9994c719640bb304c0aa9120bd3062c8c8)) by @ingeknudsen

- RA-1271 Create/Delete Environment (#9)

* createEnv

* copypaste correction

* add parameters

* add delete cmd

* correct env param read - ([07481e5](https://github.com/equinor/radix-cli/commit/07481e5d13532f0c6ffe1167cb91879d481528be)) by @JoakimHagen in [#9](https://github.com/equinor/radix-cli/pull/9)

- Restart component command (#11)

* restart component command

* update description - ([cad5491](https://github.com/equinor/radix-cli/commit/cad54915e9df48685834673d15fedd5a6038d834)) by @JoakimHagen in [#11](https://github.com/equinor/radix-cli/pull/11)


## [1.0.0](https://github.com/equinor/radix-cli/compare/v0.0.17..v1.0.0) - 2020-03-19

### 💼 Other

- Update README.md - ([977ed05](https://github.com/equinor/radix-cli/commit/977ed053609e39d99fe05d4b6ba8e706c6781567)) by @ingeknudsen

- Update README.md - ([17927eb](https://github.com/equinor/radix-cli/commit/17927ebd055a7fe268f792fcca030f2d88660e63)) by @ingeknudsen

- Ra 1313 prepare for release (#6)

* Refactored commands

* Enabled commands

* Corrected launch

* Improve suboptimal logging (#7)

* Improve suboptimal logging

* Code review

Co-authored-by: INGE KNUDSEN <ingeknudsen@INGEs-MacBook-Pro.local>

Co-authored-by: INGE KNUDSEN <ingeknudsen@INGEs-MacBook-Pro.local> - ([ce8ef77](https://github.com/equinor/radix-cli/commit/ce8ef7720ea946bd8a6b826d0dae1079fa7caed6)) by @ingeknudsen in [#6](https://github.com/equinor/radix-cli/pull/6)

- Took daft experiences into account (#8)

Co-authored-by: INGE KNUDSEN <ingeknudsen@INGEs-MacBook-Pro.local> - ([78a9536](https://github.com/equinor/radix-cli/commit/78a95362ea150edff1a24e2c84a4f33865b76265)) by @ingeknudsen in [#8](https://github.com/equinor/radix-cli/pull/8)


## [0.0.17](https://github.com/equinor/radix-cli/compare/v0.0.16..v0.0.17) - 2020-02-06

### 💼 Other

- Updated readme - ([83102e7](https://github.com/equinor/radix-cli/commit/83102e7a6fd20ada8ebf6026cead03a1e960186e))

- Create a proof of concept for creating a
application using cli - ([57481c4](https://github.com/equinor/radix-cli/commit/57481c4ac64dc4f578bda0347c6f7bb91eec6acb))

- Create application - ([b1819ca](https://github.com/equinor/radix-cli/commit/b1819cacb3de9cbcbfa109489a3feb0234b44bc0))

- Added example script - ([3a2bcb6](https://github.com/equinor/radix-cli/commit/3a2bcb697d34e898660f713c9d7ffd8f4bf17ce0))

- Update README.md - ([9fcad4c](https://github.com/equinor/radix-cli/commit/9fcad4cd0f5b262db16e113d27cdd484c0cf4bd6)) by @ingeknudsen

- Merge branch 'master' of github.com:equinor/radix-cli - ([0863580](https://github.com/equinor/radix-cli/commit/0863580eba3eb61c43f01ae8a06cc61dfb1e8257))

- Break if - if new timeout was set - ([1e5f642](https://github.com/equinor/radix-cli/commit/1e5f642c71389d38e1ffb942ced7454dbd5f5c22))

- Changed logic to log ok - ([f636de6](https://github.com/equinor/radix-cli/commit/f636de6dcdb3562930f435aaadc673b939334f2d))

- Prepare for demo - ([6bf9e61](https://github.com/equinor/radix-cli/commit/6bf9e61982f255f3f7cc5ea4e5e764e927ee0120))

- Extracted into utility function - ([e06ce45](https://github.com/equinor/radix-cli/commit/e06ce45cc15bbeb628ebcdc4df306e155ac67884))

- Added install guide for mac users - ([e4048e7](https://github.com/equinor/radix-cli/commit/e4048e7a6a8acafed567d165f52b95b79d727825))

- Link to releases - ([d01eb21](https://github.com/equinor/radix-cli/commit/d01eb21eb25b3d5493818278a691b4ecf43bf18b))

- Fix links - ([38c9abe](https://github.com/equinor/radix-cli/commit/38c9abed03ef8218815b4ace38e51c9217c7a95b))

- Feature flags + Documentation updates - ([e825494](https://github.com/equinor/radix-cli/commit/e8254943e501fb2d0f6f4f8d78c7864c84b8ed1d)) by @JoakimHagen in [#2](https://github.com/equinor/radix-cli/pull/2)

- Deploy-only command (#3) - ([01af2a4](https://github.com/equinor/radix-cli/commit/01af2a421b7aedf5f385297fb847b12efa194245)) by @JoakimHagen in [#3](https://github.com/equinor/radix-cli/pull/3)

- Readme Cleanup (#4) - ([e11cf83](https://github.com/equinor/radix-cli/commit/e11cf83f3ffa4e84bb1054fe555a014bfd9c7765)) by @JoakimHagen in [#4](https://github.com/equinor/radix-cli/pull/4)

- Swagger in Makefile (#5)

* Swagger in Makefile

* local does not work outside of functions - ([6e88617](https://github.com/equinor/radix-cli/commit/6e8861713abe4aed337be99b2a1f11db3f3fb96e)) by @JoakimHagen in [#5](https://github.com/equinor/radix-cli/pull/5)


## New Contributors ❤️

* @JoakimHagen made their first contribution in [#5](https://github.com/equinor/radix-cli/pull/5)
## [0.0.16](https://github.com/equinor/radix-cli/compare/v0.0.15..v0.0.16) - 2020-01-04

### 💼 Other

- Added check for reconciliation - ([01fb993](https://github.com/equinor/radix-cli/commit/01fb99342be6c4b4bf06385193785c1d6f01a148))

- Tested reconciliation - ([c946ddf](https://github.com/equinor/radix-cli/commit/c946ddf361d6e239fb36660edf9c2c40b0ec6de6))


## [0.0.15](https://github.com/equinor/radix-cli/compare/v0.0.14..v0.0.15) - 2020-01-03

### 💼 Other

- Added command to get environment from branch - ([4494114](https://github.com/equinor/radix-cli/commit/449411414a51527612d1c957b50c8ea54cfd9ce1))


## [0.0.14](https://github.com/equinor/radix-cli/compare/v0.0.13..v0.0.14) - 2020-01-03

### 💼 Other

- Follow logs (#1)

* Follow logs

* Added following of job

* Added proper follow job - ([e0e499a](https://github.com/equinor/radix-cli/commit/e0e499ad5254c742cc4c3c8b7634b77d0fe2e722)) by @ingeknudsen in [#1](https://github.com/equinor/radix-cli/pull/1)


## [0.0.13](https://github.com/equinor/radix-cli/compare/v0.0.12..v0.0.13) - 2019-12-23

### 💼 Other

- Try to add method to set secret - ([8a3b4ec](https://github.com/equinor/radix-cli/commit/8a3b4ec4cbe61c0e46503337fa5148622e51c546))


## [0.0.12](https://github.com/equinor/radix-cli/compare/v0.0.11..v0.0.12) - 2019-12-23

### 💼 Other

- Added documentation - ([068bdf5](https://github.com/equinor/radix-cli/commit/068bdf59f9846a876f22ec1df4fda1c95d6b064d))

- Changed to be a heading - ([3a6b034](https://github.com/equinor/radix-cli/commit/3a6b034ac74f24a1ee93cd8082e890d0689e56d4))

- Should be build-deploy - ([3f6ba1c](https://github.com/equinor/radix-cli/commit/3f6ba1cb2e30137e7ec596a07703dc1fc4beec5b))

- Should be build-deploy - ([7302020](https://github.com/equinor/radix-cli/commit/73020207ef413b5c9a5b7a05558e74e0df5d33db))


## [0.0.11](https://github.com/equinor/radix-cli/compare/v0.0.10..v0.0.11) - 2019-12-23

### 💼 Other

- Use token from environment variable - ([71bc6b9](https://github.com/equinor/radix-cli/commit/71bc6b970dba0eea8ba1ce4d8dc6f3056e2b016f))

- Create LICENCE - ([f1b3e87](https://github.com/equinor/radix-cli/commit/f1b3e8771ff0d9f923874bd686a574e9a9ccb466)) by @ingeknudsen

- Merge branch 'master' of github.com:equinor/radix-cli - ([afd084f](https://github.com/equinor/radix-cli/commit/afd084fe5b40f9e15d160287e9465c406c32e0f7))


## New Contributors ❤️

* @ingeknudsen made their first contribution
## [0.0.10](https://github.com/equinor/radix-cli/compare/v0.0.9..v0.0.10) - 2019-12-22

### 💼 Other

- Prepare release - ([97565f8](https://github.com/equinor/radix-cli/commit/97565f8c03a55501ce9ad2501cbdc9bbf57f5b14))


## [0.0.9](https://github.com/equinor/radix-cli/compare/v0.0.8..v0.0.9) - 2019-12-20

### 💼 Other

- Wrong folder - ([c98eb54](https://github.com/equinor/radix-cli/commit/c98eb5427398a1c90b44a96c92ce44dc5066d33b))


## [0.0.8](https://github.com/equinor/radix-cli/compare/v0.0.7..v0.0.8) - 2019-12-20

### 💼 Other

- Revert back to original - ([238d2b4](https://github.com/equinor/radix-cli/commit/238d2b4ce5c20ce863eca762375653e34d70c5a8))


## [0.0.7] - 2019-12-20

### 💼 Other

- The first commit - ([33c624d](https://github.com/equinor/radix-cli/commit/33c624dde9b5b1662ccbd56a9e1240f100b6a3b0))

- Added radix config - ([d13718a](https://github.com/equinor/radix-cli/commit/d13718a1d7fc485f6ba0e84c516c9eeeb373d031))

- Attempt to release only docker image - ([8343362](https://github.com/equinor/radix-cli/commit/8343362898b59e545b987eed3ec4982241d96247))


## New Contributors ❤️

* @ made their first contribution
<!-- generated by git-cliff -->
