---
title: Final checks
linktitle: Final checks
weight: 60
---

We recommend you run a final check to catch any issues before you submit your PR. Running `task` with no parameters runs a quick set of checks that will catch most issues. 

``` bash
task
```

Alternatively you can run the _exact same set of checks_ as our continuous integration (CI) builds by running

``` bash
task ci
```

This can take up to an hour to run, and will hammer your PC while it runs, so it's an optional step. 

If you run into problems with your PR getting our continuous integration tests passing, using `task ci` locally is the fastest way to reproduce the problem.

## Congratulations!

You're ready to submit your PR! We're pretty quick at reviewing contributions, so you should expect to see a review within a few days. If you don't, feel free to ping us in the PR or on Slack.


