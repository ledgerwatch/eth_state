# Ethereum 1x Definition (part 2)

## Method (continuation from part 1)
In part 1, we define agents, their contributions to the Ethereum system, and their challenges. This part deals with causes for these challenges, and the solution that have been proposed for these causes.

### Causes
We try to understand the main causes of each challenge. Where it adds to understanding, we talk about sub-causes
```graphviz
digraph causes {
rankdir=LR;
cause [shape=hexagon style=filled fillcolor=salmon];
sub_cause [label="sub-cause" shape=hexagon style=filled fillcolor=salmon];
challenge [shape=octagon style=filled fillcolor=plum];
challenge -> cause [dir=back];
cause -> sub_cause [dir=back];
}
```
### Solutions
Proposed solutions should be targeted at the causes but may have side-effects (some solutions may be causes for other challenges, though perhaps less critical than the ones they are trying to address). These side effects are shown as dotted lines.
```graphviz
digraph solutions {
rankdir=LR;
solution [shape=box style=filled fillcolor=lightblue];
cause [shape=hexagon style=filled fillcolor=salmon];
challenge [shape=octagon style=filled fillcolor=plum]
cause -> solution [dir=back];
challenge -> solution [dir=back style=dotted]
}
```

## Summary of challenges
Here we summarise the challenges we have identified in the part 1.
1. Long time to sync a new node
2. Cost of storage devices
3. High internet traffic
4. Complex DevOps to run nodes
5. Conflict between updating rules of Ethereum and maintaining backwards compatibility
6. Conflict between updating the rules of Ethereum and improving client implementations and 
7. Difficulty of writing new implementations
8. Limited transaction throughput
9. Assessing safety of transactions
10. Estimating cost of a transaction

## Description of challenges

### Long time to sync a new node
When a new computer joins the Ethereum network, it needs to obtain the current state before it can start processing arbitrary transactions. The current state can be obtained in two main ways.
Firstly, the current state can be computed from the Genesis block, by applying (executing all transactions within) all the blocks in sequence. This method is now computationally hard due to the number of blocks (and transactions in them) to apply, and due to the size of the state.
Secondly, the current state can be downloaded from the peers that already computed (or downloaded) it. This method is not as hard computationally, but it requires a lot of communication with the peers, due to the size and the structure of the state.

### Cost of storage devices
One can divide the persistent data managed by an Ethereum client into two main categories. First category is the data that rarely changes after being written. This category includes the content of downloaded block headers, block bodies, transaction receipts. Since chain reorganisations are relatively rare and "shallow" (meaning that not many blocks are getting reverted), one can store these data in an append-only form. This makes it feasible to store such data in a lower-cost devices such as HDD (Hard Disk Drives), or even tape (which is the cheapest storage device at the moment of writing). In fact, the "Freezer" implementation in go-ethereum 1.9 is based on these observations.
Second category is the data that is expected to change constantly. This includes mostly the active state of Ethereum, but also some indices for the append-only data. The Ethereum state is special in this category, because not only it constantly changes, but practically any part of it needs to be available for access at a short notice. To achieve maximum performance, one would choose storage devices with lower latency (SSD instead of HDD), and/or with a higher capacity for parallel reads (NVM instead of SSD). As the state grows, cost of these devices may grow super-linearly.

### High internet traffic
Ethereum nodes constantly communicate with their peers in the network. In any peer-to-peer distribution network, being a "fair" participant requires not only satisfy one's own needs for data but also provide data for the others. Depending on how many new nodes are joining the network, and what is current transaction activity, the internet traffic usage can be unpredictable.

### Complex DevOps to run a node
The art of DevOps is how to run the infrastructure in a way that causes the least disruption, is cost efficient, and satisfy the functional requirements of the individual, business, or any organisation owning that infrastructure. Most of the complexities of running Ethereum nodes are usually related to the management of persistent data. Here are some operations that can be challenging, considering the amount of data and its dynamic character:

* creating replicas of data to bootstrap new nodes
* maintaining backups
* pruning old histories
* recovering from database corruptions
* upgrading to new database formats
* representing data in a format convenient for analysis

### Conflict between updating rules of Ethereum and maintaining backwards compatibility
Often, the rules of Ethereum need to be updated because:

* Implement a missing but very useful feature
* Fix a flaw in the original design
* Adjust to a change in the technological landscape

Smart contracts that have been deployed using old rules, may come to rely on those rules, explicitly or implicitly. When such contracts are still in frequent use, the challenge is to balance their need to exist in their original form (often, code improvements and redeployments solve the issue), and the need to update the rules.

### Conflict between  updating rules of Ethereum and improving client implementations
For reasons mentioned earlier, it is often necessary to update rules of Ethereum. When this happens, there is mandatory work for every Ethereum client implementation team who want their implementation to include the support for the update and stay relevant in the ecosystem. This work may distract from other, equally important work on constantly improving and optimising the implementations to keep up with the growing demands of the system. The balance between these two types of work can be difficult to strike. Ideally, there should be no conflict and the two efforts should be aligned.

### Difficulty of writing new implementations
Sice Ethereum's launch in August 2015, many teams have attempted to produce working implementation of Ethereum. Initially, there was a great variety of such projects. However, as the system's usage grew, most implementations became non-viable or too difficult to maintain, although the number of rule changes is still relatively modest.

### Limited transaction throughput
For Dapp developers, it is important to understand limitations of the Ethereum as a platform. One of the main limitation is the transaction throughput. Not only it puts limits on how many users may interact with an application, but it also puts popular application in competition with each other, since this limitation is global, and not per-application. It is currently believed that the maximum theorerical throughput of Ethereum is around 50 transactions per second (assuming simple Ether transferring transaction, 10M block gas limit, and 15 seconds block time). Applications that require more throughput, usually try to utilise so-called "Layer 2" solutions, however, these ase still not very mature at present, and may require lifting other limits (for example, maximum allowed transaction size, or maximum allowed size of contract byte code).

### Assessing safety of transactions
For simple Ether transferring transaction, safety usually comes down to how well the private key is protected, and how secure is the generation of the digital signature (i.e. that it does not leak information about the private key by generating bad random numbers).
For transactions sent to smart contracts, safety assessment is harder, because the code of the smart contract may trigger undesirable actions if it is flawed or misused, intentionally or not. Often, Dapp developers attempt to assure the users of the safety of their contracts by conduction or commissioning security audits, offering bug bounties, or attempting to formally prove some safety properties.
It is believed that design of the execution engine (EVM in case of Ethereum) can be done in a way to make security analysis and formal verfication of smart contracts significantly easier and cheaper.
Another important consideration for safety is the probability of a transaction being reverted due to the chain reorganisations. Unintentional chain reorganisations happen all the time, but their depth is usually quite low. Therefore, the rule of "wait X number of blocks and consider transaction final" works quite well for dealing with unintentional reorganisations. Intentional reorganisations are often referred to as "51% attacks", though the term may be inaccurate. When dealing with transactions of large value, the safety assessment need to include the risk of intentional reorganisations.

### Estimating cost of a transaction
Gas price is a means of auctioning the limited space within the Ethereum blocks to the transaction senders. Such auctons work well when the space in the blocks is not fully utilised. However, when the blocks are becoming consistently full, higher gas price volatility is observed, and the transaction senders tend to overpay for gas, realising that they are competing with other senders. High volatility of gas prices makes the estimation of correct price to pay more complex. 

## Prioritising of challenges
Although any prioritisation of challenges would appear subjective, the approach is to compare the impact of challenges becomes overwhelming. For example, if the challenge **Long time to sync a new node** becomes overwhelming, and the sync time keeps growing, we can predict that in the future, the network will become difficult or impossible to join for the new operators. Although this will not immediately cause the system to fall, it will make it less resilient with some node operators disappearing.

### Long time to sync a new node
It this challenge becomes overwhelming, and the sync time keeps growing, we can predict that in the future, the network will become difficult or impossible to join for the new operators. Although this will not immediately cause the system to fall, it will make it less resilient with some node operators disappearing. Even insentivising such node operators is unlikely to stop this tendency.

### Cost of storage devices
If this challenge becomes overwhelming, and the cost of storage devices keeps growing, the number of node operators may reduce, unless there is a way of compensating node operators.

### High internet traffic
If this challenge becomes overwhelming, and the demands of internet traffic grows, it is unclear what consequences there will be, because the cost of internet traffic appears to be insignificant compared to other costs involved in running an Ethereum node.

### Complex DevOps to run a node
If this challenge becomes overwhelming, and the running an Ethereum node becomes operationally challenging, it is not clear what consequences will be. There are already mitigations to this challenge, in the form of pre-packages solutions to run node (like DappNode), service providers who specialise in hosting Ethereum node, and some cloud operators potentially offering Ethereum nodes "as a service".

### Conflict between updating rules of Ethereum and maintaining backwards compatibility
If this challenge becomes overwheleming, there are at least two possible scenarios:

1. Insufficient change of rules
2. Too many changes breaking backwards compatibility

Both of these require further analysis. Possible consequences of insufficient change incude:

 * eventual imbalance of gas schedule and DOS attacks on the system
 * inability to respond to highly desired feature requests
 * inability to fix design flaws and allow performance breakthroughs

Possible consequences of too many changes breaking backwards compatibility include:

 * Some popular contracts become obsolete or require redesign and redeployment

### Conflict between updating rules of Ethereum and improving client implementations
If this challenge becomes overwhelming, there are at least two scenarios:

1. Client implementations struggle to keep up with the growth of the system, while most of the time is spent on updating the rules.
2. Updating rules slows down while core developers are busy with optimisations.

In the first scenario, one potential consequence is the reduction of the block gas limit to protect the implementations from breaking. After that, either existing implementations catch up, new implementation appears, or the system fails.
In the second scenario, the downside is prolonged "feature freeze".

### Difficulty of writing new implementations
If this challenge becomes overwhelming, the consequence
is the reduction of viable implementations, and increased reliance on the resposible teams for the decision making about the rules of Ethereum.

### Limited transaction throughput
If this challenge becomes overwhelming, that can mean one of two situations:
1. Transaction throughput starts to decline
2. Transaction throughput stays the same despite an expectation that it will increase.

First situation is more serious, and it would mean the reduction in the block gas limit as response to DOS attack or low performance of client implementations.
Second situation is less serious, but lead to "disappointment" in Ethereum as a smart contract platform. Currently, there seems to be an expectation that greatly increased transaction throughput will come from Ethereum 2, and Ethereum 1 can at best deliver x10 improvement. As long as such expectation remains justified, the status quo appears to be sustainable.

### Assessing safety of transactions
This challenge can become overwhelming in two ways:

1. Researchers in formal verification and security of smart contracts come to general conclusion that it does not make economic sense to try to advance this without fundamental change of the EVM design
2. Intentional reorganisations become possible and sometimes happen in practice

If the challenge becomes overwhelming in the first way, it would lead to potentially large use cases only being deployed on Ethereum for limited use (e.g. limitation no more than $10 million in a smart contract that was suggested straight after "The DAO incident").
The second way is more immediately alarming, because it would affect even people, companies and organisation that do not utilise sophisticated smart contracts, but rely on simple Ether transfer transactions.

### Estimating cost of a transaction
If this challenge becomes overwhelming, it would mean that the gas prices are consistenty very volatile and it is hard to estimate a required fee to ensure that the transactions get confirmed in timely manner. The worst affected would be people, companies, and organisations, whole operations depend on sending and confirming large number of transactions, for example, exchanges (for deposits and withdrawals, and for trading in case of decentralised exchanges), mining pools (for payouts), popular Dapps. When gas prices are volatile, it becomes more economically attractive to perform transaction front-running and use mechanisms like GasToken to "accumulate" gas at lower prices to release it at higher prices.

## Summary of causes
Here we summarise the causes for the challenges. Most of them are technological, though some could be viewed as organisational.
1. Large (and growing) state
2. Snapshot sync algorithm implementation deficiencies
3. Increasingly intricate data management that an Ethereum node needs to perform
4. Functional coupling on implementation level
5. Functional coupling on protocol level
6. Large semantic gap between EVM and many dapps
7. EVM design favours expressiveness over tractability
8. Transaction fee market behaves sub-optimally when blocks are at the block gas limit (which is currently almost always).
9. Possibility of long chain re-orgs.
10. Over-reliance of spontaneous contributions.

## Description of causes


### Large (and growing) state
Ethereum's state is a data structure that needs to be implicitly constructed, stored and accessed in order to execute arbitrary transactions. This is because a transaction may theoretically access any item in the current state. The state size grew beyond the capacity of RAM (Random Access Memory) on average computers some time in 2017. After that point, RAM could only cache certain portions of the state, whereas the entire state resides on persistent storage devices. Any caching strategy apart from keeping a random portion of the accessible state, would be vulnerable to attack. Therefore, assuming that random caching strategy, the cache hit ratio would very close to the ratio of size of the cache to the size of the entire state. And cache miss would mean accessing devices with much higher latency.
Because of the further growth of the state, it became impractical to use HDD (Hard Disk Drives - storage devices with mechanically spinning disks) for storage of the state, due to the high latency of the access. Even SDD (Solid State Drives) are on the edge of being appropriate. Devices such as NVM (Non Volatile Memory) are now required to ensure good performance. However, such devices are still relatively expensive and their price proportional to capacity ($/Gb) is highly non linear after a certain point.
State size also places a significant burden on the new participants in the Ethereum network. Most popular way of joining the network at this moment is so-called "snapshot synchronisation". It is the process in which the new-joiner downloads the entire state from the existing peers. The sheer size of the state puts a high demand on the bandwidth quality. Dealing with the network latencies requires sophisticated algorithms for downloading. And the ever-changing nature of the state (it keeps changing during the download) either puts snapshot sync at odds with state history pruning, or requires even more sophisticated algorithms for the downloading of the state.

### Snapshot sync algorithm implementation deficiencies
Snapshot synchronisation (briefly described in the second ) can be implemented in a variety of ways. The dominant algorithm used at the moment is called "fast sync". It has been designed in the circumstances when the state size was still reasonably small. The main advantage of this algorithm is its simplicity. However, implementations may also vary in efficiency. For example, they might be sensitive to network latencies (how much slower algorithm performs with an increase in latency), or require a lot of overhead traffic (how much more data algorithm transmits with an increase of state size).

### Increasingly intricate data management that an Ethereum node needs to perform
The hardest part of a main-net capable implementation seems to be the data management. Ethereum node needs to transmit, process and store large amounts of data. To do so efficiently, it needs non-trivial techniques, like data caching strategies, partitioning data by usage type. Such techniques, however, are currently not considered in Ethereum specification documents or other literature.

### Functional coupling on the implementation level
Example - persistent storage of state as a Particia tree, which couples logical structure needed to compute state root, and physical structure needed to store the state. Such coupling makes it harder to switch to binary Merkle tree for hashing the state.

### Functional coupling on protocol level
In the design of Ethereum, there some crucial concepts with double or triple functions. First example: gas is used as a charge for resources, as well as means of restricting callbacks and recursion depth. Second example: nonce of accounts is used for both replay protection and as input for generating of the contract addresses. Third example (though this is a prevalent implementation choice rather than a requirement of the specification): Particia Merkle tree is used for both defining what the state root hash is, and for storing data. Functional coupling makes the design inflexible and causes issues when something about the concept needs to be changed. For example, changing gas cost of some operations, like reduction for `SSTORE` and increase for `SLOAD` clashes with the concept of call stipend, which is part of the restricting function of gas.

### Large semantic gap between EVM and many dapps
Definition of semantic gap from www.techopedia.com:
> The “semantic gap” as it is often referenced in IT is the difference between high-level programming sets in various computer languages, and the simple computing instructions that microprocessors work with in machine language. This classic difference has compelled engineers and designers to look at different ways of mediating between high-level language and basic machine language.
> 
> In the past, engineers tried to bridge the semantic gap by making microprocessors more complex, as with Complex Instruction Set Computing (CISC) models. However, they found that it could be just as effective, if not more so, to design Reduced Instruction Set Computing (RISC) models. The philosophy is that microprocessors do not have to do complex work, but can break the high-level instructions down into simple steps. That resonates with how semantic programming is compiled or broken down into machine language. The semantic gap illustrates the difference between humans and computers and how they process data.

After EVM has been designed and implemented, it turned out that there are many use cases, specifically involving cryptographic primitives (e.g. hashing algorithms, digital signatures) whose implementation in EVM opcodes would consume prohibitive amounts of gas.

### Over-reliance on spontaneous contributions
By "spontaneous voluntary contributions" we understand contributions to core implementation by people who are not explicitly asked to do the work, but do it, because they found it interesting and/or important.
By "managed development" we understand development in the core implementations that are directed by some leadership, according to some implementation plan.
Both ways of development have their pros and cons. It seems that in the current circumstances we mainly rely on the spontaneous voluntary contributions and that seems to leave important gaps and technical debt.

## Summary of solutions
At this moment, only solutions to the technological causes are listed and explained.
1. State rent
2. Stateless clients
3. Advanced sync protocols and algorithms (Leaf sync, Beam sync, Firehose/Red Queen)
4. Content Distribution Networks (Cloudflare, Swarm, etc.) to store blocks and/or state
5. Decoupling of two functions of gas
6. Burning part of the transaction fees (EIP-1559)
7. Alternative data layouts and databases
8. Adding pre-compiles (generic elliptic curve pre-compiles, BLAKE hash function)
9. Introducing WASM engine
10. Finality gadget to link Ethereum 1x to the Beacon chain
11. Upgrade of the Proof Of Work algorithm (ProgPOW)
12. State management component (interface design)

## Description of solutions
### State rent
State rent is the only currently proposed solution to decreasing of the state size. All other solutions are concerned with how to deal with the large (and growing) state size more efficiently and prevent some of the most obvious failure modes.

### Stateless clients
Stateless clients is one approach to improve performance of Ethereum client implementations while processing blocks of transactions. Specifically, it seeks to ease the increasing burden that is the state size. It does so by removing the need to download, or implicitly construct and maintain the state, for most of the participants in the Ethereum network. The requirement to access the state is removed by introducing another type of data packets (existing data packet types are, for example, blocks and transactions) to be gossipped around the p2p network. We call this data packets "block witnesses". For each block we have one corresponding block witness. The two main properties that block witnesses have:
1. It is possible verify efficiently that the block witness is indeed constructed from the correct version of the Ethereum state.
2. Block witness has all the required information to make it possible to execute the corresponding block.

More details can be found here: https://medium.com/@akhounov/data-from-the-ethereum-stateless-prototype-8c69479c8abc
