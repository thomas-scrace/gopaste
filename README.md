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

Download [the compiled binary](https://github.com/thomas-scrace/gopaste/releases/latest) and
put it somewhere convenient for execution.


Usage
-----

Execute the program, specifying the absolute path to the directory to
use as a store for pastes using the --store argument. This must be an
existing directory that is writeable by the user running gopaste.
Gopaste will not attempt to create a directory itself.

The optional --port argument specifies the TCP port number on which to
serve gopaste. The default value is 8000.


License
-------

This program is distributed under the 2-clause BSD license found in the
LICENSE file.

