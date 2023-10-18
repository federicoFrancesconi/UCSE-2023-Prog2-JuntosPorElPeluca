const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  obtenerEnvios();
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
      data.forEach((elemento) => {
        const row = document.createElement("tr"); //crear una fila

        row.innerHTML = ` 
                    <td>${elemento.Id}</td>
                    <td>${elemento.FechaCreacion}</td>
                    <td>${elemento.FechaUltimaActualizacion}</td>
                    <td>${elemento.PatenteCamion}</td>
                    <td>
                        <table>
                            <tr>
                                <th>Ciudad</th>
                                <th>Km Recorridos</th>
                            </tr>
                            ${elemento.Paradas.map(
                              (parada) => `
                                <tr>
                                    <td>${parada.Ciudad}</td>
                                    <td>${parada.KmRecorridos}</td>
                                </tr>
                            `
                            ).join("")}
                        </table>
                    </td>
                    <td class="acciones"></td>
                `;

        elementosTable.appendChild(row);
      });
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}
