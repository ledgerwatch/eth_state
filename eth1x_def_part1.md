# Ethereum 1x Definition (part 1)

## Objective
This document aims to define and describe the Ethereum 1x initiative, with a primary goal of being used as a guide for funding and other support that various improvement proposals and projects should receive under the auspices of Ethereum 1x.

## Definition
We define Ethereum 1x as a set of solutions that seek to address the causes of the most critical challenges faced by the agents in the Ethereum system. TODO: this hinges on the determinations of which challenges are more critical than the others.

## Methodology and terminology
For the purpose of this analysis, we view Ethereum as a set of interacting agents. Their interaction is described by the concepts of contributions and challenges.
Ethereum system itsef is an emergent entity, and only exists as long as its crucial actors keep making their contributions. The goal of the Ethereum 1x as an initiative, is to preserse and improve the "wellfare" of this emergent entity, rather than any agents specifically.

We split the definition into two parts. First parts serves as an introduction to the challenges that Etherem 1x considers. Second parts goes into more (technical) details of the challenges, and discusses their causes and solutions. It is possible that the second part will be further subdivided to separate the discussion of causes and solutions.

### Agents
Agents make contributions and thus, collectively, give rise to the Ethereum system. Their contributions also give the Etherum system its properties, such as resilience, efficiency, accessibility. Agents are shown as ellipses.
```graphviz
digraph agents {
rankdir=LR;
agent [style=filled fillcolor=khaki]
}
```
### Contributions
We attempt to enumerate the most valuable and crucial contributions of the agents to the Ethereum system.
```graphviz
digraph contributions {
rankdir=LR;
agent [style=filled fillcolor=khaki];
contribution [shape=doubleoctagon style=filled fillcolor=lawngreen];
contribution->agent [dir=back];
}
```
### Challenges
Challenges (how critical they are, i.e. if the challenge is not met, will agents's contribution degrade or stop?) that agents face. Shown as octagons.
```graphviz
digraph challenges {
rankdir=LR;
challenge [shape=octagon style=filled fillcolor=plum];
agent [style=filled fillcolor=khaki];
agent -> challenge;
}
```
## Agents
For the purpose of this document, we will consider a somewhat coarse-grained grouping of Ethereum agents. In the later revisions of the document, the grouping can be expanded or made more fine-grained.

```graphviz
digraph agents {
rankdir=LR;
end_users [label="End users" style=filled fillcolor=khaki];
node_ops [label="Node operators" style=filled fillcolor=khaki];
miners [label="Miners" style=filled fillcolor=khaki];
dapp_devs [label="Dapp developers" style=filled fillcolor=khaki];
core_devs [label="Core developers" style=filled fillcolor=khaki];
}
```
Each agent type makes certain contributions to the wellfare of the Ethereum system, and, in turn faces challenges. If these challenges become too great, then the contribution reduces or stops, and the Ethereum system suffers or potentially ceases to exist.
## End users (agent)
End users generally contribute "usage" to the system. Some usage is directly beneficial, as it brings more resources into the system (for example, purchasing Ether), some - only indirectly (intense usage attract more users).
End users are mostly concerned with lack of use cases, transactional cost of using Ethereum dapps or Ether currency, as well as the safety of their funds and data.
```graphviz
digraph end_users {
rankdir=LR;
end_users [label="End users" style=filled fillcolor=khaki];
resources [label="Resources" shape=doubleoctagon style=filled fillcolor=lawngreen];
usage [label="Usage" shape=doubleoctagon style=filled fillcolor=lawngreen];
no_usecases [label="Lack of use cases" shape=octagon style=filled fillcolor=plum];
tx_cost [label="Cost of transactions" shape=octagon style=filled fillcolor=plum];
tx_safety [label="Risk assessment" shape=octagon style=filled fillcolor=plum];
{resources usage} -> end_users [dir="back"]
end_users -> {tx_cost tx_safety no_usecases}
}
```
### Cost of transactions (challenge)
By cost of transactions for the end user we normally understand the gas cost of those transactions multiplied by the current cost of gas. Users may often overpay for gas because they cannot reliably predict which gas price level will ensure the inclusion of their transactions with required urgency.
### Risk assessment (challenge)
Often users need to weight risk of exploit against the benefit they hope to get from the transaction. Initial design of Ethereum did not put a lot of emphasys on making that risk assessment easier.
## Node operators (agent)
Cost of operating a full node is rising. Below are the main cost categories.
```graphviz
digraph node_ops {
rankdir=LR;
node_ops [label="Node operators" style=filled fillcolor=khaki];
nodes [label="Network nodes" shape=doubleoctagon style=filled fillcolor=lawngreen];
nodes -> node_ops;
storage_devices [label="Costly high end storage devices" shape=octagon style=filled fillcolor=plum];
traffic [label="High internet traffic" shape=octagon style=filled fillcolor=plum];
dev_ops [label="Complex DevOps" shape=octagon style=filled fillcolor=plum]
node_ops -> storage_devices;
node_ops -> traffic;
node_ops -> dev_ops;
}
```
### High end storage devices (challenge)
All main-net capable implementations require at least SSD to maintable acceptable operating speed, and NVMe is desirable for a initial sync (downloading all the blocks and reconstituting the current state of the blockchain).
### Internet traffic (challenge)
After initial sync, traffic usage should not be very high. However, if many new nodes joining the network, the outgoing traffic of the incument nodes may increase, to serve blocks and initial state to the new-joiners.
### DevOps (challenge)
Ethereum nodes transmit, process, and store large amounts of data. They also often have very little tolerance to downtime. These two characteristics mean that management of Ethereum nodes can be non-trivial.

## Dapp developers (agent)
Dapp developers are mainly concerned with two issues: scalability and security. For some potential dapps, Ethereum as a platform does not provide enough capacity in terms of transactional throughput, cost of transactions, capacity of data retention.
These two issued can be thought of mirroring the issues facing the end users. Higher scalability correlates to the lower cost for users. Better security measures for smart contracts translate to easier risk assessment for the users.
```graphviz
digraph dapp_devs {
rankdir=LR;
dapp_devs [label="Dapp developers" style=filled fillcolor=khaki];
use_cases [label="Use cases" shape=doubleoctagon style=filled fillcolor=lawngreen];
use_cases -> dapp_devs;
scalability [label="Scalability" shape=octagon style=filled fillcolor=plum];
security [label="Security" shape=octagon style=filled fillcolor=plum];
dapp_devs -> scalability;
dapp_devs -> security;
}
```
### Scalability (challenge)
Some use cases of smart contracts and currency require certain level of scalability, which usually translates to how quickly a transaction gets "confirmed" related to how high gas price is paid. If scalability of the system is not enough, some more complex constructions (e.g. Level 2 solutions like state channels and Plasma) might need to be employed, at the cost of increased complexity.

### Security (challenge)
Whenever users' interaction with the Ethereum system are less trivial than currency transfer, some security analysis is usually performed, to demonstrate to the potential users of smart contracts that most likely their transactions will not have unintended consequences. So far the tradition is that the additional cost of performing security audits of the code, or formal verification, or other security measures is paid by the dapp developers.

## Core developers (agent)
Contribuion of core developers is in developing and maintaining the software that, when run on the computers connected to the Internet, gives rise to the Ethereum network.
Development of a fully functional Ethereum implementation which is capable of run on the mainnet is challenging.
```graphviz
digraph core_devs {
rankdir=LR;
core_devs [label="Core developers" style=filled fillcolor=khaki];
software [label="Node software" shape=doubleoctagon style=filled fillcolor=lawngreen];
software -> core_devs;
new_implementations [label="New implementations" shape=octagon style=filled fillcolor=plum];
product_vs_system [label="Product vs System" shape=octagon style=filled fillcolor=plum];
backwards_compatibility [label="Backwards compatibility" shape=octagon style=filled fillcolor=plum]
core_devs -> new_implementations;
core_devs -> product_vs_system;
core_devs -> backwards_compatibility;
}
```
#### New implementations (challenge)
It is difficult to create new indepdendent implementations (clients) that are capable of working with the main net. Lack of implementation diversity leads to core developers being "conservative" by default.```

#### Product vs system (challenge)
It is difficult to balance product improvements (performance, usability), and the upgrades of Ethereum. Example: preparation for the Istanbul hard fork vs finishing Geth 1.9

#### Backwards compatibility (challenge)
It is difficult to balance backwards compatibility (with the existing contracts), and the rule changes that are neccessary for Ethereum system to be sustainable.

## Entire diagram
```graphviz
digraph agents {
rankdir=LR;

contributions [label="CONTRIBUTIONS" shape=box penwidth=0 fontsize=17];
agents [label="AGENTS" shape=box penwidth=0 fontsize=17];
challenges [label="CHALLENGES" shape=box penwidth=0 fontsize=17];

end_users [label="End users" href="#End-users-agent" style=filled fillcolor=khaki];

resources [label="Resources" shape=doubleoctagon style=filled fillcolor=lawngreen];
usage [label="Usage" shape=doubleoctagon style=filled fillcolor=lawngreen];
tx_cost [label="Cost of transactions" shape=octagon style=filled fillcolor=plum];
tx_safety [label="Risk assessment" shape=octagon style=filled fillcolor=plum];
{resources usage} -> end_users [dir=back]
end_users -> {tx_cost tx_safety}

node_ops [label="Node operators" style=filled fillcolor=khaki href="#Node-operators-agent"];
nodes [label="Network nodes" shape=doubleoctagon style=filled fillcolor=lawngreen dir=back];
nodes -> node_ops;
storage_devices [label="Cost of storage devices" shape=octagon style=filled fillcolor=plum];
traffic [label="High internet traffic" shape=octagon style=filled fillcolor=plum];
dev_ops [label="Complex DevOps" shape=octagon style=filled fillcolor=plum href="#DevOps-challenge"];
sync_time [label="Sync time" shape=octagon style=filled fillcolor=plum];
node_ops -> {storage_devices traffic dev_ops sync_time}

miners [label="Miners" style=filled fillcolor=khaki];
chain_security [label="Chain security" shape=doubleoctagon style=filled fillcolor=lawngreen];
chain_security -> miners [dir=back];

dapp_devs [label="Dapp developers" style=filled fillcolor=khaki];
use_cases [label="Use cases" shape=doubleoctagon style=filled fillcolor=lawngreen];
use_cases -> dapp_devs [dir=back];
scalability [label="Scalability" shape=octagon style=filled fillcolor=plum href="#Scalability-challenge"];
security [label="Security" shape=octagon style=filled fillcolor=plum href="#Security-challenge"];
dapp_devs -> scalability;
dapp_devs -> security;

core_devs [label="Core developers" style=filled fillcolor=khaki href="#Core-developers-agent"];
software [label="Node software" shape=doubleoctagon style=filled fillcolor=lawngreen];
software -> core_devs [dir=back];
new_implementations [label="New implementations" shape=octagon style=filled fillcolor=plum href="#New-implementations-challenge"];
product_vs_system [label="Product vs System" shape=octagon style=filled fillcolor=plum href="#Product-vs-system-challenge"];
backwards_compatibility [label="Backwards compatibility" shape=octagon style=filled fillcolor=plum href="#Backwards-compatibility-challenge"]
core_devs -> {new_implementations product_vs_system backwards_compatibility};

{rank=same; contributions use_cases}
{rank=same; agents dapp_devs}
{rank=same; challenges storage_devices tx_cost dev_ops tx_safety sync_time}
}
```

## TODOs
- [ ] Discuss and correct methodology, wording, and content (specically the challenges)
- [ ] Levels of criticality for challenges, perhaps expressed by colours
- [ ] Add remaining open projects to the solutions
- [ ] Place correct hyperlinks on all the nodes in the diagram


