# Overlord components overview

Overlord need to manage a wide and complicated problem. In order to achieve this we split the 
big problem into smaller pieces.

On a nutshell, from a linear big problem perspecive, the overlord components are:

~~~mermaid

graph LR
    A(Rule Interpreter) --> B(Rule Parametrizer)
    B --> C(Rule Executor)
    C --> D(Function types Finder)
    D --> E(Function Finder)
    E --> F(Component Executor)
    F --> G(Rule Result Signer)

~~~

Detailed interactions are provided next.

TODO: Input Outputs are just conceptual proposals.

# Rule Interpreter

## Input - Output
A Rule file to check -> A valid Rule File

## Funcitonal Description
The Rule DSL must be the simplest it can be. It's objetive is comunicate about an specific control at
the highest level possible.

Must define "What must be" regarding one or more controls already in place. And cannot allow the user
to specify the little details of "how to" check if control it's up and running as expected.
This requirement it's by desing, no for technical reasons. There is a lot of good reasons, but the 
biggest one is because companies can share the what but no the how they systems are. Another good
reason is about separation of concerns; the one that write the rules maybe don't know how to check it
in every technical context.

This piece of software check all the syntax details, offering hints about what are you writing.

## Implementations Problems
Finding a generic Rule DSL that describes a security program across industries it's a hard task.

There is two main issues regarding Rule DSL:
 - Users tend to add their context all the time.
 - Users try to add how to do the check step by step.

A valid Overlord Rule must avoid this two problems. Rule it's a context free communication tool. So it can't
describe any context detail or describe "how to do" anything.

The Rule primary focus it's the "What must be". And the most important property of the rules is that they 
are inmutable. If anyone can change the rule thru the process then the original rule, and their information,
will be lost.

## Solution architecture
The current DSL proposal describe an state that must be found on a control instance, by example:

```
 All $(ip@gather) the port $(port@rule) must be on $(closed_ports@gather).
```

On this example we can see 3 parameters, two of them (ip and closed_ports) must be gathered from your 
current systems to ensure they refers to the control instance under testing. The other one (port) must be defined before executión, and 
ideally, come from a Thread Model session and control assessment and definition.


# Rule Parametrizer

## Input - Output
A valid Rule File, The avaliable Typed Functions -> Parametrized Rule File

## Funcitonal Description
We can see a Paramatrized Rule File as a Valid Rule File that contains the maximum avaliable information
about the Parameters in order to generate a specific sub-set of functions.

Be aware that parametrizing a Rule it's a complete kind of interaction with Overlord. You can specify a
parameter whitout execute it.

The Rule must remain untouched so parametrizer will allways add context and never change the rule itself.
In order to ensure Rule's inmutability the Parametrization it's a complete separated solution.

## Implementations Problems
Keep the same Rule across the different environments on different organizations and even projects is not easy. Rules are intended to be a communication's
tool. They are valid across contexts and context can't "contaminate" the rule.

Not everyone can modify every parameter. There are different kinds of parameters and users to match, as it's described
in the [Actors Documentation](https://bbva.github.io/Overlord/overview/actors/).

## Solution architecture
The current set of parameters defined in Overlord are:
  - Check Params
  - Target Params
  - Gather Params
  - Params

Check Params change the higher level of abstraction in Overlord. Are reserverd for technical things that
are more relatable to Risk than technology and express the organization risk appetite/aversion or a compliance requirement. An example it's the maximun risk score obtained on an specific tool. Or the minimum version of encription found on an cerificate.

Target Params are related to the technical environemnt and technologies that you use on your context.

Gather Params are a placeholder for things that we must "Gather" during execution, in order to perform the final check or to obtain auxiliary data to reach that final check.

Params are the risk-agnostic and context-free params that you can spcify when you execute the rule. Also params that are no critical
and anyone can set.

# Function Types

## Input - Output
Parametrized Rule File, The avaliable Typed Functions (built-in and contributed) -> Program orchestration scaffold

## Funcitonal Description
A program scaffold it's an incompleted program, contains two kinds of informations inside:
 - Functions that we know (Rule Funtions and Single Candidate Functions)
 - Paths to solve (An input type and an output type), candidate function type requirements

## Implementations Problems
Translate the Rule and their parameters into a valid program query require an iterative specification to find
what current types describe the solution. The Function Finder needs to know the specific types to search into.

The Rule Executor will try to find the nearest candidate between the current parameters and the avaliable
types provided by the current avaliable functions.

It also needs to interpretate additional constraints on the program. Sometimes you need to use an specific 
function (a technical implementation/solution) on your context to provide a correct test. 

## Solution architecture

The task is to retrieve all the available information about the Rule in order to assemble the correct question
for the function finder.

It will use other Overlord components to provide an interactive Rule builder for each Overlord user.

To ensure a viable result of the function finder, this software must reduce the query space at maximum, providing a
filtered collection of functions to the function finder.

# Function Finder

## Input - Output
Program Orchestration Scaffold, filtered collection of functions, all available functions -> Program Gap Hint (function's Types to satisfy) or the Final Program

## Funcitonal Description
Usually the Function Finder will provide several valid candidate programs. In the current state of the solution
we involve the user to choose if there are several proposals. Most of the time more parametrization will 
lead you to a single Program solution, so it's expected to be executed from Rule Parametrizer and iterated to refine an adecuate
solution to your Risks requirements.

In case there's no candidate programs because the available functions can't be combined to build the required program, Function Finder will provide the expected types for the missing function 

## Implementations Problems
Search through all avaliable functions can lead us into a complete NP problem. If you think about it
it's the same problem as the salesman travel problem.

## Solution architecture
Using the state of the art software synthesis techniques you can build a valid program using an infinite-like combination of 
functions.

There are some papers discussing this problem, like [this one](https://cseweb.ucsd.edu/~npolikarpova/publications/icfp22.pdf).
Please be aware that to read this paper will requiere some advanced knowledge about program synthesis, e-graphs, Haskell and type systems.
Fortunately it's not requiered to understand and use the Overlord solution, this is the engine under the hood.

# Component Executor

## Input - Output
Final Program -> Program Result

## Functional Description
This piece of software has its own repository well explained.

TODO: set the link

In a nutshell must be capable to execute any tool and integrate results between diferent pieces of software in a cordinated way.

## Implementation problems

TODO: set the link

Mix and max several piecen of software can be a challenge problem to solve.

## Solution architecture

TODO: set the link

Thanks to NixOS and their Flakes feature you can use software from a huge variety of sources and make it to work as a single piece
of software.

# Rule Results Signer

## Input - Output
Program Result -> Signed Program Result

## Funcitonal Description
This is the piece of software that ensures integrity from the original Rule and each step of the process to the end result, and that every parameter was set by the right Actor.

It validates the final result at the end of the process, but Overlord will use it to compute, 
sign and verify thru all the rule's lifecycle.

You can learn how it works in the next [documentation](https://bbva.github.io/Overlord/en/overview/rules/).

## Implementations Problems
There is no big deal, just segregation of functions and good crypo work.

## Solution architecture
Simple check and sign program.

