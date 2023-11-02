document.addEventListener("DOMContentLoaded", function (event) {
  //verifico si tiene el parametro id
  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");
  const operacion = urlParams.get("tipo");

  if (idEnvio != "" && idEnvio != null && operacion == "INICIAR") {
    iniciarViaje(idEnvio);
  } else if (idEnvio != "" && idEnvio != null && operacion == "FINALIZAR") {
    finalizarViaje(idEnvio);
  } else {
    document
      .getElementById("form")
      .addEventListener("submit", function (event) {
        guardarEnvio(event);
      });

    obtenerPedidos();
  }
});

//obtiene los productos para mostrar en el form de crear
function obtenerPedidos() {
  const urlConFiltro = `http://localhost:8080/pedidos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorPedido
  );
}

function exitoObtenerPedidos(data) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("tablePedidos")
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

var productos = [];

function obtenerPedidosArray() {
  var PedidosSeleccionados = [];
  let checkboxes = document.querySelectorAll(".pedido-checkbox");
  checkboxes.forEach(function (checkbox) {
    if (checkbox.checked) {
      // Agregar el producto seleccionado al objeto ProductosSeleccionados
      var tr = checkbox.closest("tr");

      var idPedido = tr.cells[1].textContent;
      var ciudadDestino = tr.cells[2].textContent;
      var estado = tr.cells[3].textContent;
      var fechaCreacion = tr.cells[4].textContent;
      var fechaUltimaActualizacion = tr.cells[5].textContent;
      var idCreador = tr.cells[6].textContent;
      var idPedido = tr.cells[7].textContent;

      var pedidoSeleccionado = {
        id: idPedido,
        fecha_creacion: fechaCreacion,
        fecha_ultima_actualizacion: fechaUltimaActualizacion,
        ciudad_destino: ciudadDestino,
        productos_elegidos: productos,
        id_creador: idCreador,
        estado: estado,
      };

      PedidosSeleccionados.push(pedidoSeleccionado);
    }
  });

  return PedidosSeleccionados;
}

function obtenerProductosDelPedido(idPedido) {
  const urlConFiltro = `http://localhost:8080/pedidos`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos(data, idPedido),
    errorPedido(response)
  );
}

function exitoObtenerPedidos(data, idPedido) {
  data.forEach((elemento) => {
    if (elemento.idPedido == idPedido) {
      if (productos.length > 0) {
        productos = [];
      }
      productos.push(elemento.productos_elegidos);
    }
  });
}

const urlConFiltro = `http://localhost:8080/envios`;

function guardarEnvio() {
  const pedidosArray = obtenerPedidosArray();

  //armo la data a enviar
  const data = {
    id: "",
    fecha_creacion: "2023-10-14T12:00:00Z",
    fecha_ultima_actualizacion: "2023-10-14T12:00:00Z",
    patente_camion: document.getElementById("PatenteCamion").value,
    paradas: [],
    pedidos: pedidosArray,
    id_creador: parseInt(document.getElementById("IdCreador").value),
    estado: "ADespachar",
  };

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoEnvio,
    errorEnvio
  );
}

function exitoEnvio(data) {
  window.location = window.location.origin + "/web/envios/index.html";
}

function errorEnvio(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function iniciarViaje(id) {
  if (confirm("¿Estás seguro de que deseas iniciar el viaje?")) {
    makeRequest(
      `${urlConFiltro}/${id}/cambiarEstado?estado=EnRuta`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}

function finalizarViaje(id) {
  if (confirm("¿Estás seguro de que deseas finalizar el viaje?")) {
    makeRequest(
      `${urlConFiltro}/${id}/cambiarEstado?estado=Despachado`,
      Method.PUT,
      data,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoEnvio,
      errorEnvio
    );
  } else {
    window.location = document.location.origin + "/web/envios/index.html";
  }
}
