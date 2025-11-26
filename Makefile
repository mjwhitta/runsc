-include gomk/main.mk
-include local/Makefile

clean: clean-default
ifeq ($(unameS),windows)
ifneq ($(wildcard ./cmd/runsc/resource_windows*.syso),)
	@remove-item -force ./cmd/runsc/resource_windows*.syso
endif
else
	@rm -f ./cmd/runsc/resource_windows*.syso
endif
