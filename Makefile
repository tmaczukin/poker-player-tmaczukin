VERSION ?= $(shell date +%Y%m%d%H%M%S)
save:
	sed 's|VERSION = .*$$|VERSION = "$(VERSION)"|' -i ./leanpoker/player.go
	git add ./leanpoker/player.go
	git commit
