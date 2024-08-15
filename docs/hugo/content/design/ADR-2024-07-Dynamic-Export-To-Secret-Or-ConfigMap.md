---
title: '2024-07: Dynamic Export to Secret or ConfigMap'
---

## Context

It is a common ask for users to request a new value be exportable to `ConfigMap`, or for the values
that are currently exported to a `ConfigMap` or `Secret` to be customizable in some way.

Outstanding issues related to this topic:
1. [#2920: Feature: Support export of storage account connection strings](https://github.com/Azure/azure-service-operator/issues/3929).
2. [#2555: Feature: Support dynamic secret/configmap exports](https://github.com/Azure/azure-service-operator/issues/2555)
3. [#3711: PublicIPPrefix: Allow exporting IP range as ConfigMap](https://github.com/Azure/azure-service-operator/issues/3711)
4. [#3671: DatabaseAccount CRD: Add support to store the account Id in a config map](https://github.com/Azure/azure-service-operator/issues/3671)

(there may be more)

Digging into [#2920](https://github.com/Azure/azure-service-operator/issues/3929):

We don't currently support Azure storage connection strings, which is something that ASOv1 has
supported for quite some time.

The format of a connection string is:
`DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=%s`

Users would like the ability to export a `Secret` formatted like this from a `StorageAccount`. While we could
manually support this shape in the `OperatorSpec` of the `StorageAccount`, a better solution would be one that
scaled to other connection string formats and worked for other resources without manual effort required for each one.

## Requirements

* Support composite exports that contain a combination of static text and variable data.
* Support variable data from from spec or status.
* Support variable data from operatorSpec.secrets (only to secrets).
  * If no secrets are used, no secret APIs should be called. 
* Mechanism of export is safe in a multitenant environment.

## Options

There are 3 aspects to delivering this feature.

1. How to express the desired output string format. This is the [template language](#template-language).
2. Where to signal the desire for one or more of these dynamic outputs, and where they should go. This is 
   the [location](#location).
3. How to include secret data.

### Template Language

#### Option 1: Go Text Templates (text/template)

Templates based on the Go stdlib [text/template](https://pkg.go.dev/text/template) package.

The format of these templates is: 
`DefaultEndpointsProtocol=https;AccountName={{.Spec.AzureName}};AccountKey={{.Spec.OperatorSpec.Secrets.Key1}};EndpointSuffix=core.windows.net`

The Go code involved in parsing this template would look something like this:
```go
func renderTemplate(resource any, template string) string {
	tmpl, err := template.New("test").Parse(template)
	if err != nil { 
		panic(err) 
	}
	builder := strings.Builder{}
	err = tmpl.Execute(builder, resource)
	if err != nil { 
		panic(err) 
	}

	return builder.String()
}
```

**Pros**

* Likely familiar to users of Helm, which uses the same engine, as do many other projects like go-task.
* Simple templates are simple, string interpolation is supported.

**Cons**

* `text/template` is powerful and not very configurable. It supports loops (`range`), if-statements, and some builtin
  functions (`printf`, etc) out of the box without any way to restrict these capabilities. This puts the operator at 
  risk in a multi-tenant environment because there's nothing stopping namespace A from submitting a template 
  `range int64max`.

#### Option 2: Google Common Expression Language (CEL)

You can read more about it in the 
[language specification](https://github.com/google/cel-spec/blob/master/doc/langdef.md#overview) and at
[cel-go](https://github.com/google/cel-go).

The format of these templates is: 
`'DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net'.format([resource.spec.azureName, resource.spec.operatorSpec.secrets.key1])`

You can read more about other CEL string extensions and the format function specifically in 
[cel-go string extensions](https://github.com/google/cel-go/blob/master/ext/strings.go#L81).

If we really wanted string interpolation we could possibly write a helper for it, see 
[here](https://github.com/google/cel-go/blob/a555792723a2ae8d35cec381e8b222b5b20cb1d9/ext/strings.go#L468).

The Go code involved in parsing this template would look something like this:
```go
// renderTemplate produces a variable from the given template
func renderTemplate(resource any, template string) string {
	// The environment can be customized by providing the options cel.EnvOption to the call.
	// Those options are able to disable macros, declare custom variables and functions, etc.
	env, err := cel.NewEnv(
		ext.Strings(),
		cel.Declarations(
			decls.NewVar("resource", decls.Dyn)),
	)
	if err != nil {
		panic(err) // Simple panic-based error handling for example purposes.
	}

	ast, iss := env.Compile(template)
	if iss.Err() != nil {
		panic(iss.Err())
	}
	// Check the output type is a string.
	if !reflect.DeepEqual(ast.OutputType(), cel.StringType) {
		panic(fmt.Sprintf(
			"Got %v, wanted %v output type",
			ast.OutputType(), cel.StringType,
		)
	}

	program, err := env.Program(ast)
	if err != nil {
		panic(err)
	}


    input := map[string]any{
        "resource": resource,
    }
	
    out, _, err := program.Eval(input)
    if err != nil {
        panic(err)
    }
    outStr := out.Value().(string)
    return outStr
}
```

**Pros**

* We have a significant amount of control over what functions are accessible in the CEL expressions, 
  including defining our own if we want.
* Security: Users have little/no flow-control constructs, and CEL is hardened for use in multitenant environments such 
  as Kubernetes APIServer.
* Power: We can parse the CEL expression and examine the AST ourselves to extract useful information such as 
  which secrets they're attempting to use, which enables us to deliver on the "no secret APIs called if no secrets used"
  goal.

**Cons**

* May not be as familiar to users.
* No string interpolation, which makes the templates not quite as clean as they otherwise could be.
  * Mitigation: We could write our own, although then that function will _definitely_ not be familiar to users. 

#### Option 3: Write our own simple text-replacement engine

Supporting a format similar to `text/template`, but without any of the other looping/flow-control features.
`DefaultEndpointsProtocol=https;AccountName={{.Spec.AzureName}};AccountKey={{.Spec.OperatorSpec.Secrets.Key1}};EndpointSuffix=core.windows.net`

**Pros**

* Theoretically simple

**Cons**

* More work than using something off the shelf. Including things like getting errors right, etc.
* Not as easily extensible as something like CEL if we wanted to add more capabilities.
* Trusting our security/hardening is scarier than trusting a "real" product.

### Location

#### Option 1: Part of operatorSpec.secrets and operatorSpec.configMap

A new `dynamicValues` option in `operatorSpec.secrets` and `operatorSpec.configMaps`

**Secrets**
```yaml
  operatorSpec:
    secrets:
      hostName:
        name: redis-secret
        key: hostName
      sslPort:
        name: redis-secret
        key: port
      dynamicValues:
        - name: redis-secret # Name of the destination secret
          key: hostPort # Name of the key in the secret
          value: "'%s:%s'.format([secret.hostName, secret.sslPort])" # Value (format) of the secret
```

**ConfigMaps**
```yaml
  operatorSpec:
    configMaps:
      dynamicValues:
        - name: account-data # Name of the destination configMap
          key: id # Name of the key in the configMap
          value: "resource.status.id" # Value (format) of the configMap entry
```

**Pros**

* Fits into our existing model relatively well.

**Cons**

* Need to deal with possibility of property name collisions with `dynamicValues`,
  though practically it seems unlikely to be common.

#### Option 2: Two new operatorSpec properties: operatorSpec.dynamicSecrets and operatorSpec.dynamicConfigMaps

Other than the location, this is very similar to the
[first proposal](#option-1-part-of-operatorspecsecrets-and-operatorspecconfigmap).

**Secrets**
```yaml
  operatorSpec:
    secrets:
      hostName:
        name: redis-secret
        key: hostName
      sslPort:
        name: redis-secret
        key: port
    dynamicSecrets:
      - name: redis-secret # Name of the destination secret
        key: hostPort # Name of the key in the secret
        value: "'%s:%s'.format([secret.hostName, secret.sslPort])" # Value (format) of the secret
```

**ConfigMaps**
```yaml
  operatorSpec:
    dynamicConfigMaps:
      - name: account-data # Name of the destination configMap
        key: id # Name of the key in the configMap
        value: "resource.status.id" # Value (format) of the configMap entry
```

**Pros**

* Fits into our existing model relatively well.
* Avoids potential issue with names clashing if adding a new magically named field in 
  `operatorSpec.secrets` or `operatorSpec.configMaps`.

**Cons**


#### Option 3: A new CRD `Exporter`

The final name of this resource is TBD.

```yaml
apiVersion: serviceoperator.azure.com/v1
kind: Exporter
spec:
  connectedResources:
    - ref:
        group: documentdb.azure.com
        kind: DatabaseAccount
        name: mydb
      secrets:
        - name: redis-secret # Name of the destination secret
          key: hostPort # Name of the key in the secret
          value: "'%s:%s'.format([secret.hostName, secret.sslPort])" # Value (format) of the secret
      configMaps:
        - name: account-data # Name of the destination configMap
          key: id # Name of the key in the configMap
          value: "resource.status.id" # Value (format) of the configMap entry
```

Open question if this would also support the well-typed approach we have now with fields like:
```yaml
    secrets:
      hostName:
        name: redis-secret
        key: hostName
      sslPort:
        name: redis-secret
        key: port
```

or if everything required a dynamic expression.

**Pros**

* Could combine multiple resources data into a single ConfigMap, though this brings questions of ownership
  and deletion/propagation.

**Cons**

* Another way to do things. Possibly confusing because there will be 2 ways to export data to configmap/secret 
  for many existing resources.
* More complicated, especially setting up watches on N resource types so that changes are correctly monitored and
  propagated quickly. That additional complexity doesn't necessarily bring all that many wins.
* Access to secret data is more difficult to implement. We would need to have a static mapping of resource types and
  the `Exporter` would need to call each referenced resources `ExportKubernetesResources`, and we'd likely need to 
  refactor `ExportKubernetesResources` some to support the new pattern (though we may need to do this anyway).

### Access to secret data

Getting access to data to export to ConfigMaps is pretty easy, we'll just pass the whole `resource` to
the expression. As shown elsewhere fields will be accessible in expressions like `resource.status.id`.

Secret data is by definition _not_ on the `resource` though, and is often returned in a list format only after
calling a special API. As the [requirements](#requirements) section discusses, we need to ensure that we only
invoke this secret API if and when a user asks for secrets. Then, actual access to the secret fields should
be off of a _separate_ object: `secret.hostName` or `secret.key1` for example.

We currently have an extension for Secrets or custom ConfigMaps whose signature is:
```go
ExportKubernetesResources(
    ctx context.Context,
    obj genruntime.MetaObject,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) ([]client.Object, error)
```

This extension is used both by handcrafted extensions used to export secrets and autogenerated extensions to support
`$generatedConfigs`.
This implementation is not suitable to supporting the new approach because it returns `[]client.Object`, 
but to supply the CEL-expressions we need a return type of an object or a map whose
properties/keys line up with the secret names. Alternatively we could delegate evaluation of the CEL expression
to the extension itself but this would require relatively complex code autogenerated for each type, since every
type will support the new `operatorSpec.dynamicConfigMaps` field.

Since `ExportKubernetesResources` is already used for a variety of purposes, rather than change its signature I propose
we add a new sister interface to it called `SecretGetter` for resources that have Secret APIs to call.
This interface must be implemented manually alongside the manual implementation of the
`KubernetesExporter` resource. We _could_ add the `SecretGetter` method to the `KubernetesExporter` interface and 
just have some implementations that return nothing always (because they don't get any secrets, only ConfigMaps),
but I think it might be cleaner to have a separate interface.

The existing `KuberentesExporter` interfaces implementation should be updated to share code with the
`SecretGetter` implementation to reduce code duplication.

The signature of this interface is:
```go
type GetSecretsParameters struct {
	Obj MetaObject
	ARMClient *genericarmclient.GenericClient
	Log logr.Logger

	// Hints is the set of secrets used in the expression(s), which can be used to filter/define what 
	// if any secrets should be retrieved
	Hints set.Set[string]
}

type SecretGetter interface {
	// GetSecrets retrieves the set of secrets defined by params.
	// The resulting map is keyed by the secret key (== property name of operatorSpec.secrets.<key>)
	// which is the same as what is referenced by the user when authoring their expressions and passed in
	// via params.Hints
	GetSecrets(
		ctx context.Context,
		params *GetSecretsParameters,
	) (map[string]string, error)
}
```

The generic reconciler will check for this extension like it does for others before invoking the new generic code
to handle the `operatorSpec.dynamicSecrets` field.

## Decision

### Aspect 1: Template language

We will use [Option 2: CEL](#option-2-google-common-expression-language-cel) as it gets us the most power for the lowest 
amount of effort and is fully customizable and already hardened for use in multitenant environments.

### Aspect 2: Location

We will use [Option 2: operatorSpec.dynamicConfigMaps and operatorSpec.dynamicSecrets](#option-2-two-new-operatorspec-properties-operatorspecdynamicsecrets-and-operatorspecdynamicconfigmaps).
It avoids the problems with [Option 3](#option-3-a-new-crd-exporter) and avoids the 
(admittedly minimal) collision risk with [Option 1](#option-1-part-of-operatorspecsecrets-and-operatorspecconfigmap)

### Aspect 3: Secret access

Only one option was presented and we're using it.

### FAQ

Q: How does this new way avoid calling the secrets APIs if the user didn't ask for any secrets to be exported? 
Keep in mind that some resources such as ManagedCluster (at v2/api/containerservice/customizations/managed_cluster_extensions.go)
require the ability to differentiate _which_ secrets are being requested as well because there are multiple
secrets APIs.

A: For `DynamicSecrets`, the function signature to discover secrets will have `secretHints` passed to it, 
which will contain details about the specific secrets requested in the formatting expression. The resource-specific 
implementation can use these hints to determine which (if any) APIs to call. Empty `secretHints` means no secrets
were requested and no secret APIs should be invoked.

Q: How will `secretHints` be determined, since they're buried in an expression the users authored?

A: We will examine the CEL AST and find all instances of `SelectExpr` and extract out those that correspond to
`secret.foo`.

Q: What about the existing `$generatedConfigs` capability from `azure-arm.yaml`?

A: It won't be removed, but will be deprecated in favor of this more generic way that requires less manual effort
on our part.

Q: Can this be used to export a particular item from an arbitrary array or map?

A: Yes! The syntax for maps is: `request.myMap['hello']`, and for arrays `request.myArray[0]`. CEL has a concept of
macros and there are some useful ones such as `filter`, `exists`, and `all` 
[built in](https://github.com/google/cel-spec/blob/master/doc/langdef.md#macros), so you can craft more complex
expressions like `request.slice.exists_one(i, i.name=='foo') ? request.slice.filter(i, i.name == 'foo')[0].value : 0`
which checks if the slice contains exactly one element whose name is `foo` and if so, gets that elements value, 
otherwise returns 0. A similar construct can be done for maps using the same `exists_one` and `filter` macros.

Q: Can this be used to export entire collections of secret or configmap values?

A: Yes. We will support two flavors of entry into the `dynamicSecrets` and `dynamicConfigMaps` collections. One
flavor is the one we've seen above, where `value` is a CEL expression that returns a `string` and the `name` and `key`
fields specify which configMap or secret and what key within that to export.

The other kind of entry is:
```
  operatorSpec:
    dynamicConfigMaps:
      - name: account-data # Name of the destination configMap
        valueMap: "{"key1": "foo", "key2": "bar"}" # CEL expression that returns a map[string]string
```

The `valueMap` fields result can be directly saved as keys + values in the resulting configmap or secret, so the author
of the CEL expression has direct control over what values go into their secret/configMap.

Q: Do we want the syntax in the CEL expression to access the `spec` field to be just `spec.foo` or `resource.spec.foo`?

A: **Open question:** I believe that `resource.spec.foo` is better, because:
1. It disambiguates between the resource and its secrets, since for secret usage we'll support `secrets.myKey`. 
   If we control the top-level name of all the inputs (the main 2 right now being `resource` and `secrets`), then we 
   can ensure there are no collisions because we own the top of the path for each input. If we instead snip that 
   part out (for just resource? or for secrets too?) then we run the risk that a top level 
   field such as `secrets` collides.
2. It puts us in a better position to grow in the future by adding other top-level types going forward, 
   such as `operator` (holding operator configuration?) or `cluster` (holding cluster details?). 
   I don't know if we're ever going to want to add these but if we did, disambiguating between them at the 
   top level helps with clarity.
3. It means we're consistent between `resource.x` and `secrets.y` in terms of the pattern, as opposed to doing `x` 
   and `secrets.y`. I think including secrets as a prefix for secret values which are not on 
   the resource is useful from a security/safety point of view, making it obvious the value you're about to 
   output is a secret. Obviously we only support this in the context of saving to secrets,
   so it's not like there's risk of exposing a secret value as plaintext but still the clarity seems a 
   win for something sensitive like secrets.

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None
