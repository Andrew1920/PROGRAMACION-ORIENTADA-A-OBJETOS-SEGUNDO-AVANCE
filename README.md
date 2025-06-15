# E-commerce en Go

Este proyecto fue desarrollado como parte de la asignatura **Programación Orientada a Objetos**.

* **Estudiante:** Andrew Chavez Hernandez

---

## 1. Descripción del Proyecto

El objetivo principal de este proyecto es el desarrollo de un sistema de gestión de e-commerce robusto y escalable utilizando el lenguaje de programación Go. La API implementa las funcionalidades esenciales de una tienda en línea, sirviendo como una base sólida y extensible para futuras integraciones como pasarelas de pago, gestión de inventario y facturación.

---

## 2. Características Implementadas (Segundo Avance)

En esta fase, la API es completamente funcional y cuenta con los siguientes módulos:

#### Módulo de Productos (100% Completado)
Se ha implementado un sistema **CRUD (Create, Read, Update, Delete)** completo para la gestión de productos.
* **Crear:** `POST /api/products` (para un producto) y `POST /api/products/batch` (para múltiples productos).
* **Leer:** `GET /api/products` (todos los productos) y `GET /api/products/{id}` (un producto específico).
* **Actualizar:** `PUT /api/products/{id}`.
* **Eliminar:** `DELETE /api/products/{id}`.

#### Módulo de Carrito de Compras (100% Completado)
Se ha desarrollado un sistema de carrito de compras funcional basado en un ID de sesión único, permitiendo:
* Crear un nuevo carrito (`POST /api/cart`).
* Añadir productos a un carrito específico (`POST /api/cart/{cartId}/add`).
* Consultar el contenido del carrito (`GET /api/cart/{cartId}`).
* Vaciar el carrito para simular una compra (`DELETE /api/cart/{cartId}`).

#### Módulo de Usuarios (Funcionalidad Base)
Se construyó la base para la gestión de usuarios, que incluye:
* **Registro Seguro:** Creación de usuarios con contraseñas encriptadas mediante `bcrypt`.
* **Verificación de Credenciales:** Un endpoint de `login` que valida el usuario y la contraseña.
*(Nota: La autenticación por tokens JWT fue desarrollada y posteriormente desactivada para simplificar las pruebas de los otros módulos durante la presentación).*

---

## 3. Principios de Arquitectura y Diseño Aplicados

Este proyecto demuestra varios conceptos clave de la ingeniería de software moderna:

* **Arquitectura Modular:** El código está organizado por responsabilidades en paquetes (`handlers`, `models`, `routes`, `storage`, `utils`), aplicando el principio de encapsulación.
* **Uso de Interfaces para Desacoplamiento:** Se define una interfaz `ProductStorer` que abstrae la capa de almacenamiento. Esto permite que los *handlers* no dependan de una implementación concreta (como el mapa en memoria actual), facilitando el cambio a una base de datos real en el futuro sin modificar la lógica de negocio.
* **Inyección de Dependencias:** El archivo `main.go` actúa como el "orquestador", creando las dependencias (el almacén) y "inyectándolas" en los componentes que las necesitan (los *handlers*).
* **Manejo de Concurrencia:** Se utiliza `sync.Mutex` en la capa de almacenamiento para proteger el acceso a los datos en memoria, asegurando que la API pueda manejar múltiples peticiones simultáneas de forma segura.

---

## 4. Estructura del Proyecto
```
/tienda
├── go.mod
├── go.sum
├── main.go
├── /handlers
│   ├── cart_handler.go
│   ├── product_handler.go
│   └── user_handler.go
├── /models
│   ├── cart.go
│   ├── product.go
│   └── user.go
├── /routes
│   └── routes.go
├── /storage
│   ├── interfaces.go
│   └── memory.go
└── /utils
└── password.go
```
## 5. Endpoints de la API

| Método | Ruta                                  | Descripción                               |
| :----- | :------------------------------------ | :---------------------------------------- |
| `POST` | `/register`                           | Registra un nuevo usuario.                |
| `POST` | `/login`                              | Inicia sesión (verifica credenciales).    |
| `GET`  | `/api/products`                       | Obtiene la lista de todos los productos.  |
| `POST` | `/api/products`                       | Crea un nuevo producto.                   |
| `POST` | `/api/products/batch`                 | Crea múltiples productos a la vez.        |
| `GET`  | `/api/products/{id}`                  | Obtiene un producto por su ID.            |
| `PUT`  | `/api/products/{id}`                  | Actualiza un producto por su ID.          |
| `DELETE`| `/api/products/{id}`                  | Elimina un producto por su ID.            |
| `POST` | `/api/cart`                           | Crea un nuevo carrito de compras vacío.   |
| `GET`  | `/api/cart/{cartId}`                  | Obtiene el contenido de un carrito.       |
| `POST` | `/api/cart/{cartId}/add`              | Añade un producto a un carrito.           |
| `DELETE`| `/api/cart/{cartId}`                  | Elimina (vacía) un carrito.               |

---

## 6. Cómo Ejecutar el Proyecto

#### Requisitos
* Tener **Go**.
* link : http://localhost:8080/
* http://localhost:8080/api/products
* 
## Pruebas con Postman
Para la verificación y prueba de todos los endpoints de la API se utilizó la herramienta Postman.
Se ha creado y adjuntado en el repositorio un archivo de Colección de Postman llamado localhost-8080.postman_collection.json. Este archivo contiene todas las peticiones (GET, POST, PUT, DELETE) pre-configuradas con sus URLs y cuerpos JSON de ejemplo
