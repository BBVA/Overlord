---
title: "Programmer"
date: 2021-10-29T11:52:55+02:00
draft: false
---

From a programmer's perspective the idea of a universal tool may sound crazy. We are used to dealing with software as a whole that must solve the problem; although microservices are changing this trend slowly.

But the automatic search for paths and ancillary information that Overlord needs so that it can be developed independently is not something that can be "solved" globally. We need a smart, modular and extensible approach. Like an operating system or a compiler.

Overlord will need small pieces of code in order to carry out its information gathering tasks. In these cases, each piece of code must have two parts, gathering and execution.

## Gathering

When we create code for Overlord it needs an interface to know what data we can ask for. This is declared in the atomic code registry metadata; if for example we need our code to provide bread, we will declare it so.

{{< vid "/Overlord/programmer/query.mp4" >}}

Properties have namespace (eg ip.net) simply because Overlord can understand IP in many contexts. For example, it could be an AWS IP, the property in that case could be called something like "ip.aws.com". Don't worry too much about the name now, you don't need a central entity to assign names. Although the names are up to you, looking at the community can be very advantageous.

Now that Overlord knows what you can offer, the system will send you requests with a set of properties, for example a netmask, a domain name and other things. Keep in mind that you will receive all possible properties, not just the ones that interest you. If with these properties you can collect the information, your code will be able to pass the full name of the property you provide, and you have finished the gathering step.

{{< vid "/Overlord/programmer/resolved.mp4" >}}

In the event that you cannot provide that information because there are missing parameters, you will have to return what properties you need to carry out your task. And these can be many, since you can carry out your task in different ways.

{{< vid "/Overlord/programmer/partial.mp4" >}}

A relevant feature with respect to traditional graphs is that the gathering atomic code can manage how to consider an incoming request. It is not expected to always provide the same answer, this can be based on multiple factors, in fact that is why it is "dynamic". The code is expected to always have the same output, but it must be flexible on input. This allows us to have great flexibility for cases we do not yet know about.

Also with the same parameters, but different user, we can require different things. If the query, for example, is made by the system administrator, we do not require any additional permission. If, on the other hand, an anonymous user makes the same query, we can require an additional authentication parameter to the system.

## Executing

When Overlord invokes the execution code, it will receive, again and symmetrically to the previous case, all the parameters it has at that time. Among these, surely there is a way to get the information. If you can get the information in more than one way, it is up to you to decide how it will be executed to achieve your goal.

{{< vid "/Overlord/programmer/execution.mp4" >}}

Once the code is executed, it will deliver the result to Overlord. It's that simple (probably through an exchange API to be decided, but always as isolated as possible to don’t introduce dependencies in the atomic code)

# Putting the process in perspective

This process seems pretty straightforward, but using these small pieces of code (that can themselves invoque more complex applications like SAST), Overlord is able to compose highly complex solutions.

# Added Value

Through its checks Overlord can be a general checking framework for your own code. It does not have to be used only in a business environment, nor is it a tool that requires large amounts of infrastructure. We want to keep it simple so it can be used standalone and offer a comprehensive testing capability.

We know that there are "Property based testing" frameworks, and maybe you are very comfortable with them. They may even offer better abstractions for you. But keep two things in mind:

  * When you write code for Overlord you contribute to a community beyond your test case.
  * If that framework is so good, you can integrate it into Overlord! and thus contribute to a larger community.

We try to keep the integration effort as low as possible, adding value at a low cost.

In addition to checking security controls, you can use the Overlord abstractions as an independent tool. We believe that the dynamic graph can be adapted to a huge number of cases, specially in high uncertainty environments. 
