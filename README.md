# erg

A Gossip Protocol for Event Causality Consensus

In distributed systems, discovering the objective order of events can solve a lot of very hard problems.

`erg` is a protocol for a cluster of peers to rapidly determine a consensus of the sequence of a series of unique events
detected and reported by nodes in the network

Erg is, from the word 'ergodicity', meaning the property of a system that eventually visits every possible state within
a period of time that it takes for this to happen, on average.

## Network Saturation Delay

The time from a node sensing an event, to all N nodes hearing of it, is the Network Saturation Delay.

Given a network of N nodes that are sending a message at every opportunity, there is N/2+N%2 messages per cycle. From
this one can predict the cycles to saturation, with strong probabalistic guarantees.

The goal of erg is minimizing NSD.

## Network Membership

The primary event shared by an erg mesh is membership.

Members advertise themselves to other nodes, and record the first seen time, and update it each time they receive a
message from them.

If the node doesn't come back after 2N message cycles, it can be assumed to be stopped or disconnected, and its initial
time is purged.

By requiring continual participation, nodes that get a long way out of date are ignored off the network. This can be
determined by the amount of common events in their recent message log gossip.

If there are none, a node will drop the node's registration and ignore its messages until it receives a new one inside
the protocol joining window of a valid next registration, which is about 2N message cycles.

## Propagation Protocol

- nodes select their next node to query by two criteria:
    - they haven't been queried for N/4 cycles
    - their uptime is the highest of conflicting options
- nodes request a log by sending their previous log's hash
- a correctly functioning node will send a log that produces a different hash, or confirmation acknowledgement.
- logs are zipped into a single chain and older duplicates are removed
- nodes that have no duplicates are considered to be out of sync or byzantine, and are expired.
- their uptime value is zeroed and their contact falls to the end of the list for next query selection.

## Probabalistic Consensus

This is a probabalistic consensus with a very small rule set that guarantees a maximum N/2 cycles to convergence, to a
minimum N/4 cycles of best case convergence.

A cycle is 200ms, and 5 cycles happen per second, 300 cycles per hour, on average. In each cycle, there is N/2+N%2
messages.

If this was 100 nodes, then 15000 messages per hour, which is 3750 fully sorted messages in that time period, or
potentially around 112.5 transactions per second throughput, 20 seconds total finality, nearly 10x the nakamoto
consensus maximum of 17tps.

If the number of transactions in a time period is lower than 1/4 of the message cycles of the network, the network
converges all transactions between N/2 and N/4 cycles.

In other words, this network counters the factorial expansion with powers of two, which is still not linear but it's a
lot less than factorial.