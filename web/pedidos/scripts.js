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

  obtenerPedidos();

  document
    .getElementById("AplicarFiltros")
    .addEventListener("click", function (event) {
      obtenerPedidoFiltrado();
    });
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

var url = new URL(urlConFiltro);

function obtenerPedidoFiltrado() {
  if (document.getElementById("FiltroId").value != '') {
    url.searchParams.set(
      "idEnvio",
      document.getElementById("FiltroId").value
    );
  }

  if (document.getElementById("FiltroEstado").value != '') {
    url.searchParams.set(
      "estado",
      document.getElementById("FiltroEstado").value
    );
  }

  if (document.getElementById("FechaDesde").value != '' && document.getElementById("FechaHasta").value != '') {
    url.searchParams.set(
      "fechaCreacionComienzo",
      document.getElementById("FechaDesde").value
    );

    url.searchParams.set(
      "fechaCreacionFin",
      document.getElementById("FechaHasta").value
    );
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

  elementosTable.innerHTML = "";

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
                      ${producto.nombre_producto}
                  `
                          )
                          .join(" ")
                      : `No hay productos disponibles`
                  }</td>
                  <td>${elemento.ciudad_destino}</td>
                  <td>${elemento.estado}</td>
                  <td>${elemento.fecha_creacion}</td>
                  <td>${elemento.fecha_ultima_actualizacion}</td>
                  <td>${elemento.id_creador}</td>
                  <td class="acciones"> <a href="form.html?id=${
                    elemento.id
                  }&tipo=ACEPTAR" class="anchorVerde">Aceptar Pedido</a> <a class="anchorRojo" href="form.html?id=${
        elemento.id
      }&tipo=CANCELAR">Cancelar Pedido</a></td>
                  `;

      elementosTable.appendChild(row);
    });
  }
}

function errorObtenerPedidos(response) {
  alert(response.Error);
  console.log(response.json());
  throw new Error(response.Error);
}
