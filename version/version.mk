GOSU_MODULE_NAME?=github.com/switch-st/gosu
GOSU_VERSION_NAME?=$(shell git describe --tags --always --dirty)
GOSU_VERSION_CODE?=$(shell git rev-list --all --count)
GOSU_COMMIT_ID?=$(shell git describe --always --dirty --abbrev=0)
GOSU_BUILD_DATE?=$(shell date '+%Y-%m-%dT%H:%M:%S%z')
GOSU_BUILD_FLAG=-X $(GOSU_MODULE_NAME)/version.versionName=$(GOSU_VERSION_NAME) \
				-X $(GOSU_MODULE_NAME)/version.versionCode=$(GOSU_VERSION_CODE) \
				-X $(GOSU_MODULE_NAME)/version.commitID=$(GOSU_COMMIT_ID) \
				-X $(GOSU_MODULE_NAME)/version.buildDate=$(GOSU_BUILD_DATE)
