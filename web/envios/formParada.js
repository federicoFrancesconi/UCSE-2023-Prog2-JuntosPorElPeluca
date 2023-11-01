const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  document
    .getElementById("formParada")
    .addEventListener("submit", function (event) {
      agregarParada(event);
    });
});

function agregarParada() {
  const data = {
    ciudad: document.getElementById("Ciudad").value,
    km_recorridos: parseInt(document.getElementById("KmRecorridos").value),
  };

  const urlParams = new URLSearchParams(window.location.search);
  const idEnvio = urlParams.get("id");

  const urlConFiltro = `http://localhost:8080/${idEnvio}/nuevaParada`;

  makeRequest(
    `${urlConFiltro}`,
    Method.POST,
    data,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoAgregarParada,
    errorAgregarParada
  );
}

function exitoAgregarParada(data) {
  window.location = document.location.origin + "/web/envios/index.html";
}

function errorAgregarParada(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}
