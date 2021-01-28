# sc

**s**mall RPN **c**alculator written in Golang

## Usage

The most basic usage is just by invoking the program name:

```
sc
```

Then you can enter values and commands ([see commands](#commands)):

```
$ sc

12
        12.000000

15
        12.000000
        15.000000

+
        27.000000
```

On the left, it is your inputs. On the right, the stack.

You can enter more than one value, as long as they are separated by spaces:

```
$ sc 

100 200 add
        300.000000
```

Moreover, you can pipe commands into `sc`:

```
$ echo "2 10 pow" | sc 
1024
```

## Commands

Supported commands are the following:

* `+`,`a`,`add`: Add the last two elements of the stack
* `-`,`s`,`sub`: Substract the last two elements of the stack
* `*`,`m`,`mul`: Multiply the last two elements of the stack
* `/`,`d`,`div`: Divide the last two elements of the stack
* `sw`,`swap`: Swap the last two elements of the stack
* `p`,`pop`: Pop the last element of the stack
* `q`,`quit`: Quit the program
* `c`,`clear`: Clear the stack
* `^`,`pow`: Do a power of the last two elements of the stack
* `sum`: Sum all the stack
  
## License

MIT