# Contributing

We love your input. Whether it be:

- Reporting a bug
- Submitting a fix
- Proposing new features
- Discussing the state of the code

All feedback is welcome!

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## Official Repository

We use Github to host code, track issues, and accept pull requests. This Github repository is the "official" repository of this project. All changes/fixes/suggestions should be submitted here (i.e. github.com/coinbase/staking-client-library-go)

## Reporting Bugs

If you find a bug, please create an issue in our GitHub repository.

## Suggesting Enhancements

If you have an idea for a new feature, or just general feedback, please create an issue in our GitHub repository.

## Pull Request Process

1. Fork the repository on GitHub.
2. Clone your fork to your local machine.
3. Create a new branch on which to make your change, e.g. `git checkout -b my_new_feature`.
4. Make your change (more tips [below](#making-changes).
5. Push your change to your fork.
6. Submit a pull request.

### Making Changes

This repository uses a code generation process based off of protobuf pushed to [the public BSR](https://buf.build/cdp). To ensure the most recent updates are reflected in your changes, regenerate the code by running:

```shell
make gen
```

Before submitting the PR, ensure the linting CI will pass by running this command and checking for errors:

```shell
make lint
```

or use this command fix the linting mistakes proactively:

```shell
make lintfix
```
