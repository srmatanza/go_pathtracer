6/8/2022
Actually, the biggest performance bottleneck turned out to be the random number generator. I switched to a much faster third party implementation, and now the renderer can saturate the CPU.

1/18/2022
Because the bulk of memory used in this application is from one-off allocations of Vec3 and Ray structs, it may be a text-book use case for memory pools. First though, let's try and make a go benchmark using our renderer.

1/17/2022
I've started doing some profiling, and it looks like the biggest issue for my code is the use of large numbers of pointers (e.g. \*Vec3 and \*Ray mainly). From profiling runs, we are spending a lot of time in mallocgc and procyield, which may indicate that garbage collection pressure is extremely high. If these resources are managed more explicitly, then we should see drastic performance improvements.

