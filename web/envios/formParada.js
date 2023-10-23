const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  debugger;
  document
    .getElementById("formParada")
    .addEventListener("submit", function (event) {
      agregarParada(event);
    });
});

function agregarParada() {
  debugger;
  const data = {
    Ciudad: document.getElementById("Ciudad").value,
    KmRecorridos: parseInt(document.getElementById("KmRecorridos").value),
  };

  const idCamion = document.getElementById("IdCamion").value;

  fetch(`http://localhost:8080/${idCamion}/nuevaParada`, {
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

  return false;
}
