GoPaste
=======


What is it?
-----------

GoPaste is a very minimal clone of pastebin. You paste text in a box.
You get a permanent URL to a page showing the text with formatting
preserved. Once created, pastes are immutable.


Requirements
------------

Unix.


Installation
------------

Download the compiled binary and put it somewhere convenient for
execution.


Usage
-----

Execute the program. There are two optional arguments:

    --store
    --port

The "store" argument must be an absolute path to an existing directory
(gopaste will not attempt to create a directory) that is writeable by
the user that will be running gopaste. This directory will be used to
store paste files.

The "port" arument should be the port number on which to serve.

The store argument defaults to $HOME/gopaste. The port argument defaults
to 8000.


License
-------

This program is distributed under the 2-clause BSD license found in the
LICENSE file.

