# Rule Interpreter

## Input - Output
A Rule file to check -> A valid Rule File

## Funcitonal Description
## Implementations Problems
Finding a generic Rule DSL that describes a security program across industries it's a hard task.

There is two main issues regarding Rule DSL:
 - Users tend to add their context all the time.
 - Users try to add how to do the check step by step.

A valid Overlord Rule must avoid this two problems. Rule it's a context free communication tool. So can't
describe any context detail or describe "how to do" anything.

The Rule primary focus it's the "What must be".

## Solution architecture
The current DSL proposal describe an state that must be found on a place. Allowing to parametrize the
specific state and place by parameters.

# Rule Parametrizer

## Input - Output
'''
A valid Rule File, The avaliable typed functions -> Parametrized Rule File
'''

## Funcitonal Description
We understand a Paramatrized Rule File as a Valid Rule File that contains the maximum avaliable information
about the Parameters in order to generate a specific sub-set of functions.

Be aware that parametrizing a Rule it's a complete kind of interaction with Overlord. The Rule must remain
untouched so parametrizer will allways add context and never change the rule itself.

In order to ensure Rule's inmutability the Parametrization it's a complete separated solution.

## Implementations Problems
Keep the same Rule across the different environments on different companies. Rules are thought as a communication's
tool. Are interchangeable across contexts and context can't "contaminate" the rule.

No everyone can modify all parameters. There is diferent kind of parameters and users to match, as it's described
in the Actors Documentation.

## Solution architecture
The current set of parameters are:
  - Check Params
  - Target Params
  - Gather Params
  - Params

Check Params change the higher level of abstraction in Overlord. Are reserverd for technical things that
are more relatable to Risk than technology. An example it's the maximun risk score obtained on an specific tool.
Or the minimum version of encription found on an cerificate.

Target Params are more about the places and technologies that you must use on your context.

Gather Params are a Placeholder for things that we must Gather in order to ensure that the current state match
the expected state.

Params are the free context params that you can spcify when execute the rule. Also params that are no critical
and anyone can set.

# Rule Executor

## Input - Output
Parametrized Rule File, The avaliable typed functions -> Program Query

## Funcitonal Description
A program Query it's an incompleted program, contains two kind of information inside:
 - Functions that we know (Rule Funtions and Single candidate Funcitons)
 - Paths to solve (An input type and an output type)

## Implementations Problems
Translate The Rule and their parameters into a valid program query require an iterative specification of find
what current types describe the solution. The Function Finder need to know the specific types to search into.

The Rule Executor will try to find the most near candidate between the current parameters and the avaliable
types provided by the current avaliable functions.

Also need to interpretate the current checkpoints on the program. Sometimes you need to use an specific 
function on your context to provide a correct test. This implies to generate several queries.

## Solution architecture

# Function Finder

## Input - Output
Program Query -> Program Hint or Program

## Funcitonal Description
Most of the time the Function Finder will provide several valid programs. In current state of the solution
we involve de user to choose if there is several propositions. Most of the time more parametrization will 
lead you to one Program solution, so it's spected to execute from Rule Parametrizer and refine an adecuate
solution to your Risks demands.

The secure nature of Overlord tend to involve the human in this loop because the current limitations of
machines understanding security requirements like Trustness of the current check. It's not the same trust
on a configuration check than on a hack tool scan.

## Implementations Problems
Search among all the avaliable functions can lead us into
## Solution architecture

# Rule Result Signer

## Input - Output
Program Result -> Signed Program Result
## Funcitonal Description
## Implementations Problems
## Solution architecture

