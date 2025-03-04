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

### Using az rest

If you get an error from `az apim deletedservice list`, an alternative is to use their REST API directly. Fortunately, **az** lets you do this too

``` bash
az rest --method get --url https://management.azure.com/subscriptions/{subscription}/providers/Microsoft.ApiManagement/deletedservices?api-version=2021-08-01
```

## Purging soft-deleted services

### Using az apim

You need to specify both the service name and location.

```bash
az apim deletedservice purge --service-name "asotestqpewjd" --location "eastus"
```

Deletion is slow - it can take a few minutes for az to return. After that, give it another 10m or so before rerunning your test.

For more information, see <https://andrewilson.co.uk/post/2022/09/apim-purge-soft-deleted-instance/>

### Using az rest

Similarly, you can use **az** to purge a selected instance via the REST API. Fill in `{subscription}`, `{location}` and `{service-name}` and run this command:

``` bash
az rest --method delete --url https://management.azure.com/subscriptions/{subscription}/providers/Microsoft.ApiManagement/locations/{location}/deletedservices/{service-name}?api-version=2021-08-01
```

## Troubleshooting

If APIM are having an outage, or their API doesn't work for whatever reason, you may find you simply need to wait 48h+ for the soft-delete to turn into a hard-delete.
