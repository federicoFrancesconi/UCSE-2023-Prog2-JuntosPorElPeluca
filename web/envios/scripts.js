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

  document
    .getElementById("FiltrarPatente")
    .addEventListener("click", function (event) {
      obtenerEnvioFiltrado("patente");
    });

  document
    .getElementById("FiltrarEstado")
    .addEventListener("click", function (event) {
      obtenerEnvioFiltrado("estado");
    });

  document
    .getElementById("FiltrarCiudad")
    .addEventListener("click", function (event) {
      obtenerEnvioFiltrado("ciudad");
    });

  document
    .getElementById("FiltrarFecha")
    .addEventListener("click", function (event) {
      obtenerEnvioFiltrado("fecha");
    });

  obtenerEnvios();
});

function obtenerEnvios() {
  urlConFiltro = `http://localhost:8080/envios`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerEnvio,
    errorEnvio
  );
}

function exitoObtenerEnvio(data) {
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
                <td>${elemento.fecha_creacion}</td>
                <td>${elemento.fecha_ultima_actualizacion}</td>
                <td>${elemento.patente_camion}</td>
                <td>
                    <table>
                        <tr>
                            <th>Ciudad</th>
                            <th>Km Recorridos</th>
                        </tr>
                        ${
                          elemento.paradas
                            ? elemento.paradas
                                .map(
                                  (parada) => `
                            <tr>
                                <td>${parada.ciudad}</td>
                                <td>${parada.km_recorridos}</td>
                            </tr>
                        `
                                )
                                .join("")
                            : `<tr><td>No hay paradas disponibles</td></tr>`
                        }
                    </table>
                </td>
                <td>${
                  elemento.pedidos
                    ? elemento.pedidos
                        .map(
                          (pedido) => `
                    ${pedido}
                `
                        )
                        .join(" ")
                    : `No hay pedidos disponibles`
                }</td>
                <td>${elemento.id_creador}</td>
                <td>${elemento.estado}</td>
                <td class="acciones"><a class="anchorVerde" href="/web/envios/nuevaParada.html?id=${
                  elemento.id
                }">Nueva Parada</a> <a class="anchorVerde" href="form.html?id=${
        elemento.id
      }&tipo=INICIAR">Iniciar Viaje</a> <a class="anchorRojo" href="form.html?id=${
        elemento.id
      }&tipo=FINALIZAR">Finalizar Viaje</a></td>
                `;

      elementosTable.appendChild(row);
    });
  }
}

function errorEnvio(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function obtenerEnvioFiltrado(tipo) {
  var url = new URL(urlConFiltro);

  switch (tipo) {
    case "patente":
      url.searchParams.set(
        "patente",
        document.getElementById("FiltroPatente").value
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
        document.getElementById("FechaDesde").value + "T00:00:00.00Z"
      );
      url.searchParams.set(
        "fechaCreacionFin",
        document.getElementById("FechaHasta").value + "T00:00:00.00Z"
      );
      break;
    case "ciudad":
      url.searchParams.set(
        "ultimaParada",
        document.getElementById("FiltroCiudad").value
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
    exitoObtenerEnvio,
    errorEnvio
  );
}
