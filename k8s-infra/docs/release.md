# Releasing K8s-Infra
Releases should be done often and with as little effort as possible.

## Releasing
```bash
git tag "${version}" # v0.1.0
git push origin "${version}
```

The rest should be done via the GitHub, or other, release workflow.