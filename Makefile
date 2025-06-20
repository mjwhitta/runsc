-include gomk/main.mk
-include local/Makefile

clean: clean-default
ifeq ($(unameS),windows)
ifneq ($(wildcard resource_windows*.syso),)
	@remove-item -force ./resource_windows*.syso
endif
else
	@rm -f ./resource_windows*.syso
endif

mr: fmt
	@make GOOS=darwin reportcard spellcheck vslint
	@make GOOS=linux reportcard spellcheck vslint
	@make CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows \
	    reportcard spellcheck vslint
	@make test

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L hilighter -S ".git,*.pem"
endif
