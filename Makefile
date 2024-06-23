-include gomk/main.mk
-include local/Makefile

clean: clean-default
	@rm -f resource.syso

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L hilighter -S ".git,*.pem"
endif
