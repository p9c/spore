# OSaaS

###### Operating System as a Service

Bridging the gap between physical hardware, and the often unresolveable lock-in of vendor operating systems, second level solutions can create an environment and aggregate an application development pool together that can eventually displace the rigid base level of operating software in use.

### Why is this important?

In the context of the emergence of broad scale, open, decentralised public network systems, there is a concurrent need to change the base upon which the systems run, in order to better leverage the benefits of participating in and making use of these systems.

However, platforms have a property in which nobody can afford to not have one underneath them, and new systems, to be useful, especially distributed systems, need a large number of participants.

Writing an application to connect to these network systems therefore ends up requiring the building of multiple, often differently functioning implementations of user interface. While making good libraries that abstract the concrete implementations for such things as input and output and storage and other device access, is a big step, the field and the characteristics of the network go hand in hand with changes in how the user approaches the interface and thus also how the data interacts with the interests of the user.

## Components

1. Single thread, scheduled execution engine that accepts Go Plugin standard API 'executables' with an OSaaS API interface design that defines capabilities and privileges (user granted right to this, where it is sensitive or optional). 

2. Top level GUI controller (where the rendering and input capture will live) and run code that has been compiled as a Go plugin, the plugins will stream their GUI state updates as serialized Gio Ops, and inputs will be fed back to them where there is routing that filters by various states such as focus and regional boundaries. This is done via serializing the signals that in current Gio implementation occur over channels. OS and plugin pipes are able to provide delivery guarantees, simplifying this part of the logic.
  
3. Each process server will run within its own single thread of execution. Determining scheduling between competing child processes will be based on a static analysis of the various types of signals the code causes to be emitted - such as the difference between a memory access, cache access, disk access and network access. Consequently, also, this imposes a requirement that plugins are very small and single purpose, similar in theory to the Unix CLI interface model, but as a GUI enabled, pluggable, concurrent and parallel system, with simple integration to become distributed.
  
4. The scheduler will favour low latency jobs on a frequent basis, and attempt to pack the rest of the schedule with the remainder in progressive priority, but randomly, other than this criteria.
  
This system then represents a single thread of execution server, and it can be run on top of another system no less than it could run stand-alone. The execution environment will statically define its working bootstrap state in order to allow the garbage collector to operate.

## Initial work

The initial work is to create the basic plugin execution environment and define a designated entrypoint that runs inside the statically declared initial execution environment. This essentially mainly is about implementing channels as a sort of IPC mechanism though obviously it is always in-process and there will need to be an interruption capability that can pause a process in order to process, for example, inputs, or trigger a display paint, or feed an audio buffer.
