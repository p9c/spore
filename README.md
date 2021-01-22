# OSaaS

Operating System as a Service

Building on the principles of distributed systems, and the legacy of Unix and Plan 9 operating systems, OSaaS aims to become a testbed for fracturing the concept of operating system completely.

## Von Neumann Doesn't Live Here Anymore

Remarkably, as his model of a computation engine ages the lack of reproach towards its false and now increasingly irrelevant assumptions, the name is still very important. Yes, because there will always be small processing units, that have to be sub-nanosecond synchronised.

But for everything else, we have a spectrum of different levels of latency and synchronisation problems to deal with.

### OSaaS is that everything else.

As a means to experimenting with the topology and architecture of processing systems, this project is built on Go, uses a master controller interface GUI that runs the workers who communicate to common channels via the controlling interface.

Executables are thus the form that Go can execute without recompilation - plugins - and for this there must be a shell-like abstraction for command invocation and parameter specification.

As is conventional and for reasons of good UX, the initial implementation only works with read/write pipes (stdio) and renders strings that will be wrapped automatically if not constrained, however these outputs streams will be multiplexed and filtered for display for the user.

Parameter formats will have interface translators and an input field will allow the user to specify a given available call with tab-completion, it will function a lot like a shell but a stream of text will be side by side with any other output modality that is implemented later, which will be the pipe-ified Gio Ops and Events.
