# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/query-meta/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.4.5] - 2023-11-17

### Changed

#### Dependency Updates

- (GH-288) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-276) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-293) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds
- (GH-278) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-281) go.mod: bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20
- (GH-264) go.mod: bump github.com/rs/zerolog from 1.30.0 to 1.31.0
- (GH-266) go.mod: bump golang.org/x/crypto from 0.13.0 to 0.14.0
- (GH-290) go.mod: bump golang.org/x/crypto from 0.14.0 to 0.15.0
- (GH-267) go.mod: bump golang.org/x/sys from 0.12.0 to 0.13.0
- (GH-287) go.mod: bump golang.org/x/sys from 0.13.0 to 0.14.0
- (GH-286) go.mod: bump golang.org/x/text from 0.13.0 to 0.14.0

## [v0.4.4] - 2023-10-06

### Changed

#### Dependency Updates

- (GH-250) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-234) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-237) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-238) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-251) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-258) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-247) ghaw: bump actions/checkout from 3 to 4
- (GH-241) go.mod: bump github.com/pelletier/go-toml/v2 from 2.0.9 to 2.1.0
- (GH-248) go.mod: bump golang.org/x/crypto from 0.12.0 to 0.13.0
- (GH-243) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0
- (GH-242) go.mod: bump golang.org/x/text from 0.12.0 to 0.13.0

## [v0.4.3] - 2023-08-18

### Added

- (GH-201) Add initial automated release notes config
- (GH-203) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.4`
  - `rs/zerolog`
    - `v1.29.1` to `v1.30.0`
  - `golang.org/x/crypto`
    - `v0.11.0` to `v0.12.0`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/text`
    - `v0.11.0` to `v0.12.0`
- (GH-205) Update Dependabot config to monitor both branches
- (GH-229) Update project to Go 1.20 series

## [v0.4.2] - 2023-07-14

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.6` to `go-ci-oldstable-build-v0.11.4`
  - `pelletier/go-toml`
    - `v2.0.8` to `v2.0.9`
  - `golang.org/x/crypto`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/sys`
    - `v0.9.0` to `v0.10.0`
  - `golang.org/x/text`
    - `v0.10.0` to `v0.11.0`

### Fixed

- (GH-196) Remove deploy logic from postinstall scripts

## [v0.4.1] - 2023-06-14

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.10`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.4` to `go-ci-oldstable-build-v0.10.6`
  - `pelletier/go-toml`
    - `v2.0.7` to `v2.0.8`
  - `rs/zerolog`
    - `v1.29.0` to `v1.29.1`
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
  - `golang.org/x/crypto`
    - `v0.8.0` to `v0.10.0`
  - `golang.org/x/sys`
    - `v0.7.0` to `v0.9.0`
  - `golang.org/x/text`
    - `v0.9.0` to `v0.10.0`
- (GH-183) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-172) Remove duplicate CHANGELOG dependency listing
- (GH-180) Disable depguard linter
- (GH-187) Restore local CodeQL workflow
- (GH-188) Fix various CHANGELOG dependency entries

## [v0.4.0] - 2023-04-11

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-151) Generate RPM/DEB packages using nFPM
- (GH-152) Add version details to Windows executables

### Changed

- (GH-150) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-153) Makefile: Compress binaries & use fixed filenames
- (GH-154) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-155) Build dev/stable releases using go-ci Docker image

## [v0.3.13] - 2023-04-11

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-139) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
  - `pelletier/go-toml`
    - `v2.0.6` to `v2.0.7`
  - `rs/zerolog`
    - `v1.28.0` to `v1.29.0`
  - `mattn/go-isatty`
    - `v0.0.16` to `v0.0.18`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.7.0`
  - `golang.org/x/text`
    - `v0.5.0` to `v0.9.0`
  - `golang.org/x/crypto`
    - `v0.3.0` to `v0.8.0`
- CI
  - (GH-145) Drop `Push Validation` workflow
  - (GH-146) Rework workflow scheduling
  - (GH-148) Remove `Push Validation` workflow status badge

### Fixed

- (GH-160) Update vuln analysis GHAW to use on.push hook
- (GH-166) Fix 'if-return' revive linting error
- (GH-167) Fix year in 2022 CHANGELOG release entries

## [v0.3.12] - 2022-12-08

### Overview

- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `denisenkom/go-mssqldb`
    - `0.12.3` to `0.12.4`
  - `pelletier/go-toml`
    - `v2.0.5` to `v2.0.6`
  - `golang.org/x/text`
    - `v0.3.7` to `v0.5.0`
  - `mattn/go-colorable`
    - `v0.1.12` to `v0.1.13`
  - `mattn/go-isatty`
    - `v0.0.14` to `v0.0.16`
  - `golang.org/x/sys`
    - `v0.0.0-20220722155257-8c9f86f7a55f` to `v0.3.0`
  - `golang.org/x/crypto`
    - `v0.0.0-20220622213112-05595931fe9d` to `v0.3.0`
  - `alexflint/go-scalar`
    - `v1.1.0` to `v1.2.0`
  - `golang-sql/civil`
    - `v0.0.0-20190719163853-cb61b32ac6fe` to
      `v0.0.0-20220223132316-b832511892a9`
- (GH-122) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-130) Fix Makefile Go module base path detection

## [v0.3.11] - 2022-09-22

### Overview

- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
  - `github/codeql-action`
    - `v2.1.22` to `v2.1.25`
- (GH-111) Update project to Go 1.19
- (GH-112) Update Makefile and GitHub Actions Workflows
- (GH-113) Add CodeQL GitHub Action Workflow
- (GH-114) Follow-up CI and Makefile tweaks
- (GH-115) GHAW: Emit markdownlint version before running

## [v0.3.10] - 2022-07-29

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.17.13`
  - `pelletier/go-toml`
    - `v2.0.2` to `v2.0.5`
  - `rs/zerolog`
    - `v1.27.0` to `v1.28.0`

### Fixed

- (GH-108) Apply linting fixes for Go 1.19 release
- (GH-109) Add missing cmd doc file

## [v0.3.9] - 2022-07-21

### Overview

- Dependency updates
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.10` to `1.17.12`
  - `denisenkom/go-mssqldb`
    - `v0.12.0` to `v0.12.2`
  - `pelletier/go-toml`
    - `v2.0.1` to `v2.0.2`
  - `rs/zerolog`
    - `v1.26.1` to `v1.27.0`

### Fixed

- (GH-100) Update lintinstall Makefile recipe

## [v0.3.8] - 2022-05-13

### Overview

- Dependency updates
- built using Go 1.17.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.10`
  - `pelletier/go-toml`
    - `v2.0.0` to `v2.0.1`

## [v0.3.7] - 2022-04-28

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`
  - `pelletier/go-toml`
    - `v1.9.4` to `v2.0.0`

## [v0.3.6] - 2022-03-02

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `alexflint/go-arg`
    - `v1.4.2` to `v1.4.3`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-76) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-77) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-79) var-declaration: should omit type string from declaration (revive)

## [v0.3.5] - 2022-01-24

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-71) Update go.mod file, canary Dockerfile to reflect current
      dependencies
  - `denisenkom/go-mssqldb`
    - `v0.11.0` to `v0.12.0`
  - `actions/setup-node`
    - `v2.5.0` to `v2.5.1`

## [v0.3.4] - 2021-12-28

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `rs/zerolog`
    - `v1.26.0` to `v1.26.1`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.0`

## [v0.3.3] - 2021-11-09

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `rs/zerolog`
    - `v1.25.0` to `v1.26.0`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

### Fixed

- (GH-61) False positive `G307: Deferring unsafe method "Close" on type
  "*os.File" (gosec)` linting error

## [v0.3.2] - 2021-09-25

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
  - `denisenkom/go-mssqldb`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/text`
    - `v0.3.6` to `v0.3.7`
  - `pelletier/go-toml`
    - `v1.9.3` to `v1.9.4`
  - `rs/zerolog`
    - `v1.23.0` to `v1.25.0`

## [v0.3.1] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.3.0` to `v2.4.0`

## [v0.3.0] - 2021-07-22

### Overview

- Add new database fields
- Dependency update
- built using Go 1.16.6

### Added

- Add record fields
  - `NearFieldIphoneID`
  - `NearFieldIphoneExp`
  - `NearFieldIphoneLost`
  - `NearFieldIwatchID`
  - `NearFieldIwatchExp`
  - `NearFieldIwatchLost`
  - `NearFieldAndroidID`
  - `NearFieldAndroidExp`
  - `NearFieldAndroidLost`

### Changed

- dependencies
  - `action/setup-node`
    - `v2.2.0` to `v2.3.0`

## [v0.2.1] - 2021-07-16

### Overview

- Dependency updates
- built using Go 1.16.6

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- dependencies
  - built using `Go 1.16.6`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `pelletier/go-toml`
    - `v1.9.2` to `v1.9.3`
  - `rs/zerolog`
    - `v1.22.0` to `v1.23.0`
  - `actions/setup-node`
    - updated from `v2.1.5` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.2.0] - 2021-06-10

### Overview

- Add tests (flag parsing, config file loading)
- Add new database field
- Bug fixes
- Dependency updates
- built using Go 1.16.5

### Added

- Add tests
  - `--config-file` flag parsing
  - config file parsing
- Add `NearFieldBadgeID` record field

### Changed

- dependencies
  - built using `Go 1.16.5`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `alexflint/go-arg`
    - `v1.3.0` to `v1.4.2`
  - `pelletier/go-toml`
    - `v1.9.0` to `v1.9.2`
  - `rs/zerolog`
    - `v1.21.0` to `v1.22.0`

### Fixed

- config file-specific settings exposed as flags

## [v0.1.1] - 2021-04-15

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.16.3

### Changed

- dependencies
  - built using `Go 1.16.3`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `denisenkom/go-mssqldb`
    - `v0.0.0-20201104001602-427686ac8ec1` to `v0.10.0`
  - `golang.org/x/text`
    - `v0.3.5` to `v0.3.6`
  - `pelletier/go-toml`
    - `v1.8.1` to `v1.9.0`
  - `rs/zerolog`
    - `v1.20.0` to `v1.21.0`
  - `actions/setup-node`
    - `v2.1.4` to `v2.1.5`

### Fixed

- `denisenkom/go-mssqldb` library is incompatible with Go 1.16
- Fix environment variable names
- Minor README issues
- PURPOSE statement in `doc.go` file

## [v0.1.0] - 2021-03-24

### Added

Initial release!

This release provides an early version of a prototype tool intended to replace
the existing `query_meta.php` script used to retrieve patron records from Meta
and write out a pipe-delimited CSV file for further processing.

Due to known issues with the `denisenkom/go-mssqldb` package, Go 1.16 is not
supported at this time. Go 1.15 should be used instead until upstream
[GH-639](https://github.com/denisenkom/go-mssqldb/issues/639) is resolved.

[Unreleased]: https://github.com/atc0005/query-meta/compare/v0.4.5...HEAD
[v0.4.5]: https://github.com/atc0005/query-meta/releases/tag/v0.4.5
[v0.4.4]: https://github.com/atc0005/query-meta/releases/tag/v0.4.4
[v0.4.3]: https://github.com/atc0005/query-meta/releases/tag/v0.4.3
[v0.4.2]: https://github.com/atc0005/query-meta/releases/tag/v0.4.2
[v0.4.1]: https://github.com/atc0005/query-meta/releases/tag/v0.4.1
[v0.4.0]: https://github.com/atc0005/query-meta/releases/tag/v0.4.0
[v0.3.13]: https://github.com/atc0005/query-meta/releases/tag/v0.3.13
[v0.3.12]: https://github.com/atc0005/query-meta/releases/tag/v0.3.12
[v0.3.11]: https://github.com/atc0005/query-meta/releases/tag/v0.3.11
[v0.3.10]: https://github.com/atc0005/query-meta/releases/tag/v0.3.10
[v0.3.9]: https://github.com/atc0005/query-meta/releases/tag/v0.3.9
[v0.3.8]: https://github.com/atc0005/query-meta/releases/tag/v0.3.8
[v0.3.7]: https://github.com/atc0005/query-meta/releases/tag/v0.3.7
[v0.3.6]: https://github.com/atc0005/query-meta/releases/tag/v0.3.6
[v0.3.5]: https://github.com/atc0005/query-meta/releases/tag/v0.3.5
[v0.3.4]: https://github.com/atc0005/query-meta/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/query-meta/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/query-meta/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/query-meta/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/query-meta/releases/tag/v0.3.0
[v0.2.1]: https://github.com/atc0005/query-meta/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/query-meta/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/query-meta/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/query-meta/releases/tag/v0.1.0
