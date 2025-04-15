---
title: Final checks
linktitle: Final checks
weight: 70
---

We recommend you run a few final checks to catch any issues before you submit your PR.

## Check your PR contents

Double check that everything you expect to include in your PR is actually included. You should have:

* Configuration changes to `asure-arm.yaml`.
* Generated code changes, mostly in the `v2/api` folder.
* Generated documentation updates in the `docs` folder 
* Any hand-coded extensions you implemented, in the `v2/api/<group>/customizations` folder.
* A test for the resource in the `v2/internal/controllers` folder, with a recording for playback testing.
* A sample for the resource in the `v2/samples` folder, also with a recording.

## Check it builds

Running `task` with no parameters runs a quick set of checks that will catch most issues. 

```bash
task
```

Alternatively you can run the _exact same set of checks_ as our continuous integration (CI) builds by running

```bash
task ci
```

This can take up to an hour to run, and will hammer your PC while it runs, so it's an optional step. 

If you run into problems with your PR getting our continuous integration tests passing, using `task ci` locally is the fastest way to reproduce the problem.

## Congratulations!

You're ready to submit your PR! We're pretty quick at reviewing contributions, so you should expect to see a review within a few days. If you don't, feel free to ping us in the PR or on Slack.


