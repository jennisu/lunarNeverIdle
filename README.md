<p align="center">
  <picture>
    <img 
      width="456px"
      src="assets/lunarneveridle_icon.png"
    >
  </picture>
 </p>

# Lunar NeverIdle

[**Español**](README_ES.md) | **English** | [**简体中文**](README_CN.md)

*I love you, but do not stop my machine, could you?*

---

## Usage

Download executable file from Release. Note the distinction between amd64 and arm64.

Start a screen on the server and run it.
If you want to learn about screen command, just Google.

Command arguments：

```shell
./lunarNeverIdle -cp 0.15 -m 2 -n 4h
```

In which:

-c enables CPU periodic waste, followed by the interval between wastes.  
E.g. waste CPU every 12 hours 23 minutes and 34 seconds, then the argument would be `-c 12h23m34s`.
Just follow this template.

-cp enables coarse-grained CPU percentage waste, and the waste rate will change in real time with the usage level of the machine.  
If the maximum waste of 20% of the CPU is `-cp 0.2`. The value range of percentage is [0, 1] and be careful not to use it with `-c`.

-m enables memory waste, followed by a number in MiB.  
After startup, the specified amount of memory will be occupied and will not be released until the process is killed.

-n enables network(bandwidth) periodic waste, followed by the interval between wastes.  
Argument format is the same as `-c`. The Ookla Speed Test will be performed periodically (and the results will be output!)

-t specifics the number of concurrent connections of the network periodic waste.  
The default is 10. The larger the value, the more resources will be consumed. For most situations, it does not need to be changed.

-p specifics the process priority, followed by a priority value. If not specified, the lowest priority of the platform will be used by default.  
For UNIX-like systems (such as Linux, FreeBSD, and macOS), the value range is [-20,19], and the higher the number, the lower the priority.  
For Windows, see [the official documentation](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass).  
It is recommended not to specify because the default is the lowest priority, making way for all other processes.

*All the functions you configured will be executed once immediately when you start the program, so you can take a look at the effect.*

## Docker - Deployment
1. Download `Dockerfile`
```shell
wget https://raw.githubusercontent.com/jennisu/lunarNeverIdle/master/Dockerfile
```
2. Build the image
```shell
# ARM Machine
docker build -t lunarneveridle:latest .
# AMD machine specifications ARCH=amd64
docker build --build-arg ARCH=amd64 -t lunarneveridle:latest .
```
3. Execute
```bash
# Parameters are the same as above
docker run -d --name lunarneveridle lunarneveridle:latest -c 1h -m 2 -n 4h
```

