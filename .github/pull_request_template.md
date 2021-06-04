<!-- Before you press Create pull request:

- Did you assign yourself to the PR?
- Did you connect an issue to the PR?
- Did you verify the issue has an estimate?
- Did you request necessary code reviewers?
- Is your PR title appropriate for release notes?
-->

## Description

<!--- Describe your changes in detail -->

## Motivation and Context

<!--- Why is this change required? What problem does it solve? -->
<!--- If it fixes an open issue, please link to the issue here. -->

## Types of changes

<!--- What types of changes does your code introduce? -->
<!--- Bug fix? (non-breaking change which fixes an issue)? -->
<!--- New feature? (non-breaking change which adds functionality) -->
<!--- Breaking change? (fix or feature that would cause existing functionality to change) -->

## Checklist

<!--- Go over all the following points, and make sure they all apply. -->

**Sanity checks:**

- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation

**Verify and {label}:**

- [ ] Did you add {env}, {type}, and {scope} labels?
- [ ] Does your PR contain a Database Migration {note}?
- [ ] Does your PR have breaking API or Settings changes {note}?

**LND specific:**

- [ ] If this is your first time contributing, we recommend you read the [Code
  Contribution Guidelines](https://github.com/LN-Zap/lnd-strike/blob/master/docs/code_contribution_guidelines.md)
- [ ] All changes are Go version 1.16 compliant
- [ ] For new code: Code is accompanied by tests which exercise both
  the positive and negative (error paths) conditions (if applicable)
- [ ] For bug fixes: Code is accompanied by new tests which trigger
  the bug being fixed to prevent regressions
- [ ] Any new logging statements use an appropriate subsystem and
  logging level
- [ ] Code has been formatted with `go fmt`
- [ ] Protobuf files (`lnrpc/**/*.proto`) have been formatted with
  `make rpc-format` and compiled with `make rpc`
- [ ] New configuration flags have been added to `sample-lnd.conf`
- [ ] For code and documentation: lines are wrapped at 80 characters
  (the tab character should be counted as 8 characters, not 4, as some IDEs do
  per default)
- [ ] Running `make check` does not fail any tests
- [ ] Running `go vet` does not report any issues
- [ ] Running `make lint` does not report any **new** issues that did not
  already exist
- [ ] All commits build properly and pass tests. Only in exceptional
  cases it can be justifiable to violate this condition. In that case, the
  reason should be stated in the commit message.
