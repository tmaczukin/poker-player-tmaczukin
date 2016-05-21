VERSION ?= $(shell date +%Y%m%d%H%M%S)
save:
	sed 's|VERSION = .*$$|VERSION = "$(VERSION)"|' -i ./player/player.go
	git add ./player/player.go
	git commit
