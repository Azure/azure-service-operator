SHELL           = /bin/bash

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all:
	$(Q) cd generator && make all
	$(Q) cd generated && make all


# This is a separate target to avoid slow test runs when doing local development unless
# you opt in
.PHONY: test-int
test-int:
	$(Q) cd generated && make test-int
