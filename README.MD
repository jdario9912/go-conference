# Notas sobre tutorial Golang

Este archivo contiene algunas notas sobre el tutorial de Go.

El codigo desarrollado es sobre una aplicaciòn de consola que simula la generaciòn de tickets para una conferencia llamada "Go conference".

El objetivo era conocer los conceptos bàsicos del lenguaje: definiciòn de variables, tipos, estructuras condicionales, bucles, funciones, hilos de ejecuciòn y paquetes.

Ver el tutorial [aquì](https://www.youtube.com/watch?v=yyUHQIec83I&t=33s "Go for beginners").

## Estructuras de tipos de datos

### List:

Es un array con una longitud màxima establecida previamente y que puede contener un solo tipo de dato.

#### Pro

Evita guardar una cantidad inesperada de datos.

#### Contra

Cualquier lugar definido y no usado es un espacio vacio en memoria inutil.

#### Ejemplo

```golang
var myListStrings [10]string
var myListInts [5]int
var myListUints [17]uint
```

### Slice:

Es un array en el cual su longitud màxima no esta definida.

#### Pro

Se pueden guardar una cantidad muy grande de datos.

#### Contra

Si se guardan muchos datos afectarà el rendimiento de las operaciones que se le hagan. Por ejemplo en la busqueda de un valor.

#### Ejemplo

```golang
var mySliceStrings []string
var mySliceInts []int
var mySliceUints []uint
```

### Map

Es una estructura de tipo key-value en el que cada uno tiene un tipo definido. A su vez, la estructura se le pasa a la funcion 'make'

#### Ejemplo

```golang
var userData = make(map[string]string)

userData["firstName"] = "Joel"
userData["lastName"] = "Dario"
userData["email"] = "joel@dario.com"
```

Si se necesita guardar un nùmero en un map que tiene un tipo de dato string, deberia ser convertido a string previamente.

```golang
age := 30

userData["age"] = strconv.FormatUint(uint64(age), 10)
```

#### Lista de Maps

Se utiliza la funcion make que recibe 2 parametros: la estructura del map y la longitud inicial de la lista.

```golang
var myList = make([]map[string]string, 0)
```

### Struct

Se utiliza para definir una estructura de dato mediante key-value y puede contener cualquier tipo de dato.
Se define con la palabra 'type' y asigna un nombre.

```golang
type userData struct {
  firstName string
  lastName string
  email string
  numberOfTicket uint
}
```
