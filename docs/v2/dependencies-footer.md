
Dependencies are listed alphabetically. Refer to `install-dependencies.sh` for a recommended order of installation.

To update this file:

* Modify the file header content in `docs/v2/dependencies-header.md`;
* Modify the file footer in `docs/v2/dependencies-footer.md`; or
* Modify the dependencies installed by `.devcontainer/install-dependencies.sh`.

Regenerate the file using task:

``` bash
$ task doc:dependencies
```

Finally, submit the update as a PR.
