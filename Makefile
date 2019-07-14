# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL = test

# ----------------------------------------------------------------------------
# target

YAML_TEST_SUITE_REPO = https://github.com/yaml/yaml-test-suite
YAML_TEST_SUITE_BRANCH = data

.PHONY: yaml/test-suite/update
yaml/test-suite/update: yaml/test-suite/clean
	git clone --depth 1 --branch ${YAML_TEST_SUITE_BRANCH} --single-branch ${YAML_TEST_SUITE_REPO} $(TMPDIR)/yaml-test-suite
	cp -r $(TMPDIR)/yaml-test-suite ${CURDIR}/testdata/yaml-test-suite
	@rm -rf $(TMPDIR)/yaml-test-suite ${CURDIR}/testdata/yaml-test-suite/.git

.PHONY: yaml/test-suite/clean
yaml/test-suite/clean:
	@rm -rf ${CURDIR}/testdata/yaml-test-suite

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk

# ----------------------------------------------------------------------------
# overlays
