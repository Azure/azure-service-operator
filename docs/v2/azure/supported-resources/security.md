`DefenderForStorageSetting` is an extension resource owned by a StorageAccount. Its Azure name is always `current`, regardless of its Kubernetes name.

Azure does not support deleting these settings. To remove the Kubernetes resource while retaining the settings in Azure, set the annotation `serviceoperator.azure.com/reconcile-policy: detach-on-delete`. Detaching does not disable Defender for Storage or restore subscription-level settings.
