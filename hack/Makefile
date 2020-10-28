SHELL           = /bin/bash

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all:
	$(Q) cd generator && make all
	$(Q) cd generated && make all


