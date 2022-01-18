1/17/2022
I've started doing some profiling, and it looks like the biggest issue for my code is the use of large numbers of pointers (e.g. \*Vec3 and \*Ray mainly). From profiling runs, we are spending a lot of time in mallocgc and procyield, which may indicate that garbage collection pressure is extremely high. If these resources are managed more explicitly, then we should see drastic performance improvements.
