const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  if (!isUserLogged()) {
    window.location =
      document.location.origin + "/web/login/login.html?reason=login_required";
  }

  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const codProducto = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (codProducto != "" && codProducto != null && operacion == "ELIMINAR") {
    eliminarProducto(codProducto);
  } else if (
    codProducto != "" &&
    codProducto != null &&
    operacion == "EDITAR"
  ) {
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
        actualizarProducto(event);
      });

    obtenerProductoPorId(codProducto);
  } else {
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
        guardarProducto(event);
      });
  }
});

const urlConFiltro = `http://localhost:8080/productos`;

function guardarProducto() {
  //armo la data a enviar
  const data = {
    codigo_producto: "",
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    tipo_producto: document.getElementById("TipoProducto").value,
    nombre: document.getElementById("Nombre").value,
    peso_unitario: parseFloat(document.getElementById("PesoUnitario").value),
    precio_unitario: parseFloat(
      document.getElementById("PrecioUnitario").value
    ),
    stock_minimo: parseInt(document.getElementById("StockMinimo").value),
    stock_actual: parseInt(document.getElementById("StockActual").value),
    id_creador: "",
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoProducto,
    errorProducto
  );
}

function actualizarProducto() {
  const data = {
    codigo_producto: "",
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    tipo_producto: document.getElementById("TipoProducto").value,
    nombre: document.getElementById("Nombre").value,
    peso_unitario: parseFloat(document.getElementById("PesoUnitario").value),
    precio_unitario: parseFloat(
      document.getElementById("PrecioUnitario").value
    ),
    stock_minimo: parseInt(document.getElementById("StockMinimo").value),
    stock_actual: parseInt(document.getElementById("StockActual").value),
    id_creador: "",
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.PUT,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoProducto,
    errorProducto
  );
}

function exitoProducto(data) {
  alert("Operacion exitosa");
  window.location = window.location.origin + "/web/productos/index.html";
}

function errorProducto(response) {
  alert(`Error del servidor: ${response.error}`);
  console.log(response.json());
  throw new Error(response.Error);
}

function eliminarProducto(codProducto) {
  if (confirm("¿Estás seguro de que deseas eliminar este producto?")) {
    makeRequest(
      `${urlConFiltro}/${codProducto}`,
      Method.DELETE,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoProducto,
      errorProducto
    );
  } else {
    window.location = document.location.origin + "/web/productos/index.html";
  }
}

function obtenerProductoPorId(codProducto) {
  makeRequest(
    `${urlConFiltro}/${codProducto}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProducto,
    errorProducto
  );
}

function exitoObtenerProducto(data) {
  document.getElementById("TipoProducto").value = data.tipo_producto;
  document.getElementById("Nombre").value = data.nombre;
  document.getElementById("PesoUnitario").value = data.peso_unitario;
  document.getElementById("PrecioUnitario").value = data.precio_unitario;
  document.getElementById("StockMinimo").value = data.stock_minimo;
  document.getElementById("StockActual").value = data.stock_actual;
}
