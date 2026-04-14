# Test Tips

API Management Services are soft-deleted by default, which can make running tests difficult.

Before you can do a clean run of a test, you need to ensure any soft-deleted debris from prior test runs has been purged.

To purge a soft-deleted service, you have to use **az** - as of March 2025, it cannot be done via the Azure portal.

## Checking for soft-deleted services

### Using az apim

List all known soft-deleted services.

```bash
$ az apim deletedservice list --query "[].[name, location]"
[
  [
    "asotest-apimsvcv2-yovquu",
    "East US 2"
  ],
  [
    "asotest-apimanagementv2-yovquu",
    "East US"
  ],
  [
    "asotestqpewjd",
    "East US"
  ],
  [
    "asotest-apimanagementv2-twykac",
    "East US"
  ]
]
```

## Purging soft-deleted services

### Using az apim

You need to specify both the service name and location.

```bash
az apim deletedservice purge --service-name "asotestqpewjd" --location "eastus"
```

Deletion is slow - it can take a few minutes for az to return. After that, give it another 10m or so before rerunning your test.

For more information, see <https://andrewilson.co.uk/post/2022/09/apim-purge-soft-deleted-instance/>

## Troubleshooting

If APIM are having an outage, or their API doesn't work for whatever reason, you may find you simply need to wait 48h+ for the soft-delete to turn into a hard-delete.
