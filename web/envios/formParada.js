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

  debugger;
  fetch(`http://localhost:8080/${idEnvio}/nuevaParada`, {
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
