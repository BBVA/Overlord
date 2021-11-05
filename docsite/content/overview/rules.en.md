---
title: "Rules"
date: 2021-10-26T12:14:39+02:00
draft: true
---

In the case of the execution of a complete rule, we can see the complete abstraction process from the most general rule to the piece of code finally executed.

This is a simple example; a failed rule, an incomplete rule and other secondary scenarios are reserved for other examples. Currently the Overlord PoC is focused on viability so it develops just the "Happy Path".

The complete log is also excluded from this example, because due to its length it would be very complex to explain in detail. In the log there will be the partial results of all the intermediate executions, tools and paths of the discarded dynamic graph.


# A complete rule

<style>
.gdoc-expand {
  margin-bottom: 0rem;
  margin-top: 0rem;
}

.blue{
  background-color: #9acdff;
}
</style>

{{< expand "Recommendation / Generic Rule" >}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.

=> Rule
All ($url) must have (g$tls_version) >= (ch$version)
```

At this point a rule is an abstract high-level recommendation. Its purpose is to recommend a transversal and generic control to all the projects of the company, and it’s enough as a “definition of done” or “security feature story” to check if all the communications of the company comply with this control.


It also enables collaboration between companies/organizations on a very general level as it doesn’t need details about the targets ($url) and goals (ch$version).

In this rule we can observe some general concepts such as the reserved word * All * that indicates that this rule must be true for all the following members. We can also see some words in parentheses; those are the parameters that allow you to instantiate this rule for a specific environment.


**Don’t pay attention to the functional details of the rule, as it progresses through the system you will understand the whole process.**

## Actors

The main producers of these rules will be:

 * Security communities (Ex.: OWASP)
 * Standardization bodies and agencies (Ex.: NIST)
 * Security and infrastructure vendors at large

The main stakeholders in these rules will be:

 * CISOs, CTOs, Compliance&Audit

{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule
All ($url) must have (g$tls_version) >= (ch$version)
== rule id
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5
```

This rule has a SHA256 hash so that it can be referenced in community discussions.

{{< /expand >}}
{{< expand "General Instantiation" "" "">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
```
At this point we are specifying in the rule that in our organization we consider a TLS version of 1.3 or higher to be secure. This block allows you to set the parameters, which, as we saw in the previous point, have a "$" symbol and are also enclosed in parentheses.

If you notice this parameter has the letters ch before the "$" symbol, this tells Overlord that this parameter can only be specified in the "Check" phase; which is where the cross-organizational criteria are established.

Those criteria express the security posture and should be the result of a threat modeling or a compliance constrain


# Actors

These types of decisions are made across an organization. Each organization has its own governance model to select and approve the final criteria on each adjustable goal.


{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779
```

This rule has a SHA256 Hash to be able to refer to it in the ASR-TM (Application Security Requirements and Threat Management) or GRC (Governance, Risk management, and Compliance) of the organization.


{{< /expand >}}
{{< expand "Target" "" "">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
```

In the control testing process it is necessary to add some technical parameters of the specific project system. In this case we are adding one (dns_ip) that is not explicitly in the rule, but can be used to discover the targets.

This is so because in the process of checking security controls, how things are done (the exact mechanism used and the confidence of the results) may be relevant at a certain moment.
Due to this to obtain the data, the dynamic graph built inside Overlord will preferably use the paths it already knows.


# Actors

In this case, both the technical manager and his team are involved in the process. It may be necessary to manage the partial results of Overlord to distinguish from those possibilities of obtaining information those that fit in the current project.

It is from this review process that you can decide what extra parameters should be provided to Overlord to check the control in one way or another.

They can also see in the log how the check has been carried out without having to agree to it previously; so that the Overlord’s principle of segregation of concerns will be preserved.


{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
```

In this case the identifier of the targeting process includes the technical parameter. This will change in case the DNS server associated with this particular process changes. Overlord reinforces transparency whenever possible.

If it is not relevant, we can simply refer to the identifier of the upper level and that way we will not have notifications every time a detail of the environment changes.

{{< /expand >}}
{{< expand "Execution" "" "">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
```
For the execution we need to provide to the rule the rest of the required parameters. In this case, the domain names we want to test.

# Actors

In this case, those in charge of executing the testing code, either manually or automatically. They can include the following profiles:

 * Developers interested in improving code security
 * Auditors who want to check the status of a control
 * Security and risk analysts or operators who want to check the status of a control


{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
== execution id
26b67743880b9bc6cbdbe66d51abd2d8846822bf7325d2e2bef08c611a6023d5
```

Again, the identifier hash includes the execution parameters so that the execution conditions cannot be changed without it changing


{{< /expand >}}
{{< expand "Result" "" "">}}

```
=> Rationale
PCI DSS 4.1 Use strong cryptography and security protocols to safeguard sensitive cardholder data during transmission over open, public networks.  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
== execution id
26b67743880b9bc6cbdbe66d51abd2d8846822bf7325d2e2bef08c611a6023d5
== Execution path
dns_mask_tool
ssl_tools
rule_check
== Final Rule
All [
  ("one.com" as $url) must have ("1.3" as g$tls_version) >= ("1.3" as ch$version),
  ("two.com" as $url) must have ("1.3" as g$tls_version) >= ("1.3" as ch$version),
  ]
== Result
All Rules OK
== Result id
c04a56d97f7e8e6d506e15c9d37ef45849fdb90a06d6b21d8d41b6516adfc67b
```

In the final step, using the dynamic graph built by Overlord, we follow the following steps:

 * First of all we are executing code that, using the "DNSMask" tool, will resolve the urls using a specific DNS server (passed as a parameter in the "Target Block"). 
 * Then the ssl analysis tool is used and the searched property is extracted (tls_version).
 * The rule is expanded with this information.
 * Rule is evaluated

In this way we obtain an "All Rules OK". Our control is working.

Finally, we can see that a new hash has been generated, which includes data from all of the previous steps.

# Actors

The result of the execution is something interesting to all those involved in the process. In this case, the final hash also includes the system time to prevent previous results from being reused. Or to use cached intermediate results if a rule expiration is set (not in this example).

If some of the final (atomic) testing code doesn’t exist yet (dns_mask_tool or ssl_tools in this example), it must be designed, coded and ideally open sourced. Who and how the tools are selected and vetted is a responsibility of the organization, Overlord provides total transparency and audit to automate the trust in the final result.

{{< /expand >}}
