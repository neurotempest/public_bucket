Scripts for intracting with itcoin using python bitcoinlib.

# Prerequisits

* `bitcoinlib`:

  * This can be installed via pip, however it require gmp to also be installed. This can be installed via homebrew:

    > brew install gmp

  * Since the pip install for bitcoinlib will build a few c libs we also need to let the C compiler know where to find gmp headers + libs. So run pip install with this command (worked on osx 11.6):

    > env "CFLAGS=-I/opt/homebrew/include -L/opt/homebrew/lib"  pip3 install bitcoinlib
