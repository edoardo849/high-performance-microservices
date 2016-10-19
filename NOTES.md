High Performance Microservices 
===

## Hardware limits 

What do we do when we want to augment productivity:
- we can do things faster... until we reach a limit 
- we put more people on the job... until organization starts to slow down the process
- we better organize the job... until we reach an optimal organization 

last option:
"Are we actually doing this job right?"

hardware limitations.
application organization.
security

Let's talk about hardware limits.

## Free lunch is over 
- speed... we've hit a limit on clock speed 



There are two main benefits WebAssembly provides:

The kind of binary format being considered for WebAssembly can be natively decoded much faster than JavaScript can be parsed (experiments show more than 20× faster). On mobile, large compiled codes can easily take 20–40 seconds just to parse, so native decoding (especially when combined with other techniques like streaming for better-than-gzip compression) is critical to providing a good cold-load user experience.

By avoiding the simultaneous asm.js constraints of AOT-compilability and good performance even on engines without specific asm.js optimizations, a new standard makes it much easier to add the features required to reach native levels of performance.

Put on a graph new Programming Languages: after 2005.

Also on the web there is this trend: WebAssembly.

- WebAssembly is going into the same direction for the web.

[ First, free lunch is over ]

[ How this relates to Microservices ]

Faster CPUs, more transistors = Scale vertically (more powaaa)
Multi Core CPUs = Scale horizontally (orchestration)

Scaling Microservices
- X-axis: multiple copies of an application behin a load balancer
- Y-axis: splits the application into multiple, different services
- Z-axis: each server is responsible for only a subset of the data (i.e. customer type)

Complexity requires orchestration (Kubernetes, CoreOS, Containers)

What is the reason for scaling?
- Isn't there anything that we can do first?
- Hope for the best, prepare for the worst: set up an auto-scaling group

Everything is a tradeoff
- Cost of power for vertical scaling
- Cost of orchestration for horizontal scaling (and higher skilled teams) *
- Cost of fine-tuning for optimization (and higher skilled teams) *

* consultancy opportunities :-)

The Art of Scalability by Martin L. Abbott and Michael T. Fisher 

from the customer or end-user perspective, what matters is the performance of the app itself

- Is compiled, not interpreted.
- Permits efficient code to be written.
- Lets programmers talk about memory effectively, think structs vs java objects
- Has a compiler that produces efficient code, it has to be small code as well, because cache.