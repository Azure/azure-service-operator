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

The chances of any particular test failing was remote - but as the number of supported resources grew, our test suite got larger and larger and the probability of all the tests passing at the same time declined. We also noticed that the cost of running our test suite was growing, as was the length of time for a successful run.

To address these issues, one of our developers ([George Pollard](https://github.com/porges)) introduced HTTP recordings, improving the reliabilty and performance of our integration tests. These recordings capture the HTTP REST interactions between ASO and Azure.

Once a test has passed once, the recording allows us to replay that test in isolation - without any need to interact with Azure proper. We get to verify that ASO issues the same REST API requests, and that it behaves consistently as it receives the responses.

Benefits include

* **Consistency** - tests will (mostly) only fail if we break ASO, not if there's a problem elsewhere
* **Performance** - we can replay in a minute or so a test that might take 90 minutes to run "in the real world"  
  (The system's not perfect - timing differences from run to run very occasionally cause a test failure.)
* **Cost** - because we're not creating actual Azure resources, we don't incur much in the way of actual cost

We extended the recordings to our samples to verify that our samples work and can be used by new ASO users as a reference. I ([@theunrepentantgeek](https://github.com/theunrepentantgeek)) have had some poor experiences with other projects where I spent several days trying to use their samples, failing, only to be told "Yeah, those never worked" when I reached out for help.

We figured ASO could do better than that.
