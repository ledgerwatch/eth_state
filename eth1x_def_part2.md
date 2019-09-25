# Ethereum 1x Definition (part 2)

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
Proposed solutions should be targeted at the causes but may have side-effects (some solutions may be causes for other challenges, though perhaps less critical than the ones they are trying to address)
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


```graphviz
digraph new_implementations {
rankdir=LR;
new_implementations [label="New implementations" shape=octagon style=filled fillcolor=plum];
data_management [label="Data management" shape=hexagon style=filled fillcolor=salmon];

new_implementations -> data_management [dir=back];
}
```
#### Data management (cause)
The hardest part of a mainnet capable implementation seems to be data management. Ethereum node needs to transmit, process and store large amounts of data. To do so efficiently, it needs non-trivial techniques. Such techniques, however, are currently not considered in Ethereum specification documents or other literature.

#### Large state (sub-cause)
```graphviz
digraph causes {
rankdir=LR;
data_management [label="Data management" shape=hexagon style=filled fillcolor=salmon];
large_state [label="Large state" shape=hexagon style=filled fillcolor=salmon];
data_management -> large_state [dir=back];
}
```

#### Functional coupling (cause)
In the design of Ethereum, there some crucial concepts with double of tripple functions. First example: gas is used as charge for resources, as well as means of restricting callbacks and recursion depth. Second example: nonce of accounts is used for both replay protection and as input for generating of the contract addresses. Third example (though this is a prevalent implementation choice rather than a requirement of the specification): particia merkle tree is used for both defining what the state root hash is, and for storing data. Functional coupling makes the design inflexible and causes issues when something about the concept needs to be changed. For example, changing gas cost of some operations, like reduction for `SSTORE` and increase for `SLOAD` clashes with the concept of call stipend, which is part of the restricting function of gas.

#### Spontaneous voluntary contributions vs Managed development (cause)
By "spontaneous voluntary contributions" we understand contributions to core implementation by people who are not explicitely asked to do the work, but decide to do it, because they found it interesting and/or important.
By "managed development" we understand development in the core implementations that are directed by some leadership, according to some implementation plan.
Both ways of development have their pros and cons. It seems that in the current circumstances we mainly rely on the spontaneous voluntary contributions and that seems to leave important gaps and technical debt.

## Entire diagram
```graphviz
digraph agents {
rankdir=LR;

challenges [label="CHALLENGES" shape=box penwidth=0 fontsize=17];
causes [label="CAUSES" shape=box penwidth=0 fontsize=17];
solutions [label="SOLUTIONS" shape=box penwidth=0 fontsize=17];

tx_cost [label="Cost of transactions" shape=octagon style=filled fillcolor=plum];
tx_safety [label="Risk assessment" shape=octagon style=filled fillcolor=plum];

gas_limit [label="Block gas limit" shape=hexagon style=filled fillcolor=salmon];
tx_cost -> gas_limit [dir=back];
fee_market [label="Fee market" shape=hexagon style=filled fillcolor=salmon];
tx_cost -> fee_market [dir=back];

fee_burn [label="Fee burn" shape=box style=filled fillcolor=lightblue];
fee_market -> fee_burn [dir=back];

reorgs [label="Long reorgs" shape=hexagon style=filled fillcolor=salmon]
tx_safety -> reorgs [dir=back];
finality_gadget [label="Finality gadget" shape=box style=filled fillcolor=lightblue];
reorgs -> finality_gadget [dir=back];

storage_devices [label="Cost of storage devices" shape=octagon style=filled fillcolor=plum];
traffic [label="High internet traffic" shape=octagon style=filled fillcolor=plum];
dev_ops [label="Complex DevOps" shape=octagon style=filled fillcolor=plum href="#DevOps-challenge"];
sync_time [label="Sync time" shape=octagon style=filled fillcolor=plum];
sync_time -> large_state [dir=back];

scalability [label="Scalability" shape=octagon style=filled fillcolor=plum href="#Scalability-challenge"];
security [label="Security" shape=octagon style=filled fillcolor=plum href="#Security-challenge"];

new_implementations [label="New implementations" shape=octagon style=filled fillcolor=plum href="#New-implementations-challenge"];
product_vs_system [label="Product vs System" shape=octagon style=filled fillcolor=plum href="#Product-vs-system-challenge"];
backwards_compatibility [label="Backwards compatibility" shape=octagon style=filled fillcolor=plum href="#Backwards-compatibility-challenge"]

data_management [label="Data management" shape=hexagon style=filled fillcolor=salmon];
dev_ops -> data_management [dir=back];
new_implementations -> data_management [dir=back];
large_state [label="Large_state" shape=hexagon style=filled fillcolor=salmon];
data_management -> large_state [dir=back];
storage_devices -> large_state [dir=back];
functional_coupling [label="Functional coupling" shape=hexagon style=filled fillcolor=salmon];
backwards_compatibility -> functional_coupling [dir=back];
spontaneous_vs_managed [label="Spontaneous vs managed" shape=hexagon style=filled fillcolor=salmon];
product_vs_system -> spontaneous_vs_managed [dir=back];

state_rent [label="State rent" shape=box style=filled fillcolor=lightblue];
large_state -> state_rent [dir=back];
backwards_compatibility -> state_rent [dir=back style=dotted label="worsen"];

stateless [label="Stateless clients" shape=box style=filled fillcolor=lightblue];
data_management -> stateless [dir=back];

{rank=same; challenges storage_devices tx_cost dev_ops tx_safety sync_time backwards_compatibility}
{rank=same; causes gas_limit}
{rank=same; solutions stateless state_rent fee_burn finality_gadget}
}
```

## TODOs
- [ ] Discuss and correct methodology, wording, and content (specically the challenges)
- [ ] Levels of criticality for challenges, perhaps expressed by colours
- [ ] Add remaining open projects to the solutions
- [ ] Place correct hyperlinks on all the nodes in the diagram
