<p align="center">
  <picture>
    <img 
      width="456px"
      src="assets/lunarneveridle_icon.png"
    >
  </picture>
 </p>

# 🌙 Lunar NeverIdle

**Español** | [**English**](README.md) | [**简体中文**](README_CN.md)

*Te quiero, pero ¿podrías no detener mi máquina?*

---

Hola, soy @jennisu y esto es **Lunar NeverIdle**.

**Lunar NeverIdle** es un *fork* de [NeverIdle](https://github.com/layou233/NeverIdle), creado por @layou233.

**Lunar NeverIdle** se ejecuta en segundo plano y genera el uso de recursos necesario para evitar que las instancias de **OCI Free Tier** sean eliminadas.  
Durante varios años utilicé **NeverIdle**, y mi instancia nunca fue eliminada. Sin embargo, con el tiempo me di cuenta de dos problemas:

- Alto consumo de RAM desde el inicio, ya que no era posible ajustar el uso de memoria a menos de 1 GiB.  
- Con el tiempo, el uso de memoria aumentaba debido a que los bloques no se limpiaban y se acumulaban progresivamente.

Así nació **Lunar NeverIdle**, con la idea de resolver estos problemas. Se diferencia del original en que:

- El uso de memoria puede especificarse en **MiB**, lo que permite una mayor flexibilidad en la asignación de recursos.  
- Limpia los bloques de memoria periódicamente, evitando que el consumo se acumule y manteniendo estable el uso de RAM con el tiempo.

Espero que **Lunar NeverIdle** sea tan útil para quienes lo usen como lo ha sido para mí.

---

## Uso

Descarga el archivo ejecutable desde "Release". Ten en cuenta la diferencia entre las versiones para amd64 y arm64.

Inicia una sesión "screen" en el servidor y ejecútalo.
Si deseas aprender sobre el comando "screen", simplemente busca en Google.

Argumentos del comando:

```shell
./lunarNeverIdle -cp 0.15 -m 2 -n 4h
```

Donde:

-c activa el desperdicio periódico de la CPU, seguido del intervalo entre los desperdicios.
Por ejemplo, para desperdiciar CPU cada 12 horas, 23 minutos y 34 segundos, el argumento sería `-c 12h23m34s`.
Solo sigue esta plantilla.

-cp activa el desperdicio de porcentaje de CPU de granulación gruesa, y la tasa de desperdicio cambiará en tiempo real según el nivel de uso de la máquina.
Si el desperdicio máximo del 20% de la CPU es `-cp 0.2`. El rango de valores del porcentaje es [0, 1], y ten cuidado de no usarlo junto con `-c`.

-m activa el desperdicio de memoria, seguido de un número en MiB.
Después de iniciarse, se ocupará la cantidad de memoria especificada y no se liberará hasta que el proceso sea detenido.

-n activa el desperdicio periódico de la red (ancho de banda), seguido del intervalo entre los desperdicios.
El formato del argumento es igual al de `-c`. ¡Se realizará una prueba de velocidad de Ookla periódicamente (y los resultados serán mostrados)!

-t especifica el número de conexiones simultáneas para el desperdicio periódico de la red.
El valor predeterminado es 10. Cuanto mayor sea el valor, más recursos se consumirán. En la mayoría de las situaciones, no es necesario cambiarlo.

-p especifica la prioridad del proceso, seguida de un valor de prioridad. Si no se especifica, se utilizará la prioridad más baja de la plataforma.
Para sistemas similares a UNIX (como Linux, FreeBSD y macOS), el rango de valores es [-20, 19], y cuanto mayor sea el número, menor será la prioridad.
Para Windows, consulta [la documentación oficial](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass).
Se recomienda no especificar un valor, ya que el valor predeterminado es la prioridad más baja, lo que permitirá que todos los demás procesos tengan prioridad.

*Todas las funciones que hayas configurado se ejecutarán inmediatamente una vez que inicies el programa, para que puedas ver el efecto.*

## Docker - Implementación
1. Descarga `Dockerfile`
```shell
wget https://raw.githubusercontent.com/jennisu/lunarNeverIdle/master/Dockerfile
```
2. Contruye la imagen
```shell
# Máquina ARM
docker build -t lunarneveridle:latest .
# Especificaciones de máquina AMD ARCH=amd64
docker build --build-arg ARCH=amd64 -t lunarneveridle:latest .
```
3. Ejecuta
```bash
# Los parámetros de comando son iguales a los de arriba
docker run -d --name lunarneveridle lunarneveridle:latest -c 1h -m 2 -n 4h
```
