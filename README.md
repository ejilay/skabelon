# SKABELON

Simple tool for preprocessing files with environment variables. May be useful for Docker containers run scripts for config files preprocessing.

Currently uses all available environment variables. 

Usage:

```bash
$ ENV_VAR=value1 skabelon input.envtmpl output.conf
```

For input file like this
```
The variable $ENV_VAR is {{ .ENV_VAR }} and $SHELL is {{ .SHELL }}
```

Will be produced output
```
The variable $ENV_VAR is value1 and $SHELL is /bin/bash
```

## IMPROVEMENTS
    
 [x] Simple file processing
 
 [ ] stdin to stdout processing
    
 [ ] multiple file processing
  
 [ ] multiple template support, nested templates
 
 [ ] variables overrdiding and filtering
