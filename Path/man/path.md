% Path(1) Path user manual
% Christopher Lillthors
% Stockholm, Sweden.
% April 23, 2014

# NAME
Path - pathfinder

# SYNOPSIS
Path [-filename=<true/false>][-profile=<true/false>][-stack=<true/false>] from to

*Warning* if no file is given, then *input.txt* is standard.

# DESCRIPTION
This program gives you a path from a node to another. It is not guaranteed to be the shortest path.

# OPTIONS
-filename file
-profile true/false
-stack true/false

# Examples
Path -filename="routes.txt" -stack=true 0 1

Path 1 4

Path -filename="routes.txt" -profile=true 3 2

# License
The MIT License (MIT)

Copyright (c) 2014 Christopher Lillthors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

# SEE ALSO
The path source code and all documentation may be downloaded from
<http://github.com/christopherL91/GoInda13/path>
