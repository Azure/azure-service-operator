---
title: "Why record our tests?"
date: 2026-02-03
description: "Why do we use HTTP recordings for our tests?"
type: blog
---

Azure Service Operator v2 uses recorded HTTP interactions for many of its integration tests. This blog post explains why we do this, and the benefits it brings.

Early in the development of ASO v2, we ran into problems with some of our end-to-end tests.

Sometimes a test that successfully worked one day would fail the next. We'd spend time investigating, assuming that this was due to something we'd changed, only to find the failure was due to outside factors.

We encountered quota issues, capacity limits, outages, and more.

The chances of any particular test failing was remote - but as the number of supported resources grew, our test suite got larger and larger and the probability of all the tests passing at the same time declined. We also noticed that the cost of running our test suite was growing, as was the length of time for a successful run, to the point where a full run might take many hours.

To address these issues, one of our developers ([George Pollard](https://github.com/porges)) introduced HTTP recordings, improving the reliabilty and performance of our integration tests. These recordings capture the HTTP REST interactions between ASO and Azure.

Once a test has passed once, the recording allows us to replay that test in isolation - without any need to interact with Azure proper. We get to verify that ASO issues the same REST API requests, and that it behaves consistently as it receives the responses.

We've also extended the recordings to our samples to verify that the samples work and can be used by new ASO users as a reference. I ([@theunrepentantgeek](https://github.com/theunrepentantgeek)) have had some poor experiences with other projects where I spent several days trying to use their samples, failing, only to be told "Yeah, those never worked" when I reached out for help. We figured ASO could do better than that.

Benefits of the approach include:

* **Consistency** - tests will (mostly) only fail if we break ASO, not if there's a problem elsewhere
* **Performance** - we can replay in a minute or so a test that might take 90 minutes to run "in the real world"  
  (The system's not perfect - timing differences from run to run very occasionally cause a test failure.)
* **Cost** - because we're not creating actual Azure resources, we don't incur much in the way of actual cost

Of course, nothing comes without a cost, and we've had to tweak a few things to make this work well:

* **Controllers aren't inherently deterministic** - different test runs can do things slightly differently, due to variations in goroutine scheduling and timing. To allow for this, we've implemented a [`replayRoundTripper`](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/testcommon/vcr/v3/replay_roundtripper.go) to allow some flexibility in matching requests to recorded responses.

* **Caching can break our tests** - any global cache we add (such as the one for Azure built-in RoleDefinitions in [`role_assignment_extensions.go`](https://github.com/Azure/azure-service-operator/blob/main/v2/api/authorization/customizations/role_assignment_extensions.go)) can result in different behavior between test runs, forcing us to add controls so we can disable them during tests.

* **Disconnect from Azure** - these tests ensure that ASO behaves correctly given a particular sequence of HTTP interactions, but they don't catch issues when changes to Azure itself happen. For example, our tests of `ScheduledQueryRules` from Microsoft.Insights version `2024-01-01preview` continued to pass even after the preview API was removed from Azure.

Overall though, the benefits have far outweighed the costs, and we're pleased with how this approach has worked out for us.
