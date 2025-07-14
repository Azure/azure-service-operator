---
title: CEL Expressions
linktitle: CEL Expressions
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## CEL Expressions in ASO

ASO offers some support for [Common Expression Language (CEL)](https://github.com/google/cel-spec) expressions.

Expressions in ASO can be used to build your own Secret or ConfigMap output. This helps to solve problems like:

- Applications that need connection strings with specific (possibly non-standard) formats.
- Exporting arbitrary fields from a resource for use by other applications.
- Exporting static text alongside dynamic data.

CEL expressions are currently supported on most resources via the `spec.operatorSpec.configMapExpressions` and
`spec.operatorSpec.secretExpressions` fields, which allow configuring expressions output 
to ConfigMaps and Secrets, respectively.

### Inputs

Expressions have access to the following inputs:

| Type      | Variable(s)      |
|-----------|------------------|
| ConfigMap | `self`           |
| Secret    | `self`, `secret` |

`self`: The resource itself.

{{% alert title="Note" %}}
`self` represents the resource at the version you last applied the resource at. So for example if `v1api20200101`
has a `self.spec.foo` field that has been _removed_ in `v1api20220101`, you can still refer to `self.spec.foo`
in your expressions as long as you're `kubectl apply`-ing the resource with the `v1api20200101` version.

This also holds for new fields, you can refer to newly added fields _only in the API versions that have those fields_. 
{{% /alert %}}

`secret`: A set of secrets associated with the resource. The specific secrets supported vary by resource type, 
but can be found on the `spec.operatorSpec.secrets` type. 
For example, [eventhub Namespace](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespaceOperatorSecrets)
supports the following 4 secrets: 
- `primaryConnectionString`
- `primaryKey`
- `secondaryConnectionString`
- `secondaryKey`

### Outputs

Expressions for `spec.operatorSpec.configMapExpressions` and `spec.operatorSpec.secretExpressions`
must output `string` or `map[string]string` data. Any other expression output type will be rejected.

### CEL options, language features, and libraries

Compare with [Kubernetes supported options](https://kubernetes.io/docs/reference/using-api/cel/#cel-options-language-features-and-libraries)

| Feature                                                                                                             |
|---------------------------------------------------------------------------------------------------------------------|
| [Standard macros](https://github.com/google/cel-spec/blob/master/doc/langdef.md#macros)                             |
| [Standard functions](https://github.com/google/cel-spec/blob/master/doc/langdef.md#list-of-standard-definitions)    |
| [Default UTC Time Zone](https://pkg.go.dev/github.com/google/cel-go/cel#DefaultUTCTimeZone)                         |
| [Eagerly Validate Declarations](https://pkg.go.dev/github.com/google/cel-go/cel#EagerlyValidateDeclarations)        |
| [Extended String Library, version 3](https://pkg.go.dev/github.com/google/cel-go/ext#Strings)                       |
| [Extended Two Var Comprehensions, version 0](https://pkg.go.dev/github.com/google/cel-go/ext#TwoVarComprehensions)  |
| [Optional Types](https://pkg.go.dev/github.com/google/cel-go/cel#OptionalTypes)                                     |
| [Cross Type Numeric comparisons](https://pkg.go.dev/github.com/google/cel-go/cel#CrossTypeNumericComparisons)       |


### Escaping

ASO follows the [same escaping rules as Kubernetes](https://kubernetes.io/docs/reference/using-api/cel/#escaping),
except we don't escape CEL keywords.

While escaping isn't required for most properties, CEL doesn't support the `.`, `-`, or `/` characters in expressions, so
properties whose names contain these characters must be escaped. The escaping rules are:

| Escape sequence   | Property name equivalent                                                                                             | Unescaped example | Escaped example          |
|-------------------|----------------------------------------------------------------------------------------------------------------------|-------------------|--------------------------|
| \_\_underscores__ | __                                                                                                                   | my__field         | my\_\_underscores__field |
| \_\_dot__         | .                                                                                                                    | my.field          | my\_\_dot__field         |
| \_\_dash__        | -                                                                                                                    | my-field          | my\_\_dash__field        |
| \_\_slash__       | /                                                                                                                    | my/field          | my\_\_slash__field       |
| \_\_{keyword}__   | [CEL RESERVED keyword (false, true, in, null)](https://github.com/google/cel-spec/blob/master/doc/langdef.md#syntax) | false             | \_\_false__              |

### Examples

These examples are all written for a simple 
[cache.Redis](https://azure.github.io/azure-service-operator/reference/cache/v1api20230801/#cache.azure.com/v1api20230801.Redis).
There's nothing special about Redis it's just used as an example.

In each of these examples, we start with the given Redis and apply the expression as mentioned in the example itself.

It might be useful to refer to 
the [DestinationExpression](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)
documentation to understand what the `name`/`key`/`value` fields are doing in all of these examples.

For more advanced examples, refer to the 
[CEL language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md).

**Sample Redis:**
```yaml
apiVersion: cache.azure.com/v1api20230801
kind: Redis
metadata:
  name: sampleredis1
  namespace: default
  annotations:
    foo: bar
    baz: qux
spec:
  location: westus2
  owner:
    name: aso-sample-rg
  tags:
    "tag one": hello
    "tag two": goodbye
    three: four
  sku:
    family: P
    name: Premium
    capacity: 1
  enableNonSslPort: false
  minimumTlsVersion: "1.2"
  redisConfiguration:
    maxmemory-delta: "10"
    maxmemory-policy: allkeys-lru
  redisVersion: "6"
  zones:
  - "1"
  - "2"
  - "3" 
```

{{% alert title="Note" %}}
Just to keep it simple, the sample YAML doesn't include status. Real resources have status, and you can build your
CEL expressions with `self.status` in addition to `self.spec` or `self.metadata`.
{{% /alert %}}

#### A full example

Expression snippet:
```yaml
spec:
  operatorSpec:
    configMapExpressions:
    - name: my-configmap
      key: location
      value: self.spec.location
```

Output (ConfigMap):
```yaml
metadata:
  name: my-configmap
data:
  location: westus2
```

From here on, we'll elide the supporting YAML of the above example and just focus on the CEL expression input (`value`)
and the resulting output, still based on the example `Redis` above.

| Description                                                | Expression                                                                    | Result                                                                         |
| ---------------------------------------------------------- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| Hardcoded string                                           | `"helloworld"`                                                                | `helloworld`                                                                   |
| Formatted string                                           | `"%s:%d".format([self.spec.location, 7])`                                     | `westus2:7`                                                                    |
| String math                                                | `self.metadata.namespace + ":" + self.metadata.name`                          | `default:sampleredis1`                                                         |
| Int output (error)                                         | `self.spec.sku.capacity`                                                      | Error, expression must return one of [string,map(string, string)], but was int |
| Coerce to string                                           | `string(self.spec.sku.capacity)`                                              | `1`                                                                            |
| Map output                                                 | `self.metadata.annotations`                                                   | `{"foo": "bar", "baz": "qux"}`                                                 |
| Array macro                                                | `self.spec.zones.filter(a, int(a) % 2 == 0).join("-")`                        | `2-4`                                                                          |
| Map macro (TwoVarComprehension) replace spaces with dashes | `self.spec.tags.transformMapEntry(k, v, {k.replace(" ", "-"): v})`            | `{"tag-one": "hello", "tag-two": "goodbye", "three": "four"}`                  |
| Select array item                                          | `self.spec.zones[0]`                                                          | `1`                                                                            |
| Select map item                                            | `self.metadata.annotations["foo"]`                                            | `bar`                                                                          |

