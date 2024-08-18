# go-farm
Este es un proyecto con objetivos didácticos, se estudian los siguientes puntos:

1. Goroutines
2. Paso por valor
3. Paso por referencia
4. Librería Tview (k9s) para mostrar datos.
5. Compilación del proyecto y lanzado del mismo con dagger.io
6. Compilación del proyecto en wasm y lanzamiento en navegador

## modo terminal
```shell
dagger -m dagger-wasm/ call desktop-mode --project-dir="."
```
![image](https://github.com/user-attachments/assets/0644c220-25d9-433d-b084-84abe314e1b7)

## modo wasm

Wasm tiene ciertas limitaciones, por tanto se ha realizado una implementación específica para ver como corre 
en la consola del navegador.

```shell
dagger -m ./dagger-wasm call wasm-mode --project-dir="." as-service up --ports=8080:80
```

Al lanzar este comando y ver que el proyecto se lanza, podemos ver que se expone al puerto `8080` de nuestro `localhost`
![image](https://github.com/user-attachments/assets/c7ba2ffd-5b16-4b06-b400-d20a4e6b6094)

Debemos acceder a nuestro `http://localhost:8080/` y abrir las herramientas de desarrollador y ver cómo corre el programa.

[Screencast from 2024-08-18 19-30-40.webm](https://github.com/user-attachments/assets/3e622377-dcc9-4390-929b-b8bd4ece4f02)
