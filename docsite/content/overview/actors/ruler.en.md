---
title: "Ruler"
date: 2021-11-05T12:34:51+01:00
draft: true
---

The rule creation process for Overlord is fairly straightforward, but requires an understanding of the rules-specific DSL. The creator of Rules is the Overlord user who understands security problems and wants to add value by designing generic tests on controls (what needs to be verified).

In this case, the community will be its main objective, since the creation of rules adds value within it in the form of consolidation of how to recognize that a control is effectively working as designed, and that both the generic test and the adaptation for a specific environment can be shared and contrasted with the rest of the industry, facilitating the discussion about something factual and verifiable

The value she offers to the community is its specific mastery of a specific technology platform and how to specify and instantiate a generic rule. In many of the examples in this documentation we use a network environment specific rule that the TLS version of a web service must be above 1.3. This is a perfect example of a rule created by a security expert. This rule is an instantiation that belongs to a general recommendation: "Communication channels with customers must be secure."

The task of the rule creator is vital in that he has the technical security knowledge necessary to materialize the tests in generic environments. Understanding as a general environment for example the public cloud and as a specific environment a specific project of the public cloud; that would be the responsibility of the technical managers of each specific project. Although the Security Risk Manager should establish in which cases controls must exist, the rule creator is a necessary ally that says which rules are necessary to verify the existence and effectiveness of the controls in the different environments.

In many cases the rule maker does not have to understand the technical details down to the last detail. So the DSL that defines the rules should not describe the process at the level of what steps are necessary; since it would be very difficult to reuse for different environments, as happens nowadays to the tests of controls written directly in code. Rather than establishing how things are checked, it is more important to know WHAT needs to be checked.

In other cases, technical knowledge is vital for the correct resolution and the rule must specify what programs and procedures must be carried out. In this case the rules provide the ability to add these technical details from a high level of abstraction.

## Parameterizing

For the rules to be truly reusable and useful for the community it is sometimes necessary that they can already be adjusted at a high level of abstraction. To do this, the rules have the goal parameters. These allow the rule builder to tailor the security rule to a certain level of risk based on its risk analysis.

Putting a simple example we can discuss a rule with the following free text (final syntax is not closed yet).

```
Every https service must have TLS version at least ($tls_min)
```

In this case the "$ tls_min" parameter can be set to various values depending on the business or use case. For example, if we have high security requirements it could be the latest version.

Also, if this detail is strategic, it can be forced to be defined, and that it cannot be rewritten later, only in the first abstraction steps. For this, the technical parameters have a prefix that can be put before the "$" symbol; this is tg (from target).

The previous rule would look like this:

```
Every https service must have TLS version at least (tg$tls_min)
```

For more information on the rules, they will be described in their own documentation.
