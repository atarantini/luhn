====
Luhn
====

Simple golang application and library to validate strings with the Luhn (mod-10) algorithm.


Usage
-----

From the command line, will validate arguments::

    $ ./luhn 4242424242424242
    4242424242424242 is valid


Will work with multiple arguments::

    $ ./luhn 4242424242424242 4242424242424240
    4242424242424242 is valid
    4242424242424240 is NOT valid

Library
-------

Import it and call the ``Validate`` function:

.. code-block:: go

    package main

    import(
        "fmt"
        "github.com/atarantini/luhn/lib"
    )


    func main() {
        fmt.Println(luhn.Validate("4242424242424242"))
    }


Author
------

Andres Tarantini (atarantini@gmail.com)

License
-------

Released under MIT license, see https://opensource.org/licenses/MIT
