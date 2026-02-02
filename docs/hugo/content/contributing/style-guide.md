---
title: Documentation Style Guide
linktitle: Style Guide
weight: 5
cascade:
- type: docs
- render: always
description: "Guidelines for writing consistent ASO documentation, for both human and AI authors"
---

This style guide ensures consistency across all Azure Service Operator documentation. Whether you're a human author or an AI assistant contributing to our docs, following these guidelines helps maintain a unified voice and makes our documentation more accessible to users.

## Audience and Tone

Our documentation targets experienced Kubernetes and Azure practitioners who understand cloud-native concepts but may be new to ASO. Write for an audience that knows what a CRD is, but may not know how ASO implements one.

- **Be professional but approachable.** Write as if explaining to a knowledgeable colleague—informed and helpful, but never condescending or overly casual.
- **Assume intermediate Kubernetes knowledge.** Readers understand pods, secrets, namespaces, and `kubectl`, so don't explain these basics.
- **Provide context for ASO-specific concepts.** Explain how ASO works, including conditions, ownership, reconcile policies, and credential scopes.
- **Focus on practical outcomes.** Emphasise what users can accomplish rather than abstract descriptions of features.
- **Use an inclusive "we" and "you".** Address readers directly with "you" for instructions and use "we" when discussing ASO team recommendations.

## Document Structure

Each documentation page should have a clear purpose and logical flow. Start with the essential information and progressively add detail for those who need it.

- **Begin with a brief overview.** The first paragraph should explain what the page covers and why readers should care.
- **Use descriptive headings.** Headings should tell readers what they'll learn, not just label sections. Prefer "How to provide secrets to Azure" over just "Secrets".
- **Organise content hierarchically.** Start with the most common use cases, then cover edge cases and advanced scenarios.
- **Keep paragraphs focused.** Each paragraph should address a single concept in 2-4 sentences.
- **End with next steps or related links.** Point readers to related documentation or logical next actions.

## Writing Style

Clear, direct writing helps users find answers quickly. Every sentence should serve a purpose.

- **Lead with the action or outcome.** Write "Set the annotation to skip deletion" rather than "There is an annotation that can be set to skip deletion".
- **Use present tense and active voice.** Write "ASO creates the resource" rather than "The resource will be created by ASO".
- **Be precise about requirements.** Distinguish between "must", "should", and "can" to indicate what is required versus recommended versus optional.
- **Avoid jargon without explanation.** If you must use terms like "reconciliation" or "hub version", ensure they're explained on first use or link to a glossary.
- **Keep sentences concise but complete.** Aim for clarity over brevity—don't sacrifice understanding for fewer words.

## Code Examples

Code samples are central to ASO documentation. They should be realistic, complete, and immediately usable.

- **Provide complete, working examples.** Users should be able to copy-paste YAML directly into their cluster with minimal modification.
- **Use realistic resource names and values.** Prefer names like `aso-sample-rg` or `sampleredis` over abstract placeholders like `foo` or `example`.
- **Include comments for non-obvious fields.** Annotate YAML with comments explaining why particular values were chosen.
- **Show both the YAML and the kubectl command.** When demonstrating resource creation, show the YAML structure and how to apply or check it.
- **Specify the API version explicitly.** Always include the full `apiVersion` in YAML samples so readers know exactly which version is being demonstrated.

## YAML Formatting

Consistent YAML formatting improves readability and reduces confusion.

```yaml
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
spec:
  location: westus2
```

- **Use 2-space indentation.** This is the Kubernetes convention.
- **Place `apiVersion` and `kind` first.** Follow with `metadata`, then `spec`, then `status` if shown.
- **Quote strings only when necessary.** YAML strings don't need quotes unless they contain special characters or could be misinterpreted.
- **Redact sensitive values.** Use placeholder patterns like `00000000-0000-0000-0000-000000000000` for subscription IDs or `$AZURE_CLIENT_SECRET` for secrets.

## Terminal Commands

Command-line examples should be easy to follow and adapt.

- **Use bash syntax by default.** Provide PowerShell alternatives in tabs when commands differ significantly.
- **Prefix commands with `$`.** This distinguishes commands from their output and follows common convention.
- **Show expected output where helpful.** Include truncated or representative output to help users verify their commands worked.
- **Explain long or complex commands.** Break down multi-part commands or explain significant flags.

```bash
$ kubectl get resourcegroups.resources.azure.com 
NAME            READY     SEVERITY   REASON          MESSAGE
aso-sample-rg   True
```

## Formatting Conventions

Consistent formatting makes documentation scannable and professional.

- **Use bold for UI elements and emphasis.** Write "Click **Create**" for UI actions.
- **Use backticks for code elements inline.** Field names, resource kinds, file paths, and commands should appear as `spec.owner.name`, `ResourceGroup`, or `kubectl apply`.
- **Use admonitions for important callouts.** Employ Note, Warning, and Tip blocks to highlight critical information without breaking the reading flow.
- **Format links with descriptive text.** Write "[authentication documentation](link)" rather than "click [here](link)".
- **Use tables for structured comparisons.** Present options, environment variables, or feature comparisons in tables for easy scanning.

## Admonitions and Callouts

Hugo shortcodes provide formatted callout boxes for important information. Use them to draw attention without disrupting the main content flow.

```
{{% alert title="Note" %}}
Content for a general note or tip.
{{% /alert %}}

{{% alert title="Warning" color="warning" %}}
Content for a warning about potential issues.
{{% /alert %}}
```

- **Use Note for supplementary information.** Tips, additional context, or "good to know" details.
- **Use Warning for potential pitfalls.** Situations where users might encounter problems or data loss.
- **Keep callout content brief.** If the callout needs more than 2-3 sentences, consider making it part of the main text.
- **Don't overuse callouts.** Too many callouts dilute their impact. Reserve them for genuinely important information.

## Cross-References and Links

Effective linking helps users navigate related content and find additional detail.

- **Use relative links for internal documentation.** Use Hugo's `relref` shortcode for links within the documentation: `[authentication]( {{< relref "authentication" >}} )`.
- **Link to specific sections when relevant.** Deep-link to specific headings rather than just the page: `[credential scope]( {{< relref "authentication#credential-scope" >}} )`.
- **Provide context for external links.** Explain what readers will find at external URLs.
- **Link to samples in the repository.** Reference working samples in `v2/samples/` to give users complete examples.
- **Keep link text descriptive.** The linked text should describe the destination, not generic phrases like "click here" or "this page".

## FAQ and Troubleshooting Content

Question-based documentation helps users find answers to specific problems quickly.

- **Use the actual question as the heading.** Write "How can I protect against accidentally deleting an important resource?" rather than "Deletion Protection".
- **Start with the direct answer.** Don't bury the solution in background information.
- **Provide numbered steps for multi-step solutions.** Complex fixes should be broken into clear, sequential steps.
- **Link to related issues or discussions.** Reference GitHub issues where users can find more context or report problems.
- **Include error messages verbatim.** When documenting error conditions, show the exact error text users will see.

## Design Documents and ADRs

Architecture Decision Records follow a specific structure to capture the reasoning behind significant decisions.

- **Start with Context.** Explain the problem or situation that prompted the decision.
- **Present the Decision clearly.** State what was decided and how it will be implemented.
- **Explain the reasoning.** Include the trade-offs considered and why this approach was chosen.
- **Use code examples to illustrate.** Show how the decision manifests in actual code or configuration.
- **Date the document.** ADR filenames include the date (e.g., `ADR-2022-01-Reconciler-Extensions.md`).

## Version-Specific Content

ASO supports multiple API versions for resources, so documentation must be clear about which versions apply.

- **Specify API versions in examples.** Always include the complete `apiVersion` field in YAML samples.
- **Note when features are version-specific.** If a feature is only available in certain versions, state this clearly.
- **Link to the reference documentation.** Point users to the generated reference docs for complete schema information.
- **Use version-neutral language when possible.** Write instructions that work across versions unless discussing version-specific behaviour.

## Accessibility and Inclusivity

Documentation should be accessible to all users.

- **Use descriptive alt text for images.** When including diagrams or screenshots, describe what they show.
- **Don't rely on colour alone.** If using colour to convey meaning, also use text or symbols.
- **Use plain language.** Avoid idioms, metaphors, or cultural references that may not translate well.
- **Structure content for screen readers.** Use proper heading hierarchy and meaningful link text.
