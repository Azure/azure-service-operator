# Issue Management

## New Issue Triage

On a regular basis (target: weekly) the development team will triage new issues matching the following [GitHub search](https://github.com/Azure/azure-service-operator/issues?q=is%3Aissue+is%3Aopen+label%3A%22needs-triage+%3Amag%3A%22+sort%3Aupdated-desc+):

```
is:issue is:open label:needs-triage label:"needs-triage :mag:" 
```

When new issues are created, our bot will automatically add the `needs-triage` tag so they show up in this list.

## Old Issue Review

On a regular basis (target: weekly) the development team will review old issues using the following [GitHub search](https://github.com/Azure/azure-service-operator/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-asc):

```
is:issue is:open sort:updated-asc
```

To ensure we're not reviewing the same issues every week, a comment will be added indicating the outcome of the review, resetting the updated timestamp and moving the issue to the end of the list.
