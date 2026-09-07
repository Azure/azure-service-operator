---
title: "Database Watchers"
linktitle: "Database Watchers"
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

A `Watcher` collects performance data from the targets you give it. When a target has public network access
disabled, the watcher reaches it over a *managed private endpoint*, which you declare as a
`SharedPrivateLink`. This page covers how ASO completes that private endpoint, and what to expect while it
does.

## How a shared private link becomes usable

Creating a `SharedPrivateLink` asks Azure to open a private endpoint connection on the resource the link
points at. That connection starts out `Pending`: Azure creates it from a Microsoft-owned subscription rather
than yours, so no role assignment of yours can approve it. Until somebody approves it, the watcher cannot
collect from the target.

Azure reports nothing about that connection on the link itself - `status` on the `SharedPrivateLink` is never
populated - so ASO reads the connection on the target resource, which is the only place its state appears.
The link's `Ready` condition follows what it finds there:

| Connection state | `Ready` condition |
| --- | --- |
| Not opened yet | `False`, waiting for Azure to open it |
| `Pending` | `False`, and ASO approves it if it can (see below) |
| `Approved` | `True` |
| `Rejected`, `Disconnected` | `False`, reporting the decision and leaving it alone |

A link that stays `False` is telling you the private endpoint behind it carries no traffic yet. Check the
condition message with `kubectl describe` for the reason.

## When ASO approves the connection for you

ASO approves the connection only when the link names its target as a Kubernetes resource:

```yaml
apiVersion: databasewatcher.azure.com/v20241001preview
kind: SharedPrivateLink
metadata:
  name: aso-sample-spl
  namespace: default
spec:
  owner:
    name: aso-sample-watcher
  groupId: sqlServer
  requestMessage: Please approve the connection from the database watcher
  # Naming the target as a Kubernetes resource, rather than by armId, is what lets ASO approve
  # the private endpoint connection this link opens on it
  privateLinkResourceReference:
    group: sql.azure.com
    kind: Server
    name: aso-sample-server
```

That reference is what supplies the API version the connection is read and written with - the link's own
`databasewatcher` API version is not valid for `Microsoft.Sql` or `Microsoft.Kusto` - and it means ASO is
completing a handshake on a resource you already manage with it.

A link that names its target by `armId` instead keeps the behaviour it has always had. With no API version
its connections cannot even be read, so ASO reports nothing about them and the link goes `Ready` on its own
provisioning alone. Approve those connections yourself, in the portal or with `az`.

ASO also declines to approve, and holds the link `False` with the reason, when:

- **The target is managed by a different operator.** Operators sharing a cluster have their own credentials
  and policies, and none of that is visible from the link.
- **The target asks for a different credential.** Approving writes to the target with the link's credential,
  so anything short of an identical `serviceoperator.azure.com/credential-from` is refused.
- **Either resource's [reconcile policy]({{< relref "annotations" >}}) forbids modification.** The policy
  suppresses the approval, not the reporting, so the link still tells you the connection is pending.
- **Two connections on the target are named after the link.** A link name is unique only under its own
  watcher, so ASO refuses to guess which connection is this link's.

## Permissions

Approving a private endpoint connection requires `Owner`, or another role granting
`Microsoft.Sql/servers/privateEndpointConnections/write` on the target, which the credential that created
the link does not necessarily hold. When ASO is refused, the link reports that approval is required and that
its credential cannot give it - approve the connection yourself, or grant the operator the role.

## Restarting a watcher after approval

A watcher that is already running when a connection is approved
[must be restarted](https://learn.microsoft.com/en-us/azure/azure-sql/database-watcher-manage) before it
uses that connection. ASO does not restart one for you, because stopping a watcher interrupts collection for
every target on it. Nothing sequences a watcher's start against a link's approval either, so this can happen
on first creation as well. If a target reports no data over a private endpoint you know to be approved, stop
and start its watcher.
