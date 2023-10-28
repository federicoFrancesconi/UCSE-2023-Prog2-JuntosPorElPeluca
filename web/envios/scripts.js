const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerEnvios();
  document.getElementById("form").addEventListener("submit", function (event) {
    obtenerBeneficioEntreFechas(event);
  });
});

function obtenerEnvios() {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  urlConFiltro = `http://localhost:8080/envios`;

  fetch(urlConFiltro, {
    method: "GET",
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        alert("Error en la solicitud al servidor.");
        console.log(response.json());
        throw new Error("Error en la solicitud al servidor.");
      }
      return response.json();
    })
    .then((data) => {
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
                    <td class="acciones"><a href="/web/envios/nuevaParada.html?id=${
                      elemento.id
                    }">Nueva Parada</a> | <a href="form.html?id=${
            elemento.id
          }&tipo=INICIAR">Iniciar Viaje</a> | <a href="form.html?id=${
            elemento.id
          }&tipo=FINALIZAR">Finalizar Viaje</a></td>
                    `;

          elementosTable.appendChild(row);
        });
      }
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}

function obtenerEnvioPorId() {
  console.log("obtenerEnvioPorId");

  const id = document.getElementById("FiltroId").value;

  const elementosTable = document //tabla en la que se colocan los camiones que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  urlConFiltro = `http://localhost:8080/envios/${id}`; //ver que url colocariamos

  fetch(urlConFiltro, {
    method: "GET",
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        alert("Error en la solicitud al servidor.");
        throw new Error("Error en la solicitud al servidor.");
      }
      return response.json();
    })
    .then((data) => {
      // Llenar la tabla con los datos obtenidos
      console.log(data);
      elementosTable.innerHTML = "";

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
                            ${elemento.paradas
                              .map(
                                (parada) => `
                                <tr>
                                    <td>${parada.ciudad}</td>
                                    <td>${parada.km_recorridos}</td>
                                </tr>
                            `
                              )
                              .join("")}
                        </table>
                    </td>
                    <td>${elemento.pedidos
                      .map(
                        (pedido) => `
                        ${pedido}
                    `
                      )
                      .join(" ")}</td>
                    <td>${elemento.id_creador}</td>
                    <td>${elemento.estado}</td>
                    <td class="acciones"><a href="form.html?id=${
                      elemento.id
                    }&tipo=PARADA">Nueva Parada</a> | <a href="form.html?id=${
        elemento.id
      }&tipo=INICIAR">Iniciar Viaje</a> | <a href="form.html?id=${
        elemento.id
      }&tipo=FINALIZAR">Finalizar Viaje</a></td>
                    `;

      elementosTable.appendChild(row);
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}

function obtenerBeneficioEntreFechas() {
  var fechaDesde = document.getElementById("FechaDesde").value;
  var fechaHasta = document.getElementById("FechaHasta").value;

  fetch(
    `http://localhost:8080/envios/beneficioEntreFechas?fechaDesde=${fechaDesde}Z&fechaHasta=${fechaHasta}Z`,
    {
      method: "POST",
      body: JSON.stringify(data),
      headers: customHeaders,
    }
  ) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      document.getElementById("beneficio").innerHTML = response.json();
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
      window.location = "/web/envios/form.html";
    });
}
