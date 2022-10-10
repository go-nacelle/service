# Changelog

## [Unreleased]

## [v2.0.1] - 2022-10-10

### Added

- Added `WithContainer` and `FromContext`. [#11](https://github.com/go-nacelle/service/pull/11)

## [v2.0.0] - 2021-05-31

### Added

- Added `InjectableServiceKey`. [#4](https://github.com/go-nacelle/service/pull/4)
- Exported top-level `Inject` method. [#9](https://github.com/go-nacelle/service/pull/9)
- Added `WithValues` to `Container`. [#9](https://github.com/go-nacelle/service/pull/9)

### Changed

- The `Inject` function and `PostInject` interface now receives a context parameter. [#10](https://github.com/go-nacelle/service/pull/10)
- Change type of service keys from `string` to `interface{}`. [#4](https://github.com/go-nacelle/service/pull/4)
- Replaced the `ServiceContainer` interface with `Container`, a struct with the same name and set of methods. [#7](https://github.com/go-nacelle/service/pull/7)
- Renamed `NewServiceContainer` to `New`. [#7](https://github.com/go-nacelle/service/pull/7)
- Removed `Inject` method from `Container`. [#9](https://github.com/go-nacelle/service/pull/9)

### Removed

- Removed `MustGet` and `MustSet` methods. [#3](https://github.com/go-nacelle/service/pull/3)
- Removed mocks package. [#6](https://github.com/go-nacelle/service/pull/6)
- Removed `Overlay`. [#9](https://github.com/go-nacelle/service/pull/9)

## [v1.0.2] - 2020-09-30

### Removed

- Removed dependency on [aphistic/sweet](https://github.com/aphistic/sweet) by rewriting tests to use [testify](https://github.com/stretchr/testify). [#2](https://github.com/go-nacelle/service/pull/2)

## [v1.0.1] - 2020-09-07

### Added

- Added overlay container. [#1](https://github.com/go-nacelle/service/pull/1)

## [v1.0.0] - 2019-06-17

### Changed

- Migrated from [efritz/bussard](https://github.com/efritz/bussard).

[Unreleased]: https://github.com/go-nacelle/service/compare/v2.0.0...HEAD
[v1.0.0]: https://github.com/go-nacelle/service/releases/tag/v1.0.0
[v1.0.1]: https://github.com/go-nacelle/service/compare/v1.0.0...v1.0.1
[v1.0.2]: https://github.com/go-nacelle/service/compare/v1.0.1...v1.0.2
[v2.0.0]: https://github.com/go-nacelle/service/compare/v1.0.2...v2.0.0
