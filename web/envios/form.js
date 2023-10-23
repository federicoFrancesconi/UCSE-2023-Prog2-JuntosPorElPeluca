const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

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
  }
});

function guardarEnvio() {
  const data = {
    Id: 0,
    FechaCreacion: "2023-10-14T12:00:00Z",
    FechaUltimaActualizacion: "2023-10-14T12:00:00Z",
    PatenteCamion: document.getElementById("PatenteCamion").value,
    Paradas: [],
    Pedidos: [],
    IdCreador: parseInt(document.getElementById("IdCreador").value),
    Estado: 0,
  };

  fetch(`http://localhost:8080/envios`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/envios/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}

function iniciarViaje(id) {
  fetch(`http://localhost:8080/${id}/iniciar`, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/envios/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}

function finalizarViaje(id) {
  fetch(`http://localhost:8080/${id}/finalizar`, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: customHeaders,
  }) // Realizar la solicitud de búsqueda (fetch) al servidor
    .then((response) => {
      if (!response.ok) {
        throw new Error("Error en la solicitud al servidor.");
      }

      window.location = "/web/envios/index.html";
    })
    .catch((error) => {
      console.error("Error:", error);
      alert(error);
    });
}
