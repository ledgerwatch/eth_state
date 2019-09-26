# Ethereum 1x Definition (part 2 diagram WIP)

The diagram for the part 2 is moved here for rework

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
