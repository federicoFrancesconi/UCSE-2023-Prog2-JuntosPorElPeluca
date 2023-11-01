const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerPedidos();
});

var urlConFiltro = `http://localhost:8080/pedidos`;

function obtenerPedidos() {
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function obtenerPedidoFiltrado(tipo) {
  var url = new URL(urlConFiltro);

  switch (tipo) {
    case "id":
      url.searchParams.set(
        "idEnvio",
        document.getElementById("FiltroId").value
      );
      break;
    case "estado":
      url.searchParams.set(
        "estado",
        document.getElementById("FiltroEstado").value
      );
      break;
    case "fecha":
      url.searchParams.set(
        "fechaCreacionComienzo",
        document.getElementById("FechaDesde").value
      );
      url.searchParams.set(
        "fechaCreacionFin",
        document.getElementById("FechaHasta").value
      );
      break;
    default:
      url = `http://localhost:8080/pedidos`;
      break;
  }

  makeRequest(
    `${url}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function exitoObtenerPedidos(data) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  // Llenar la tabla con los datos obtenidos
  if (data != null) {
    data.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila

      row.innerHTML = ` 
                  <td>${elemento.id}</td>
                  <td>${
                    elemento.productos_elegidos
                      ? elemento.productos_elegidos
                          .map(
                            (producto) => `
                      ${producto.nombre}
                  `
                          )
                          .join(" ")
                      : `No hay productos disponibles`
                  }</td>
                  <td>${elemento.ciudad_destino}</td>
                  <td>${elemento.estado}</td>
                  <td>${elemento.fecha_creacion}</td>
                  <td>${elemento.fecha_utlima_actualizacion}</td>
                  <td>${elemento.id_creador}</td>
                  <td class="acciones"> <a href="form.html?id=${
                    elemento.id
                  }&tipo=ACEPTAR">Aceptar Pedido</a> | <a href="form.html?id=${
        elemento.id
      }&tipo=CANCELAR">Cancelar Pedido</a></td>
                  `;

      elementosTable.appendChild(row);
    });
  }
}

function errorObtenerPedidos(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}
