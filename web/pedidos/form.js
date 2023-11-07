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
  const idPedido = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idPedido != "" && idPedido != null && operacion == "ACEPTAR") {
    aceptarPedido(idPedido);
  } else if (idPedido != "" && idPedido != null && operacion == "CANCELAR") {
    cancelarPedido(idPedido);
  } else {
    document
      .getElementById("buttonSave")
      .addEventListener("click", function (event) {
        guardarPedido(event);
      });

    obtenerProductos();
  }
});

//obtiene los productos para mostrar en el form de crear
function obtenerProductos() {
  const urlConFiltro = `http://localhost:8080/productos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProductos,
    errorPedido
  );
}

function exitoObtenerProductos(data) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("tableProductos")
    .querySelector("tbody");

  data.forEach((elemento) => {
    const row = document.createElement("tr"); //crear una fila

    row.innerHTML = ` 
                  <td><input type="checkbox" class="producto-checkbox"></td>
                  <td>${elemento.codigo_producto}</td>
                  <td>${elemento.nombre}</td>
                  <td><input type="text" placeholder="Ingrese la cantidad"></td>
                  <td>${elemento.precio_unitario}</td>
                  <td>${elemento.peso_unitario}</td>
                 `;

    elementosTable.appendChild(row);
  });
}

function obtenerProductosElegidos() {
  var ProductosSeleccionados = [];
  let checkboxes = document.querySelectorAll(".producto-checkbox");
  checkboxes.forEach(function (checkbox) {
    if (checkbox.checked) {
      // Agregar el producto seleccionado al objeto ProductosSeleccionados
      var tr = checkbox.closest("tr");
      var codigoProducto = tr.cells[1].textContent;
      var nombreProducto = tr.cells[2].textContent;
      var cantidad = parseInt(
        tr.cells[3].getElementsByTagName("input")[0].value
      );
      var precioUnitario = parseFloat(tr.cells[4].textContent);
      var pesoUnitario = parseFloat(tr.cells[5].textContent);

      var productoSeleccionado = {
        codigo_producto: codigoProducto,
        nombre_producto: nombreProducto,
        cantidad: cantidad,
        precio_unitario: precioUnitario,
        peso_unitario: pesoUnitario,
      };

      ProductosSeleccionados.push(productoSeleccionado);
    }
  });

  return ProductosSeleccionados;
}

function guardarPedido() {
  //armo la data a enviar
  const data = {
    id: "",
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    ciudad_destino: document.getElementById("CiudadDestino").value,
    productos_elegidos: obtenerProductosElegidos(),
    id_creador: 0,
    estado: "Pendiente",
  };

  const urlConFiltro = `http://localhost:8080/pedidos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoPedido,
    errorPedido
  );
}

function exitoPedido(data) {
  window.location = window.location.origin + "/web/pedidos/index.html";
}

function errorPedido(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function aceptarPedido(id) {
  if (confirm("¿Estás seguro de que deseas aceptar este pedido?")) {
    const urlConFiltro = `http://localhost:8080/pedidos/${id}/aceptar`;

    data = [];

    makeRequest(
      `${urlConFiltro}`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoPedido,
      errorPedido
    );
  } else {
    window.location = document.location.origin + "/web/pedidos/index.html";
  }
}

function cancelarPedido(id) {
  if (confirm("¿Estás seguro de que deseas cancelar este pedido?")) {
    const urlConFiltro = `http://localhost:8080/pedidos/${id}/cancelar`;
    data = [];
    makeRequest(
      `${urlConFiltro}`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoPedido,
      errorPedido
    );
  } else {
    window.location = document.location.origin + "/web/pedidos/index.html";
  }
}
