                Runtime Safety	    Optimizations
Debug	        Yes	                No
ReleaseSafe	    Yes	                Yes, Speed
ReleaseSmall	No	                Yes, Size
ReleaseFast	    No	                Yes, Speed

These are the four build modes of Zig:

1. `Debug` is the default mode when you run these commands without any flags:
`zig run file_name.zig` or `zig test file_name.zig` 

For the others just add `-O ReleaseSafe` 
or `-O ReleaseSmall` or `-O ReleaseFast`
