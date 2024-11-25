A Tool to automate the process of creating c++ Class + it header file

# Usage
first you need go on your system, the :
```bash
git clone https://github.com/flutterde/cppcreate.git
cd cppcreate
make
```

## Add the command as an Alias
for bash:
```
# go to your home
cd ~
# edit .bashrc file
nano .bashrc
# and Past this:
alias cppcreate='/PATH_TO_COMMAND/cppcreate'
# Start using it:
cppcreate -ch MyClassName
```

## Flags

```
-h		Help
-ch		Create Class and it Header file
-c		Create Class only
-mlx	install MLX

```


