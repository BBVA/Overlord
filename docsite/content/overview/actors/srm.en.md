---
title: "SRM"
date: 2021-11-03T11:40:50+01:00
draft: false
---

Security Risk Manager focuses on threat modeling, risks, and controls to
mitigate them. However, they usually lose the last mile of verification that
the controls they specified are operational or not, especially in Continuous
Deployment and Infrastructure as Code environments, so Overlord should be one
of the tools that can provide them the most value.

Having a complete view of all controls across the company is a task addressed
by risk management software (ASR&TM: Application Security Requirements and
Threat Modeling, or GRC/IRM). In some of those programs it is possible to
check, using code integrated in the tool, what is the status of some typical
controls.

But for now, no tool offers the integration and ease to be able to check, in
the most automated way possible, all the controls of the organization. The
ability to federate all systems and tools into a single process of actual test
execution can be a critical step to achieve continuous risk management.

Current development trends have posed a real technology challenge for Security
Risk Managers. Automation and technological proliferation have pushed the
demand for security controls and their implementation to the limit.

Although Overlord is not a risk management software, its integration
possibilities with one of it are wide and extensible. Opening the door to a
symbiosis that can change a typical trend of this type of work; from checking
only the most essential at any given time, to being able to have a true,
continuous and updated overview of each change.

Overlord offers you the possibility of managing both a generic catalog of rules
that can be associated with any technical reality, as well as a codebase in
which you can establish relevant risk level verification details.

## Rules

Overlord’s Rules are intended to be easily written by any security expert. It
is not necessary to learn to code in any programming language, nor is it
necessary to learn a long list of commands to establish how things should be
done in each environment. They define the what but not the how, although the
how will be available for approval (test effectiveness) and audit (verification
integrity).

A rule can be specified and adapted to any environment, because it establishes
the security control check objective in an abstract way. In this way it can be
shared and adapted.

On the other hand, both the auxiliary code that allows Overlord to obtain the
intermediate parameters with which to configure the atomic test code (atomic
code), as well as these same atomic tests already developed in the form of
products, open source tools or small internal developments can be integrated
into Overlord independently and decoupled from the rules.

An Overlord system with enough built-in Atomic code may be able to execute a
large set of new rules without the need to add more code. This is because the
Overlord system will try to find its way through the existing code to the test
target.

And if it cannot find their way, it can give you some alternatives so that the
amount of code to add to the system is always the minimum necessary, creating a
backlog of elements to add to maximize the test coverage that can be
prioritized according to the risk of failing to verify the associated controls.

## Atomic Code

As a Risk expert, it is not necessary to know how to code, but it is necessary
to be familiar with the differences between different control verification
processes.

It is not the same level of assurance to check a configuration file than, for
example, to use a network analysis tool to verify in the real system that a
condition is met or to launch said tool from the same network or from another
one.

The Atomic code is always associated with a description of how the information
is being obtained. Using the check parameters (Check Block) you can force
Overlord to use a certain piece of code in an execution; thus ensuring that the
way things are done meets the security requirements of a specific project.
