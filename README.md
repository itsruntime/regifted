Regifted
========

This is an mpeg2 transport stream (mpeg2 ts) to mpeg4 repackaging utility.

Build Process  
=============  

#### Installing GO and Setting the GO Environment
##### 1. Download and Install GO using instructions from http://golang.org/doc/install
##### 2. Create a directory where you would like to keep your source, bin, and package files. This is typically named "go".
##### 3. In the folder made in step 2, create three folders: src, pkg, and bin.
##### 4. Set-up your go environment.
 					i.  export GOPATH=(the path to the file in step 2.)
 					ii. export GOBIN=$GOPATH/bin

#### Pulling the source code from Github
##### 5. Pull the source code from github into the src folder created in step 3.
					i.   navigate to the $GOPATH/src folder
					ii.  type 'git clone https://github.com/itsruntime/regifted.git'

#### Options for building and running the program using 'go install', 'go run', and the resultant binaries.
##### 6.  Move to the regifted folder using the CMD or Terminal. Type both:
					i.  go install regifted/moof_main.go
					ii. go install regifted/ts.go
##### 7.  You can run the file from this location.
					i.  go run ts.go [ts file location]
					ii. go run moof_main.go [mp4 fragment file location]
##### 8. In step 6 you built binaries for moof_main.go and ts.go. These can be found in $GOPATH/bin.
					i. From the $GOPATH/bin directory you can type:
						a. ts [ts file location]
						b. moof_main [mp4 fragment file location]

Test Process  
============  
#### Running the unit tests
##### 1. Visit http://golang.org/cmd/go/#hdr-Test_packages for more on the GO testing suite.
##### 2. From the regifted folder you make these calls:
					i.  go test -v regifted/moof
					ii. go test ts_test.go


Implementation Documentation  
============================
##### 1. Steps to generate godocs. From the regifted folder:
					i.  godoc regifted/moof
					ii. godoc regifted/ts
##### 2. You can find a link to a prezi-based diagrams demonstrating the activity of ts.go here:
##### https://github.com/itsruntime/regifted/wiki/Design-Documentation
(A similar diagram for moof.go is forthcoming.)
