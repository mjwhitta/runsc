-include gomk/main.mk
-include local/Makefile

ifneq ($(unameS),Windows)
spellcheck:
	@codespell -f -L hilighter -S ".git,*.pem"
endif
