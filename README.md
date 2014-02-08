Regifted
========

This is an mpeg2 transport stream (mpeg2 ts) to mpeg4 repackaging utility.

Build Process  
=============  
The structure of the GO "workspace" is language-defined. There is a $GOPATH  defined by the local administrator A.K.A. end-user and the directories $GOPATH/{src,pkg,bin}.  
This is described here:  
http://golang.org/doc/code.html  

export GOPATH=/home/user/go  
cd $GOPATH  
mkdir {src,pkg,bin}  
go get github.com/itsruntime/regifted  
$GOPATH/bin/regifted  

Test Process  
============  
We make use of the language testing library that ships with go. It is ran by running the very verbose command:  
go test github.com/itsruntime/regifted  

Implementation Documentation  
============================  
The language documentation with go is connected to libraries. I think maybe that libraries are defined by functions that start with a capital letter. The documentation is sourced by regular comments directly above the function. e.g.;  
  
// foo says bar  
func foo() { ...  

License  
=======  
This software is released under the Apache license. My non-legal, interpertation of the license is that it allows use, redistribution, and modification to the source with no requirement to redistribute derivatives in a similar way (i.e. you can close source what you add on). The two major restrictions are that the license has to stay bundled with derivative works with the full original source and credit must be appropriated correctly to authors.