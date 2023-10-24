const url = ""; //definir la url

const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerCamiones();
});

function obtenerCamiones() {
  const elementosTable = document //tabla en la que se colocan los camiones que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  urlConFiltro = `http://localhost:8080/camiones`; //ver que url colocariamos

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
      data.forEach((elemento) => {
        const row = document.createElement("tr"); //crear una fila

        row.innerHTML = ` 
                  <td>${elemento.patente}</td>
                  <td>${elemento.pesoMaximo}</td>
                  <td>${elemento.fechaCreacion}</td>
                  <td>${elemento.fechaUltimaActualizacion}</td>
                  <td>${elemento.costoPorKilometro}</td>
                  <td>${elemento.idCreador}</td>
                  <td class="acciones"><a href="form.html?patente=${elemento.patente}&tipo=EDITAR">Editar</a> | <a href="form.html?patente=${elemento.patente}&tipo=ELIMINAR">Eliminar</a></td>
              `;

        elementosTable.appendChild(row);
      });
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}

function obtenerCamionPorPatente() {
  console.log("obtenerCamionPorPatente");

  const patente = document.getElementById("FiltroPatente").value;

  const elementosTable = document //tabla en la que se colocan los camiones que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  urlConFiltro = `http://localhost:8080/camiones/${patente}`; //ver que url colocariamos

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
                  <td>${data.Patente}</td>
                  <td>${data.PesoMaximo}</td>
                  <td>${data.FechaCreacion}</td>
                  <td>${data.FechaUltimaActualizacion}</td>
                  <td>${data.IdCreador}</td>
                  <td class="acciones"><a href="form.html?patente=${data.Patente}&tipo=EDITAR">Editar</a> | <a href="#" onclick="eliminarCamion('${data.Patente}')">Eliminar</a></td>
              `; //crear una celda por cada campo que quiera mostrar

      elementosTable.appendChild(row);
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}
