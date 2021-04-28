# ocp-release

This is one of the ways to get the best candidate release for the ocp installation based on the version type you need (nightly or CI) and filtered by a condition which must be success.

## Install 

- Just to install the binary in your system:

[https://github.com/alknopfler/ocp-release/releases/latest](https://github.com/alknopfler/ocp-release/releases/latest)

- Download the file for your OS
- Unzip/unTar the file to locate into a folder inside your $PATH


## Use it

To show the instructions or the help message 
```commandline
./ocp-release -h
```

To get the best release candidate tag:
```commandline
./ocp-release -v <version> -c <condition>
```

For example:
```commandline
./ocp-release -v nightly -c assisted-metal
```



