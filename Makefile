VERSION ?= $(shell date +%Y%m%d%H%M%S)

bump_version:
	sed 's|VERSION = .*$$|VERSION = "$(VERSION)"|' -i ./player/player.go
	git add ./player/player.go
	git commit -m "Update version number to $(VERSION)"
