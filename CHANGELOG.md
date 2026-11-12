# Changelog

All notable changes to this project will be documented in this file.

## [1.12.2](https://github.com/microsoft/kiota-abstractions-go/compare/v1.12.1...v1.12.2) (2026-11-12)


### Bug Fixes

* **abstractions, store:** ensure backing store enablement idempotency and support int scalar serialization ([d0569a4](https://github.com/microsoft/kiota-abstractions-go/commit/d0569a4cc7a28bbbbf06a0543f73f022d885440c))
* **abstractions:** normalize direct numeric scalar parameters ([8a10b6c](https://github.com/microsoft/kiota-abstractions-go/commit/8a10b6cc93b9ef51a37a93d15e5d8b4c50d13952))
* **abstractions:** prevent double-wrapping in EnableBackingStoreForParseNodeFactory ([5a5341e](https://github.com/microsoft/kiota-abstractions-go/commit/5a5341e5b13149c46177439d68d6b534c93c3585))
* **abstractions:** prevent double-wrapping in EnableBackingStoreForSerializationWriterFactory ([d2d17b8](https://github.com/microsoft/kiota-abstractions-go/commit/d2d17b8fe6545a42c402342f88e256fbe6e14e64))
* **abstractions:** support direct scalar fields in AddQueryParameters ([43b8b50](https://github.com/microsoft/kiota-abstractions-go/commit/43b8b508bf00fd97f6280952a357f8cca6c38b8c))
* **abstractions:** support direct string, bool, and int32 query parameters ([bf5a589](https://github.com/microsoft/kiota-abstractions-go/commit/bf5a58902aedc1615b949a819c0e1c967e9a7375))
* **abstractions:** support int and int pointer in RequestInformation.SetContentFromScalar ([b3ad9cc](https://github.com/microsoft/kiota-abstractions-go/commit/b3ad9cc308d727e4fff12ac89753e790225fc357))
* **abstractions:** support int slices in RequestInformation.SetContentFromScalarCollection ([a729863](https://github.com/microsoft/kiota-abstractions-go/commit/a729863dc345ab20ef68dc5f599aa3d8fa1b3925))
* **serialization:** account for milliseconds in duration.hasTimePart ([19791c9](https://github.com/microsoft/kiota-abstractions-go/commit/19791c9da46a5f3f9e718ee394ba3375850a1270))
* **serialization:** anchor ISO duration regexes, support integer seconds, and format zero duration ([1193d38](https://github.com/microsoft/kiota-abstractions-go/commit/1193d38e62d0951c3debf6161c28c9d4de11a3f5))
* **serialization:** emit fractional seconds when milliseconds are present in duration string ([199428b](https://github.com/microsoft/kiota-abstractions-go/commit/199428b46a6d7b3eca767ba6a3d0b03755701a9a))
* **serialization:** handle integer seconds in ISO durations and reject component-free strings ([3d1c59e](https://github.com/microsoft/kiota-abstractions-go/commit/3d1c59e9c38da97b5ced693967060fa5236324f1))
* **serialization:** preserve millisecond precision in ISODuration.ToDuration ([c2ef044](https://github.com/microsoft/kiota-abstractions-go/commit/c2ef044011cdccc5b0ef1c6b3855c88ad4e1e2b5))
* **serialization:** prevent invalid week normalization and preserve millisecond precision in durations ([6102fb1](https://github.com/microsoft/kiota-abstractions-go/commit/6102fb171e9e0cb70f8109661b854bc1b7959c1c))
* **serialization:** prevent invalid week normalization on durations with time components ([99c14f4](https://github.com/microsoft/kiota-abstractions-go/commit/99c14f40e5f7e93afe331d1de3e2bef598287728))
* **store:** make NewBackingStoreParseNodeFactory idempotent ([9605a71](https://github.com/microsoft/kiota-abstractions-go/commit/9605a714c903bd4950e9f3bc26348e15d6102101))
* **store:** make NewBackingStoreSerializationWriterProxyFactory idempotent ([6b03eee](https://github.com/microsoft/kiota-abstractions-go/commit/6b03eee7608efdbfaac952fcde7beeccf2e7cd20))

## [1.12.1](https://github.com/microsoft/kiota-abstractions-go/compare/v1.12.0...v1.12.1) (2026-11-05)


### Bug Fixes

* **abstractions:** add pointer support and nil safety to query parameters and add RemoveRequestOptions ([#308](https://github.com/microsoft/kiota-abstractions-go/issues/308)) ([a69d22c](https://github.com/microsoft/kiota-abstractions-go/commit/a69d22ced3cb5a2f3a30343f90eb521888dd4b83))
* **abstractions:** clear PathParametersAny in SetUri, support raw_url_key in PathParametersAny, and harden typed nil guards ([#311](https://github.com/microsoft/kiota-abstractions-go/issues/311)) ([5b9790e](https://github.com/microsoft/kiota-abstractions-go/commit/5b9790ebaffec774a21ebc4e45a11b97dadd43a9))
* **headers:** validate addition in RequestHeaders.TryAdd and delegate AddAll ([#303](https://github.com/microsoft/kiota-abstractions-go/issues/303)) ([840be4e](https://github.com/microsoft/kiota-abstractions-go/commit/840be4ef2049766d97c07fab12b1b6b588e6cbc6))
* **serialization, store:** invoke onAfterAction in ParseNodeProxyFactory and add nil safety guards ([#309](https://github.com/microsoft/kiota-abstractions-go/issues/309)) ([0871a32](https://github.com/microsoft/kiota-abstractions-go/commit/0871a32918abe9e4a595c2d1ea66dc6fed9668db))
* **serialization:** add safe type assertions and nil guards to untyped node models ([#305](https://github.com/microsoft/kiota-abstractions-go/issues/305)) ([041c28b](https://github.com/microsoft/kiota-abstractions-go/commit/041c28b1898a677767853d327f2f7834066aaf1f))
* **serialization:** use pointer receiver on ISODuration setters and normalize 12 months ([#306](https://github.com/microsoft/kiota-abstractions-go/issues/306)) ([d44fa7d](https://github.com/microsoft/kiota-abstractions-go/commit/d44fa7d86df87813e5fdef5ed15c365176f80842))

## [1.12.0](https://github.com/microsoft/kiota-abstractions-go/compare/v1.11.4...v1.12.0) (2026-10-28)


### Features

* **headers:** add TryAdd to ResponseHeaders and add comprehensive tests for response headers and errors ([#301](https://github.com/microsoft/kiota-abstractions-go/issues/301)) ([6f450ec](https://github.com/microsoft/kiota-abstractions-go/commit/6f450ece9773ad7a4d18418880d23c58d51dee01))


### Bug Fixes

* **abstractions:** add nil guards to multipartBody and support case-insensitive media types in registries ([#299](https://github.com/microsoft/kiota-abstractions-go/issues/299)) ([e27abf7](https://github.com/microsoft/kiota-abstractions-go/commit/e27abf77c6f54cdec9a2e1db4cf6452dfe4814f2))
* **abstractions:** add nil guards to RequestInformation and support direct values in SetContentFromScalar ([#298](https://github.com/microsoft/kiota-abstractions-go/issues/298)) ([eed598c](https://github.com/microsoft/kiota-abstractions-go/commit/eed598cb6e69c6c144e825ef25988de1506a8ba0))
* **store, serialization:** add nil guards to InMemoryBackingStore, api_client_builder, and trim whitespace in ParseDateOnly ([#296](https://github.com/microsoft/kiota-abstractions-go/issues/296)) ([6c50f38](https://github.com/microsoft/kiota-abstractions-go/commit/6c50f38da78fe8769727b47127e4a7a70f184bcd))

## [1.11.4](https://github.com/microsoft/kiota-abstractions-go/compare/v1.11.3...v1.11.4) (2026-10-21)


### Bug Fixes

* **abstractions:** add nil receiver guards to headers and fix multipart close error check ([#294](https://github.com/microsoft/kiota-abstractions-go/issues/294)) ([a8125d6](https://github.com/microsoft/kiota-abstractions-go/commit/a8125d62282c4cc55162481da54f601b83b38e30))
* **auth, abstractions:** add nil receiver checks, host trimming, and request options safety ([#295](https://github.com/microsoft/kiota-abstractions-go/issues/295)) ([8c9bd9a](https://github.com/microsoft/kiota-abstractions-go/commit/8c9bd9aa995a0a50c7a1482185b599373c203ddf))
* **headers:** add nil receiver checks and improve whitespace trimming and additional value handling ([#292](https://github.com/microsoft/kiota-abstractions-go/issues/292)) ([83e3471](https://github.com/microsoft/kiota-abstractions-go/commit/83e3471416a0a7eba7637cab36331173bf3a1f5e))

## [1.11.3](https://github.com/microsoft/kiota-abstractions-go/compare/v1.11.2...v1.11.3) (2026-10-19)


### Bug Fixes

* track typed nil collections in backing store ([#289](https://github.com/microsoft/kiota-abstractions-go/issues/289)) ([3f8d682](https://github.com/microsoft/kiota-abstractions-go/commit/3f8d6820f5bd34574430a23062a0011b8285f5bf))

## [1.11.2](https://github.com/microsoft/kiota-abstractions-go/compare/v1.11.1...v1.11.2) (2026-10-16)


### Bug Fixes

* preserve original serialization start callbacks ([#287](https://github.com/microsoft/kiota-abstractions-go/issues/287)) ([4c759ff](https://github.com/microsoft/kiota-abstractions-go/commit/4c759ff17b8c6e84eff0097ebd23d97b2868603e))

## [1.11.1](https://github.com/microsoft/kiota-abstractions-go/compare/v1.11.0...v1.11.1) (2026-09-23)


### Bug Fixes

* Harden host comparison in AllowedHostsValidator ([#277](https://github.com/microsoft/kiota-abstractions-go/issues/277)) ([38675e4](https://github.com/microsoft/kiota-abstractions-go/commit/38675e4070cd50d25ed192d769f73e113a675322))

## [1.11.0](https://github.com/microsoft/kiota-abstractions-go/compare/v1.10.1...v1.11.0) (2026-09-17)


### Features

* Add support for the HTTP QUERY verb (RFC 10008) ([#273](https://github.com/microsoft/kiota-abstractions-go/issues/273)) ([9a1d756](https://github.com/microsoft/kiota-abstractions-go/commit/9a1d756af48aa2ea2c978803b4425e40ab41c402))

## [1.10.1](https://github.com/microsoft/kiota-abstractions-go/compare/v1.10.0...v1.10.1) (2026-08-31)


### Bug Fixes

* stop dropping scalar numeric query parameters ([#267](https://github.com/microsoft/kiota-abstractions-go/issues/267)) ([6bcceab](https://github.com/microsoft/kiota-abstractions-go/commit/6bcceab3ddf49d0710c1352d33c793d5428141b3))

## [1.10.0](https://github.com/microsoft/kiota-abstractions-go/compare/v1.9.4...v1.10.0) (2026-08-27)


### Features

* adds support for wildcard host validation ([#265](https://github.com/microsoft/kiota-abstractions-go/issues/265)) ([3458a73](https://github.com/microsoft/kiota-abstractions-go/commit/3458a737025051cb24873f8b5f8e20b69b8e7699))
* Handle map-style query parameters with nil values ([#247](https://github.com/microsoft/kiota-abstractions-go/issues/247)) ([dc2d48e](https://github.com/microsoft/kiota-abstractions-go/commit/dc2d48e594ae8d0562280d2b9c5b06a7a623085e))

## [1.9.4](https://github.com/microsoft/kiota-abstractions-go/compare/v1.9.3...v1.9.4) (2026-03-09)


### Bug Fixes

* duration normalization no longer produces invalid ISO 8601 week combinations ([7285e50](https://github.com/microsoft/kiota-abstractions-go/commit/7285e5075231e31a38d3d638c4c76a7aca190723))
* ISO 8601 duration normalization must not produce partial-week combinations ([aa3357f](https://github.com/microsoft/kiota-abstractions-go/commit/aa3357f68d1dfdd9e1d39d037ddce8e046c8268d))

## [1.9.3](https://github.com/microsoft/kiota-abstractions-go/compare/v1.9.2...v1.9.3) (2025-07-16)


### Bug Fixes

* for timeOnly to support string with nanosecond precision ([#206](https://github.com/microsoft/kiota-abstractions-go/issues/206)) ([fbc2113](https://github.com/microsoft/kiota-abstractions-go/commit/fbc21133d031ef3e09b6aeda2de54033bd8533ce))

## [1.9.2](https://github.com/microsoft/kiota-abstractions-go/compare/v1.9.1...v1.9.2) (2025-04-02)


### Bug Fixes

* drop dependency on common-go ([#199](https://github.com/microsoft/kiota-abstractions-go/issues/199)) ([5702488](https://github.com/microsoft/kiota-abstractions-go/commit/57024882f125ff719a2155293d5e28716f2ba9db))

## [1.9.1](https://github.com/microsoft/kiota-abstractions-go/compare/v1.9.0...v1.9.1) (2025-03-24)


### Bug Fixes

* upgrades common go dependency to solve triming issues ([5769aa7](https://github.com/microsoft/kiota-abstractions-go/commit/5769aa7c12ddcd01862ded353ffe4b966da25b55))
* upgrades common go dependency to solve triming issues ([73349ca](https://github.com/microsoft/kiota-abstractions-go/commit/73349ca6ebe5122fad13eb85fd957490ea8cfeea))

## [1.9.0](https://github.com/microsoft/kiota-abstractions-go/compare/v1.8.1...v1.9.0) (2025-03-13)


### Features

* upgrades required go version from go1.18 to go 1.22 ([526145f](https://github.com/microsoft/kiota-abstractions-go/commit/526145f3f219287feffe814d2f3efb2d9b7e233d))

## [1.8.1] - 2024-07-09

### Changed

- Fixed outdated reference to std uri template.

## [1.8.0] - 2024-07-09

### Changed

-  Removed native support for date in `std-uritemplate`  [#1.8.0](https://github.com/microsoft/kiota-abstractions-go/issues/183)

### Added

## [1.7.0] - 2024-07-09

-  Added accessors for headers and status to `ApiErrorable`  [#177](https://github.com/microsoft/kiota-abstractions-go/issues/177)

### Changed

## [1.6.1] - 2024-07-09

- Corrected two instances of `octet-steam` to `octet-stream` [#173](https://github.com/microsoft/kiota-abstractions-go/pull/173), [#174](https://github.com/microsoft/kiota-abstractions-go/pull/174)

## [1.6.0] - 2024-02-29

### Added

- Added support for untyped nodes. (https://github.com/microsoft/kiota/pull/4095)

## [1.5.6] - 2024-01-18

### Changed

- The input contains http or https which function will return an error. [#130](https://github.com/microsoft/kiota-abstractions-go/issues/130)

## [1.5.5] - 2024-01-17

### Changed

- Fixed a bug where reseting properties to null would be impossible with the in memory backing store. [microsoftgraph/msgraph-sdk-go#643](https://github.com/microsoftgraph/msgraph-sdk-go/issues/643)

## [1.5.4] - 2024-01-16

### Changed

- Fix bug where empty string query parameters are added to the request. [#133](https://github.com/microsoft/kiota-abstractions-go/issues/133)

## [1.5.3] - 2023-11-24

### Added

- Added support for multi valued query and path parameters of type other than string. [#124](https://github.com/microsoft/kiota-abstractions-go/pull/124)

## [1.5.2] - 2023-11-22

### Added

- Added ApiErrorable interface. [microsoft/kiota-http-go#110](https://github.com/microsoft/kiota-http-go/issues/110)

## [1.5.1] - 2023-11-15

### Added

- Added support for query an path parameters of enum type. [microsoft/kiota#3693](https://github.com/microsoft/kiota/issues/3693)

## [1.5.0] - 2023-11-08

### Added

- Added request information methods to reduce the amount of generated code.

## [1.4.0] - 2023-11-01

### Added

- Added serialization helpers. [microsoft/kiota#3406](https://github.com/microsoft/kiota/issues/3406)

## [1.3.1] - 2023-10-31

### Changed

- Fixed an issue where query parameters of type array of anything else than string would not be expanded properly. [#114](https://github.com/microsoft/kiota-abstractions-go/issues/114)

## [1.3.0] - 2023-10-12

### Added

- Added an overload method to set binary content with their content type.

## [1.2.3] - 2023-10-05

### Added

- A tryAdd method to RequestHeaders

## [1.2.2] - 2023-09-21

### Changed

- Switched the RFC 6570 implementation to std-uritemplate

## [1.2.1] - 2023-09-06

### Changed

- Fixed a bug where serialization registries would always replace existing values. [#95](https://github.com/microsoft/kiota-abstractions-go/issues/95)

## [1.2.0] - 2023-07-26

### Added

- Added support for multipart request body.

## [1.1.0] - 2023-05-04

### Added

- Added an interface to represent composed types.

## [1.0.0] - 2023-05-04

### Changed

- GA Release.

## [0.20.0] - 2023-04-12

### Added

- Adds response headers to Api Error class

### Changed

## [0.19.1] - 2023-04-12

### Added

### Changed

- Fixes concurrent map write panics when enabling backing stores.

## [0.19.0] - 2023-03-22

### Added

- Adds base request builder class to reduce generated code duplication.

## [0.18.0] - 2023-03-20

### Added

- Adds utility functions `CopyMap` and `CopyStringMap` that returns a copy of the passed map.

## [0.17.3] - 2023-03-15

### Changed

- Fixes panic when updating in-memory slices, maps or structs .

## [0.17.2] - 2023-03-01

### Added

- Adds ResponseStatusCode field in ApiError struct.

## [0.17.1] - 2023-01-28

### Added

- Adds a type qualifier for backing store instance type to be `BackingStoreFactory`.

### Changed

## [0.17.0] - 2023-01-23

### Added

- Added support for backing store.

## [0.16.0] - 2023-01-10

### Added

- Added a method to convert abstract requests to native requests in the request adapter interface.

## [0.15.2] - 2023-01-09

### Changed

- Fix bug where empty string query parameters are added to the request.

## [0.15.1] - 2022-12-15

### Changed

- Fix bug preventing adding authentication key to header requests.

## [0.15.0] - 2022-12-15

### Added

- Added support for multi-valued request headers.

## [0.14.0] - 2022-10-28

### Changed

- Fixed a bug where request bodies collections with single elements would not serialize properly

## [0.13.0] - 2022-10-18

### Added

- Added an API key authentication provider.

## [0.12.0] - 2022-09-27

### Added

- Added tracing support through OpenTelemetry.

## [0.11.0] - 2022-09-22

### Add
- Adds generic helper methods to reduce code duplication for serializer and deserializers
- Adds `WriteAnyValue` to support serialization of objects with undetermined properties at execution time e.g maps.
- Adds `GetRawValue` to allow returning an `interface{}` from the parse-node

## [0.10.1] - 2022-09-14

### Changed

- Fix: Add getter and setter on `ResponseHandler` pointer .

## [0.10.0] - 2022-09-02

### Added

- Added support for composed types serialization.

## [0.9.1] - 2022-09-01

### Changed

- Add `ResponseHandler` to request information struct

## [0.9.0] - 2022-08-24

### Changed

- Changes RequestAdapter contract passing a `Context` object as the first parameter for SendAsync

## [0.8.2] - 2022-08-11

### Added

- Add tests to verify DateTime and DateTimeOffsets default to ISO 8601.
- Adds check to return error when the baseUrl path parameter is not set when needed.

## [0.8.1] - 2022-06-07

### Changed

- Updated yaml package version through testify dependency.

## [0.8.0] - 2022-05-26

### Added

- Adds support for enum and enum collections responses.

## [0.7.0] - 2022-05-18

### Changed

- Breaking: adds support for continuous access evaluation.

## [0.6.0] - 2022-05-16

- Added a method to set the content from a scalar value in request information.

## [0.5.0] - 2022-04-21

### Added

- Added vanity methods to request options to add headers and options to simplify code generation.

## [0.4.0] - 2022-04-19

### Changed

- Upgraded uri template library for quotes in template fix.
- Upgraded to Go 18

## [0.3.0] - 2022-04-08

### Added

- Added support for query parameters with special characters in the name.

## [0.2.0] - 2022-04-04

### Changed

- Breaking: simplifies the field deserializers.

## [0.1.0] - 2022-03-30

### Added

- Initial tagged release of the library.
