<p align="center">
  <picture>
    <img 
      width="456px"
      src="assets/lunarneveridle_icon.png"
    >
  </picture>
 </p>

# 🌙 Lunar NeverIdle

[**Español**](README_ES.md) | [**English**](README.md) | **简体中文**

*我喜欢你，但别删我机，好么？*

本程序随手写的，下面介绍也是随心写的，不喜勿碰。

---

大家好，我是 @jennisu，这里是 **Lunar NeverIdle**。

**Lunar NeverIdle** 是由 @layou233 创建的 [NeverIdle](https://github.com/layou233/NeverIdle) 的一个分支（fork）。

**Lunar NeverIdle** 在后台运行，用于生成必要的资源使用量，以防止 **OCI Free Tier** 实例被删除。  
我使用 **NeverIdle** 已经好几年了，在此期间我的实例从未被删除。但随着时间推移，我意识到存在两个问题：

- 从启动开始就占用较高的 RAM，因为无法将内存使用量调整到低于 1 GiB。  
- 随着时间的推移，内存使用量会逐渐增加，因为内存块未被清理而不断累积。

因此，**Lunar NeverIdle** 诞生了，用来解决这些问题。它与原始版本的区别在于：

- 可以以 **MiB** 为单位指定内存使用量，从而在资源分配上具有更大的灵活性。  
- 它会定期清理内存块，使资源使用保持恒定，防止 RAM 使用量随时间增长。

希望 **Lunar NeverIdle** 对其他使用者来说，也能像对我一样有帮助。

---

## 一键脚本 One click to go

```shell
bash <(curl -s -L https://gist.githubusercontent.com/Ansen/e45320205faf5786d3282ac880f20bab/raw/onekey-NeverIdle.sh)
```

MJJ 们估计会喜欢这个。感谢脚本作者 @Ansen

默认执行下面的命令，当然肯定没法覆盖所有的需求。  
比如 AMD 没有 2G 内存，也没有浪费内存的要求。  
所以依然建议各位自己安装，也是非常便捷迅速的。

## Usage

从 Release 下载可执行文件。注意区分 amd64 和 arm64。

在服务器上启动一个 screen，然后执行本程序，用法自己搜。

命令参数：

```shell
./lunarNeverIdle -cp 0.15 -m 2 -n 4h
```

其中：

-c 指启用 CPU 定期浪费，后面跟随每次浪费的间隔时间。  
如每 12 小时 23 分钟 34 秒浪费一次，则为 `-c 12h23m34s`。按照格式填。

-cp 指启用粗粒度的 CPU 百分比浪费，浪费率将随机器的使用水平实时变化。  
如最大浪费20%的CPU，则为 `-cp 0.2`。百分比的取值范围 [0, 1] 并且注意不要和 `-c` 一起使用。

-m 指启用浪费的内存量，后面是一个数字，单位为 MiB。  
启动后会占用对应量的内存，并且保持不会释放，直到手动杀死进程。

-n 指启用网络定期浪费，后面跟随每次浪费的间隔时间。  
格式同 CPU。会定期执行一次 Ookla Speed Test（还会输出结果哦！）

-t 指设置网络定期浪费的并发连接数。  
默认为10个，值越大消耗的资源越多，一般情况不需要更改。

-p 指设置该进程优先级，后跟随一个优先级数值。不指定则默认使用本平台的最低优先级。  
对于 UNIX-like 系统（如 Linux、FreeBSD 和 macOS），数值取值范围为 [-20,19] ，数字越大优先级越低。  
对于 Windows ，参见 [官方文档](https://learn.microsoft.com/zh-cn/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass)。  
建议不进行指定，默认即为最低优先级，为其它所有进程让路。

*启动该程序后即立刻执行一次你配置的所有功能，可以观察效果。*

## docker 部署
1. 下载 `Dockerfile`
```shell
wget https://raw.githubusercontent.com/jennisu/lunarNeverIdle/master/Dockerfile
```
2. 构建镜像
```shell
# arm机器
docker build -t lunarneveridle:latest .
# amd机器指定 ARCH=amd64
docker build --build-arg ARCH=amd64 -t lunarneveridle:latest .
```
3. 运行
```bash
# 命令参数同上
docker run -d --name lunarneveridle lunarneveridle:latest -c 1h -m 2 -n 4h
```