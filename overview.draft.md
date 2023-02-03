# Overlord components overview

Overlord need to manage a wide and complicated problem. In order to achieve this we split the 
big problem into smaller pieces.

On a nutshell, from a linear big problem perspecive, the overlord components are:

~~~mermaid

graph LR
    A(Rule Interpreter) --> B(Rule Parametrizer)
    B --> C(Rule Executor)
    C --> D(Function Finder)
    D --> E(Component executor)
    E --> F(Rule Result Signer)

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
specify the little details of "how to" check if control it's there.
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

A valid Overlord Rule must avoid this two problems. Rule it's a context free communication tool. So can't
describe any context detail or describe "how to do" anything.

The Rule primary focus it's the "What must be". And the most important property of the rules is they 
are inmutable. If anyone can change the rule thru the process the original rule, and their information,\
will be lost.

## Solution architecture
The current DSL proposal describe an state that must be found on a place, by example:

> All $(ip@gather) the port $(port@rule) must be on $(closed_ports@gather).

On this example we can see 3 parameters, two of them (ip and closed_ports) must be gathered from your 
current systems to ensure the control. The other one (port) must be defined before executión, and 
ideally, come from Thread Model session and control definitión.


# Rule Parametrizer

## Input - Output
A valid Rule File, The avaliable typed functions -> Parametrized Rule File

## Funcitonal Description
We understand a Paramatrized Rule File as a Valid Rule File that contains the maximum avaliable information
about the Parameters in order to generate a specific sub-set of functions.

Be aware that parametrizing a Rule it's a complete kind of interaction with Overlord. You can specify a
parameter whitout execute it.

The Rule must remain untouched so parametrizer will allways add context and never change the rule itself.
In order to ensure Rule's inmutability the Parametrization it's a complete separated solution.

## Implementations Problems
Keep the same Rule across the different environments on different organizations and even projects. Rules are thought as a communication's
tool. They are interchangeable across contexts and context can't "contaminate" the rule.

Not everyone can modify every parameter. There is diferent kind of parameters and users to match, as it's described
in the [Actors Documentation](https://bbva.github.io/Overlord/overview/actors/).

## Solution architecture
The current set of parameters are:
  - Check Params
  - Target Params
  - Gather Params
  - Params

Check Params change the higher level of abstraction in Overlord. Are reserverd for technical things that
are more relatable to Risk than technology and express the organization risk appetite/aversion or a compliance requirement. An example it's the maximun risk score obtained on an specific tool. Or the minimum version of encription found on an cerificate.

Target Params are related to the technical environemnt and technologies that you must use on your context.

Gather Params are a placeholder for things that we must Gather (TODO complete).

Params are the risk-agnostic and context-free params that you can spcify when you execute the rule. Also params that are no critical
and anyone can set.

# Rule to Function Types finder (TODO, change name  hints: precompiler, scafolding, preprocesing, ...)

## Input - Output
Parametrized Rule File, The avaliable typed functions (built-in Rule Funtions and contributed) -> Program orchestration scafold

## Funcitonal Description
A program Query it's an incompleted program, contains two kind of informations inside:
 - Functions that we know (Rule Funtions and Single candidate Functions)
 - Paths to solve (An input type and an output type), candidate function type requirements

## Implementations Problems
Translate The Rule and their parameters into a valid program query require an iterative specification to find
what current types describe the solution. The Function Finder needs to know the specific types to search into.

The Rule Executor will try to find the nearest candidate between the current parameters and the avaliable
types provided by the current avaliable functions.

It also needs to interpretate additional constraints on the program. Sometimes you need to use an specific 
function (a technical implementation/solution) on your context to provide a correct test. This implies to generate several queries.

## Solution architecture

The task is to retrieve all the available information about the Rule in order to assemble the correct question
for the function finder.

It will uses the other pieces of software to provide an interactive Rule builder to the final user.

To grant a viable result of the function finder, this software must reduce the query space at maximum. Keep in
mind that this problem grows quickly to a complete NP problem.

# Function Finder

## Input - Output
Program Query, filtered collection of functions, all available functions -> Program Gap (Functions or Types) Hint or Final Program

## Funcitonal Description
Most of the time the Function Finder will provide several valid composed programs. In the current state of the solution
we involve the user to choose if there are several proposals. Most of the time more parametrization will 
lead you to a single Program solution, so it's expected to be executed from Rule Parametrizer and iterated to refine an adecuate
solution to your Risks requirements.

TODO: Review paragraph (The secure nature of Overlord tend to involve the human in the loop because the current limitations of
machines understanding security requirements like Trustness of the current check. It's not the same trust
on a configuration check than on a hack tool scan.)

## Implementations Problems
Search through all avaliable functions can lead us into a complete NP problem (TODO: check literature). If you think about it
it's the same problem as the salesman travel problem. TODO (link to the papers)

## Solution architecture
Using the state of art software sintesis techniques you can compute a valid program using an infinite-like ammount of 
functions.

There is some papers talking about this, like [this one](https://cseweb.ucsd.edu/~npolikarpova/publications/icfp22.pdf).  
Please be aware that this paper requieres some knowledge about functional programming and Theory of computation.

# Rule Result Signer

## Input - Output
Program Result -> Signed Program Result

## Funcitonal Description
This is the piece of software that ensures integrity using hashes. It validates at the end, but you can use it to compute,
sign and verify thru all the rule's livecycle.

## Implementations Problems
There is no big deal, just good cripo work.

## Solution architecture
Simple check and sign program, the more options the best.

