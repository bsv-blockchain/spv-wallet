# Dependency Upgrade Log

**Date**: 2026-01-08
**Branch**: deps/upgrade
**Goal**: Upgrade all 50+ direct dependencies (except gorm.io/gorm v1.25.12)

---

## Summary Statistics
- **Total Dependencies Processed**: 28 (2 models + 6 Tier 1 + 6 Tier 2 + 5 Tier 3 + 9 Tier 4)
- **Successfully Upgraded**: 1 (go-logger)
- **Failed (Reverted)**: 0
- **Skipped (Pinned)**: 3 (gorm.io/gorm, bsm/redislock, gomodule/redigo)
- **No Update Available**: 27 (direct dependencies already latest)
- **Transitive Upgrades**: 14+ packages (security, validation, API libs)
- **Packages with Updates Available (baseline)**: 192

---

## Pre-Upgrade Verification
- ✅ On `deps/upgrade` branch
- ✅ GORM pinned at v1.25.12
- ✅ Replace directives intact (models, bsm/redislock, gomodule/redigo)
- ✅ 192 packages have updates available

---

## Module 1: models/

Path: `/Users/mrz/projects/spv-wallet/models/`

### github.com/pkg/errors
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.9.1 (already latest)
- **Test Result**: All tests passed
- **Notes**: No upgrade needed

### github.com/stretchr/testify
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.11.1 (already latest)
- **Test Result**: All tests passed (cached)
- **Notes**: No upgrade needed

**Module Summary**: Both dependencies already at latest versions. All tests passing.

---

## Module 2: Root Module

Path: `/Users/mrz/projects/spv-wallet/`

Starting tier-based upgrade of 50 direct dependencies...

### Tier 1: Low-Risk Utilities

#### github.com/google/uuid
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.6.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/iancoleman/strcase
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.3.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/samber/lo
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.52.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/pkg/errors
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.9.1 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/coocood/freecache
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.2.4 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/robfig/cron/v3
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v3.0.1 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

**Tier 1 Summary**: All 6 packages already at latest versions. All tests passing.

### Tier 2: Configuration & Logging

#### github.com/rs/zerolog
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.34.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/spf13/viper
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.21.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/go-viper/mapstructure/v2
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v2.4.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/mrz1836/go-logger
- **Status**: ✅ SUCCESS
- **Version Change**: v1.0.0 → v1.0.2
- **Test Result**: make test-unit PASSED
- **Notes**: Upgraded 2 minor versions

#### go.elastic.co/ecszerolog
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.2.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/prometheus/client_golang
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.23.2 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

**Tier 2 Summary**: 1 package upgraded (go-logger), 5 already at latest. All tests passing.
**Bonus**: golang.org/x/sys v0.39.0 → v0.40.0 (security update from Tier 3)

### Tier 3: Security & Stdlib Extensions

#### golang.org/x/crypto
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.46.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### golang.org/x/net
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.48.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### golang.org/x/sys
- **Status**: ✅ ALREADY UPGRADED (Tier 2)
- **Version**: v0.40.0 (upgraded in Tier 2)
- **Test Result**: make test-unit PASSED
- **Notes**: Automatically upgraded as transitive dependency in Tier 2

#### golang.org/x/text
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.32.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### golang.org/x/sync
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.19.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

**Tier 3 Summary**: All 5 security packages already at latest versions. All tests passing.

### Tier 4: HTTP & API Libraries

#### github.com/gin-gonic/gin
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.11.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/gin-contrib/pprof
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.5.3 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/getkin/kin-openapi
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.133.0 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/oapi-codegen/oapi-codegen/v2
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v2.5.1 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/oapi-codegen/runtime
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.1.2 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/swaggo/swag
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.16.6 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/swaggo/files
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.0.1 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/swaggo/gin-swagger
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v1.6.1 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

#### github.com/99designs/gqlgen
- **Status**: ⏭️ NO UPDATE AVAILABLE
- **Version**: v0.17.85 (already latest)
- **Test Result**: make test-unit PASSED
- **Notes**: No upgrade needed

**Tier 4 Summary**: All 9 packages already at latest versions. All tests passing.

**Bonus Transitive Upgrades**:
- bytedance/sonic v1.14.0 → v1.14.2
- gabriel-vasile/mimetype v1.4.8 → v1.4.12
- go-openapi/jsonpointer v0.21.0 → v0.22.4
- go-openapi/jsonreference v0.21.0 → v0.21.4
- go-openapi/spec v0.21.0 → v0.22.3
- go-openapi/swag v0.23.0 → v0.25.4 (plus submodules)
- go-playground/validator/v10 v10.27.0 → v10.30.1
- goccy/go-yaml v1.19.0 → v1.19.2
- mailru/easyjson v0.7.7 → v0.9.1
- quic-go/quic-go v0.57.0 → v0.58.0
- ugorji/go/codec v1.3.0 → v1.3.1
- golang.org/x/arch v0.20.0 → v0.23.0

