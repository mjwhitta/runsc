-include gomk/main.mk
-include local/Makefile

clean: clean-default
ifeq ($(unameS),windows)
ifneq ($(wildcard resource.syso),)
	@remove-item -force ./resource_windows*.syso
endif
else
	@rm -f ./resource_windows*.syso
endif

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L hilighter -S ".git,*.pem"
endif
