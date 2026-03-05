# Upgradable Resources

The following resources have newer versions available in the Azure REST API specifications. Versions in **bold** are recommended for upgrade.

| Group                   | Resource                 | Supported Stable | Supported Preview    | Available Stable  | Available Preview        |
|-------------------------|--------------------------|------------------|----------------------|-------------------|--------------------------|
| app                     | ContainerApp             | v1api20250101    | -                    | v1api20250701     | -                        |
| app                     | Job                      | v1api20250101    | -                    | v1api20250701     | -                        |
| app                     | ManagedEnvironment       | v1api20250101    | -                    | v1api20250701     | -                        |
| cache                   | RedisEnterprise          | v1api20250401    | -                    | v1api20250701     | -                        |
| cdn                     | Profile                  | v1api20230501    | -                    | **v1api20250601** | -                        |
| cognitiveservices       | Account                  | v1api20250601    | -                    | v1api20250901     | -                        |
| compute                 | AvailabilitySet          | v1api20241101    | -                    | v1api20250401     | -                        |
| compute                 | Disk                     | v1api20240302    | -                    | v1api20250102     | -                        |
| compute                 | DiskAccess               | v1api20240302    | -                    | v1api20250102     | -                        |
| compute                 | DiskEncryptionSet        | v1api20240302    | -                    | v1api20250102     | -                        |
| compute                 | Image                    | v1api20220301    | -                    | **v1api20250401** | -                        |
| compute                 | Snapshot                 | v1api20240302    | -                    | v1api20250102     | -                        |
| compute                 | VirtualMachine           | v1api20220301    | -                    | **v1api20250401** | -                        |
| compute                 | VirtualMachineScaleSet   | v1api20220301    | -                    | **v1api20250401** | -                        |
| containerinstance       | ContainerGroup           | v1api20211001    | -                    | **v1api20250901** | -                        |
| containerregistry       | Registry                 | v1api20230701    | -                    | **v1api20251101** | -                        |
| containerservice        | Fleet                    | v1api20250301    | -                    | v1api20250301     | **v1api20250801preview** |
| containerservice        | ManagedCluster           | v1api20250801    | v20251002preview     | v1api20251001     | v1api20251002preview     |
| dataprotection          | BackupVault              | v1api20231101    | -                    | **v1api20250901** | -                        |
| dbformysql              | FlexibleServer           | v1api20231230    | -                    | **v1api20241230** | -                        |
| devices                 | IotHub                   | v1api20210702    | -                    | **v1api20230630** | -                        |
| documentdb              | DatabaseAccount          | v1api20240815    | -                    | **v1api20251015** | -                        |
| documentdb              | MongoCluster             | v1api20240701    | -                    | **v1api20250901** | -                        |
| eventgrid               | Domain                   | v1api20200601    | -                    | **v1api20250215** | -                        |
| eventgrid               | EventSubscription        | v1api20200601    | -                    | **v1api20250215** | -                        |
| eventgrid               | Topic                    | v1api20200601    | -                    | **v1api20250215** | -                        |
| insights                | Component                | v1api20200202    | -                    | v1api20200202     | **v1api20180501preview** |
| insights                | ScheduledQueryRule       | v1api20220615    | v1api20250101preview | **v1api20231201** | v1api20250101preview     |
| keyvault                | Vault                    | v1api20230701    | v1api20210401preview | **v1api20250501** | v1api20210401preview     |
| kubernetesconfiguration | FluxConfiguration        | v1api20241101    | -                    | v1api20250401     | -                        |
| machinelearningservices | Registry                 | v1api20240401    | -                    | **v1api20250901** | -                        |
| machinelearningservices | Workspace                | v1api20240401    | -                    | **v1api20250901** | -                        |
| managedidentity         | UserAssignedIdentity     | v1api20230131    | -                    | **v1api20241130** | **v1api20220131preview** |
| network                 | ApplicationGateway       | v1api20220701    | -                    | **v1api20250501** | -                        |
| network                 | ApplicationSecurityGroup | v1api20240101    | -                    | **v1api20250501** | -                        |
| network                 | AzureFirewall            | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | BastionHost              | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | DnsForwardingRuleset     | v1api20220701    | -                    | **v1api20250501** | -                        |
| network                 | DnsResolver              | v1api20220701    | -                    | **v1api20250501** | -                        |
| network                 | FirewallPolicy           | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | LoadBalancer             | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | NatGateway               | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | NetworkInterface         | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | NetworkSecurityGroup     | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | NetworkWatcher           | v1api20241001    | -                    | v1api20250501     | -                        |
| network                 | PrivateEndpoint          | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | PrivateLinkService       | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | PublicIPAddress          | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | PublicIPPrefix           | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | RouteTable               | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | VirtualNetwork           | v1api20240301    | -                    | **v1api20250501** | -                        |
| network                 | VirtualNetworkGateway    | v1api20240301    | -                    | **v1api20250501** | -                        |
| redhatopenshift         | OpenShiftCluster         | v1api20231122    | -                    | **v1api20250725** | -                        |
| resources               | ResourceGroup            | v1api20200601    | -                    | **v1api20250401** | -                        |
| search                  | SearchService            | v1api20231101    | -                    | **v1api20250501** | -                        |
| sql                     | Server                   | v1api20211101    | -                    | **v1api20230801** | -                        |
| web                     | Site                     | v1api20220301    | -                    | **v1api20250501** | -                        |

