#
# These variables should not need tweaking.
#

# ==============================================================================
# Includes

# include the common make file
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif

APIROOT := $(ROOT_DIR)/pkg/api