# Dependency Upgrade Log

**Date**: 2026-01-08
**Branch**: deps/upgrade
**Goal**: Upgrade all 50+ direct dependencies (except gorm.io/gorm v1.25.12)

---

## Summary Statistics
- **Total Dependencies Processed**: 14 (2 models + 6 Tier 1 + 6 Tier 2)
- **Successfully Upgraded**: 1 (go-logger)
- **Failed (Reverted)**: 0
- **Skipped (Pinned)**: 3 (gorm.io/gorm, bsm/redislock, gomodule/redigo)
- **No Update Available**: 13
- **Bonus Upgrades**: 1 (golang.org/x/sys from Tier 3)
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

